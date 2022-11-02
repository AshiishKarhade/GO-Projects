package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AshiishKarhade/GO-Projects/mongo-go/models"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	client *mongo.Client
}

func NewUserController(s *mongo.Client) *UserController {
	return &UserController{s}
}

func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	u := models.User{}

	if err := uc.client.Database("go-mongo").Collection("users").FindOne(context.TODO(), bson.D{{"_id", oid}}).Decode(&u); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	// Getting JSON Response from POST Request and giving it to struct variable
	json.NewDecoder(r.Body).Decode(&u)

	u.Id = bson.NewObjectId()

	if _, err := uc.client.Database("go-mongo").Collection("users").InsertOne(context.TODO(), bson.D{{"_id", u.Id}, {"name", u.Name}, {"gender", u.Gender}, {"age", u.Age}}); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Sending uploaded data back to request
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	if _, err := uc.client.Database("go-mongo").Collection("users").DeleteOne(context.TODO(), bson.D{{"_id", oid}}); err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted User %v\n", oid)
}
