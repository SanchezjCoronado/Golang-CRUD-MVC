package model

import "gopkg.in/mgo.v2/bson"

type Shopping struct {
	ID bson.ObjectId `bson:"_id" json:"id"`
	User int `bson:"user" json:"user"`
	Products []string `bson:"products" json:"products"`
	Payment string `bson:"payment" json:"payment"`
	PriceTotal int `bson:"pricetotal" json:"price_total"`
}

type ShoppingID struct {
	ID string `json:"id"`
}