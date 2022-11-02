package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/AshiishKarhade/GO-Projects/mongo-go/controllers"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	r := httprouter.New()
	uc := controllers.NewUserController(getSession())

	r.GET("/user/id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/id", uc.DeleteUser)

	http.ListenAndServe("localhost:9000", r)
}

func getSession() *mongo.Client {
	//uri := fmt.Sprintf("mongodb://%s:%s/%s", "127.0.0.1", "27017", "go-mongo")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		fmt.Println("Mongo Server Error")
		panic(err)
	}
	return client
}
