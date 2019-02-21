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
	mockStore := InitMockStore()

	mockStore.On("GetBooks").Return([]*Book{{"Sprint", "Sprint Design Knowledge"}}, nil).Once()
	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	recorder := http.NewRecorder()

	hf := http.HandlerFunc(getBooksHandler)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.ErrorF("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := Book{"Sprint", "Sprint Design Knowledge"}
	b := []Book{}
	err = json.NewDecoder(recorder.Body).Decode(&b)

	if err != nil {
		t.Fatal(err)
	}

	actual := b[0]

	if actual != expected {
		t.ErrorF("handler returned unexpected body: got %v want %v", actual, expected)
	}

	mockStore.AssertExpectations(t)
}

func TestCreateBooksHandler(t *testing.T) {
	mockStore := InitMockStore()

	mockStore.On("CreateBook", &Book{"Zero to One", "Make anything from zero"}).Return(nil)

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
		t.ErrorF("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	mockStore.AssertExpectations(t)
}

func newCreateBookForm() *url.Values {
	form := url.Values{}
	form.Set("title", "Zero to One")
	form.Set("description", "Make anything from zero")
	return &form
}
