package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AshiishKarhade/GO-Projects/go-postgres-stocksapi/router"
)

func main() {
	PORT := "8081"
	r := router.Router()
	fmt.Println("Starting server on the port...", PORT)

	log.Fatal(http.ListenAndServe(":"+PORT, r))
}
