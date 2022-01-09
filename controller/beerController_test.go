package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/darias-developer/test-ms-beer/data"
	"github.com/darias-developer/test-ms-beer/model"
	"github.com/darias-developer/test-ms-beer/util"
)

func beerModel() model.BeerModel {

	var beerModel model.BeerModel
	beerModel.Id = 1
	beerModel.Name = "test"
	beerModel.Price = 1000
	beerModel.Currency = "CLP"
	beerModel.Brewery = "test"
	beerModel.Country = "Chile"

	return beerModel
}

func boxpriceRequestData() data.BoxpriceRequest {

	var boxpriceRequest data.BoxpriceRequest
	boxpriceRequest.Currency = "CLP"
	boxpriceRequest.Quantity = 6

	return boxpriceRequest
}

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

func muck_list_success(get util.TypeGet) (data.ListResponse, error) {
	var listResponse data.ListResponse

	m := make(map[string]string)
	m["USD"] = "USD"
	m["EUR"] = "EUR"
	m["CLP"] = "CLP"

	listResponse.Currencies = m

	return listResponse, nil
}

func mock_list_success(get util.TypeGet) (int, []byte) {

	var listResponse data.ListResponse

	m := make(map[string]string)
	m["USD"] = "USD"
	m["EUR"] = "EUR"
	m["CLP"] = "CLP"

	listResponse.Currencies = m
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(listResponse)

	return 200, reqBodyBytes.Bytes()
}

func muck_list_error(get util.TypeGet) (data.ListResponse, error) {
	var listResponse data.ListResponse
	return listResponse, errors.New("error al llamar servicio list")
}

func muck_live_success(currencies string, get util.TypeGet) (data.LiveResponse, error) {
	var liveResponse data.LiveResponse

	m := make(map[string]float32)

	m["USDEUR"] = 0.884795
	m["USDUSD"] = 1
	m["USDCLP"] = 837.598714

	liveResponse.Quotes = m

	return liveResponse, nil
}

func muck_live_error(currencies string, get util.TypeGet) (data.LiveResponse, error) {
	var liveResponse data.LiveResponse
	return liveResponse, errors.New("error al llamar servicio live")
}

func muck_beerFindAll_success() ([]model.BeerModel, error) {

	list := []model.BeerModel{
		beerModel(),
	}
	return list, nil
}

func muck_beerFindAll_error() ([]model.BeerModel, error) {
	var list []model.BeerModel
	return list, errors.New("error_en_db")
}

func muck_beerFindById_success(id int) (model.BeerModel, error) {
	var beerModel model.BeerModel
	return beerModel, nil
}

func muck_beerFindById_error(id int) (model.BeerModel, error) {
	var beerModel model.BeerModel
	return beerModel, errors.New("error_en_db")
}

