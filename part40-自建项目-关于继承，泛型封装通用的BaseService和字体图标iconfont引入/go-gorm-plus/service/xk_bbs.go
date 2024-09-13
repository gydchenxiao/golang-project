package service

import (
	"go-gorm-plus/model"
	service "go-gorm-plus/service/common"
)

// bbs业务层
type BbsService struct {
	service.BaseServiceImpl[uint, model.XkBbs]
}
