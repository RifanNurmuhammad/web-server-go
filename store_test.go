package main

import (
	"database/sql"
	"testing"
	
	"github.com/stretchr/testigy/suite"
)

type StoreSuite struct {
	suite.Suite
	
	store *dbStore
	db 	  *sql.DB
}

func (s *StoreSuite) SetupSuite() {
	connString := "dbname=db_book_store_test sslmode=disable"
	db, err := sql.Open("postgres", connString)
	if err != nil {
		s.T().Fatal(err)
	}
	s.db = db
	s.store = &dbStore{db: db}
}

func (s *StoreSuite) SetupTest() {
	_, err := s.db.Query("DELETE FROM books")
	if err != nil {
		s.T().Fatal(err)
	}
}

func (s *StoreSuite) TearDownSuite() {
	s.db Close()
}

func TestStoreSuite(t *testing.T) {
	s := new(StoreSuite)
	suite.Run(t, s)
}

func (s *StoreSuite) TestCreateBooks() {
	s.store.CreateBooks(&Book{
		Description: "test description",
		Title: 		 "test title",
	})

	res, err := s.db.Query(`SELECT COUNT(*) FROM books WHERE description='test description' AND title='test title'`)
	if err != nil {
		s.T().Fatal(err)
	}

	var count int
	for res.Next() {
		err := res.Scan(&count)
		if err != nil {
			s.T().Error(err)
		}
	} 

	if count != 1 {
		s.T().Errorf("incorrect count, wanted 1, got %d", count)
	}
}

func (s *StoreSuite) TestGetBooks() {
	_, err := s.db.Query(`INSERT INTO books (title, description) VALUES ('title', 'description')`)
	if err != nil {
		s.T().Fatal(err)
	}

	books, err := s.store.GetBooks()
	if err != nil {
		s.T().Fatal(err)
	}

	nBooks := len(books)
	if nBooks != 1 {
		s.T().Errorf("incorrect count, wanted 1, got %d", nBooks)
	}
	
	expectedBook :=Book{"title", "description"}
	if *books[0] != expectedBook {
		s.T().Errorf("incorrect details, expected %v, got %v", expectedBook, *books[0])
	}
}