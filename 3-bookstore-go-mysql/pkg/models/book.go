package models

import (
	"github.com/AshiishKarhade/bookstore-go-mysql/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publicatoin string `json:"publicaiton"`
}

// function that initialises mysql database
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// function to create a new book in mysql database
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

// function to get all the books from mysql database
func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

// funciton to get a book by it's ID from mysql database
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id)
	db.Find(&getBook)
	return &getBook, db
}

// function to delete a book by its ID from mysql database
func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book
}
