package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type DataType struct {
	ID      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Content string             `json:"content"`
	Link    string             `json:"link" bson:"link,omitempty"`
}

type ErrorType struct {
	Error string `json:"error"`
}

type GType interface {
	string | ErrorType | DataType
}

type ResponseType[T any] struct {
	Success bool `json:"success"`
	Data    T    `json:"data"`
}

type RequestBody struct {
	Data string `json:"data"`
}
