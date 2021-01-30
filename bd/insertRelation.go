package bd

import (
	"context"
	"log"
	"time"

	"github.com/jonathanbs9/tweety/models"
)

// InsertRelation => inserto relacion a la base de datos
func InsertRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("tweety")
	col := db.Collection("relation")

	_, err := col.InsertOne(ctx, t)

	// si hubo un error
	if err != nil {
		log.Fatal("Error BD => " + err.Error())
		return false, err
	}

	// Inserto realacion
	return true, nil
}
