package login

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginApi struct {
}

func (e *LoginApi) Login(c *gin.Context) {
	// session ---- 是一种所有请求之间的数据共享机制，为什么会出现session，是因为http请求是一种无状态。
	// 什么叫无状态：就是指，用户在浏览器输入方位地址的时候，地址请求到服务区，到响应服务，并不会存储任何数据在客户端或者服务端，
	// 也是就：一次request---response就意味着内存消亡，也就以为整个过程请求和响应过程结束。(数据没有存储，就没了)
	// 但是往往在开发中，我们可能要存存储一些信息，让各个请求之间进行共享。所有就出现了session会话机制（存储的数据量过多的话就会占用大量的服务端资源，就出现了jwt）
	// session会话机制其实是一种服务端存储技术，底层原理是一个map
	// 比如：我登录的时候，要把用户信息存储session中，然后给 map[key]any =
	// key = sdf365454klsdflsd -- sessionid （cookie 会把 sessionid 存储到浏览器，cookie是一种客户端存储技术）

	// 初始化session对象
	session := sessions.Default(c)
	// 存放用户信息到session
	session.Set("user", "feige") //map[sessionid] == map[user][feige]
	// 记住一定调用save方法，否则内存不会写入进去
	session.Save()
	c.JSON(http.StatusOK, "我是gin")
}
