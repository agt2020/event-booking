package main

import (
	"agt2020/event-booking/db"
	"agt2020/event-booking/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Initdb()

	server := gin.Default()
	server.GET("/events", getEvents)
	// server.GET("/dbtest", checkDbConnection)
	server.POST("/events", createEvent)
	server.Run(":8000")
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "This is a bad request !"})
		return
	}

	event.ID = 1
	event.UserID = 1
	event.Save()

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created !",
		"event":   event,
	})
}

// func checkDbConnection(context *gin.Context) {
// 	db.Initdb()
// 	rows, err := db.RunQuery("SELECT user_id,username,email FROM users")
// 	if err != nil {
// 		context.JSON(500, gin.H{
// 			"connection status": "Failed to fetch",
// 			"error":             err,
// 		})
// 	}
// 	result := db.FetchRows(rows)
// 	context.JSON(200, gin.H{
// 		"data":  result,
// 		"error": nil,
// 	})
// }
