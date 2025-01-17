package middlewares

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"os"
	"store/src/models"
	"store/src/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func CheckAuth(c *gin.Context) {

	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	authToken := strings.Split(authHeader, " ")
	if len(authToken) != 2 || authToken[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tokenString := authToken[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var user models.User
	models.Database.Where("ID=?", claims["id"]).Find(&user)

	if user.ID == uuid.Nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("currentUser", user)

	c.Next()
}

func CheckUser(context *gin.Context) (string, bool) {
	token, _ := utils.GetToken(context)

	claims, _ := token.Claims.(jwt.MapClaims)

	userID := claims["id"].(string)
	role := uint(claims["role"].(float64))

	if role == 0 {
		return "", true
	}

	return userID, false
}
