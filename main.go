package main

import (
	"github.com/gin-gonic/gin"

	"example.com/event_booking_api/db"
	"example.com/event_booking_api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
