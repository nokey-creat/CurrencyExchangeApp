package router

import (
	"CurrencyExchangeApp/controllers"

	"github.com/gin-gonic/gin"
)

// 注册路由的封装
// 在这里注册路由组，返回gin.Engine示例（gin的路由器）
func SetupRouter() *gin.Engine {

	router := gin.Default() //返回带两个默认中间件的 *gin.Engine

	auth := router.Group("/api/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}

	return router
}
