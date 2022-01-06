package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* estructura para el registro de cervezas */
type BeerModel struct {
	OID      primitive.ObjectID `bson:"_id,omitempty" json:"oid"`
	Id       int                `bson:"id" json:"id,omitempty"`
	Name     string             `bson:"name" json:"name,omitempty"`
	Brewery  string             `bson:"brewery" json:"brewery,omitempty"`
	Country  string             `bson:"country" json:"country,omitempty"`
	Price    float32            `bson:"price" json:"price,omitempty"`
	Currency string             `bson:"currency" json:"currency,omitempty"`
}
