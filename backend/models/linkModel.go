package models

import (
	"context"
	"time"

	"github.com/kazion500/secure-share/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Link struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	IsUsed      bool               `json:"isUsed" bson:"isUsed"`
	Link        string             `json:"link" bson:"link"`
	LinkContent string             `json:"linkContent" bson:"linkContent"`
	CreatedAt   time.Time          `json:"createdAt,omitempty" bson:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt,omitempty" bson:"updatedAt"`
}

func InsertLink(link Link) (Link, error) {
	linkCollection := database.DB.Collection("links")

	insertedLink, err := linkCollection.InsertOne(context.Background(), link)

	if err != nil {
		return Link{}, err
	}

	link.ID = insertedLink.InsertedID.(primitive.ObjectID)

	return link, nil
}

func FindLinkByShortLink(id primitive.ObjectID) (Link, error) {
	linkCollection := database.DB.Collection("links")

	var link Link
	filter := bson.M{"_id": id}
	err := linkCollection.FindOne(context.Background(), filter).Decode(&link)

	if err != nil {
		return Link{}, err
	}

	return link, nil
}

func UpdateLink(id primitive.ObjectID, link Link) (Link, error) {
	linkCollection := database.DB.Collection("links")
	filter := bson.M{"_id": id}
	payload := bson.M{"$set": bson.M{"isUsed": link.IsUsed}}

	_, err := linkCollection.UpdateOne(context.Background(), filter, payload)

	if err != nil {
		return Link{}, err
	}

	return link, nil
}
