package config

import (
	"log"

	"github.com/spf13/viper"
)

// 储存配置信息的结构体指针
type Config struct {
	App struct {
		Name string
		Port string
	}
	Database struct {
		Addr     string
		User     string
		Password string
		Name     string

		//连接池中空闲连接的最大数量
		MaxIdleConns int

		//打开数据库连接的最大数量
		MaxOpenConns int

		//连接可复用的最大时间(小时)
		ConnMaxLifetime int
	}
}

// 储存配置信息的结构体指针
var AppConfig *Config

// 初始化并读取配置信息
func InitConfig() {
	//viper读取配置文件
	//单个 Viper 实例仅支持一个配置文件
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config") //./是目前工作目录
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v\n", err)
	}

	//读取后，将其反序列化为结构体
	AppConfig = &Config{} //分配内存
	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("unable to decode into struct: %v\n", err)
	}

	//根据配置初始化数据库连接
	initDB()
	initRedis()
}
