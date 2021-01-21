package bd

import (
	"context"
	"time"

	"github.com/jonathanbs9/tweety/models"
	"go.mongodb.org/mongo-driver/bson"
)

// CheckUserExists func => consulta por email si ya existe en base de datos
func CheckUserExists(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("tweety")
	col := db.Collection("users")

	condition := bson.M{"email": email}

	var result models.User
	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}
