package jwtgo

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"xkginweb/global"
	"xkginweb/model/entity/jwt"
	"xkginweb/utils"
)

type JwtService struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: JsonInBlacklist
//@description: 拉黑jwt
//@param: jwtList model.JwtBlacklist
//@return: err error

func (jwtService *JwtService) JsonInBlacklist(jwtList jwt.JwtBlacklist) (err error) {
	err = global.KSD_DB.Create(&jwtList).Error
	if err != nil {
		return
	}
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: IsBlacklist
//@description: 判断JWT是否在黑名单内部
//@param: jwt string
//@return: bool

func (jwtService *JwtService) IsBlacklist(jwttoken string) bool {
	//_, ok := global.BlackCache.Get(jwt)
	//return ok
	err := global.KSD_DB.Where("jwt = ?", jwttoken).First(&jwt.JwtBlacklist{}).Error
	isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	return !isNotFound
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetRedisJWT
//@description: 从redis取jwt
//@param: userName string
//@return: redisJWT string, err error

func (jwtService *JwtService) GetRedisJWT(userName string) (redisJWT string, err error) {
	redisJWT, err = global.REDIS.Get(context.Background(), userName).Result()
	return redisJWT, err
}

// @author: [piexlmax](https://github.com/piexlmax)
// @function: SetRedisJWT
// @description: jwt存入redis并设置过期时间
// @param: jwt string, userName string
// @return: err error
func (jwtService *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	dr, err := utils.ParseDuration("7d")
	if err != nil {
		return err
	}
	timer := dr
	err = global.REDIS.Set(context.Background(), userName, jwt, timer).Err()
	return err
}

func LoadAll() {
	var data []string
	err := global.KSD_DB.Model(&jwt.JwtBlacklist{}).Select("jwt").Find(&data).Error
	if err != nil {
		fmt.Println("加载数据库jwt黑名单失败!", zap.Error(err))
		return
	}
	for i := 0; i < len(data); i++ {
		global.BlackCache.SetDefault(data[i], struct{}{})
	} // jwt黑名单 加入 BlackCache 中
}