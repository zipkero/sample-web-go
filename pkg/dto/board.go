package dto

import "time"

type BoardBase struct {
	Title    string    `json:"title" binding:"required"`
	Content  string    `json:"content" binding:"required"`
	CreateAt time.Time `json:"create_at"`
	CreateBy string    `json:"create_by"`
	UpdateAt time.Time `json:"update_at"`
	UpdateBy string    `json:"update_by"`
}

type BoardInsert struct {
	ID int `json:"id,omitempty"`
	BoardBase
}

type BoardUpdate struct {
	ID int `json:"id" binding:"required"`
	BoardBase
}
