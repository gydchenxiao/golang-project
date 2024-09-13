package middle

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 直接用 session 鉴权的话代码上是不太优雅的，就出现了 session 中间件鉴权
func LoginInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取会话
		session := sessions.Default(c)
		// 获取登录用户信息
		user := session.Get("user")
		//  如果用户没有登录，直接重定向返回登录
		if user == nil {
			c.Redirect(http.StatusMovedPermanently, "/login")
			c.Abort() // 拦截,到这里就不会往下执行请求了
		}
		// 取出会话信息
		username := user.(string)
		// 把session用户信息，放入到context文中，个后续路由进行使用
		// 好处就是：router中方法不需要再次获取session在来拿会话中的信息（全局性，后面就可以通过 c.Get("username") 来获取了，而不用再初始化session对象获取了）
		c.Set("username", username) // username --- user --- feige
		c.Next()                    // 放行，默认就会放行
	}
}
