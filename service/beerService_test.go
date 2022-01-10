package service

import (
	"context"
	"errors"
	"testing"

	"github.com/darias-developer/test-ms-beer/model"
	"github.com/darias-developer/test-ms-beer/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func mock_connectDB_success() (*mongo.Client, error) {
	var client *mongo.Client
	return client, nil
}

func mock_connectDB_error() (*mongo.Client, error) {
	return nil, errors.New("error connectDB")
}

func mock_insertOneBeer_success(client *mongo.Client, ctx context.Context, beer model.BeerModel) (string, error) {
	return "1", nil
}

func mock_insertOneBeer_error(client *mongo.Client, ctx context.Context, beer model.BeerModel) (string, error) {
	return "", errors.New("ha ocurrido un error al agregar un registro")
}

func mock_findOneBeer_success(client *mongo.Client, ctx context.Context, condition primitive.M) (model.BeerModel, error) {
	return beerModel(), nil
}

func mock_findOneBeer_error(client *mongo.Client, ctx context.Context, condition primitive.M) (model.BeerModel, error) {
	return beerModel(), errors.New("ha ocurrido un error al encontrar un registro")
}

func mock_beerFindAll_success(client *mongo.Client, ctx context.Context, condition primitive.M) ([]model.BeerModel, error) {

	list := []model.BeerModel{
		beerModel(),
	}
	return list, nil
}

func mock_beerFindAll_error(client *mongo.Client, ctx context.Context, condition primitive.M) ([]model.BeerModel, error) {
	return nil, errors.New("ha ocurrido un error al listar los registros")
}

func TestBeerAdd(t *testing.T) {

	util.LoggerInit()

	t.Run("test BeerAdd: ConnectDB error, insertOneBeer: error", func(t *testing.T) {

		oid, err := BeerAdd(beerModel(), mock_connectDB_error, mock_insertOneBeer_error)

		t.Log(oid)
		t.Log(err)

		if err.Error() != "error connectDB" {
			t.Errorf("Expected: %v, got: %v", "error connectDB", err.Error())
		}
	})

	t.Run("test BeerAdd: ConnectDB success, insertOneBeer: error", func(t *testing.T) {

		oid, err := BeerAdd(beerModel(), mock_connectDB_success, mock_insertOneBeer_error)

		t.Log(oid)
		t.Log(err)

		if err.Error() != "ha ocurrido un error al agregar un registro" {
			t.Errorf("Expected: %v, got: %v", "ha ocurrido un error al agregar un registro", err.Error())
		}
	})

	t.Run("test BeerAdd: ConnectDB success, insertOneBeer: success", func(t *testing.T) {

		oid, err := BeerAdd(beerModel(), mock_connectDB_success, mock_insertOneBeer_success)

		t.Log(oid)
		t.Log(err)

		if oid != "1" {
			t.Errorf("Expected: %v, got: %v", "1", oid)
		}
	})
}

func TestBeerFindById(t *testing.T) {

	util.LoggerInit()

	t.Run("test BeerFindById: ConnectDB error, findOneBeer: error", func(t *testing.T) {

		beer, err := BeerFindById(1, mock_connectDB_error, mock_findOneBeer_error)

		t.Log(beer)
		t.Log(err)

		if err.Error() != "error connectDB" {
			t.Errorf("Expected: %v, got: %v", "error connectDB", err.Error())
		}
	})

	t.Run("test BeerFindById: ConnectDB success, findOneBeer: error", func(t *testing.T) {

		beer, err := BeerFindById(1, mock_connectDB_success, mock_findOneBeer_error)

		t.Log(beer)
		t.Log(err)

		if err.Error() != "ha ocurrido un error al encontrar un registro" {
			t.Errorf("Expected: %v, got: %v", "ha ocurrido un error al encontrar un registro", err.Error())
		}
	})

	t.Run("test BeerFindById: ConnectDB success, findOneBeer: success", func(t *testing.T) {

		beer, err := BeerFindById(1, mock_connectDB_success, mock_findOneBeer_success)

		t.Log(beer)
		t.Log(err)

		if beer.Id != 1 {
			t.Errorf("Expected: %v, got: %v", 1, beer.Id)
		}
	})
}

func TestBeerFindAll(t *testing.T) {

	util.LoggerInit()

	t.Run("test BeerFindAll: ConnectDB error, beerFindAll: error", func(t *testing.T) {

		beer, err := BeerFindAll(mock_connectDB_error, mock_beerFindAll_error)

		t.Log(beer)
		t.Log(err)

		if err.Error() != "error connectDB" {
			t.Errorf("Expected: %v, got: %v", "error connectDB", err.Error())
		}
	})

	t.Run("test BeerFindAll: ConnectDB success, beerFindAll: error", func(t *testing.T) {

		beer, err := BeerFindAll(mock_connectDB_success, mock_beerFindAll_error)

		t.Log(beer)
		t.Log(err)

		if err.Error() != "ha ocurrido un error al listar los registros" {
			t.Errorf("Expected: %v, got: %v", "ha ocurrido un error al listar los registros", err.Error())
		}
	})

	t.Run("test BeerFindAll: ConnectDB success, beerFindAll: success", func(t *testing.T) {

		beers, err := BeerFindAll(mock_connectDB_success, mock_beerFindAll_success)

		t.Log(beers)
		t.Log(err)

		if len(beers) != 1 {
			t.Errorf("Expected: %v, got: %v", 1, len(beers))
		}
	})
}
