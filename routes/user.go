package routes

import (
	"agt2020/event-booking/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// SAVE EVENT
	id, err := user.SaveUser()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"id":     id,
		"result": "User successfuly created !",
	})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Login
	err = user.Auth()

	if err == nil {
		context.JSON(http.StatusOK, nil)
	} else {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}
}
