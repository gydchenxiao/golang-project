package service

import (
	"test/core"
	"test/model"
)

type UserService struct {
	core.IBaseService[uint, model.User]
}
