package bd

import (
	"context"
	"time"

	"github.com/jonathanbs9/tweety/models"
	"go.mongodb.org/mongo-driver/bson"
)

// GetFollowedTweets => Leo los tweets de la gente que estoy siguiendo
func GetFollowedTweets(ID string, page int) ([]models.ResponseFollowedTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("tweety")
	col := db.Collection("relation")

	// Primero nos basamos en la tabla relacion. De ahi sacamos  con quien estamos relacionados
	// Skip de registros
	skip := (page - 1) * 20
	conditions := make([]bson.M, 0)

	// Framework Aggregate. Comando match para unir relaciones
	conditions = append(conditions, bson.M{"$match": bson.M{"userid": ID}})
	// LookUp unir 2 tablas
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "userRelationId",
			"foreignField": "userID",
			"as":           "tweet",
		}})

	// $unwind => Para poder procesar los resultados.
	conditions = append(conditions, bson.M{"$unwind": "$tweet"})
	conditions = append(conditions, bson.M{"$sort": bson.M{"tweet.date": -1}})
	conditions = append(conditions, bson.M{"$skip": skip})
	conditions = append(conditions, bson.M{"$limit": 20})

	// Creo un cursor. Uso aggregate
	cursor, err := col.Aggregate(ctx, conditions)
	var result []models.ResponseFollowedTweets

	// Procesamos el cursor
	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}

	return result, true

}
