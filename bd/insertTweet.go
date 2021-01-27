package bd

import (
	"context"
	"time"

	"github.com/jonathanbs9/tweety/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertTweet => Inserta tweet en la base de datos
func InsertTweet(t models.SaveTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("tweety")
	col := db.Collection("tweet")

	record := bson.M{
		"userID":  t.UserID,
		"message": t.Message,
		"date":    t.Date,
	}

	result, err := col.InsertOne(ctx, record)
	if err != nil {
		return "", false, err
	}

	// obtengo el ulitmo campo insertado y lo convierte a objectID
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil

}
