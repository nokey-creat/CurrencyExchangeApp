package config

import (
	"CurrencyExchangeApp/global"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化数据库配置
func initDB() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		AppConfig.Database.User, AppConfig.Database.Password, AppConfig.Database.Addr, AppConfig.Database.Name)

	//创建gorm.DB实例，即gorm的会话入口
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to initialize database, got err: %v", err)
	}

	//获取标准库 database/sql 的连接池句柄，用于配置和管理底层连接池
	sqlDB, err := db.DB()

	if err != nil {
		log.Fatalf("failed to configure database, got err: %v", err)
	}

	sqlDB.SetMaxIdleConns(AppConfig.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(AppConfig.Database.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(AppConfig.Database.ConnMaxLifetime) * time.Hour)

	//将初始化完成的gorm.DB实例指针放在全局变量中
	global.Db = db

}
