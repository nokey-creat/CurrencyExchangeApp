package models

import "time"

//货币汇率表的模型
type ExchangeRate struct {
	ID           uint    `gorm:"primarykey" json:"_id"`
	FromCurrency string  `json:"fromCurrency" binding:"required"`
	ToCurrency   string  `json:"toCurrency" binding:"required"`
	Rate         float64 `json:"rate" binding:"required"`

	//最后修改时间
	Date time.Time `json:"date"`
}
