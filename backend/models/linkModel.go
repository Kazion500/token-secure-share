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
	IsViewed    bool               `json:"isViewed" bson:"isViewed"`
	ShortCode   string             `json:"shortCode" bson:"shortCode"`
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

func FindLinkByShortCode(shortCode string) (Link, error) {
	linkCollection := database.DB.Collection("links")

	var link Link
	filter := bson.M{"shortCode": shortCode}
	err := linkCollection.FindOne(context.Background(), filter).Decode(&link)

	if err != nil {
		return Link{}, err
	}

	return link, nil
}

func UpdateLink(id primitive.ObjectID, link Link) (Link, error) {
	linkCollection := database.DB.Collection("links")
	filter := bson.M{"_id": id}
	payload := bson.M{"$set": bson.M{"isViewed": link.IsViewed}}

	_, err := linkCollection.UpdateOne(context.Background(), filter, payload)

	if err != nil {
		return Link{}, err
	}

	return link, nil
}
