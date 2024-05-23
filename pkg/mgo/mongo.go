package mgo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func Connect(dsn string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		panic(any(err))
	}
	ctx, ff := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer ff()
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	return client
}
