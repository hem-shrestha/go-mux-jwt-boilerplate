package db

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/jinzhu/gorm"
	"os"
)

//database global
var DB *gorm.DB

func SetupDB() *gorm.DB {
	//connect to db
	db, dbError := gorm.Open("sqlite3", "data.db")
	if dbError != nil {
		panic("Failed to connect to database")
	}
	return db
}
