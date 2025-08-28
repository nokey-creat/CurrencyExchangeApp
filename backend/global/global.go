// 放一些全局变量,方便项目其他地方访问
package global

import "gorm.io/gorm"

var (
	Db *gorm.DB
)
