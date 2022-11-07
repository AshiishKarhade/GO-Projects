package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

const (
	PhotoApi = "https://api.pexels.com/v1"
	VideoApi = "https://api.pexels.com/videos"
)

type Client struct {
	Token          string
	hc             http.Client
	RemainingTimes int32
}

func NewClient(token string) *Client {
	c := http.Client{}
	return &Client{Token: token, hc: c}
}

type Photo struct {
	Id              int32       `json:"id"`
	Width           int32       `json:"width"`
	Height          int32       `json:"height"`
	Url             string      `json:"url"`
	Photographer    string      `json:"photographer"`
	PhotographerUrl string      `json:"photographer_url"`
	Src             PhotoSource `json:"src"`
}

type PhotoSource struct {
	Original  string `json:"original"`
	Large     string `json:"large"`
	Large2x   string `json:"large2x"`
	Medium    string `json:"medium"`
	Small     string `json:"small"`
	Portrait  string `json:"portrait"`
	Square    string `json:"square"`
	Landscape string `json:"landscape"`
	Tiny      string `json:"tiny"`
}

type Search struct {
	Page         int32   `json:"page"`
	PerPage      int32   `json:"per_page"`
	TotalResults int32   `json:"total_results"`
	NextPage     int32   `json:"next_page"`
	Photos       []Photo `json:"photos"`
}

func (c Client) SearchPhotos(query string, perpage int, page int) (*SearchResult, error) {

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
