package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Product struct {
	ID    string  `bson:"id"`
	NAME  string  `bson:"name"`
	PRICE float64 `bson:"price"`
}

// var connectobj *mongo.Client
// var col *mongo.Collection
func main() {
	connectobj, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = connectobj.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	col := connectobj.Database("products").Collection("productscollection")
	inserting(ctx, col)
	updatingdoc(ctx, col)
	deletingdoc(ctx, col)
	defer disconnecting(connectobj)

}

//func mongocreation() {}

func inserting(ctx context.Context, col *mongo.Collection) {
	pro := Product{
		ID:    "1122",
		NAME:  "truck123",
		PRICE: 100.9,
	}
	insertdata, err := col.InsertOne(ctx, pro)
	if err != nil {
		log.Fatal(err)

	} else {
		fmt.Println("data has been inserted succesfully")
	}

	fmt.Println("the doc inserted at the object :", insertdata.InsertedID)

}

func updatingdoc(ctx context.Context, col *mongo.Collection) {

	filtering := bson.M{"id": "1122"}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "name", Value: "varaiationcar"},
			{Key: "price", Value: 2222.0},
		}},
	}

	_, err := col.UpdateOne(ctx, filtering, update)
	if err != nil {
		fmt.Println("error ", err)
	} else {
		fmt.Println("the updation has been completed")
	}

}

func deletingdoc(ctx context.Context, col *mongo.Collection) {
	deletedoc := bson.M{"id": "209"}
	_, err := col.DeleteOne(ctx, deletedoc)
	if err != nil {
		log.Fatal("error", err)
	} else {
		fmt.Println("The doc has been deleted succesfully ")
	}

}
func disconnecting(connectobj *mongo.Client) {
	err := connectobj.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection closed.")
	}

}
