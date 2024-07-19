package routers

import (
	"github.com/gin-gonic/gin"
	"store/src/controllers"
	"store/src/utils"
)

func CommentRouter(baseRouter *gin.RouterGroup) {
	comment := baseRouter.Group("comment")

	comment.Use(utils.JWTAuthCustomer())
	comment.POST("/addComment", controllers.AddComment)
	comment.GET("/getCommentByIdProduct/:id", controllers.GetCommentByIdProduct)
	comment.DELETE("/deleteComment/:id", controllers.DeleteComment)
}
