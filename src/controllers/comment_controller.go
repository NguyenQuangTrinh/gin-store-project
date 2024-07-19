package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"store/src/middlewares"
	"store/src/models"
)

func AddComment(c *gin.Context) {
	var input models.Comment

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error(), "data": nil})
		return
	}
	saveComment, err := input.SaveComment()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Startup saved successfully", "data": saveComment})
}

func GetCommentByIdProduct(c *gin.Context) {
	productId := c.Param("id")
	if productId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Startup ID is required"})
		return
	}
	comments, err := models.GetCommentByIdProduct(productId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Startup fetched successfully", "status": "success", "data": comments})
}

func DeleteComment(c *gin.Context) {
	commentId := c.Param("id")
	var err error
	if commentId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Startup ID is required"})
	}
	userId, checkAdmin := middlewares.CheckUser(c)
	if checkAdmin {
		err = models.DeleteCommentAdmin(userId)
	} else {
		err = models.DeleteComment(commentId, userId)
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully", "status": "success"})
}
