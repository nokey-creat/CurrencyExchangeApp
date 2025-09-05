package models

import "gorm.io/gorm"

//文章的模型
type Article struct {
	gorm.Model
	Title   string `binding:"required" json:"title"`
	Content string `binding:"required" json:"content"`
	Preview string `binding:"required" json:"preview"`
	//Likes   int    `gorm:"default:0" json:"Likes"` //点赞通过redis实现了
	Author string `json:"author"`
}
