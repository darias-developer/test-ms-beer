package config

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* MongoCN es la funcion que realiza la conexion  al db */
//var MongoCN = ConectDB()

/* ConectDB() es la funcion que realiza la conexion  al db */
func ConectDB() *mongo.Client {

	var clientOptions = options.Client().ApplyURI(os.Getenv("DB_SOURCE"))

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal("ERROR: " + err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	return client
}

/* CheckConnection() es la funcion que realiza un ping a la db */
func CheckConnection() int {

	err := ConectDB().Ping(context.TODO(), nil)

	if err != nil {
		return 0
	}

	return 1
}
