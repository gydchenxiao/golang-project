package upload

import (
	"github.com/gin-gonic/gin"
	v1 "xkginweb/api/v1"
)

type OssUploadRouter struct{}

func (e *OssUploadRouter) InitOssUploadRouter(Router *gin.RouterGroup) {

	ossUploadApi := v1.WebApiGroupApp.Upload.OSSUploadApi
	// 这个路由多了一个对对post，put请求的中间件处理，而这个中间件做了一些对post和put的参数的处理和一些公共信息的处理
	xkBbsCustomerRouterWithoutRecord := Router.Group("oss") //.Use(middleware.OperationRecord())
	{
		// 保存
		xkBbsCustomerRouterWithoutRecord.POST("upload/file", ossUploadApi.UploadFile)
	}

}
