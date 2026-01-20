package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Order struct {
	Id        primitive.ObjectID `bson:"_id, omitempty"`
	Name      string             `bson:"name"`
	Product   string             `bson:"product"`
	Quantity  int                `bson:"quantity"`
	CreatedAt string             `bson:"createdAt"`
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("Could not connect to mongo: ", err)
		return
	}
	defer client.Disconnect(ctx)
	db := client.Database("orders")
	ordersCollection := db.Collection("orders")

	newOrder := Order{
		Id:        primitive.NewObjectID(),
		Name:      "Vladimir",
		Product:   "Chicken Burger",
		Quantity:  1,
		CreatedAt: time.Now().Format("2006/01/02 15:04:05"),
	}
	_, err = ordersCollection.InsertOne(ctx, newOrder)
	if err != nil {
		fmt.Println("Error with Insert: ", err)
		return
	}

	cursor, err := ordersCollection.Find(ctx, bson.M{"name": "Vladimir"})
	if err != nil {
		fmt.Println("Find error: ", err)
		return
	}
	defer cursor.Close(ctx)

	var orders []Order
	err = cursor.All(ctx, &orders)
	if err != nil {
		fmt.Println("Cursor.All error: ", err)
	}
	for _, oneOrder := range orders {
		fmt.Println(oneOrder)
	}

}
