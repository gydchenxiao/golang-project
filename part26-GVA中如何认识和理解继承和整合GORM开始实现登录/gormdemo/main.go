package main

import (
	"fmt"
	"gormdemo/commons"
	"gormdemo/global"
	"gormdemo/global/dborm"
	"gormdemo/model/user"
	"time"
)

func main() {
	// 先提前把需要管理的对象全部在InitIoc进行提前注册，好处什么：不用到处去实例化
	// 你需要对象在map去拿就行了
	global.InitIoc()
	// 获取数据连接对象
	dborm.InitMySQL()

	fmt.Println("commons.KSD_DB", commons.KSD_DB)

	dbUser := user.User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	result := commons.KSD_DB.Create(&dbUser) // 通过数据的指针来创建
	fmt.Println("用户添加成功", result)
}
