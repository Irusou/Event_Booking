package routes

import (
	"fmt"
	"net/http"

	"example.com/event_booking_api/models"
	"example.com/event_booking_api/utils"
	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data."})
		return
	}

	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not save user."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user created successfully"})
}

func login(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data."})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "could not authenticate user."})
		return
	}

	token, err := utils.GenerateJWT(user.Email, user.ID)

	fmt.Println(token)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not authenticate user."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "login successful", "token": token})
}
