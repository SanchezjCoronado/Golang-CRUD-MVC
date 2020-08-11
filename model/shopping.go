package model

import "gopkg.in/mgo.v2/bson"

type Shopping struct {
	ID bson.ObjectId `bson:"_id" json:"id"`
	User int `bson:"user" json:"user"`
	Products []string `bson:"products" json:"products"`

}

type ShoppingID struct {
	ID string `json:"id"`
}