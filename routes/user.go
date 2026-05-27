package routes

import (
	"net/http"

	"example.com/event_booking_api/models"
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
