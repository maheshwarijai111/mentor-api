package controllers

import (
	"mentor/database"
	"mentor/models"
	"mentor/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetProfile(c *gin.Context) {

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr[1:])
	var user []models.UserList
	if id == 0 {
		if err := database.DB.Table("users").Find(&user).Error; err != nil {
			utils.RespondJSON(c, http.StatusNotFound, "User not found", nil, gin.H{"details": err.Error()})
			return
		}
	} else {
		if err := database.DB.Table("users").Where("user_id = ?", id).Find(&user).Error; err != nil {
			utils.RespondJSON(c, http.StatusNotFound, "User not found", nil, gin.H{"details": err.Error()})
			return
		}
	}

	utils.RespondJSON(c, http.StatusOK, "Users retrieved successfully", user, nil)
}

func CreateProfile(c *gin.Context) {
	var user models.User
	// Bind JSON input to the User struct
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "Validation error", nil, gin.H{"details": err.Error()})
		return
	}

	// Hash password before saving (use bcrypt or similar)
	hashedPassword, _ := HashPassword(user.PasswordHash)
	user.PasswordHash = hashedPassword

	// Save the user in the database
	if err := database.DB.Create(&user).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to create user", nil, gin.H{"details": err.Error()})
		return
	}
	utils.RespondJSON(c, http.StatusCreated, "User registered successfully!", user, nil)
}

func UpdateProfile(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello, world! test",
	})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
