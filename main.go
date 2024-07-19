package main

import (
	"github.com/gin-contrib/cors"
	"net/http"
	"store/src/models"
	"store/src/routers"
	"store/src/utils"
	"time"
)

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	router.Use(cors.New(config))

	utils.LoadEnv()

	models.OpenDatabaseConnection()
	models.AutoMigrateModels()
	//models.SeedData()

	router.GET("/gin", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello Gin!"})
	})

	v1 := router.Group("api/v1")

	routers.UserGroupRouter(v1)
	routers.AuthGroupRouter(v1)
	routers.ProductRouter(v1)
	routers.CommentRouter(v1)
	routers.AdminGroupRouter(v1)
	routers.CategoryGroupRouter(v1)
	// serve and listen to localhost:8080
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
