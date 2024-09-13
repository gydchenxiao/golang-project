package commons

import "gorm.io/gorm"

var (
	KSD_DB *gorm.DB // gorm 框架中 gorm.Open() 方法返回的就是 *gorm.DB 类型
)
