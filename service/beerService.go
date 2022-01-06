package service

import (
	"context"
	"time"

	"github.com/darias-developer/test-ms-beer/config"
	"github.com/darias-developer/test-ms-beer/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* CreateUser crea usuario en la db */
func CreateBeer(beerModel model.BeerModel) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := config.ConectDB().Database("beer-test")
	collection := db.Collection("beer")

	result, err := collection.InsertOne(ctx, beerModel)

	if err != nil {
		return "", err
	}

	oid, _ := result.InsertedID.(primitive.ObjectID)

	return oid.String(), nil
}

/* FindBeerById busca una cerveza en la db por medio del id */
func FindBeerById(id int) (model.BeerModel, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := config.ConectDB().Database("beer-test")
	collection := db.Collection("beer")

	condition := bson.M{"id": id}

	var beerModel model.BeerModel

	err := collection.FindOne(ctx, condition).Decode(&beerModel)

	if err != nil {
		return beerModel, err
	}

	return beerModel, nil
}

/* SearchBeers obtiene todas las cervezas registradas */
func SearchBeers() ([]*model.BeerModel, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := config.ConectDB().Database("beer-test")
	collection := db.Collection("beer")

	var results []*model.BeerModel

	cur, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return results, err
	}

	err = cur.All(ctx, &results)

	if err != nil {
		return results, err
	}

	err = cur.Err()

	if err != nil {
		return results, err
	}

	cur.Close(ctx)

	return results, nil
}
