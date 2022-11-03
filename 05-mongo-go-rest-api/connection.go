package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getCollection() *mongo.Collection {
	client := getClient()

	collection := client.Database("go-mongo").Collection("users")

	return collection
}

func getClient() *mongo.Client {
	//uri := fmt.Sprintf("mongodb://%s:%s/%s", "127.0.0.1", "27017", "go-mongo")
	log.Println("Connecting to mongo client . . .")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		fmt.Println("Mongo Server Error")
		panic(err)
	} else {
		log.Println("Mongo Connection Established")
	}
	return client
}
