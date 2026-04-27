package main

import (
	"event-bookings/db"
	"event-bookings/routers"
	"fmt"
)

func main() {
	fmt.Println("Welcome to the Event Booking System!")

	db.InitDB()
	server := routers.GetServer()

	if err := server.Run("localhost:8080"); err != nil {
		fmt.Printf("Failed to start server: %s\n", err.Error())
	}
}
