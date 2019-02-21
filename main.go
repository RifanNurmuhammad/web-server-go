// This is the name of our package
package main

// These are the libraries we are going to use
import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// new router to initiate and test the router outside of the main function
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")

	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	r.HandleFunc("/books", getBooksHandler).Methods("GET")
	r.HandleFunc("/books", createBooksHandler).Methods("POST")
	return r
}

func main() {
	fmt.Println("starting server..")
	connString := "dbname=db_book_store user=postgres password=bukapassword sslmode=disable"
	db, err := sql.Open("postgres", connString)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	InitStore(&dbStore{db: db})
	// calling `newRouter` constructor function
	r := newRouter()
	fmt.Println("Serving on port 8080")
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}
