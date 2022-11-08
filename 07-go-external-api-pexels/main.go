package main

import (
	"fmt"
	"github.com/AshiishKarhade/GO-Projects/go-external-api/models"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

const (
	PhotoApi = "https://api.pexels.com/v1"
	VideoApi = "https://api.pexels.com/videos"
)

func NewClient(token string) *models.Client {
	c := http.Client{}
	return &models.Client{Token: token, Hc: c}
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file -", err)
	}
	TOKEN := os.Getenv("API_KEY")

	client := NewClient(TOKEN)

	result, err := client.SearchPhotos("waves", 10, 1)
	if err != nil {
		log.Println("search error", err)
	}
	if result.Page == 0 {
		log.Println("Search resulted wrong")
	}

	fmt.Println(result)
}
