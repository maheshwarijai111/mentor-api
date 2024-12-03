package controllers

import (
	"mentor/database"
	"mentor/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProfile(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello, world! test",
	})
}
func CreateProfile(c *gin.Context) {
	var user models.User
	// Bind JSON input to the User struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	// Save the user in the database
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user.ID})
}
func UpdateProfile(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello, world! test",
	})
}
