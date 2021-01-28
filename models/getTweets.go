package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetTweets => Estructura de devolución de un tweet.
type GetTweets struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID  string             `bson:"user_id" json:"user_id,omitempty"`
	Message string             `bson:"message" json:"message,omitempty"`
	Date    time.Time          `bson:"date" json:"date,omitempty"`
}
