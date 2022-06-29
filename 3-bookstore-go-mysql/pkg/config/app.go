package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// function to connect gorm to mysql database securely
func Connect() {
	d, err := gorm.Open("mysql", "ashishkarhade:ashish7/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

// function to return our db vairable to other packages
func GetDB() *gorm.DB {
	return db
}
