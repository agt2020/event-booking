package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// USERS
	server.POST("/user/signup", createUser)
	// EVENTS
	// Get list of events
	server.GET("/events", getEvents)
	// Get Single event by ID
	server.GET("/event/:id", getEvent)
	// Create Event
	server.POST("/event", createEvent)
	// Update Event
	server.PUT("/event", updateEvent)
	// Delete an event
	server.DELETE("/event/delete/:id", deleteEvent)
}
