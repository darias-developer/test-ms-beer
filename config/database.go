package config

import (
	"context"
	"os"

	"github.com/darias-developer/test-ms-beer/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TypeConnectDB func() (*mongo.Client, error)

/* ConnectDB es la funcion que realiza la conexion  al db */
func ConnectDB() (*mongo.Client, error) {
	var clientOptions = options.Client().ApplyURI(os.Getenv("DB_SOURCE"))
	return mongo.Connect(context.TODO(), clientOptions)
}

/* CheckConnection es la funcion que realiza un check a la db */
func CheckConn(typeConnectDB TypeConnectDB, typeMakePing util.TypeMakePing) error {

	client, err := typeConnectDB()

	if err != nil {
		return err
	}

	err = typeMakePing(client)

	if err != nil {
		return err
	}

	return nil
}
