package service

import (
	"context"
	"errors"
	"github.com/zipkero/sample-web-go/internal/db"
	"github.com/zipkero/sample-web-go/internal/entity"
	"github.com/zipkero/sample-web-go/pkg/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BoardService struct {
	provider *db.MongoProvider
}

var (
	dbName         = "web"
	collectionName = "board"
)

func NewBoardService(provider *db.MongoProvider) *BoardService {
	return &BoardService{provider}
}

func (b *BoardService) FindAll(ctx context.Context) ([]*dto.BoardResponse, error) {
	var boards []entity.Board
	cursor, err := b.provider.Find(ctx, dbName, collectionName, bson.D{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &boards); err != nil {
		return nil, err
	}

	var boardResponses []*dto.BoardResponse
	for _, board := range boards {
		boardResponses = append(boardResponses, b.ToDto(&board))
	}

	return boardResponses, nil
}

func (b *BoardService) FindOne(ctx context.Context, id string) (*entity.Board, error) {
	var board entity.Board
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = b.provider.FindOne(ctx, dbName, collectionName, bson.D{{
		"_id", objectId,
	}}, &board)
	return &board, err
}

func (b *BoardService) InsertOne(ctx context.Context, board *entity.Board) (string, error) {
	result, err := b.provider.InsertOne(ctx, dbName, collectionName, board)
	if err != nil {
		return "", err
	}
	objectId := result.InsertedID.(primitive.ObjectID)
	return objectId.Hex(), err
}

func (b *BoardService) UpdateOne(ctx context.Context, id string, board *entity.Board) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	result, err := b.provider.UpdateOne(ctx, dbName, collectionName, bson.D{{
		"_id", objectId,
	}}, bson.D{{
		"$set", board,
	}})
	if err != nil {
		return err
	}
	if result.ModifiedCount == 0 {
		return errors.New("not found")
	}
	return nil
}

func (b *BoardService) DeleteOne(ctx context.Context, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	result, err := b.provider.DeleteOne(ctx, dbName, collectionName, bson.D{{
		"_id", objectId,
	}})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("not found")
	}
	return nil
}

func (b *BoardService) InsertMany(ctx context.Context, boards []*entity.Board) ([]string, error) {
	var documents []interface{}
	for _, board := range boards {
		documents = append(documents, board)
	}

	result, err := b.provider.InsertMany(ctx, dbName, collectionName, documents)
	if err != nil {
		return nil, err
	}

	var ids []string
	for _, id := range result.InsertedIDs {
		ids = append(ids, id.(primitive.ObjectID).Hex())
	}

	return ids, err
}

func (b *BoardService) DeleteMany(ctx context.Context, ids []string) error {
	var objectIds []primitive.ObjectID
	for _, id := range ids {
		objectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return err
		}
		objectIds = append(objectIds, objectId)
	}

	filter := bson.M{"_id": bson.M{"$in": objectIds}}
	_, err := b.provider.DeleteMany(ctx, dbName, collectionName, filter)

	return err
}

func (b *BoardService) ToDto(board *entity.Board) *dto.BoardResponse {
	return &dto.BoardResponse{
		ID: board.ID.Hex(),
		BoardBase: dto.BoardBase{
			Title:    board.Title,
			Content:  board.Content,
			CreateAt: board.CreateAt,
			CreateBy: board.CreateBy,
			UpdateAt: board.UpdateAt,
			UpdateBy: board.UpdateBy,
		},
	}
}
