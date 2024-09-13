package core

import "test/model"

type IBaseService[D int | uint | string, T model.BaseModel[int]] struct {
}

func (service *IBaseService[D, T]) GetById(id D) T {
	return nil
}

func (service *IBaseService[D, T]) DelById(id D) bool {
	return true
}
