package login

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"xkginweb/commons/jwtgo"
	"xkginweb/commons/response"
	"xkginweb/global"
	"xkginweb/model/entity/jwt"
)

// 登录业务
type LogOutApi struct{}

var jwtService = jwtgo.JwtService{}

// 退出接口
func (api *LogOutApi) ToLogout(c *gin.Context) {
	// 获取头部的token信息
	token := c.GetHeader("Authorization")
	if token == "" {
		response.Fail(401, "请求未携带token，无权限访问", c)
		return
	}

	// 同时删除缓存中的uuid的信息
	customClaims, _ := jwtgo.GetClaims(c)
	userIdStr := strconv.FormatUint(uint64(customClaims.UserId), 10)
	global.Cache.Delete("LocalCache:Login:" + userIdStr)

	// 退出的token,加入到黑名单中
	err := jwtService.JsonInBlacklist(jwt.JwtBlacklist{Jwt: token})
	// 保存失败会进到到错误
	if err != nil {
		response.Fail(401, "token作废失败", c)
		return
	}
	// 如果保存到黑名单中说明,已经可以告知前端可以进行执行清理动作了
	response.Ok("token作废成功!", c)
}
