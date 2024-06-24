package code

import (
	"github.com/gin-gonic/gin"
	"xkginweb/api/v1/code"
)

type CodeRouter struct{}

func (e *CodeRouter) InitCodeRouter(Router *gin.RouterGroup) {

	codeApi := code.WebApiGroup{}
	// 这个路由多了一个对对post，put请求的中间件处理，而这个中间件做了一些对post和put的参数的处理和一些公共信息的处理
	courseRouter := Router.Group("code") //.Use(middleware.OperationRecord())
	{
		courseRouter.GET("init", codeApi.Captcha)
		//courseRouter.GET("verify", controller.VerfifyCaptcha)
	}
}
