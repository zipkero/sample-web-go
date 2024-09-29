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
		Username: config.Mongo.Username,
		Password: config.Mongo.Password,
	})

	client, err := mongo.Connect(context.TODO(), opt)

	if err != nil {
		return nil, err
	}

	return &MongoProvider{client}, nil
}

func (m *MongoProvider) FindOne(ctx context.Context, db, collection string, filter interface{}, result interface{}) error {
	c := m.client.Database(db).Collection(collection)
	return c.FindOne(ctx, filter).Decode(result)
}

func (m *MongoProvider) Find(ctx context.Context, db, collection string, filter interface{}) (*mongo.Cursor, error) {
	c := m.client.Database(db).Collection(collection)
	return c.Find(ctx, filter)
}
