package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Welcome to the Event Booking System!")

	server := gin.Default()

	server.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	})

	server.GET("/info", getInfo)

	if err := server.Run("localhost:8080"); err != nil {
		fmt.Printf("Failed to start server: %s\n", err.Error())
	}
}

func getInfo(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"service": "Event Booking System",
		"version": "1.0.0",
	})
}
