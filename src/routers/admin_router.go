package routers

import (
	"github.com/gin-gonic/gin"
	"store/src/controllers"
	"store/src/utils"
)

func AdminGroupRouter(c *gin.RouterGroup) {
	admin := c.Group("admin")

	admin.Use(utils.JWTAuth())
	admin.GET("/getProfile", controllers.ProfileAdmin)
	admin.GET("/createUserManager", controllers.CreateUser)
}
