package routes

import (
	"comment/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/comments", controllers.AddComment) // 新增评论
	r.GET("/comments", controllers.GetComments) // 查询评论

	return r
}
