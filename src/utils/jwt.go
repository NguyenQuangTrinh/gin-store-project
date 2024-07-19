package utils

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"store/src/models"
	"strconv"
	"strings"
	"time"
)

var privateKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(user *models.User) (string, error) {
	tokenTTL, _ := strconv.Atoi(os.Getenv("" +
		"" +
		"func ValidateJWT(context *gin.Context) error {"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"role":     user.RoleID,
		"username": user.Username,
		"iat":      time.Now().Unix(),
		"eat":      time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})

	return token.SignedString(privateKey)
}

func ValidateJWT(context *gin.Context) error {
	token, err := GetToken(context)
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
		context.AbortWithStatus(http.StatusUnauthorized)
		return errors.New("token expired")
	}

	return errors.New("invalid token provided")
}

// fetch user details from the token
//func CurrentUser(context *gin.Context) models.User {
//	err := ValidateJWT(context)
//	if err != nil {
//		return models.User{}
//	}
//	token, _ := getToken(context)
//	claims, _ := token.Claims.(jwt.MapClaims)
//	userId := claims["id"]
//
//	user, err := models.FetchUser(userId)
//	if err != nil {
//		return models.User{}
//	}
//	return user
//}

func ValidateCustomerRoleJWT(context *gin.Context) error {
	token, err := GetToken(context)
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	userRole := uint(claims["role"].(float64))
	if ok && token.Valid && userRole == 1 || userRole == 2 || userRole == 3 {
		return nil
	}
	return errors.New("invalid author token provided")
}

func ValidateManagerRoleJWT(context *gin.Context) error {
	token, err := GetToken(context)
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	userRole := uint(claims["role"].(float64))
	if ok && token.Valid && userRole == 2 || userRole == 3 {
		return nil
	}
	return errors.New("invalid author token provided")
}

func ValidateAdminRoleJWT(context *gin.Context) error {
	token, err := GetToken(context)
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	userRole := uint(claims["role"].(float64))
	if ok && token.Valid && userRole == 1 {
		return nil
	}
	return errors.New("invalid admin token provided")
}

func GetToken(context *gin.Context) (*jwt.Token, error) {
	tokenString := getTokenFromRequest(context)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return privateKey, nil
	})
	return token, err
}

func getTokenFromRequest(context *gin.Context) string {
	bearerToken := context.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}
