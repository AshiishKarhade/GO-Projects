package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	// "strconv"
	"github.com/gorilla/mux"
)

type Movie struct {
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func main() {
	fmt.Println("CRUD API")
	r := mux.NewRouter()

	movies = append(movies, Movie{Id: "1", Isbn: "8231", Title: "The great king", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{Id: "2", Isbn: "5423", Title: "The worst king", Director: &Director{Firstname: "Jane", Lastname: "Cuo"}})
	movies = append(movies, Movie{Id: "3", Isbn: "64574", Title: "Not so good thing", Director: &Director{Firstname: "John", Lastname: "Cena"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port:8080 . . .")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, item := range movies {
		if item.Id == params["id"] {
			// first delete the id of the movie
			movies = append(movies[:idx], movies[idx+1:]...)
			// then add new id with updated value
			var movie Movie
			movie.Id = item.Id
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:idx], movies[idx+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}
