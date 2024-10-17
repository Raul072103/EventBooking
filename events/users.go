package events

import (
	"EventBooking/models"
	"EventBooking/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func signup(context *gin.Context) {
	var user models.User
	fmt.Println(context.Request)
	err := context.ShouldBindJSON(&user)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request data"})
		return
	}

	err = models.Save(&user)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user. Try again later"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created!", "user": user})
}

func login(context *gin.Context) {
	var user models.User

	err := context.BindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse requested data!"})
		return
	}

	err = models.ValidateCredentials(&user)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user!"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.Id)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token})
}
