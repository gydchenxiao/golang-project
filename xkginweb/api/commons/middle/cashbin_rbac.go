package middle

import (
	"github.com/gin-gonic/gin"
	"strings"
	"xkginweb/commons/jwtgo"
	"xkginweb/commons/response"
	"xkginweb/global"
	"xkginweb/service/sys"
)

var casbinService = sys.CasbinService{}

// 定义一个Casbin的中间件
func CashBin_RBAC() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取sub角色信息
		sub := jwtgo.GetUserRoleCode(c)
		// 请求路径
		path := c.Request.URL.Path
		obj := strings.TrimPrefix(path, "")
		// 请求方式
		act := c.Request.Method

		e := casbinService.CasbinFile() // 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		if !success {
			global.Log.Info("权限不足的地址是：" + path + ",角色是：" + sub + ",请求方式: " + act)
			response.FailWithPermission(c)
			c.Abort()
			return
		}
		c.Next()
	}
}
