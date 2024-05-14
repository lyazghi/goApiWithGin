package main

import (
	respositories "go-api/repositories"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/books", respositories.GetBooks)
	router.POST("/books", respositories.CreateBook)
	router.Run("localhost:8080")
}
