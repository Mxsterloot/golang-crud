package main

import (
	"golang-crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var books []models.Book

func main() {
	router := gin.Default()

	// สร้างข้อมูลตัวอย่าง
	books = append(books, models.Book{
		ID:     "1",
		Title:  "Harry Potter",
		Author: "J.K. Rowling",
		Price:  99.99,
	})

	// CRUD Endpoints
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookByID)
	router.POST("/books", createBook)
	router.PUT("/books/:id", updateBook)
	router.DELETE("/books/:id", deleteBook)

	router.Run(":8080")
}

// Get all books
func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

// Get book by ID
func getBookByID(c *gin.Context) {
	id := c.Param("id")
	for _, book := range books {
		if book.ID == id {
			c.JSON(http.StatusOK, book)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "หาหนังสือไม่พบ"})
}

// Create new book
func createBook(c *gin.Context) {
	var newBook models.Book
	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	books = append(books, newBook)
	c.JSON(http.StatusCreated, newBook)
}

// Update book
func updateBook(c *gin.Context) {
	id := c.Param("id")
	var updatedBook models.Book

	if err := c.BindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, book := range books {
		if book.ID == id {
			books[i] = updatedBook
			c.JSON(http.StatusOK, updatedBook)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "หาหนังสือไม่พบ"})
}

// Delete book
func deleteBook(c *gin.Context) {
	id := c.Param("id")
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "ลบหนังสือสำเร็จ"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "หาหนังสือไม่พบ"})
}
