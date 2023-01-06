package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN guarda la función ConnectDB */
var MongoCN = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://root:root@gotwitt.u2s5oq1.mongodb.net/test?retryWrites=true&w=majority") //?retryWrites=true&w=majority

/*ConnectDB es la funcion que me permite conectar con la BD */
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Éxito al conectar con la BD")
	return client
}

/*CheckConnection es el Ping a la base de datos */
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
