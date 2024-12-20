package main

import (
	"agt2020/event-booking/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/events", getEvents)
	server.GET("/event/id=?", getEvent)
	server.POST("/events", createEvent)
	server.Run(":8000")
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	id := context.Query("id")
	event, err := models.GetEvent(id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
	}
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	// SAVE EVENT
	id, err := event.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"id":      id,
		"message": "Event successfuly created !",
	})
}
