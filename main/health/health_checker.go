package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "UP",
	})
}

func GetInfo(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"service": "Event Booking System",
		"version": "1.0.0",
	})
}
