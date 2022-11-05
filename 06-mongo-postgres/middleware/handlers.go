package middleware

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/AshiishKarhade/GO-Projects/go-postgres-stocksapi/models"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
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

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert ID to int : %v", err)
	}

	stock, err := getStock(int64(id))
	if err != nil {
		log.Fatalf("Unable to get stock : %v", err)
	}

	json.NewEncoder(w).Encode(stock)
}

func GetAllStocks(w http.ResponseWriter, r *http.Request) {
	stocks, err := getAll()
	if err != nil {
		log.Fatalf("Unable to get all stocks: %v", err)
	}
	json.NewEncoder(w).Encode(stocks)
}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock
	err := json.NewDecoder(r.Body).Decode(&stock)
	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	insertId := insertStock(stock)

	res := response{
		ID:      insertId,
		Message: "stock inserted successfully.",
	}
	json.NewEncoder(w).Encode(res)
}

func UpdateStock(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Not able to convert ID to int: %v", err)
	}

	var stock models.Stock
	err = json.NewDecoder(r.Body).Decode(&stock)

	if err != nil {
		log.Fatalf("Unable to decode stock: %v", err)
	}

	updatedRows := updateStock(int64(id), stock)
	log.Printf("Stock updated. Total rows affected: %v\n", updatedRows)
	res := response{
		ID:      int64(id),
		Message: "Stock updated successfully",
	}
	json.NewEncoder(w).Encode(res)
}

func DeleteStock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert ID to int: %v", err)
	}

	deleteRows := deleteStock(int64(id))
	log.Printf("Stock deleted. Total rows affected: %v\n", deleteRows)

	res := response{
		ID:      int64(id),
		Message: "Stock deleted succesfully",
	}
	json.NewEncoder(w).Encode(res)
}







