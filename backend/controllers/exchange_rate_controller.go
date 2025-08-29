package controllers

import (
	"CurrencyExchangeApp/global"
	"CurrencyExchangeApp/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 新增货币汇率记录
func CreatExchangeRate(ctx *gin.Context) {

	var exchangeRate models.ExchangeRate

	//从请求中获取货币汇率
	if err := ctx.ShouldBindJSON(&exchangeRate); err != nil {
		//请求不匹配，无法绑定的情况
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//增加修改时间项
	exchangeRate.Date = time.Now()

	//自动迁移
	if err := global.Db.AutoMigrate(&exchangeRate); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//创建记录 gorm会自己生成id
	resDb := global.Db.Create(&exchangeRate)
	if resDb.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": resDb.Error.Error()})
		return
	}

	//创建成功，返回响应
	ctx.JSON(http.StatusCreated, exchangeRate)

}

// 获取所有货币汇率记录
func GetExchangeRates(ctx *gin.Context) {

	var exchangeRates []models.ExchangeRate

	//传入切片批量查询，这里返回表的全部结果
	if err := global.Db.Find(&exchangeRates).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//以json格式返回切片（查询结果）
	ctx.JSON(http.StatusOK, exchangeRates)
}
