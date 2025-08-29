package models

import "gorm.io/gorm"

//文章的模型
type Article struct {
	gorm.Model
	Title   string `binding:"required" json:"Title"`
	Content string `binding:"required" json:"Content"`
	Preview string `binding:"required" json:"Preview"`
	Likes   int    `gorm:"default:0" json:"Likes"`
}
