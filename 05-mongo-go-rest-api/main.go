package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/AshiishKarhade/GO-Projects/mongo-go/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {

	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	r := httprouter.New()
	uc := controllers.NewUserController(getCollection())

	r.GET("/", uc.Home)
	r.GET("/user/id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/id", uc.DeleteUser)

	fmt.Println("Server Starting at PORT:9002")
	log.Fatal(http.ListenAndServe("localhost:9002", r))
}
