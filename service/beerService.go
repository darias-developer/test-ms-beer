package service

import (
	"context"
	"time"

	"github.com/darias-developer/test-ms-beer/config"
	"github.com/darias-developer/test-ms-beer/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* BeerAddService crea usuario en la db */
func BeerAddService(beerModel model.BeerModel) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	conn, err := config.ConectDB()

	if err != nil {
		return "", err
	}

	db := conn.Database("beer-test")
	collection := db.Collection("beer")

	result, err := collection.InsertOne(ctx, beerModel)

	if err != nil {
		return "", err
	}

	oid, _ := result.InsertedID.(primitive.ObjectID)

	return oid.String(), nil
}

/* BeerFindByIdService busca una cerveza en la db por medio del id */
func BeerFindByIdService(id int) (model.BeerModel, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	var beerModel model.BeerModel

	conn, err := config.ConectDB()

	if err != nil {
		return beerModel, err
	}

	db := conn.Database("beer-test")
	collection := db.Collection("beer")

	condition := bson.M{"id": id}

	err = collection.FindOne(ctx, condition).Decode(&beerModel)

	if err != nil {
		return beerModel, err
	}

	return beerModel, nil
}

/* BeerFindAllService obtiene todas las cervezas registradas */
func BeerFindAllService() ([]model.BeerModel, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	var results []model.BeerModel

	conn, err := config.ConectDB()

	if err != nil {
		return results, err
	}

	db := conn.Database("beer-test")
	collection := db.Collection("beer")

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
