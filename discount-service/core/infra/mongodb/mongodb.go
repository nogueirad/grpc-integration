package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func Connect(uri string) *mongo.Client {
	var ctx = context.TODO()
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return client
}