func TestBeerAdd(t *testing.T) {

	util.LoggerInit()

	t.Run("test BeerAdd valida el parametro id", func(t *testing.T) {

		beerModel := beerModel()
		beerModel.Id = 0

		status, desc := BeerAdd(mock_createBeer_success, muck_findBeerById_error, muck_list_success, beerModel)

		t.Log(status)
		t.Log(desc)

		if status != http.StatusBadRequest {
			t.Errorf("Expected: %v, got: %v", status, http.StatusBadRequest)
		}
	})

	t.Run("test BeerAdd valida el parametro name", func(t *testing.T) {

		beerModel := beerModel()
		beerModel.Name = ""

		status, desc := BeerAdd(mock_createBeer_success, muck_findBeerById_error, muck_list_success, beerModel)

		t.Log(status)
		t.Log(desc)

		if status != http.StatusBadRequest {
			t.Errorf("Expected: %v, got: %v", status, http.StatusBadRequest)
		}
	})

	t.Run("test BeerAdd valida el parametro brewery", func(t *testing.T) {

		beerModel := beerModel()
		beerModel.Brewery = ""

		status, desc := BeerAdd(mock_createBeer_success, muck_findBeerById_error, muck_list_success, beerModel)

		t.Log(status)
		t.Log(desc)

		if status != http.StatusBadRequest {
			t.Errorf("Expected: %v, got: %v", status, http.StatusBadRequest)
		}
	})

	t.Run("test BeerAdd valida el parametro country", func(t *testing.T) {

		beerModel := beerModel()
		beerModel.Country = ""

		status, desc := BeerAdd(mock_createBeer_success, muck_findBeerById_error, muck_list_success, beerModel)

		t.Log(status)
		t.Log(desc)

		if status != http.StatusBadRequest {
			t.Errorf("Expected: %v, got: %v", status, http.StatusBadRequest)
		}
	})

	t.Run("test BeerAdd valida el parametro country", func(t *testing.T) {

		beerModel := beerModel()
		beerModel.Country = ""

		status, desc := BeerAdd(mock_createBeer_success, muck_findBeerById_error, muck_list_success, beerModel)

		t.Log(status)
		t.Log(desc)

		if status != http.StatusBadRequest {
			t.Errorf("Expected: %v, got: %v", status, http.StatusBadRequest)
		}
	})

	t.Run("test BeerAdd valida el parametro price", func(t *testing.T) {

		beerModel := beerModel()
		beerModel.Price = 0

		status, desc := BeerAdd(mock_createBeer_success, muck_findBeerById_error, muck_list_success, beerModel)

		t.Log(status)
		t.Log(desc)

		if status != http.StatusBadRequest {
			t.Errorf("Expected: %v, got: %v", status, http.StatusBadRequest)
		}
	})

	t.Run("test BeerAdd valida el parametro currency", func(t *testing.T) {

		beerModel := beerModel()
		beerModel.Currency = ""

		status, desc := BeerAdd(mock_createBeer_success, muck_findBeerById_error, muck_list_success, beerModel)

		t.Log(status)
		t.Log(desc)

		if status != http.StatusBadRequest {
			t.Errorf("Expected: %v, got: %v", status, http.StatusBadRequest)
		}
	})

	t.Run("test BeerAdd valida el parametro currency", func(t *testing.T) {

		beerModel := beerModel()
		beerModel.Currency = ""

		status, desc := BeerAdd(mock_createBeer_success, muck_findBeerById_error, muck_list_success, beerModel)

		t.Log(status)
		t.Log(desc)

		if status != http.StatusBadRequest {
			t.Errorf("Expected: %v, got: %v", status, http.StatusBadRequest)
		}
	})

	t.Run("test BeerAdd genera usuario de forma exitosa", func(t *testing.T) {

		beerModel := beerModel()

		status, desc := BeerAdd(mock_createBeer_success, muck_findBeerById_error, muck_list_success, beerModel)

		t.Log(status)
		t.Log(desc)

		if status != http.StatusOK {
			t.Errorf("Expected: %v, got: %v", status, http.StatusOK)
		}
	})

	t.Run("test BeerAdd falla al crear un registro con un id ya creado", func(t *testing.T) {

		beerModel := beerModel()

		status, desc := BeerAdd(mock_createBeer_success, muck_findBeerById_success, muck_list_success, beerModel)

		t.Log(status)
		t.Log(desc)

		if status != http.StatusConflict {
			t.Errorf("Expected: %v, got: %v", status, http.StatusConflict)
		}
	})

	t.Run("test BeerAdd falla al crear un registro debido que el servicio list esta abajo", func(t *testing.T) {

		beerModel := beerModel()

		status, desc := BeerAdd(mock_createBeer_success, muck_findBeerById_error, muck_list_error, beerModel)

		t.Log(status)
		t.Log(desc)

		if status != http.StatusBadRequest {
			t.Errorf("Expected: %v, got: %v", status, http.StatusBadRequest)
		}
	})

	t.Run("test BeerAdd falla al crear un registro debido que el servicio list no trae data para el ingreso", func(t *testing.T) {

		beerModel := beerModel()
		beerModel.Currency = "COM"

		status, desc := BeerAdd(mock_createBeer_success, muck_findBeerById_error, muck_list_success, beerModel)

		t.Log(status)
		t.Log(desc)

		if status != http.StatusBadRequest {
			t.Errorf("Expected: %v, got: %v", status, http.StatusBadRequest)
		}
	})

	t.Run("test BeerAdd falla al crear un registro desde db", func(t *testing.T) {

		beerModel := beerModel()

		status, desc := BeerAdd(mock_createBeer_error, muck_findBeerById_error, muck_list_success, beerModel)

		t.Log(status)
		t.Log(desc)

		if status != http.StatusInternalServerError {
			t.Errorf("Expected: %v, got: %v", status, http.StatusInternalServerError)
		}
	})
}

