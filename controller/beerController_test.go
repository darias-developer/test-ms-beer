package controller

import (
	"errors"
	"net/http"
	"testing"

	"github.com/darias-developer/test-ms-beer/data"
	"github.com/darias-developer/test-ms-beer/middleware"
	"github.com/darias-developer/test-ms-beer/model"
)

func mock_createBeer_success(beerModel model.BeerModel) (string, error) {
	return "11111111", nil
}

func mock_createBeer_error(beerModel model.BeerModel) (string, error) {
	return "", errors.New("el registro no se ha podido crear")
}

func muck_findBeerById_success(id int) (model.BeerModel, error) {
	var beerModel model.BeerModel
	return beerModel, nil
}

func muck_findBeerById_error(id int) (model.BeerModel, error) {
	var beerModel model.BeerModel
	return beerModel, errors.New("cerveza no encontrada")
}

func muck_list_success() (data.ListResponse, error) {
	var listResponse data.ListResponse

	m := make(map[string]string)
	m["USD"] = "USD"

	listResponse.Currencies = m

	return listResponse, nil
}

func muck_list_error() (data.ListResponse, error) {
	var listResponse data.ListResponse
	return listResponse, errors.New("error al llamar servicio list")
}

func beerModel() model.BeerModel {

	var beerModel model.BeerModel
	beerModel.Id = 1
	beerModel.Name = "test"
	beerModel.Price = 1000.0
	beerModel.Currency = "USD"
	beerModel.Brewery = "test"
	beerModel.Country = "Chile"

	return beerModel
}
func TestAdd(t *testing.T) {

	middleware.LoggerInit()

	t.Run("test Add valida el parametro id", func(t *testing.T) {

		beerModel := beerModel()
		beerModel.Id = 0

		status, desc := BeerAdd(mock_createBeer_success, muck_findBeerById_error, muck_list_success, beerModel)

		t.Log(status)
		t.Log(desc)

		if status != http.StatusBadRequest {
			t.Errorf("Expected: %v, got: %v", status, http.StatusBadRequest)
		}
	})

	t.Run("test Add valida el parametro name", func(t *testing.T) {

		beerModel := beerModel()
		beerModel.Name = ""

		status, desc := BeerAdd(mock_createBeer_success, muck_findBeerById_error, muck_list_success, beerModel)

		t.Log(status)
		t.Log(desc)

		if status != http.StatusBadRequest {
			t.Errorf("Expected: %v, got: %v", status, http.StatusBadRequest)
		}
	})

	t.Run("test Add valida el parametro brewery", func(t *testing.T) {

		beerModel := beerModel()
		beerModel.Brewery = ""

		status, desc := BeerAdd(mock_createBeer_success, muck_findBeerById_error, muck_list_success, beerModel)

		t.Log(status)
		t.Log(desc)

		if status != http.StatusBadRequest {
			t.Errorf("Expected: %v, got: %v", status, http.StatusBadRequest)
		}
	})

	t.Run("test Add valida el parametro country", func(t *testing.T) {

		beerModel := beerModel()
		beerModel.Country = ""

		status, desc := BeerAdd(mock_createBeer_success, muck_findBeerById_error, muck_list_success, beerModel)

		t.Log(status)
		t.Log(desc)

		if status != http.StatusBadRequest {
			t.Errorf("Expected: %v, got: %v", status, http.StatusBadRequest)
		}
	})

	t.Run("test Add valida el parametro country", func(t *testing.T) {

		beerModel := beerModel()
		beerModel.Country = ""

		status, desc := BeerAdd(mock_createBeer_success, muck_findBeerById_error, muck_list_success, beerModel)

		t.Log(status)
		t.Log(desc)

		if status != http.StatusBadRequest {
			t.Errorf("Expected: %v, got: %v", status, http.StatusBadRequest)
		}
	})

	t.Run("test Add valida el parametro price", func(t *testing.T) {

		beerModel := beerModel()
		beerModel.Price = 0

		status, desc := BeerAdd(mock_createBeer_success, muck_findBeerById_error, muck_list_success, beerModel)

		t.Log(status)
		t.Log(desc)

		if status != http.StatusBadRequest {
			t.Errorf("Expected: %v, got: %v", status, http.StatusBadRequest)
		}
	})

	t.Run("test Add valida el parametro currency", func(t *testing.T) {

		beerModel := beerModel()
		beerModel.Currency = ""

		status, desc := BeerAdd(mock_createBeer_success, muck_findBeerById_error, muck_list_success, beerModel)

		t.Log(status)
		t.Log(desc)

		if status != http.StatusBadRequest {
			t.Errorf("Expected: %v, got: %v", status, http.StatusBadRequest)
		}
	})

	t.Run("test Add valida el parametro currency", func(t *testing.T) {

		beerModel := beerModel()
		beerModel.Currency = ""

		status, desc := BeerAdd(mock_createBeer_success, muck_findBeerById_error, muck_list_success, beerModel)

		t.Log(status)
		t.Log(desc)

		if status != http.StatusBadRequest {
			t.Errorf("Expected: %v, got: %v", status, http.StatusBadRequest)
		}
	})

	t.Run("test Add genera usuario de forma exitosa", func(t *testing.T) {

		beerModel := beerModel()

		status, desc := BeerAdd(mock_createBeer_success, muck_findBeerById_error, muck_list_success, beerModel)

		t.Log(status)
		t.Log(desc)

		if status != http.StatusOK {
			t.Errorf("Expected: %v, got: %v", status, http.StatusOK)
		}
	})

	t.Run("test Add falla al crear un registro con un id ya creado", func(t *testing.T) {

		beerModel := beerModel()

		status, desc := BeerAdd(mock_createBeer_success, muck_findBeerById_success, muck_list_success, beerModel)

		t.Log(status)
		t.Log(desc)

		if status != http.StatusConflict {
			t.Errorf("Expected: %v, got: %v", status, http.StatusConflict)
		}
	})

	t.Run("test Add falla al crear un registro debido que el servicio list esta abajo", func(t *testing.T) {

		beerModel := beerModel()

		status, desc := BeerAdd(mock_createBeer_success, muck_findBeerById_error, muck_list_error, beerModel)

		t.Log(status)
		t.Log(desc)

		if status != http.StatusBadRequest {
			t.Errorf("Expected: %v, got: %v", status, http.StatusBadRequest)
		}
	})

	t.Run("test Add falla al crear un registro debido que el servicio list no trae data para el ingreso", func(t *testing.T) {

		beerModel := beerModel()
		beerModel.Currency = "CLP"

		status, desc := BeerAdd(mock_createBeer_success, muck_findBeerById_error, muck_list_success, beerModel)

		t.Log(status)
		t.Log(desc)

		if status != http.StatusBadRequest {
			t.Errorf("Expected: %v, got: %v", status, http.StatusBadRequest)
		}
	})

	t.Run("test Add falla al crear un registro desde db", func(t *testing.T) {

		beerModel := beerModel()

		status, desc := BeerAdd(mock_createBeer_error, muck_findBeerById_error, muck_list_success, beerModel)

		t.Log(status)
		t.Log(desc)

		if status != http.StatusInternalServerError {
			t.Errorf("Expected: %v, got: %v", status, http.StatusInternalServerError)
		}
	})
}
