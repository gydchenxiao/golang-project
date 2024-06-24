package initilization

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
	"xkginweb/commons/filter"
	"xkginweb/commons/middle"
	"xkginweb/global"
	"xkginweb/router"
	"xkginweb/router/code"
	"xkginweb/router/login"
	"xkginweb/router/sys"
)

func InitGinRouter() *gin.Engine {
	// 打印gin的时候日志是否用颜色标出
	//gin.ForceConsoleColor()
	//gin.DisableConsoleColor()
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 创建gin服务
	ginServer := gin.Default()
	// 提供服务组
	courseRouter := router.RouterWebGroupApp.Course.CourseRouter
	videoRouter := router.RouterWebGroupApp.Video.XkVideoRouter
	menusRouter := router.RouterWebGroupApp.SysMenu.SysMenusRouter
	userStateRouter := router.RouterWebGroupApp.State.UserStateRouter
	bbsRouter := router.RouterWebGroupApp.BBs.XkBbsRouter
	bbsCategoryRouter := router.RouterWebGroupApp.BBs.BBSCategoryRouter

	// 解决接口的跨域问题
	ginServer.Use(filter.Cors())

	loginRouter := login.LoginRouter{}
	logoutRouter := login.LogoutRouter{}
	codeRouter := code.CodeRouter{}
	sysUserRouter := sys.SysUsersRouter{}
	sysRoleRouter := sys.SysRolesRouter{}
	sysApiRouter := sys.SysApisRouter{}
	sysUserRoleRouter := sys.SysUserRolesRouter{}
	sysRoleMenusRouter := sys.SysRoleMenusRouter{}
	sysRoleApisRouter := sys.SysRoleApisRouter{}
	// 接口隔离，比如登录，健康检查都不需要拦截和做任何的处理
	// 业务模块接口，
	privateGroup := ginServer.Group("/api")
	// 无需jwt拦截
	{
		loginRouter.InitLoginRouter(privateGroup)
		codeRouter.InitCodeRouter(privateGroup)
	}
	// 会被jwt拦截
	privateGroup.Use(middle.JWTAuth()).Use(middle.CashBin_RBAC())
	{
		logoutRouter.InitLogoutRouter(privateGroup)
		videoRouter.InitXkVideoRouter(privateGroup)
		courseRouter.InitCourseRouter(privateGroup)
		menusRouter.InitSysMenusRouter(privateGroup)
		userStateRouter.InitUserStateRouter(privateGroup)
		bbsRouter.InitXkBbsRouter(privateGroup)
		bbsCategoryRouter.InitBBSCategoryRouter(privateGroup)

		sysUserRouter.InitSysUsersRouter(privateGroup)
		sysRoleRouter.InitSysRoleRouter(privateGroup)
		sysUserRoleRouter.InitSysUserRolesRouter(privateGroup)
		sysApiRouter.InitSysApisRouter(privateGroup)
		sysRoleApisRouter.InitSysRoleApisRouter(privateGroup)
		sysRoleMenusRouter.InitSysRoleMenusRouter(privateGroup)
	}

	fmt.Println("router register success")
	return ginServer
}

func RunServer() {

	// 初始化路由
	Router := InitGinRouter()
	// 为用户头像和文件提供静态地址
	Router.StaticFS("/static", http.Dir("/static"))
	address := fmt.Sprintf(":%d", global.Yaml["server.port"])
	// 启动HTTP服务,courseController
	s := initServer(address, Router)
	global.Log.Debug("服务启动成功：端口是：", zap.String("port", address))
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)

	s2 := s.ListenAndServe().Error()
	global.Log.Info("服务启动完毕", zap.Any("s2", s2))
}
