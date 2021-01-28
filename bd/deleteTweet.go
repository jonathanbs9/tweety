package bd

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DeleteTweet => Recibe ID del tweet a borrar e ID del usuario que contiene ese tweet
func DeleteTweet(ID string, UserID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoC.Database("tweety")
	col := db.Collection("tweet")

	objID, _ := primitive.ObjectIDFromHex(ID)

	// mongo : parametros
	condition := bson.M{
		"_id":    objID,
		"userID": UserID,
	}

	// Realizo la instrucción
	//resp, err := col.DeleteOne(ctx, condition)
	resp, _ := col.DeleteOne(ctx, condition)
	if resp.DeletedCount == 0 {
		fmt.Println("Base de datos no trajo resultado para borrar => ", resp)
		return false, nil
	}

	return true, nil

}
