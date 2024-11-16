package routes

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	MongoDB := "mongodb://localhost:27017/caloriesdb"
	client, err := mongo.NewClient( options.Client().ApplyURI(MongoDB))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	
	
	if err != nil {
		log.Fatal(err)
	}
	
	

	fmt.Println("Connected to mongodb successfully...")

	return client

}

var client *mongo.Client = DBinstance()

func openCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection  = client.Database("caloriesdb").Collection(collectionName)

	return collection
}
