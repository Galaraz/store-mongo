package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//Product describes an electronic produt e.g phone
type Product struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Name        string             `json:"Product_name" bson:"produt_name"`
	price       int                `json:"price" bson:"price"`
	Currency    string             `json:"currency" bson:"currency"`
	Quantity    string             `json:"quantity" bson:"quantity"`
	Discount    int                `json:"discount,omitempy" bson:"discount,omitempy"`
	Vendor      string             `json:"vendor" bson:"vendor"`
	Accessories []string           `json:"accessories, omitempty" bson:"accessories, omitempty"`
	SkuID       string             `json:"sku_id" bson:"sku_id"`
}

var iphone = Product{
	ID:          primitive.ObjectID.NewObjectID(),
	Name:        "iphone10",
	price:       900,
	Currency:    "",
	Quantity:    "USD",
	Discount:    "40",
	Vendor:      "apple",
	Accessories: []string{"charger", "headset", "slotopener"},
	SkuID:       "123",
}

var trimmer = Product{
	ID:          primitive.ObjectID.NewObjectID(),
	Name:        "easy trimmer",
	price:       120,
	Currency:    "",
	Quantity:    "USD",
	Discount:    "300",
	Vendor:      "philips",
	Accessories: []string{"charger", "comb", "bladeset", "cleaning oil"},
	SkuID:       "2345",
}

func main() {

	client, err := mongo.NewClient(options.client().applyUri("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println(err)
	}
	ctx, cancel := context.withTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}

	db := client.Database("tronics")
	collection := db.Collection("products")
	res,err := collection.Insert(context.Background(), iphone10)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.InsertdID,(primitive,ObjectID),Timestamp())
}
