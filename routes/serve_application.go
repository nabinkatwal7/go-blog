package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ServeApplication() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server responded with status code 200",
		})
	})
	router.Run("localhost:8080")
	fmt.Println("Server started on Port 8080")
}
