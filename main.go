package main

import (
	repositories "go-api/repositories"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/books", repositories.GetBooks)
	router.POST("/books", repositories.CreateBook)
	router.Run("localhost:8080")
}
