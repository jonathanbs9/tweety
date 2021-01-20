package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User struct
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FirstName string             `bson:"firstName" json:"first_name,omitempty"`
	LastName  string             `bson:"lastName," json:"last_name,omitempty"`
	DateBirth time.Time          `bson:"dateBirth," json:"date_birth,omitempty"`
	Email     string             `bson:"email," json:"email"`
	Password  string             `bson:"password," json:"password,omitempty"`
	Avatar    string             `bson:"avatar," json:"avatar,omitempty"`
	Banner    string             `bson:"banner," json:"banner,omitempty"`
	Biography string             `bson:"biography," json:"biography,omitempty"`
	Location  string             `bson:"location," json:"location,omitempty"`
	Website   string             `bson:"website," json:"website,omitempty"`
}
