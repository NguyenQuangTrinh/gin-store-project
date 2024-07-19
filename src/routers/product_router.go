package routers

import (
	"github.com/gin-gonic/gin"
	"store/src/controllers"
)

func ProductRouter(baseRouter *gin.RouterGroup) {
	product := baseRouter.Group("product")

	product.GET("/getAll", controllers.GetProduct)
	product.POST("/create", controllers.CreateProduct)
	product.GET("/getProduct", controllers.GetProduct)
	product.GET("/getProductById/:id", controllers.GetProductById)
}
