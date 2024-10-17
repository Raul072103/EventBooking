package utils

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"strconv"
	"time"
)

const secretKey = "supersecret"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": strconv.FormatInt(userId, 10),
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Unexpected signing method!")
		}

		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, errors.New("Could not parse token!")
	}

	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	//email := claims["email"].(string)
	userId, err := strconv.ParseInt(claims["userId"].(string), 10, 64)
	return userId, err
}
