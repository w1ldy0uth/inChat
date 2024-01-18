package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Server is running at " + c.Request.Host,
		})
	})

	fmt.Println("Application started using " + port + " port")
	router.Run()
}
