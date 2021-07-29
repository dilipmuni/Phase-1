package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id bson.ObjectId `json:"id" bson:"_id"`

	Name   string `json:"name" bson:"name"`
	Gender string `json:"gender" bson:"gender"`
	Age    int    `json:"age" bson:"age"`

	PhoneNo string `json:"phoneno" bson:"phoneno"`
	EmailId string `json:"emailid" bson:"emailid"`
	Address string `json:"address" bson:"address"`
}

type Product struct {
	Id bson.ObjectId `json:"id" bson:"_id"`

	ProductName  string  `json:"pname" bson:"pname"`
	ProductQty   int     `json:"pqty" bson:"pqty"`
	ProductPrice float64 `json:"pprice" bson:"pprice"`
}

type Cart struct {
	Id           bson.ObjectId `json:"id" bson:"_id"`
	CartProducts []CartProduct `json:"cartproducts" bson:"cartproducts"`
	Username     string        `json:"uname" bson:"uname"` //Add user Id here
	TotalPrice   float64       `json:"totalprice" bson:"totalprice"`
}

type CartProduct struct {
	ProductName  string  `json:"pname" bson:"pname"`
	ProductPrice float64 `json:"pprice" bson:"pprice"` //Add product ID instead of product name and price here
	ProductQty   int     `json:"pqty" bson:"pqty"`
}

type Order struct {
	Id             bson.ObjectId `json:"id" bson:"_id"`
	OrderProducts  []CartProduct `json:"orderproducts" bson:"orderproducts"`
	Username       string        `json:"uname" bson:"uname"`
	TotalPrice     float64
	OrderDate      time.Time
	DeliveryStatus string
}

type Payment struct {
	Id         bson.ObjectId `json:"id" bson:"_id"`
	Username   string        `json:"uname" bson:"uname"`
	TotalPrice float64
}
