package routes

import (
	"agt2020/event-booking/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvent(context *gin.Context) {
	id := context.Param("id")
	if !isValidID(id) {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Event ID"})
		return
	}

	event, err := models.GetEvent(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, event)
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// SAVE EVENT
	id, err := event.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"id":     id,
		"result": "Event successfuly created !",
	})
}

func updateEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := event.Update()
	if err != nil || id == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	fetchedEvent, err := models.GetEvent(strconv.Itoa(id))
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"result": "Event updated successfuly",
		"data":   fetchedEvent,
	})
}

func deleteEvent(context *gin.Context) {
	id := context.Param("id")
	if !isValidID(id) {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Event ID"})
		return
	}

	err := models.DeleteEvent(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, nil)
}

func isValidID(id string) bool {
	if id == "" {
		return false
	}
	_, err := strconv.Atoi(id)
	return err == nil
}
