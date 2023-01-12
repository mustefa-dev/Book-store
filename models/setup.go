package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func ConnectDatabase() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "postgres", "postgres")
	database, _ := gorm.Open("postgres", psqlInfo)
	database.AutoMigrate(&Book{})
	DB = database

	//database, _ := gorm.Open("sqlite3", "aa.db")
	//database.AutoMigrate(&Book{})
	//DB = database
}

//ENV Variable
