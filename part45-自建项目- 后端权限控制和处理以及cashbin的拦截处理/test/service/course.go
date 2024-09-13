package service

import (
	"test/core"
	"test/model"
)

type CourseService struct {
	core.IBaseService[int, model.Course]
}
