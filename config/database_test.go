package config

import (
	"errors"
	"strings"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

func mock_connectDB_success() (*mongo.Client, error) {
	var client *mongo.Client
	return client, nil
}

func mock_connectDB_error() (*mongo.Client, error) {
	return nil, errors.New("error connectDB")
}

func TestConectDB(t *testing.T) {

	t.Run("test ConectDB valida la conexion a la db", func(t *testing.T) {

		_, err := ConnectDB()

		if !strings.Contains(err.Error(), "error parsing uri") {
			t.Errorf("Expected: %v, got: %v", "error parsing uri...", err.Error())
		}
	})
}

func TestCheckConnection(t *testing.T) {

	t.Run("test CheckConnection: conn error", func(t *testing.T) {

		err := CheckConn(mock_connectDB_error)

		if !strings.Contains(err.Error(), "error connectDB") {
			t.Errorf("Expected: %v, got: %v", "the Command operation...", err.Error())
		}
	})

	t.Run("test CheckConnection: cliente no valido", func(t *testing.T) {

		err := CheckConn(mock_connectDB_success)

		if !strings.Contains(err.Error(), "cliente no valido") {
			t.Errorf("Expected: %v, got: %v", "cliente no valido", err.Error())
		}
	})
}
