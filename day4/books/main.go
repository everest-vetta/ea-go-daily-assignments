package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id    string  `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

type UpdateBookRequest struct {
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

var books = []Book{
	{Id: "1", Title: "Book1", Price: 100.0},
	{Id: "2", Title: "Book2", Price: 200.0},
}

func main() {

	router := gin.New()

	router.GET("/books", getBooks)
	router.GET("/book/:id", getBook)
	router.POST("/book", createBook)
	router.PUT("book/:id", updateBook)
	router.DELETE("book/:id", deleteBook)

	fmt.Println("Starting server at port 8080")
	http.ListenAndServe("localhost:8080", router)
}

func getBooks(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, books)
}

func getBook(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, book := range books {
		if book.Id == id {
			ctx.IndentedJSON(http.StatusOK, book)
			return
		}
	}

	ctx.IndentedJSON(http.StatusNotFound, gin.H{"messaage": "book not found"})
}

func createBook(ctx *gin.Context) {
	var newBook Book

	if err := ctx.BindJSON(&newBook); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err})
		return
	}

	books = append(books, newBook)
	ctx.JSON(http.StatusCreated, newBook)
}

func updateBook(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedBook Book
	if err := ctx.BindJSON(&updatedBook); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err})
	}

	i := -1
	for index, book := range books {
		if book.Id == id {
			i = index
			break
		}
	}

	if i >= 0 && i < len(books) {
		updatedBook.Id = id
		books[i] = updatedBook
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
		return
	}

	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book doesnot exist"})
}

func deleteBook(ctx *gin.Context) {
	id := ctx.Param("id")

	for index, book := range books {
		if book.Id == id {
			if index != len(books)-1 {
				books[index] = books[len(books)-1]
			}
			books = books[:len(books)-1]
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book doesnot exist"})
}
