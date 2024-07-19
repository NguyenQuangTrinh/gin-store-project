package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"store/src/models"
)

func GetAllStartups(c *gin.Context) {
	startups, err := models.FetchAllUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Startups fetched successfully", "status": "success", "data": startups})
}

func CreateUser(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error(), "data": nil})
		return
	}
	savedStartup, err := input.Save()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "User saved successfully", "data": savedStartup})
}

func GetUserByID(c *gin.Context) {
	userId := c.Param("id")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Startup ID is required"})
		return
	}
	startup, err := models.FetchUser(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Startup fetched successfully", "status": "success", "data": startup})
}

func DeleteUser(c *gin.Context) {
	userId := c.Param("id")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Startup ID is required"})
		return
	}

	err := models.DeleteUser(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Startup deleted successfully", "status": "success", "data": nil})
}

func GetUserByUserName(c *gin.Context) {
	username := c.Param("username")

	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Startup username is required"})
		return
	}
	user, err := models.GetUserByUserName(username)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Startup fetched successfully", "data": user})
}
