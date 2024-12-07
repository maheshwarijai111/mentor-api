package controllers

import (
	"mentor/database"
	"mentor/models"
	"mentor/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateLead(c *gin.Context) {
	var lead models.Lead
	// Bind JSON input to the User struct
	if err := c.ShouldBindJSON(&lead); err != nil {
		utils.RespondJSON(c, http.StatusBadRequest, "Validation error", nil, gin.H{"details": err.Error()})
		return
	}
	// Save the user in the database
	if err := database.DB.Create(&lead).Error; err != nil {
		utils.RespondJSON(c, http.StatusInternalServerError, "Failed to create lead", nil, gin.H{"details": err.Error()})
		return
	}
	utils.RespondJSON(c, http.StatusCreated, "Lead created successfully!", lead, nil)
}

func GetLeadList(c *gin.Context) {
	var lead []models.Lead
	if err := database.DB.Table("leads").Find(&lead).Error; err != nil {
		utils.RespondJSON(c, http.StatusNotFound, "lead not found", nil, gin.H{"details": err.Error()})
		return
	}

	utils.RespondJSON(c, http.StatusOK, "Leads retrieved successfully", lead, nil)
}
