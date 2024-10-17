package middlewares

import (
	"EventBooking/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized!"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		fmt.Println(err)
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized!"})
		return
	}

	// Ensures the next request handler in line executes correctly
	context.Set("userId", userId)
	context.Next()
}
