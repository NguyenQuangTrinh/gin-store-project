package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"html"
	"net/http"
	"store/src/models"
	"store/src/utils"
	"strings"
)

func SignUp(c *gin.Context) {

	var authInput models.User

	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound models.User
	models.Database.Where("username=?", authInput.Username).Find(&userFound)

	if userFound.ID != uuid.Nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username already used"})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(authInput.Password), bcrypt.DefaultCost)
	escapeStringUserName := html.EscapeString(strings.TrimSpace(authInput.Username))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: escapeStringUserName,
		Password: string(passwordHash),
		Age:      authInput.Age,
		Email:    authInput.Email,
		Name:     authInput.Name,
		RoleID:   authInput.RoleID,
	}

	saveUser, err := user.Save()

	if err != nil {
		c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "User create account successfully", "data": err.Error()})
	}
	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Startup saved successfully", "data": saveUser})

}

func LoginUser(context *gin.Context) {
	var authInput models.AuthInput
	if err := context.ShouldBindJSON(&authInput); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := models.GetUserByUserName(authInput.Username)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	err = user.ValidateUserPassword(authInput.Password)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Wrong username or password", "error": err.Error(), "status": ""})
		return
	}

	jwt, err := utils.GenerateJWT(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": jwt, "username": authInput.Username, "message": "Successfully logged in"})
}
