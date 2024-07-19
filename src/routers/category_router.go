package routers

import (
	"github.com/gin-gonic/gin"
	"store/src/controllers"
)

func CategoryGroupRouter(router *gin.RouterGroup) {
	category := router.Group("/category")

	category.GET("/createCategory", controllers.CreateCategory)
}
