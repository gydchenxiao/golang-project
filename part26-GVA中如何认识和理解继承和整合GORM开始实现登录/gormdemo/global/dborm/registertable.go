package dborm

import (
	"gormdemo/commons"
	"gormdemo/model/user"
)

func RegisterTable() {
	db := commons.KSD_DB
	// 注册和声明 model, 自动在对应的数据库建立 xk_test_user 表
	db.AutoMigrate(&user.User{})
}
