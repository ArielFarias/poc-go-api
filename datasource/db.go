package datasource

import (
	"github.com/ArielFarias/poc-go-api/book"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	database *gorm.DB
}

func Connect() Database {
	var DB Database

	db, err := gorm.Open(postgres.Open("postgres://postgres@localhost:5432/postgres"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&book.Book{})
	DB.database = db

	return DB
}
