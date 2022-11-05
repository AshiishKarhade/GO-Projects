package middleware

import (
	"database/sql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

type response struct {
	ID      int    `json:"id",omitempty`
	Message string `json:"message",omitempty`
}

func CreateConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	
	log.Println("Successfully connected to Postgres")
	return db
}

func GetStock(w http.ResponseWriter, r *http.Request) {

}

func GetAllStocks(w http.ResponseWriter, r *http.Request) {

}

func CreateStock(w http.ResponseWriter, r *http.Request) {

}

func UpdateStock(w http.ResponseWriter, r *http.Request) {

}

func DeleteStock(w http.ResponseWriter, r *http.Request) {

}
