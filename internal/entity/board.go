package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Board struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
	CreateAt time.Time          `bson:"create_at"`
	CreateBy string             `bson:"create_by"`
	UpdateAt time.Time          `bson:"update_at"`
	UpdateBy string             `bson:"update_by"`
}
