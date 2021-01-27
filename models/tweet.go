package models

// Tweet struct => estructura modelo de Tweet
type Tweet struct {
	Message string `bson:"message" json:"message"`
}
