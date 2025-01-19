package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func InitMongo(ctx *context.Context) (*mongo.Collection, *mongo.Client) {
	_ = godotenv.Load("./.env")

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		panic(fmt.Sprintf("Mongo DB Connect issue %s", err))
	}

	err = client.Ping(*ctx, nil)
	if err != nil {
		panic(fmt.Sprintf("Mongo DB ping issue %s", err))
	}

	log.Println("Connected to MongoDB")
	collection := client.Database("files").Collection("Files")
	return collection, client
}
