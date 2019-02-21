package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

func TestGetBooksHandler(t *testing.T) {
	// Initialize the mock store
	mockStore := InitMockStore()

	/* Define the data that we want to return when the mocks `GetBooks` method is
	called
	Also, we expect it to be called only once
	*/
	mockStore.On("GetBooks").Return([]*Book{{"sparrow", "A small harmless Book"}}, nil).Once()

	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(getBooksHandler)

	// Now, when the handler is called, it should cal our mock store, instead of
	// the actual one
	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := Book{"sparrow", "A small harmless Book"}
	b := []Book{}
	err = json.NewDecoder(recorder.Body).Decode(&b)

	if err != nil {
		t.Fatal(err)
	}

	actual := b[0]

	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}

	// the expectations that we defined in the `On` method are asserted here
	mockStore.AssertExpectations(t)
}

func TestCreateBooksHandler(t *testing.T) {

	mockStore := InitMockStore()
	/*
		Similarly, we define our expectations for th `CreateBooks` method.
		We expect the first argument to the method to be the Book struct
		defined below, and tell the mock to return a `nil` error
	*/
	mockStore.On("CreateBooks", &Book{"Sprint", "Sprint Design Book"}).Return(nil)

	form := newCreateBookForm()
	req, err := http.NewRequest("POST", "", bytes.NewBufferString(form.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(createBooksHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	mockStore.AssertExpectations(t)
}

func newCreateBookForm() *url.Values {
	form := url.Values{}
	form.Set("title", "Sprint")
	form.Set("description", "Sprint Design Book")
	return &form
}
