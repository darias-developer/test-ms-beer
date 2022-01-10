package util

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

type TypeMakePing func(client *mongo.Client) error

func MakePing(client *mongo.Client) error {
	if client == nil {
		return errors.New("cliente no valido")
	}
	return client.Ping(context.TODO(), nil)
}
