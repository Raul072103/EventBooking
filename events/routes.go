package events

import (
	"EventBooking/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	eventsRouter := server.Group("/events")
	{
		eventsRouter.GET("", getEvents)
		eventsRouter.POST("", middlewares.Authenticate, createEvent)
		eventsRouter.GET("/:id", getEvent) //events/1
		eventsRouter.PUT("/:id", middlewares.Authenticate, updateEvent)
		eventsRouter.DELETE("/:id", middlewares.Authenticate, deleteEvent)
	}

	server.POST("/signup", signup)
	server.POST("/login", login)
}
