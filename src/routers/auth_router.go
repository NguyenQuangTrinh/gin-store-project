package routers

import (
	"github.com/gin-gonic/gin"
	"store/src/controllers"
)

func AuthGroupRouter(baseRouter *gin.RouterGroup) {
	auth := baseRouter.Group("auth")

	auth.POST("/signUp", controllers.SignUp)
	auth.POST("/login", controllers.LoginUser)
}
