package main

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"github.com/zipkero/sample-web-go/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
)

func main() {
	cf, err := config.NewConfig("config.local.toml")
	if err != nil {
		log.Fatal(err)
	}

	mongoUri := cf.Mongo.URI
	redisUri := cf.Redis.URI

	if mongoUri == "" {
		log.Fatal("Mongo URI is empty")
	}

	if redisUri == "" {
		log.Fatal("Redis URI is empty")
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		opt := options.Client().ApplyURI(mongoUri).SetAuth(options.Credential{
			Username: "root",
			Password: "1234",
		})

		client, err := mongo.Connect(context.TODO(), opt)
		if err != nil {
			log.Fatal(err)
		}

		err = client.Ping(context.TODO(), nil)
		if err != nil {
			log.Fatal(err)
		}

		collection := client.Database("web").Collection("sample")

		// InsertOne
		_, err = collection.InsertOne(context.TODO(), map[string]string{"key": "value"})
		if err != nil {
			log.Fatal(err)
		}

		// GetOne
		var res map[string]string
		err = collection.FindOne(context.TODO(), map[string]string{"key": "value"}).Decode(&res)
		if err != nil {
			log.Fatal(err)
		}

		// DeleteOne
		_, err = collection.DeleteOne(context.TODO(), map[string]string{"key": "value"})
		if err != nil {
			log.Fatal(err)
		}

		// FindOne
		result := collection.FindOne(context.TODO(), map[string]string{"key": "value"})
		if result.Err() != nil && !errors.Is(result.Err(), mongo.ErrNoDocuments) {
			log.Fatal(result.Err())
		}

		log.Print(result.Decode(&res))

		defer func() {
			if err := client.Disconnect(context.TODO()); err != nil {
				log.Print(err)
			}
			wg.Done()
		}()
	}()

	go func() {
		opts, err := redis.ParseURL(redisUri)
		if err != nil {
			log.Fatal(err)
		}

		rdb := redis.NewClient(opts)

		ctx := context.Background()

		err = rdb.Set(ctx, "key", "value", 0).Err()
		if err != nil {
			log.Fatal(err)
		}

		val, err := rdb.Get(ctx, "key").Result()
		if err != nil {
			log.Fatal(err)
		}

		log.Println(val)

		defer func() {
			if err := rdb.Close(); err != nil {
				log.Print(err)
			}
			wg.Done()
		}()
	}()

	wg.Wait()
}
