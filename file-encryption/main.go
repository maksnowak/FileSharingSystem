package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"os"
	"time"
)

func main() {
	_ = godotenv.Load("../.env")

	client, err := mongo.Connect(options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		log.Fatalln("could not connect to mongo")
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := client.Database("mongodb").Collection("files-metadata")
	result, err := collection.InsertOne(ctx, map[string]any{
		"path": "example.file",
		"size": 123,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
