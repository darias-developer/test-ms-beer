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

func mock_makePing_success(client *mongo.Client) error {
	return nil
}

func mock_makePing_error(client *mongo.Client) error {
	return errors.New("error pingDB")
}

func TestConectDB(t *testing.T) {

	t.Run("test ConectDB valida la conexion a la db", func(t *testing.T) {

		_, err := ConnectDB()

		if !strings.Contains(err.Error(), "error parsing uri") {
			t.Errorf("Expected: %v, got: %v", "error parsing uri...", err.Error())
		}
	})
}

func TestCheckConn(t *testing.T) {

	t.Run("test CheckConn: conn error, ping error", func(t *testing.T) {

		err := CheckConn(mock_connectDB_error, mock_makePing_error)

		if !strings.Contains(err.Error(), "error connectDB") {
			t.Errorf("Expected: %v, got: %v", "the Command operation...", err.Error())
		}
	})

	t.Run("test CheckConn: conn success, ping error", func(t *testing.T) {

		err := CheckConn(mock_connectDB_success, mock_makePing_error)

		if !strings.Contains(err.Error(), "error pingDB") {
			t.Errorf("Expected: %v, got: %v", "error pingDB", err.Error())
		}
	})

	t.Run("test CheckConn: conn success, ping success", func(t *testing.T) {

		err := CheckConn(mock_connectDB_success, mock_makePing_success)

		if err != nil {
			t.Errorf("Expected: %v, got: %v", nil, err)
		}
	})
}
