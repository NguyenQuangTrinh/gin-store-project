package routers

import (
	"github.com/gin-gonic/gin"
	"store/src/controllers"
)

func UserGroupRouter(baseRouter *gin.RouterGroup) {

	users := baseRouter.Group("users")

	users.GET("/all", controllers.GetAllStartups)
	users.POST("/create", controllers.CreateUser)
	users.GET("getUserById/:id", controllers.GetUserByID)
	users.DELETE("/deleteUser/:id", controllers.DeleteUser)
	users.GET("getUserByUsername/:username", controllers.GetUserByUserName)
}
