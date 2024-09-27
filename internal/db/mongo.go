package db

import (
	"context"
	"github.com/zipkero/sample-web-go/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoProvider struct {
	client *mongo.Client
}

func NewMongoProvider(config *config.Config) (*MongoProvider, error) {
	opt := options.Client().ApplyURI(config.Mongo.URI).SetAuth(options.Credential{
		Username: "root",
		Password: "1234",
	})

	client, err := mongo.Connect(context.TODO(), opt)

	if err != nil {
		return nil, err
	}

	return &MongoProvider{client}, nil
}
