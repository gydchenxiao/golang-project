package request

import (
	"xkginweb/model/entity/comms/request"
)

type UserStatePageInfo struct {
	request.PageInfo
	Ym string `form:"ym" json:"ym"`
}
