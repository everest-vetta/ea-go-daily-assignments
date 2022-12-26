package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func setUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestSuccessfulListingOfBooks(t *testing.T) {
	// Setting Up Web Server
	r := setUpRouter()
	r.GET("/books", getBooks)

	//Client making request
	req, _ := http.NewRequest("GET", "/books", nil)

	//collecting response
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	resBody, _ := io.ReadAll(w.Body)

	var books []Book
	json.Unmarshal(resBody, &books)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSuccessfulCreationOfBook(t *testing.T) {
	// Setting Up Web Server
	r := setUpRouter()
	r.POST("/book", createBook)

	//Client making request
	book := Book{Id: "11", Title: "TestTitle11", Price: 100.0}
	jsonValue, _ := json.Marshal(book)
	req, _ := http.NewRequest("POST", "/book", bytes.NewBuffer(jsonValue))

	//collecting response
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
	resBody, _ := io.ReadAll(w.Body)

	var response Book
	json.Unmarshal(resBody, &response)

	assert.Equal(t, "11", response.Id)
	assert.Equal(t, "TestTitle11", response.Title)
}

func TestSuccessfulUpdateOfBook(t *testing.T) {
	// Setting Up Web Server
	r := setUpRouter()
	r.PUT("book/:id", updateBook)

	//Client making request
	book := Book{Title: "TestTitle12", Price: 100.0}
	jsonValue, _ := json.Marshal(book)
	req, _ := http.NewRequest("PUT", "book/1", bytes.NewBuffer(jsonValue))

	//collecting response
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateForNonExistingBook(t *testing.T) {
	// Setting Up Web Server
	r := setUpRouter()
	r.PUT("/book/:id", updateBook)

	//Client making request
	book := Book{Title: "TestTitle12", Price: 100.0}
	jsonValue, _ := json.Marshal(book)
	req, _ := http.NewRequest("PUT", "/book/100", bytes.NewBuffer(jsonValue))

	//collecting response
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestSuccessfulDeletionOfBook(t *testing.T) {
	// Setting Up Web Server
	r := setUpRouter()
	r.DELETE("/book/:id", deleteBook)

	//Client making request
	req, _ := http.NewRequest("DELETE", "/book/1", nil)

	//collecting response
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeletionForNonExistingBook(t *testing.T) {
	// Setting Up Web Server
	r := setUpRouter()
	r.DELETE("/book/:id", deleteBook)

	//Client making request
	req, _ := http.NewRequest("DELETE", "/book/100", nil)

	//collecting response
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
