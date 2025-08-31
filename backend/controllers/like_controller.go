package controllers

import (
	"CurrencyExchangeApp/global"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

// 为文章点赞的控制器
func LikeArticle(ctx *gin.Context) {

	//获取文章id
	var articleURI struct {
		ID string `uri:"id" binding:"required"`
	}
	if err := ctx.ShouldBindUri(&articleURI); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := articleURI.ID

	likeKey := "article:" + id + ":likes"

	//key对应的value+1，如果不存在，初始化为0再+1
	val, err := global.RedisDB.Incr(ctx.Request.Context(), likeKey).Result()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "successfully like", "likes": val})

}

// 获取文章点赞数
func GetArticleLikesById(ctx *gin.Context) {

	//获取文章id
	var articleURI struct {
		ID string `uri:"id" binding:"required"`
	}
	if err := ctx.ShouldBindUri(&articleURI); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := articleURI.ID

	likeKey := "article:" + id + ":likes"

	val, err := global.RedisDB.Get(ctx.Request.Context(), likeKey).Result()
	if err == redis.Nil {
		val = "0"
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"likes": val})
}
