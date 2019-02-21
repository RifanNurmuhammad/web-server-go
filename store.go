package main

import (
	"database/sql"
)

// Store have two methods, to add a new book
type Store interface {
	CreateBooks(book *Book) error
	GetBooks() ([]*Book, error)
}

type dbStore struct {
	db *sql.DB
}

func (store *dbStore) CreateBooks(book *Book) error {
	_, err := store.db.Query("INSERT INTO books(title, description) VALUS($1,$2)", book.Title, book.Description)
	return err
}

func (store *dbStore) GetBooks() ([]*Book, error) {
	rows, err := store.db.Query("SELECT title, description from books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := []*Book{}
	for rows.Next() {
		book := &Book{}
		if err := rows.Scan(&book.Title, &book.Description); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

var store Store

// InitStore method mwill need to call to initialize the store
func InitStore(s Store) {
	store = s
}
