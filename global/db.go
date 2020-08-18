package global

import "gorm.io/gorm"

// db 全局变量
var (
	DBEngine *gorm.DB
)
