package main

import (
	"net/http"

	"example.com/event_booking_api/db"
	"example.com/event_booking_api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events."})
		return
	}
	c.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}

	event.ID = 1
	event.UserID = 10
	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create events."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})
}
