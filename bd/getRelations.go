package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/jonathanbs9/tweety/models"
	"go.mongodb.org/mongo-driver/bson"
)

// GetRelations => Consulto relación entre ususarios
func GetRelations(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("tweety")
	col := db.Collection("relation")

	// Creo condición de busqueda
	condition := bson.M{
		"userid":         t.UserID,
		"userRelationId": t.UserRelationID,
	}

	var result models.Relation
	fmt.Println(result)
	err := col.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		fmt.Println("Error al buscar relación | " + err.Error())
		return false, err
	}

	return true, nil
}
