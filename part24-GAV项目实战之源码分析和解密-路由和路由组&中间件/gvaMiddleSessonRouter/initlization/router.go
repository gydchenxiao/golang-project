package initlization

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"xkgvaweb/api/login"
	"xkgvaweb/middle"
	"xkgvaweb/router/bbs"
	"xkgvaweb/router/video"
)

func WebRouterInit() {
	//  初始化 gin服务
	rootRouter := gin.Default()

	// 创建cookie存储
	store := cookie.NewStore([]byte("secret"))
	//路由上加入session中间件
	rootRouter.Use(sessions.Sessions("mysession", store))

	bbsRouter := bbs.BbsRouter{}
	videoRouter := video.VideoRouter{}
	loginApi := login.LoginApi{}

	// 登录路由（不用拦截）
	rootRouter.GET("/login", loginApi.Login)

	// 首页路由
	// 路由组 -----> 结构清晰，代码优雅
	// rootRouter --------------- /
	//		AdminGroup ---------- /admin
	//			bbsApiGroup ----- /admin/bbs
	//							  /admin/bbs/index
	//							  /admin/bbs/get/:id
	//			videoGroup  ----- /admin/video
	//							  /admin/video/index
	//							  /admin/video/get/:id
	AdminGroup := rootRouter.Group("admin")
	AdminGroup.Use(middle.LoginInterceptor()) //    /admin 加了登陆鉴权中间件拦截，那么后面的都会加上
	{
		bbsRouter.InitBBsRouter(AdminGroup)
		videoRouter.InitVideoRouter(AdminGroup)
	}
	// 启动HTTP服务,可以修改端口
	address := fmt.Sprintf(":%d", 8088)
	rootRouter.Run(address)
}
