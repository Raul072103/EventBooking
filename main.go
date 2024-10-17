package main

import (
	"EventBooking/db"
	"EventBooking/events"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	events.RegisterRoutes(server)

	err := server.Run(":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
}
