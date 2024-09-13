package bbs

import (
	"github.com/gin-gonic/gin"
	"xkgvaweb/api/bbs"
)

type BbsRouter struct {
}

func (bBsRouter BbsRouter) InitBBsRouter(group *gin.RouterGroup) {
	// 帖子路由
	bbsApi := bbs.BbsApi{}
	bbsApiGroup := group.Group("bbs")
	{
		// 函数封装
		bbsApiGroup.GET("/index", bbsApi.BbsIndex)
		bbsApiGroup.GET("/get/:id", bbsApi.GetBbsDetailById)
	}
}
