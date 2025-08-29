package controllers

import (
	"CurrencyExchangeApp/global"
	"CurrencyExchangeApp/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 新增文章
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

	if err := global.Db.Create(&article).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//创建成功，返回201和新创建的文章数据
	ctx.JSON(http.StatusCreated, article)
}

// 获取全部文章数据
func GetArticles(ctx *gin.Context) {

	var articles []models.Article

	if err := global.Db.Find(&articles).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, articles)
}

// 获取指定id的文章
func GetArticleById(ctx *gin.Context) {

	//通过结构体绑定获取url中的id参数
	var articleURI struct {
		ID uint64 `uri:"id" binding:"required"`
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
