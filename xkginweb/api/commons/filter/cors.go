package filter

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 中间件解决跨域的问题
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		// 注意这一行，不能配置为通配符“*”号 比如未来写域名或者你想授权的域名都可以
		//c.Header("Access-Control-Allow-Origin", "http://localhost:8088")
		// 所有的路由都是可以通过的哦（解决跨域的问题）
		c.Header("Access-Control-Allow-Origin", "*")

		// 响应头表示是否可以将对请求的响应暴露给页面。返回true则可以，其他值均不可以。
		c.Header("Access-Control-Allow-Credentials", "true")

		// 表示此次请求中可以使用那些header字段
		c.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Headers,Cookie, Content-Length,Origin,cache-control,X-Requested-With, Content-Type,  Accept, Authorization, KsdUUID,New-Authorization, Token, Timestamp, UserId, New-Token,New-Expires-At") // 我们自定义的header字段都需要在这里声明
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type,New-Authorization, New-Expires-At, New-Token")
		c.Header("Access-Control-Allow-Credentials", "true")
		// 表示此次请求中可以使用那些请求方法 GET/POST(多个使用逗号隔开)
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			//c.AbortWithStatus(http.StatusNoContent)
			c.AbortWithStatus(http.StatusOK)
		}
		// 处理请求
		c.Next()
	}
}
