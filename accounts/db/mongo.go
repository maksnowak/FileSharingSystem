package db

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log"
	"os"
	"time"
)

// # Client
//
// A MongoDB client handle
var Client *mongo.Client

// # Connect
//
// A function to establish a connection with the database
func Connect() {
	// create context
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// open connection
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}
	client, err := mongo.Connect(options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		log.Fatalln("could not connect to MongoDB")
	}
	Client = client
	log.Println("Connected to MongoDB")
}

// # Disconnect
//
// A function to close database connection
func Disconnect() {
	if Client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err := Client.Disconnect(ctx)
		if err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}
		log.Println("Disconnected from MongoDB")
	}
}

// # GetCollection
//
// A function to get a certain collection from the database
func GetCollection(collectionName string) *mongo.Collection {
	return Client.Database("mongodb").Collection(collectionName)
}
