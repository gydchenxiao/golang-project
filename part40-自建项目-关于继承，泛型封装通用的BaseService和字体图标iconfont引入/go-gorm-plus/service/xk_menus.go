package service

import (
	"go-gorm-plus/model"
	service "go-gorm-plus/service/common"
)

// 菜单业务层
type MenusService struct {
	// 继续的时候就确定真正数据类型
	service.BaseServiceImpl[int, model.SysMenus]
}
