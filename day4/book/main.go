package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/books", getBooks)
	router.POST("/book", addBook)
	router.DELETE("/book/:id", deleteBook)
	router.Run()
}
