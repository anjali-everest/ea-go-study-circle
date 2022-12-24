package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Book struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

var books = []Book{
	{Id: "546", Title: "Data Structures"},
	{Id: "894", Title: "Design Patterns"},
	{Id: "326", Title: "Networking"},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func addBook(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	books = append(books, book)
	c.JSON(http.StatusCreated, book)
}

func deleteBook(c *gin.Context) {
	id := c.Param("id")

	for index, book := range books {
		if book.Id == id {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}

	c.Status(http.StatusNoContent)
}