func TestBeerFindAll(t *testing.T) {

	util.LoggerInit()

	t.Run("test BeerFindAll success without data", func(t *testing.T) {

		status, desc, arr := BeerFindAll(muck_beerFindAll_error)

		t.Log(status)
		t.Log(desc)

		if len(arr) != 0 {
			t.Errorf("Expected: %v, got: %v", len(arr), 0)
		}
	})

	t.Run("test BeerFindAll success with data", func(t *testing.T) {

		status, desc, arr := BeerFindAll(muck_beerFindAll_success)

		t.Log(status)
		t.Log(desc)

		if len(arr) == 0 {
			t.Errorf("Expected: %v, got: %v", len(arr), 0)
		}
	})
}

func TestBeerFindById(t *testing.T) {

	util.LoggerInit()

	t.Run("test BeerFindById parametro id no valido", func(t *testing.T) {

		status, desc, _ := BeerFindById(muck_beerFindById_error, 0)

		t.Log(status)
		t.Log(desc)

		if status != http.StatusNotFound {
			t.Errorf("Expected: %v, got: %v", http.StatusNotFound, status)
		}
	})

	t.Run("test BeerFindAll success", func(t *testing.T) {

		status, desc, _ := BeerFindById(muck_beerFindById_success, 1)

		t.Log(status)
		t.Log(desc)

		if status != http.StatusOK {
			t.Errorf("Expected: %v, got: %v", http.StatusOK, status)
		}
	})
}

func TestBeerBoxPriceById(t *testing.T) {

	util.LoggerInit()

	t.Run("test BeerBoxPriceById parametro id no valido", func(t *testing.T) {

		data := boxpriceRequestData()

		status, desc, _ := BeerBoxPriceById(
			muck_beerFindById_error, muck_list_success, muck_live_success, data, 0)

		t.Log(status)
		t.Log(desc)

		if status != http.StatusNotFound {
			t.Errorf("Expected: %v, got: %v", http.StatusNotFound, status)
		}
	})

	t.Run("test BeerBoxPriceById error en servicio list", func(t *testing.T) {

		data := boxpriceRequestData()

		status, desc, _ := BeerBoxPriceById(
			muck_beerFindById_success, muck_list_error, muck_live_success, data, 1)

		t.Log(status)
		t.Log(desc)

		if status != http.StatusInternalServerError {
			t.Errorf("Expected: %v, got: %v", http.StatusInternalServerError, status)
		}
	})

	t.Run("test BeerBoxPriceById error curreny no es valida", func(t *testing.T) {

		data := boxpriceRequestData()
		data.Currency = "COM"

		status, desc, _ := BeerBoxPriceById(
			muck_beerFindById_success, muck_list_success, muck_live_success, data, 1)

		t.Log(status)
		t.Log(desc)

		if status != http.StatusNotFound {
			t.Errorf("Expected: %v, got: %v", http.StatusNotFound, status)
		}
	})

	t.Run("test BeerBoxPriceById error en servicio live", func(t *testing.T) {

		data := boxpriceRequestData()
		data.Currency = "EUR"

		status, desc, _ := BeerBoxPriceById(
			muck_beerFindById_success, muck_list_success, muck_live_error, data, 1)

		t.Log(status)
		t.Log(desc)

		if status != http.StatusInternalServerError {
			t.Errorf("Expected: %v, got: %v", http.StatusInternalServerError, status)
		}
	})

	t.Run("test BeerBoxPriceById success", func(t *testing.T) {

		data := boxpriceRequestData()
		data.Currency = "EUR"

		status, desc, _ := BeerBoxPriceById(
			muck_beerFindById_success, muck_list_success, muck_live_success, data, 1)

		t.Log(status)
		t.Log(desc)

		if status != http.StatusOK {
			t.Errorf("Expected: %v, got: %v", http.StatusOK, status)
		}
	})

	t.Run("test BeerBoxPriceById curreny por defecto", func(t *testing.T) {

		data := boxpriceRequestData()
		data.Currency = "EUR"

		status, desc, _ := BeerBoxPriceById(
			muck_beerFindById_success, muck_list_success, muck_live_success, data, 0)

		t.Log(status)
		t.Log(desc)

		if status != http.StatusOK {
			t.Errorf("Expected: %v, got: %v", http.StatusOK, status)
		}
	})
}
