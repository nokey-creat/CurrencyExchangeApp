package models

import "gorm.io/gorm"

//用户注册信息表的模型，与users表交互要用到
type User struct {
	gorm.Model
	Username string `gorm:"unique" json:"username"`
	Password string `json:"password"`
}
