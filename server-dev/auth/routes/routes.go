package routes

import (
	"auth/controllers"
	"auth/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	// 使用 JWT 验证中间件保护的路由
	auth := r.Group("/user", middlewares.AuthenticateJWT())
	{
		auth.POST("/update", controllers.UpdateUser)
	}

	return r
}
