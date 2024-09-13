package dborm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gormdemo/commons"
)

func InitMySQL() {
	// 打开数据库
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:020913@tcp(127.0.0.1:3306)/gva?charset=utf8&parseTime=True&loc=Local", // 登陆数据库的用户名:密码@tcp(127.0.0.1:3306)/数据库 / DSN data source name
		DefaultStringSize:         256,                                                                         // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                        // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                        // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                        // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                       // 根据当前 MySQL 版本自动配置
	}))

	// 如果报错
	if err != nil {
		panic("数据连接出错了" + err.Error()) // 把程序直接阻断，把数据连接好了再启动
	}

	commons.KSD_DB = db

	// 初始化数据库表
	RegisterTable()

	fmt.Println("数据库连接成功。开始运行", db)
}
