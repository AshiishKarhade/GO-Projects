package router

import (
	"net/http"

	"github.com/AshiishKarhade/GO-Projects/go-postgres-stocksapi/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// Home Handler
	router.HandleFunc("/", HomeHandler)

	// GET Stock details from DB
	router.HandleFunc("/api/stock/{id}", middleware.GetStock).Methods("GET", "OPTIONS")

	// GET All stock details from DB
	router.HandleFunc("/api/stocks", middleware.GetAllStocks).Methods("GET", "OPTIONS")

	// POST stock details to DB
	router.HandleFunc("/api/newstock", middleware.CreateStock).Methods("POST", "OPTIONS")

	// UPDATE stock details in DB
	router.HandleFunc("/api/stock/{id}", middleware.UpdateStock).Methods("PUT", "OPTIONS")

	// DELETE stock details from DB
	router.HandleFunc("/api/stock/{id}", middleware.DeleteStock).Methods("DELETE", "OPTIONS")

	return router
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

}
