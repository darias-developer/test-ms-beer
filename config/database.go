package config

import (
	"context"
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ConnectDBType func() (*mongo.Client, error)

/* ConnectDB es la funcion que realiza la conexion  al db */
func ConnectDB() (*mongo.Client, error) {

	var clientOptions = options.Client().ApplyURI(os.Getenv("DB_SOURCE"))

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return client, errors.New("ERROR_CONN: " + err.Error())
	}

	return client, nil
}

func PingDB(client mongo.Client) error {

	err := client.Ping(context.TODO(), nil)

	if err != nil {
		return err
	}

	return nil
}

func CheckConnection() error {

	client, err := ConnectDB()

	if err != nil {
		return err
	}

	err = PingDB(*client)

	if err != nil {
		return err
	}

	return nil
}
