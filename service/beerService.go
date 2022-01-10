package service

import (
	"context"
	"time"

	"github.com/darias-developer/test-ms-beer/config"
	"github.com/darias-developer/test-ms-beer/model"
	"github.com/darias-developer/test-ms-beer/util"
	"go.mongodb.org/mongo-driver/bson"
)

type BeerAddType func(beerModel model.BeerModel, typeConnectDB config.TypeConnectDB, insertOneBeer util.InsertOneBeerType) (string, error)
type BeerFindByIdType func(id int, typeConnectDB config.TypeConnectDB, findOneBeerType util.FindOneBeerType) (model.BeerModel, error)
type BeerFindAllType func(typeConnectDB config.TypeConnectDB, findAllBeerType util.FindAllBeerType) ([]model.BeerModel, error)

/* BeerAdd crea una cerveza en la db */
func BeerAdd(beerModel model.BeerModel, typeConnectDB config.TypeConnectDB, insertOneBeer util.InsertOneBeerType) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	conn, err := typeConnectDB()

	if err != nil {
		return "", err
	}

	oid, err := insertOneBeer(conn, ctx, beerModel)

	if err != nil {
		return "", err
	}

	return oid, nil
}

/* BeerFindById busca una cerveza en la db por medio del id */
func BeerFindById(id int, typeConnectDB config.TypeConnectDB, findOneBeerType util.FindOneBeerType) (model.BeerModel, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	var beer model.BeerModel

	conn, err := typeConnectDB()

	if err != nil {
		return beer, err
	}

	condition := bson.M{"id": id}

	beer, err = findOneBeerType(conn, ctx, condition)

	if err != nil {
		return beer, err
	}

	return beer, nil
}

/* BeerFindAll obtiene todas las cervezas registradas */
func BeerFindAll(typeConnectDB config.TypeConnectDB, findAllBeerType util.FindAllBeerType) ([]model.BeerModel, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	var results []model.BeerModel

	conn, err := typeConnectDB()

	if err != nil {
		return results, err
	}

	condition := bson.M{}

	results, err = findAllBeerType(conn, ctx, condition)

	if err != nil {
		return results, err
	}

	return results, nil
}
