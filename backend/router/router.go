package router

import (
	"CurrencyExchangeApp/config"
	"CurrencyExchangeApp/controllers"
	"CurrencyExchangeApp/middlewares"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// 注册路由的封装
// 在这里注册路由组，返回gin.Engine示例（gin的路由器）
func SetupRouter() *gin.Engine {

	router := gin.Default() //返回带两个默认中间件的 *gin.Engine

	//在根路由添加core中间件，处理跨域请求
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{config.AppConfig.CoreConfig.AllowOrigins},  //允许的来源，这里是前端服务器addr
		AllowMethods:     []string{"GET", "POST"},                             //允许的方法
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"}, //除了简单请求头外，允许的头部
		ExposeHeaders:    []string{"Content-Length"},                          //允许js获取的响应头部（不需要Authorization，这是请求头中的）
		AllowCredentials: true,                                                // 允许携带凭证，不直接影响 JWT 发送
		MaxAge:           12 * time.Hour,                                      //本次预检请求的有效期
	}))

	auth := router.Group("/api/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}

	api := router.Group("/api")

	api.GET("/exchangeRates", controllers.GetExchangeRates) //Gin 在注册路由时就将当时组上的中间件链拷贝进该路由。authmiddleware在这个路由注册后use，所以该条路由不包含auth中间件

	api.Use(middlewares.AuthMiddleware())
	{
		api.POST("/exchangeRates", controllers.CreatExchangeRate) //登录后才能修改

		api.POST("/articles", controllers.CreateArticle)
		api.GET("/articles", controllers.GetArticles)
		api.GET("/articles/:id", controllers.GetArticleById)

		api.POST("/articles/:id/likes", controllers.LikeArticle)
		api.GET("/articles/:id/likes", controllers.GetArticleLikesById)

	}

	// 自定义 404 处理
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":  "API endpoint not found",
			"status": 404,
		})
	})

	return router
}
