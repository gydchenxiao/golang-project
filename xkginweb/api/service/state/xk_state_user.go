package state

import (
	"time"
	"xkginweb/global"
	"xkginweb/model/entity/state/request"
)

type UserStateService struct {
}

type UserRegState struct {
	Cyear  string `json:"cyear"`
	Cmonth string `json:"cmonth"`
	Ym     string `json:"ym"`
	Cnum   int    `json:"cnum"`
}

type UserRegInfo struct {
	Id         uint      `json:"id"`
	Username   string    `json:"username"`
	Nickname   string    `json:"nickname"`
	Age        int       `json:"age"`
	Sex        int       `json:"sex"`
	Address    string    `json:"address"`
	Job        string    `json:"job"`
	CreateTime time.Time `json:"createTime"`
}

// @author: feige
// 查询某年度的平台用户注册量
func (userService *UserStateService) UserRegStateData(year string) (userRegStates *[]UserRegState, err error) {
	sql := "SELECT DATE_FORMAT(create_time,'%Y') as Cyear,DATE_FORMAT(create_time,'%m') as Cmonth,DATE_FORMAT(create_time,'%Y-%m') as Ym,count(1) as Cnum FROM xk_user  WHERE DATE_FORMAT(create_time,'%Y') = ? GROUP BY DATE_FORMAT(create_time,'%Y-%m')"
	err = global.KSD_DB.Raw(sql, year).Scan(&userRegStates).Error
	return
}

// 查询某年度的平台用户注册量对应的明细信息
func (userService *UserStateService) FindUserRegStateDetail(result request.UserStatePageInfo) (list interface{}, total int64, err error) {
	offset := (result.Page - 1) * result.PageSize
	limit := result.PageSize

	db := global.KSD_DB
	// 执行查询---count
	countSql := "SELECT count(1) FROM xk_user  WHERE DATE_FORMAT(create_time,'%Y-%m') = ?"
	err = db.Raw(countSql, result.Ym).Scan(&total).Error
	// 执行查询-具体查询
	var userRegInfos []UserRegInfo
	sql := "SELECT id,username,age,nickname,sex,address,job,create_time FROM xk_user  WHERE DATE_FORMAT(create_time,'%Y-%m') = ? limit ?,?"
	err = db.Raw(sql, result.Ym, offset, limit).Scan(&userRegInfos).Error
	// 查询返回
	if err != nil {
		return nil, 0, err
	} else {
		return userRegInfos, total, err
	}
}
