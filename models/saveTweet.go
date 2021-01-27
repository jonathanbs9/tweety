package models

import "time"

// SaveTweet struct => Modelo de tweet
type SaveTweet struct {
	UserID  string    `bson:"userid" json:"user_id,omitempty"`
	Message string    `bson:"message" json:"message,omitempty"`
	Date    time.Time `bson:"date" json:"date,omitempty"`
}
