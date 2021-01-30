package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/jonathanbs9/tweety/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetAllRelation => Consulta y devuelve las relaciones que tiene un determinado usuario
func GetAllRelation(ID string, page int64, search string, typo string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("tweety")
	col := db.Collection("users")

	// Creo un slice de modelo de usuario
	var results []*models.User

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	// Le seteo un limite de 20 resultados
	findOptions.SetLimit(20)

	// La i refiere a que no se va a fijar si la expresion regular va con mayus o minus
	query := bson.M{
		"firstName": bson.M{"$regex": `(?i)` + search},
	}

	// Realizamos la búsqueda y nos devuelve un cursor.
	cursor, err := col.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println("Error al realizar la consulta | " + err.Error())
		return results, false
	}

	var found, include bool
	// Recorro el cursor
	for cursor.Next(ctx) {
		var s models.User
		err := cursor.Decode(&s)
		if err != nil {
			fmt.Println("Error en los resultados de users (BD) | " + err.Error())
			return results, false
		}

		var r models.Relation
		r.UserID = ID
		r.UserRelationID = s.ID.Hex()

		// Por cada iteracion tengo que saber si voy a incluir a ese usuario en una rta o no
		include = false

		found, err = GetRelations(r)
		if typo == "new" && found == false {
			include = true
		}
		if typo == "follow" && found == true {
			include = true
		}
		if r.UserRelationID == ID {
			include = false
		}

		// En caso que haya que incluirlos, se hace el append a la lista
		if include == true {
			s.Password = ""
			s.Biography = ""
			s.Website = ""
			s.Location = ""
			s.Banner = ""
			s.Email = ""

			results = append(results, &s)
		}
	}
	// Al final del cursor preguntamos si existe algun error.
	err = cursor.Err()
	if err != nil {
		fmt.Println("Error => " + err.Error())
		return results, false
	}
	cursor.Close(ctx)
	return results, true
}
