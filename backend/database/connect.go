package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func Connect() error {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	err = client.Ping(context.Background(), nil)

	if err != nil {
		return err
	}

	log.Println("Connected to MongoDB!")

	DB = client.Database("secure-share")

	return nil
}

func GetLinkCollection() *mongo.Collection {
	collection := DB.Collection("links")

	return collection
}
