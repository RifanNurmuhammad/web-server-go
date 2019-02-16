package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Book definition of the Book struct & initialize variable
type Book struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

var books []Book

func getBooksHandler(w http.ResponseWriter, r *http.Request) {

	books, err := store.GetBooks()

	// convert books variable to json
	booksListBytes, err := json.Marshal(books)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// write the JSON List of birds to the response
	w.Write(booksListBytes)
}

func createBooksHandler(w http.ResponseWriter, r *http.Request) {
	// new instance of book
	book := Book{}

	// Parses the form values
	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// get the information from the form info
	book.Title = r.Form.Get("title")
	book.Description = r.Form.Get("description")

	err = store.CreateBooks(&book)
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/assets/", http.StatusFound)
}
