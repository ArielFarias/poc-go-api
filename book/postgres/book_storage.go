package repository

import (
	"github.com/ArielFarias/poc-go-api/book"
	"github.com/ArielFarias/poc-go-api/datasource"
)

type BookPostgreSQL struct {
	db *datasource.Database
}

func NewBookPostgreSQL(db *datasource.Database) *BookPostgreSQL {
	return &BookPostgreSQL{
		db: db,
	}
}

func (r *BookPostgreSQL) Create(e *book.Book) (int, error) {
	return 1, nil
}

func (r *BookPostgreSQL) Get(id int) (book.Book, error) {
	return book.Book{}, nil
}

func (r *BookPostgreSQL) GetAll() ([]book.Book, error) {
	return []book.Book{}, nil
}
