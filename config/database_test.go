package config

import (
	"strings"
	"testing"
)

func TestConectDB(t *testing.T) {

	t.Run("test ConectDB valida el parametro DB_SOURCE", func(t *testing.T) {

		_, err := ConnectDB()

		if !strings.Contains(err.Error(), "ERROR_CONN") {
			t.Errorf("Expected: %v, got: %v", "ERROR_CONN....", err.Error())
		}
	})
}

func TestCheckConnection(t *testing.T) {

	t.Run("test CheckConnection valida el parametro DB_SOURCE", func(t *testing.T) {

		err := CheckConnection()

		if !strings.Contains(err.Error(), "ERROR_CONN") {
			t.Errorf("Expected: %v, got: %v", "ERROR_CONN....", err.Error())
		}
	})
}
