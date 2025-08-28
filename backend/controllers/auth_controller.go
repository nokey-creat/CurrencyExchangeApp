package controllers

import (
	"CurrencyExchangeApp/global"
	"CurrencyExchangeApp/models"
	"CurrencyExchangeApp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 用户注册的处理器
func Register(ctx *gin.Context) {

	//用户注册信息的模型
	var user models.User

	//将请求体绑定到用户注册信息的模型中 (根据User的tag对应,反序列化时大小写不敏感，优先匹配准确的)
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//需加密储存，将user中的password加密
	HashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user.Password = HashPassword

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	/* 下面是写入数据库部分 */

	//自动迁移，保持schema是最新的。如果没有表，则会创建
	//只有创建初期可以这么用，生产是要手动维护
	if err := global.Db.AutoMigrate(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//创建记录
	resDb := global.Db.Create(&user)
	if resDb.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": resDb.Error.Error()})
		return
	}

	//返回JWT的token
	ctx.JSON(http.StatusOK, gin.H{"token": token, "message": "Successful registration"})
}

// 用户登录的处理器
// 处理逻辑为：比对请求和数据库中的用户名和对应的密码，如果成功返回jwt令牌
func Login(ctx *gin.Context) {

	//来自用户登录请求的输入
	//不能绑定到原来的User结构体中
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindBodyWithJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//与数据库交互的模型
	var user models.User

	//现在数据库中查找名字，并获取对应的密码
	resDb := global.Db.Where("username = ?", input.Username).First(&user)
	//名字不存在的情况
	if resDb.Error != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "wrong username"})
		return
	}

	//再与请求的密码比对

	//密码错误的情况
	if !utils.CheckPassword(input.Password, user.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "wrong password"})
		return
	}

	//密码正确的情况，返回jwt
	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token, "message": "successfully logining"})

}
