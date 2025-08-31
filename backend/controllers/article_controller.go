package controllers

import (
	"CurrencyExchangeApp/global"
	"CurrencyExchangeApp/models"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var cacheKey = "articles"

// 新增文章
// 解决缓存失效问题：更新后删去缓存
func CreateArticle(ctx *gin.Context) {

	var article models.Article

	if err := ctx.ShouldBindBodyWithJSON(&article); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := global.Db.AutoMigrate(&article); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//先更新数据库
	if err := global.Db.Create(&article).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//然后删除缓存，使用del命令
	if err := global.RedisDB.Del(ctx.Request.Context(), cacheKey).Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//创建成功，返回201和新创建的文章数据
	ctx.JSON(http.StatusCreated, article)
}

// 获取全部文章数据
func GetArticles(ctx *gin.Context) {

	//读取时采用旁路缓存
	//先尝试从缓存数据库中读取
	cacheData, err := global.RedisDB.Get(ctx.Request.Context(), cacheKey).Result()

	//缓存未命中
	if err == redis.Nil {

		//从sql数据库中读取
		var articles []models.Article
		if err := global.Db.Find(&articles).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}

		//存入redis中
		//Set只能接受字符串或者数字，要先将切片序列化
		articlesJson, err := json.Marshal(articles)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to serialize data"})
			return
		}

		//存入缓存数据库中，并设置过期时间，过期后不可用
		if err := global.RedisDB.Set(ctx.Request.Context(), cacheKey, articlesJson, 10*time.Minute).Err(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		//返回文章
		ctx.JSON(http.StatusOK, articles)

	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	} else {
		//缓存命中
		//反序列化为切片，再返回响应

		var articles []models.Article

		err := json.Unmarshal([]byte(cacheData), &articles)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, articles)

	}

}

// 获取指定id的文章
func GetArticleById(ctx *gin.Context) {

	//通过结构体绑定获取url中的id参数
	var articleURI struct {
		ID string `uri:"id" binding:"required"`
	}
	if err := ctx.ShouldBindUri(&articleURI); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := articleURI.ID

	//下面是查询id对应文章
	var article models.Article

	if err := global.Db.Where("id = ?", id).First(&article).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, article)

}
