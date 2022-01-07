package config

import (
	"context"
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* ConectDB() es la funcion que realiza la conexion  al db */
func ConectDB() (*mongo.Client, error) {

	var clientOptions = options.Client().ApplyURI(os.Getenv("DB_SOURCE"))

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return client, errors.New("ERROR_CONN: " + err.Error())
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return client, errors.New("ERROR_PING: " + err.Error())
	}

	return client, nil
}

/* CheckConnection() es la funcion que realiza un ping a la db */
func CheckConnection() error {

	conn, err := ConectDB()

	if err != nil {
		return err
	}

	err = conn.Ping(context.TODO(), nil)

	if err != nil {
		return err
	}

	return nil
}
