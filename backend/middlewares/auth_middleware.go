package middlewares

import (
	"CurrencyExchangeApp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 权限验证的中间件
func AuthMiddleware() gin.HandlerFunc {
	//Gin 框架通过 gin.Context 结构体在中间件和处理函数之间传递请求信息
	return func(ctx *gin.Context) {

		//检查Authorization请求头是否存在
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization Header"})
			ctx.Abort() //停止调用链
			return      //结束该函数
		}

		//校验jwt
		username, err := utils.ParseJWT(token)
		if err != nil {
			//过期或签名错误
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		//传递username
		ctx.Set("username", username)
		ctx.Next()
	}
}
