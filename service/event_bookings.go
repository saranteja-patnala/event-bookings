package service

import (
	"event-bookings/db"
	"event-bookings/models/events"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllEvents(ctx *gin.Context) {
	eventsArray := db.GetAllEvents()
	ctx.JSON(200, gin.H{"events": eventsArray})
}

func GetEventById(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "Invalid event ID", "error": err.Error()})
		return
	}
	eventObj, err := db.GetEventById(id)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Failed to retrieve event", "error": err.Error()})
		return
	}
	if eventObj.ID == 0 {
		ctx.JSON(404, gin.H{"message": "Event not found"})
		return
	}
	ctx.JSON(200, gin.H{"event": eventObj})

}

func CreateBooking(ctx *gin.Context) {
	var eventObj events.Event
	err := ctx.ShouldBindBodyWithJSON(&eventObj)
	if err != nil {
		ctx.JSON(400, gin.H{"message": "Invalid request body", "error": err.Error()})
		return
	}
	eventObj.DateTime = time.Now()
	eventObj = db.SaveEvent(eventObj)
	ctx.JSON(201, gin.H{"message": "Event booking created successfully", "event": eventObj})
}
