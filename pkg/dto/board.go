package dto

import "time"

type Board struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	CreateAt time.Time `json:"create_at"`
	CreateBy string    `json:"create_by"`
	UpdateAt time.Time `json:"update_at"`
	UpdateBy string    `json:"update_by"`
}
