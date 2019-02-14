// main package
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)

	// in case there is an error in forming the request
	if err != nil {
		t.Fatal(err)
	}

	// use Go httptest library to create an http recorder
	// this recorder will act as the target of our http request
	recorder := httptest.NewRecorder()

	// create an HTTP handler from handler function.
	// handler is the function defined in our main.go
	hf := http.HandlerFunc(handler)

	// serve the HTTP request to our recorder.
	// actually executes our the handler that we want to test
	hf.ServeHTTP(recorder, req)

	// check the status code is what we expect.
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// check the response body is what we expect.
	expected := `Hello Go!`
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}
