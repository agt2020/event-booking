package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// USERS
	server.POST("/user/signup", signup)
	server.POST("/user/login", login)
	// EVENTS
	server.GET("/events", getEvents)
	server.GET("/event/:id", getEvent)
	server.POST("/event", createEvent)
	server.PUT("/event", updateEvent)
	server.DELETE("/event/delete/:id", deleteEvent)
}
