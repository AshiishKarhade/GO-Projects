package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AshiishKarhade/GO-Projects/mongo-go/models"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	collection *mongo.Collection
}

func NewUserController(c *mongo.Collection) *UserController {
	return &UserController{c}
}

func (uc *UserController) Home(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uj := "Hello, World"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	// if !bson.IsObjectIdHex(id) {
	// 	w.WriteHeader(http.StatusNotFound)
	// }
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal("Invalid Hex ID")
	}

	fmt.Println("GET USER : ID", oid)
	u := models.User{}

	filter := bson.D{{"_id", oid}}
	if err := uc.collection.FindOne(context.TODO(), filter).Decode(&u); err != nil {
		fmt.Println("No record found")
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

	// Creating new MONGO Object ID
	u.Id = primitive.NewObjectID()

	if _, err := uc.collection.InsertOne(context.TODO(), bson.D{{"_id", u.Id}, {"name", u.Name}, {"gender", u.Gender}, {"age", u.Age}}); err != nil {
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

	// if !bson.IsObjectIdHex(id) {
	// 	w.WriteHeader(http.StatusNotFound)
	// }

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal("Invalid Hex ID")
	}

	if _, err := uc.collection.DeleteOne(context.TODO(), bson.D{{"_id", oid}}); err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted User %v\n", id)
}
