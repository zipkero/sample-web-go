package dto

import (
	"github.com/zipkero/sample-web-go/internal/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type BoardBase struct {
	Title    string    `json:"title" binding:"required"`
	Content  string    `json:"content" binding:"required"`
	CreateAt time.Time `json:"create_at"`
	CreateBy string    `json:"create_by"`
	UpdateAt time.Time `json:"update_at"`
	UpdateBy string    `json:"update_by"`
}

type BoardInsert struct {
	ID string `json:"id,omitempty"`
	BoardBase
}

type BoardUpdate struct {
	ID string `json:"id" binding:"required"`
	BoardBase
}

type BoardResponse struct {
	ID string `json:"id"`
	BoardBase
}

type BoardDeleteMany struct {
	IDs []string `json:"ids" binding:"required"`
}

func (b BoardInsert) ToEntity() *entity.Board {
	return &entity.Board{
		Title:    b.Title,
		Content:  b.Content,
		CreateAt: b.CreateAt,
		CreateBy: b.CreateBy,
		UpdateAt: b.UpdateAt,
		UpdateBy: b.UpdateBy,
	}
}

func (b BoardUpdate) ToEntity() (*entity.Board, error) {
	id, err := primitive.ObjectIDFromHex(b.ID)
	if err != nil {
		return nil, err
	}

	return &entity.Board{
		ID:       id,
		Title:    b.Title,
		Content:  b.Content,
		CreateAt: b.CreateAt,
		CreateBy: b.CreateBy,
		UpdateAt: b.UpdateAt,
		UpdateBy: b.UpdateBy,
	}, nil
}
