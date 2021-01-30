package bd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jonathanbs9/tweety/models"
)

// DeleteRelation => Borro la relación con el usuario que estoy siguiendo
func DeleteRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("tweety")
	col := db.Collection("relation")

	resp, err := col.DeleteOne(ctx, t)
	// Si la instrucción no se realizó ( == 0)
	if resp.DeletedCount == 0 {
		fmt.Println("Base de datos no trajo resultado para borrar relación => ", resp)
		return false, nil
	}
	// Si hubo error..
	if err != nil {
		log.Fatal("Error BD => " + err.Error())
		return false, err
	}
	return true, nil
}
