package events

import (
	"EventBooking/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		fmt.Println(err)
		return
	}

	var event models.Event
	event, err = models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		fmt.Println(err)
		return
	}

	context.JSON(http.StatusOK, event)
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later"})
		fmt.Println(err)
		return
	}

	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	fmt.Println(context.Request)
	err := context.ShouldBindJSON(&event)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request data"})
		return
	}

	userId := context.GetInt64("userId")
	event.UserID = userId

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later"})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Even created!", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		fmt.Println(err)
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the request event!"})
		fmt.Println(err)
		return
	}

	if event.UserID != userId {
		fmt.Println(userId)
		fmt.Println(event.UserID)
		context.JSON(http.StatusUnauthorized, gin.H{"message": "You didn't create an event"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request data"})
		return
	}

	updatedEvent.ID = eventId
	err = models.Update(updatedEvent)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event!"})
		fmt.Println(err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully!"})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		fmt.Println(err)
		return
	}

	userId := context.GetInt64("userID")
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find the requested event!"})
		fmt.Println(err)
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "You didnt create this event!"})
		return
	}

	err = models.DeleteEvent(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Deleting the event failed!!"})
		fmt.Println(err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Even deleted successfully!"})
}
