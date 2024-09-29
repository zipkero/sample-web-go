package service

import (
	"context"
	"github.com/zipkero/sample-web-go/internal/db"
	"github.com/zipkero/sample-web-go/pkg/entity"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type BoardService struct {
	provider *db.MongoProvider
}

func init() {
	log.Println("Board Service Init")
}

var (
	dbName         = "web"
	collectionName = "board"
)

func NewBoardService(provider *db.MongoProvider) *BoardService {
	return &BoardService{provider}
}

func (b *BoardService) FindOne(ctx context.Context, id int) (entity.Board, error) {
	var board entity.Board
	err := b.provider.FindOne(ctx, dbName, collectionName, bson.D{{
		"id", id,
	}}, &board)
	return board, err
}
