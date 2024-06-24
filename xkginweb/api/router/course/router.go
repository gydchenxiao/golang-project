package course

import (
	"github.com/gin-gonic/gin"
	v1 "xkginweb/api/v1"
)

type CourseRouter struct{}

func (e *CourseRouter) InitCourseRouter(Router *gin.RouterGroup) {

	controller := v1.WebApiGroupApp.Course.CourseController
	// 这个路由多了一个对对post，put请求的中间件处理，而这个中间件做了一些对post和put的参数的处理和一些公共信息的处理
	courseRouter := Router.Group("course") //.Use(middleware.OperationRecord())
	{
		courseRouter.GET("find", controller.FindVideos)
		courseRouter.GET("get", controller.GetByID)
	}
}
