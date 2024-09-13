package model

type BaseModel[T int | uint | string] struct {
	Id       T `json:"id"`
	ParentId T `json:"parentId"`
}
