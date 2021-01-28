package bd

import (
	"context"
	"log"
	"time"

	"github.com/jonathanbs9/tweety/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ReadTweets func  => Lee tweets de un determinado user por ID. Utilizo paginado. Devuelvo un slice.
func ReadTweets(ID string, page int64) ([]*models.GetTweets, bool) {
	// Seteo el contexto
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("tweety")
	col := db.Collection("tweet")

	var results []*models.GetTweets
	// La condición es que el userID (mongoDB) sea igual al Id que le paso por parametro
	condition := bson.M{
		"userID": ID,
	}

	// Le seteo propiedades al options, que intervienen durante el find
	opts := options.Find()
	opts.SetLimit(20)
	// le marco que me ordene por fecha descendente (-1)
	opts.SetSort(bson.D{{Key: "date", Value: -1}})
	opts.SetSkip((page - 1) * 20)

	// Creo variable cursor (puntero), donde se graban los resultados y los puedo ir recorriendo
	cursor, err := col.Find(ctx, condition, opts)
	if err != nil {
		log.Fatal("Error al procesar los tweets => " + err.Error())
		return results, false
	}

	// Recorro el cursor. Creo un contexto vacío
	for cursor.Next(context.TODO()) {
		var record models.GetTweets
		err := cursor.Decode(&record)

		if err != nil {
			return results, false
		}
		results = append(results, &record)
	}
	return results, true

}
