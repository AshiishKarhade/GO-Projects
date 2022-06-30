package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AshiishKarhade/bookstore-go-mysql/pkg/models"
	"github.com/AshiishKarhade/bookstore-go-mysql/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func CreateBook(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("create book")
	createBook := &models.Book{}
	utils.ParseBody(r, createBook)
	b := createBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// function to get ALL the books from DB and send a reponse
func GetBook(w http.ResponseWriter, r *http.Request) {
	newbooks := models.GetAllBooks()
	res, _ := json.Marshal(newbooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// function to get a book by it's ID from DB and send a response
func GetBookById(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("get book by id")
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing in GetBookById()")
	}
	book, _ := models.GetBookById(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// function to UPDATE a book by it's ID from DB and send a response
func UpdateBookById(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("update book by id")
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing in DeleteBookById()")
	}
	bookDetails, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publicatoin != "" {
		bookDetails.Publicatoin = updateBook.Publicatoin
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// function to DELETE a book by it's ID from DB and send a response
func DeleteBookById(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("delete book by id")
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing in DeleteBookById()")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
