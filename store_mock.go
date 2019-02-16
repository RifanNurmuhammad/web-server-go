package main

import (
	"github.com/stretchr/testify/mock"
)

// MockStore contains additional method for inspection
type MockStore struct {
	mock.Mock
}

func (m *MockStore) CreateBook(book *Book) error {
	rets := m.Called(book)
	return rets.Error(0)
}

func (m *MockStore) GetBooks() ([]*Book, error) {
	rets := m.Called()
	return rets.Get(0).([]*Book), rets.Error(1)
}

func InitMockStore() *MockStore {
	s := new(MockStore)
	store = s
	return s
}
