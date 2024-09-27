package service

import "github.com/zipkero/sample-web-go/internal/db"

type BoardService struct {
	provider *db.MongoProvider
}

func NewBoardService(provider *db.MongoProvider) *BoardService {
	return &BoardService{provider}
}
