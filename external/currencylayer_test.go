package external

import (
	"os"
	"testing"
)

func TestLive(t *testing.T) {

	t.Run("test Live valida el parametro currency", func(t *testing.T) {

		os.Setenv("ACCESS_KEY", "test")
		defer os.Unsetenv("ACCESS_KEY")

		_, err := Live("")

		if err.Error() != "el currency es requerido" {
			t.Errorf("Expected: %s, got: %s", "el currency es requerido", err.Error())
		}
	})

	t.Run("test Live valida el parametro ACCESS_KEY", func(t *testing.T) {

		_, err := Live("USD")

		if err.Error() != "el ACCESS_KEY es requerido" {
			t.Errorf("Expected: %v, got: %v", "el ACCESS_KEY es requerido", err.Error())
		}
	})
}

func TestList(t *testing.T) {

	t.Run("test List valida el parametro currency", func(t *testing.T) {

		_, err := List()

		if err.Error() != "el ACCESS_KEY es requerido" {
			t.Errorf("Expected: %s, got: %s", "el ACCESS_KEY es requerido", err.Error())
		}
	})
}
