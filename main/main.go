package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Welcome to the Event Booking System!")
	gin.New().Run(":8080")
}
