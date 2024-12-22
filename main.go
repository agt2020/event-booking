package main

import (
	"agt2020/event-booking/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	// Routes registration
	routes.RegisterRoutes(server)

	// Serving a server
	server.Run(":8000")
}
