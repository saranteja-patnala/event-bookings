package routers

import (
	"event-bookings/main/health"
	"event-bookings/service"

	"github.com/gin-gonic/gin"
)

func GetServer() *gin.Engine {
	server := gin.Default()
	registerServiceHealthRoutes(server)
	registerEventBookingRoutes(server)
	return server
}

func registerServiceHealthRoutes(server *gin.Engine) {
	server.GET("/health", health.GetHealth)
	server.GET("/info", health.GetInfo)
}

func registerEventBookingRoutes(server *gin.Engine) {

	server.GET("/events", service.GetAllEvents)
	server.GET("/events/:id", service.GetEventById)
	server.POST("/events/bookings", service.CreateBooking)

}
