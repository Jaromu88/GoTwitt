package bd

import (
	"context"
	"time"

	"github.com/Jaromu88/GoTwitt/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoRegistroBD es la función final que registrará al usuario en la BD */
func InsertoRegistroBD(u models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	//Esto me devuelve el ID del usuario insertado
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
