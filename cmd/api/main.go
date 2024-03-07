package main

import (
	"log"

	"github.com/ArielFarias/poc-go-api/book"
	"github.com/ArielFarias/poc-go-api/book/api"
	gin "github.com/ArielFarias/poc-go-api/book/http"
	repository "github.com/ArielFarias/poc-go-api/book/postgres"
	"github.com/ArielFarias/poc-go-api/datasource"
)

func main() {
	db := datasource.Connect()
	r := repository.NewBookPostgreSQL(&db)
	s := book.NewService(r)
	h := gin.Handlers(s)

	err := api.Start("8080", h)

	if err != nil {
		log.Fatalf("error running api: %s", err)
	}
}
