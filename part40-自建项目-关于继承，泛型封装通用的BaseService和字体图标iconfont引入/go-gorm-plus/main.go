package main

import (
	"fmt"
	initilization "go-gorm-plus/init"
	"go-gorm-plus/service"
)

func main() {
	fmt.Println("ok")
	initilization.InitMySQL()

	bbsService := service.BbsService{} //-------------这行代码确定T =model.xkBbs

	//bool, rows := bbsService.UpdateStatus(1, "status", 0)
	//bool, rows := bbsService.UpdateStatus(1, "is_deleted", 1)
	//bool, rows := bbsService.DecrByIdNum(1, "view_count", 10)
	ids := []uint{1, 2, 3}
	bool, rows := bbsService.Incrs(ids, "view_count")
	fmt.Println(bool, rows)
	//// 获取明细
	//xkbbs, _ := bbsService.UnGetByID(1)
	//fmt.Println(xkbbs)
	//
	//menuService := service.MenusService{} //-------------这行代码确定T =model.SysMenus
	//// 获取明细
	//menu, _ := menuService.UnGetByID(1)
	//fmt.Println(menu)

}
