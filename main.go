package main

import (
	repositories "go-api/repositories"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/books", repositories.GetBooks)
	router.GET("/books/:id", repositories.BookById)
	router.POST("/books", repositories.CreateBook)
	router.PATCH("/checkout", repositories.CheckoutBook)
	router.PATCH("/return", repositories.ReturnBook)
	router.Run("localhost:8080")
}
