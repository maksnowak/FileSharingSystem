package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func initMongo(ctx *context.Context) (*mongo.Collection, *mongo.Client) {
	_ = godotenv.Load("../.env")

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	client, err := mongo.Connect(*ctx, clientOptions)
	if err != nil {
		panic(fmt.Sprintf("Mongo DB Connect issue %s", err))
	}

	err = client.Ping(*ctx, nil)
	if err != nil {
		panic(fmt.Sprintf("Mongo DB ping issue %s", err))
	}

	fmt.Print("Connected to MongoDB")
	collection := client.Database("files").Collection("Files")
	return collection, client
}
