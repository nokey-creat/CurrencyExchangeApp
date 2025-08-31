// 放一些全局变量,方便项目其他地方访问
package global

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Db      *gorm.DB
	RedisDB *redis.Client
)
