package request

import (
	"xkginweb/model/entity/comms/request"
)

type BbsPageInfo struct {
	request.PageInfo
	CategoryId int8 `json:"categoryId" gorm:"not null;default:0;comment:文章分类ID"`
	Status     int8 `json:"status" gorm:"not null;default:1;comment:0 未发布 1 发布"`
}

type BbsCategorySaveReq struct {
	ID          uint   `json:"id" form:"id"`
	Title       string `form:"title" json:"title"`
	Description string `form:"description" json:"description"`
	Parent_id   uint   `form:"parentId" json:"parentId"`
	Sorted      int    `form:"sorted" json:"sorted"`
	Status      int8   `form:"status" json:"status"`
	IsDelete    int8   `form:"isDelete" json:"isDelete"`
}

type StatusReq struct {
	ID    uint   `json:"id"`
	Value int8   `json:"value"`
	Field string `json:"field"`
}
