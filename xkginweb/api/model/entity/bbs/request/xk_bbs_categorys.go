package request

import (
	"xkginweb/model/entity/comms/request"
)

type BbsCategoryPageInfo struct {
	request.PageInfo
	Status int8 `form:"status" json:"status"`
}
