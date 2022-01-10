package util

import (
	"context"

	"github.com/darias-developer/test-ms-beer/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FindOneBeerType func(client *mongo.Client, ctx context.Context, condition primitive.M) (model.BeerModel, error)
type InsertOneBeerType func(client *mongo.Client, ctx context.Context, beer model.BeerModel) (string, error)
type FindAllBeerType func(client *mongo.Client, ctx context.Context, condition primitive.M) ([]model.BeerModel, error)

func FindOneBeer(client *mongo.Client, ctx context.Context, condition primitive.M) (model.BeerModel, error) {
	var beer model.BeerModel
	collection := client.Database("beer-test").Collection("beer")
	err := collection.FindOne(ctx, condition).Decode(&beer)
	return beer, err
}

func InsertOneBeer(client *mongo.Client, ctx context.Context, beer model.BeerModel) (string, error) {
	collection := client.Database("beer-test").Collection("beer")
	result, err := collection.InsertOne(ctx, beer)

	if err != nil {
		return "", err
	}

	oid, _ := result.InsertedID.(primitive.ObjectID)

	return oid.String(), err
}

func FindAllBeer(client *mongo.Client, ctx context.Context, condition primitive.M) ([]model.BeerModel, error) {

	var results []model.BeerModel

	collection := client.Database("beer-test").Collection("beer")
	cur, err := collection.Find(ctx, condition)

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

	err = cur.Close(ctx)

	if err != nil {
		return results, err
	}

	return results, nil
}
