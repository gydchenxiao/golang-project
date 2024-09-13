# GVA 项目社区平台管理系统开发



## 01、swagger文档

https://github.com/swaggo/swag/blob/master/README_zh-CN.md

## 安装 swagger [#](https://www.gin-vue-admin.com/guide/start-quickly/swagger.html#_1-安装-swagger)

- 可以翻墙

```
go install github.com/swaggo/swag/cmd/swag@version
```

- 无法翻墙 由于国内没法安装 [go.org/x](http://go.org/x) 包下面的东西，推荐使用 [goproxy.cn](https://goproxy.cn/zh/) 或者 [goproxy.cn/](https://goproxy.cn/)

```
# Go 版本 是 1.16 ~ 最新版 可以忽略以下步骤一
# 步骤一、启用 Go Modules 功能
go env -w GO111MODULE=on 
# 步骤二、配置 GOPROXY 环境变量
go env -w GOPROXY=https://goproxy.cn,direct

# 使用如下命令下载swag
go install github.com/swaggo/swag/cmd/swag@version
```

## 2 生成API文档[#](https://www.gin-vue-admin.com/guide/start-quickly/swagger.html#_2-生成api文档)



```
cd server
swag init
```

执行上面的命令后，server目录下会出现docs文件夹，打开浏览器复制如下地址，即可查看swagger文档

```
http://localhost:8888/swagger/index.html
```

## 如果遇到安装报错可以执行下面的命令

```sh
## 如果不行，获取更新包：
go get github.com/swaggo/echo-swagger

##以上还是不行，在更新
go get github.com/alecthomas/template
```



# GVA开发流程

gva已经提供一个基础的系统架构和开发很多后台项目管理的一些基础功能。我们只需要把我们的业务和相关模块进行二次开发即可。就可以形成一个完整的关于你自己的后台管理系统。

```sh
系统：web3比特币社区管理项目
后台技术架构：gva + mysql 
前端技术架构：vue3 + vite + elemenetplus
开发周期：6个月
开始人数：11个
系统介绍：
GIN-VUE-ADMIN是一个基于vue和gin开发的全栈前后端分离的开发基础平台，拥有jwt鉴权，动态路由，动态菜单，casbin鉴权，表单生成器，代码生成器等功能，提供了多种示例文件，让大家把更多时间专注在业务开发上。
我的负责的模块：
1: 插件中心 NEW :基于 Gva自己的一套设计风格，独创 go的插件中心，现已支持 ：微信支付、登录等，K8s相关操作 ，第三方登录 等等插件
2: 权限管理：基于jwt和casbin实现的权限管理
3: 文件上传下载：实现基于七牛云的文件上传操作（为了方便大家测试，我公开了自己的七牛测试号的各种重要token，恳请大家不要乱传东西）
4: 用户管理：系统管理员分配用户角色和角色权限。
5: 角色管理：创建权限控制的主要对象，可以给角色分配不同api权限和菜单权限。
6: 菜单管理：实现用户动态菜单配置，实现不同角色不同菜单。
7: api管理：不同用户可调用的api接口的权限不同。
8: 配置管理：配置文件可前台修改（测试环境不开放此功能）。
9: 富文本编辑器：MarkDown编辑器功能嵌入。
10: 条件搜索：增加条件搜索示例。

```

使用gva 0 ~ 0.5 ~1 



# GVA整体的项目架构怎么分析

## 01、前后端开发的本质

服务端—-写接口—写路由—访问某个模块的数据看表—-使用gorm映射go对应的结构体中—–返回

客户端-–掉接口–根据服务端提供接口路由—-发起请求—ginserver接收请求获取参数—-执行路由方法—-返回路由定义模块的数据 —-结束

## 02、gva项目架构分析

如何在gva项目融入我们的业务。这里我就飞哥自己的思考告诉大家。

- gva它是一个项目搭建好的项目架构

  - 写接口

- 确定它技术架构栈

  - gin

    - web服务 端口 8888
    - 如何定义路由 
      - user/list,user/get.user/save,user/update,user/delete	
      - course/list,course/get.course/save,course/update,course/delete	 —中间件
      - ….
    - 路由如何配置中间件（统一拦截）
    - 命名空间（路由组）
      - PrivateRouterGroup —中间件
        - userRouterGroup
          - user/list,user/get.user/save,user/update,user/delete	
        - courseRouterGroup	
          - course/list,course/get.course/save,course/update,course/delete	
    - 参数的获取–request————–Req
      - 单参数 /user/get?id=1&name=zhangsan
      - 对象参数 body {id:1,name:”zhangsan”}——-
      - 路径参数  /user/get/1/zhangsan
    - 返回—response—————-VO
      - 统一返回

  - gorm

    - 数据的持久层框架 对数据库表进行curd和分页等相关的操作
    - 比如对某个表的CURD的操作要非常的熟悉。

  - api 

    - 暴露接口

  - docs

    - swagger暴露接口的目录，这个需要执行命令生成doc.go的文件

  - global 

    - 这个主要是用来做全局参数的一些获取，比如：配置文件解析的参数，全部可以使用

    ```sh
    package global
    
    import (
    	"sync"
    
    	"github.com/flipped-aurora/gin-vue-admin/server/utils/timer"
    	"github.com/songzhibin97/gkit/cache/local_cache"
    
    	"golang.org/x/sync/singleflight"
    
    	"go.uber.org/zap"
    
    	"github.com/flipped-aurora/gin-vue-admin/server/config"
    
    	"github.com/go-redis/redis/v8"
    	"github.com/spf13/viper"
    	"gorm.io/gorm"
    )
    
    var (
    	GVA_DB     *gorm.DB
    	GVA_DBList map[string]*gorm.DB
    	GVA_REDIS  *redis.Client
    	GVA_CONFIG config.Server
    	GVA_VP     *viper.Viper
    	// GVA_LOG    *oplogging.Logger
    	GVA_LOG                 *zap.Logger
    	GVA_Timer               timer.Timer = timer.NewTimerTask()
    	GVA_Concurrency_Control             = &singleflight.Group{}
    
    	BlackCache local_cache.Cache
    	lock       sync.RWMutex
    )
    
    // GetGlobalDBByDBName 通过名称获取db list中的db
    func GetGlobalDBByDBName(dbname string) *gorm.DB {
    	lock.RLock()
    	defer lock.RUnlock()
    	return GVA_DBList[dbname]
    }
    
    // MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
    func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
    	lock.RLock()
    	defer lock.RUnlock()
    	db, ok := GVA_DBList[dbname]
    	if !ok || db == nil {
    		panic("db no init")
    	}
    	return db
    }
    
    ```

    不管任何框架，未来初始化肯定只有一个地方：main函数，或者用模块的就是init方法，然后导入main方法所在的包中来进行初始化。

  - initialize

    - 数据库
    - 日志
    - 缓存
    - 等等

    这里定义的所以的中间件的初始化都在这里。但是它的初始化必须是：main函数来驱动的。

  - middleware——————-gin的路由

    - 中间件，包含路由资源，然后做统一处理
    - 鉴权拦截 jwt
    - casbin的权限拦截

  - model

    - 和数据库表进行映射

  - router

    - 提供对外的接口

  - service

    - 提供CURD操作给api进行调用

  - source

    - 数据初始化

  - utils

    - 你要没事多看看和收集起来，

  - resoruce 

    - 代码自动构建的模块

  

  

  

  # GVA开发一个文章模块

  ## 基本流程如下

  ### 数据库的初始化链接

  - 新建一个表xk_bbs
    - 自动创建
    - 手动在navicat创建，然后在映射

  ### model

  ```go
  package bbs
  
  import "github.com/flipped-aurora/gin-vue-admin/server/global"
  
  type XkBbs struct {
  	global.GVA_MODEL
  	Title        string `json:"title" gorm:"not null;index;comment:标题"`
  	Img          string `json:"img"  gorm:"not null;default:'';comment:封面图"`
  	Description  string `json:"description" gorm:"not null;default:'';comment:描述"`
  	Content      string `json:"content" gorm:"not null;default:'';comment:文章内容--MD格式"`
  	HtmlContent  string `json:"htmlContent" gorm:"not null;default:'';comment:文章内容--MD格式"`
  	CategoryId   uint   `json:"categoryId" gorm:"not null;default:0;comment:文章分类ID"`
  	CategoryName string `json:"categoryName" gorm:"not null;default:'';comment:文章分类名称"`
  	ViewCount    int8   `json:"viewCount" gorm:"not null;default:0;comment:文章阅读数"`
  	Comments     int8   `json:"comments" gorm:"not null;default:0;comment:评论数"`
  	CommentsOpen int8   `json:"commentsOpen" gorm:"not null;default:1;comment:是否允许评论 0 不允许  1 允许"`
  	Status       int8   `json:"status" gorm:"not null;default:1;comment:0 未发布 1 发布"`
  	IsDelete     int8   `json:"isDelete" gorm:"not null;default:0;comment:0 未删除 1 删除"`
  	UserId       uint   `json:"userId" gorm:"not null;comment:文章作者ID"`
  	Username     string `json:"username"  gorm:"not null;default:'';comment:文章发布者用户名"`
  	Avatar       string `json:"avatar"  gorm:"not null;default:'';comment:文章发布者头像"`
  }
  
  func (XkBbs) TableName() string {
  	return "xk_bbs"
  }
  
  ```

  #### 注册模型

  目的是让gorm能够加载对应model。然后可以自动生成数据库表，然后结构体发生变化，它会自动增加列。然后在 initialize/gorm.go 如下：

  ```go
  package initialize
  
  import (
  	"github.com/flipped-aurora/gin-vue-admin/server/model/bbs"
  	"github.com/flipped-aurora/gin-vue-admin/server/model/videocategory"
  	"os"
  
  	"github.com/flipped-aurora/gin-vue-admin/server/global"
  	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
  	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
  
  	"go.uber.org/zap"
  	"gorm.io/gorm"
  )
  
  func Gorm() *gorm.DB {
  	switch global.GVA_CONFIG.System.DbType {
  	case "mysql":
  		return GormMysql()
  	case "pgsql":
  		return GormPgSql()
  	case "oracle":
  		return GormOracle()
  	case "mssql":
  		return GormMssql()
  	default:
  		return GormMysql()
  	}
  }
  
  func RegisterTables() {
  	db := global.GVA_DB
  	err := db.AutoMigrate(
  
  		system.SysApi{},
  		system.SysUser{},
  		system.SysBaseMenu{},
  		system.JwtBlacklist{},
  		system.SysAuthority{},
  		system.SysDictionary{},
  		system.SysOperationRecord{},
  		system.SysAutoCodeHistory{},
  		system.SysDictionaryDetail{},
  		system.SysBaseMenuParameter{},
  		system.SysBaseMenuBtn{},
  		system.SysAuthorityBtn{},
  		system.SysAutoCode{},
  		system.SysChatGptOption{},
  
  		example.ExaFile{},
  		example.ExaCustomer{},
  		example.ExaFileChunk{},
  		// 生成视频分类表
  		bbs.XkBbs{},
  	)
  	if err != nil {
  		global.GVA_LOG.Error("register table failed", zap.Error(err))
  		os.Exit(0)
  	}
  	global.GVA_LOG.Info("register table success")
  }
  
  ```

  重新启动main.go。查看数据库中是否生成xk-bbs和xk_video_categorys

  

  ### service

  就是对当前xk-bbs的数据做curd操作。目的就是curd

  - create 查询的意思，使用gorm把结构体中的数据，同步到保存数据库表中
  - update更新的意思，使用gorm把结构体中的数据，同步到更新数据库表中
  - read 读取的意思，含义使用gorm框架把数据库的数据映射到go语言中结构体中。
  - drop 删除的意思，把数据库中的记录删除

  

  ```java
  package bbs
  
  import (
  	"github.com/flipped-aurora/gin-vue-admin/server/global"
  	"github.com/flipped-aurora/gin-vue-admin/server/model/bbs"
  	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
  )
  
  // 定义bbs的service提供xkbbs的数据curd的操作
  
  type XkBssService struct{}
  
  // @author: feige
  // @function: CreateXkBbs
  // @description: 创建文章
  // @param: e bbs.XkBbs
  // @return: err error
  func (cbbs *XkBssService) CreateXkBbs(xkBbs *bbs.XkBbs) (err error) {
  	// 1： 获取数据的连接对象
  	err = global.GVA_DB.Create(xkBbs).Error
  	return err
  }
  
  //@author: feige
  //@function: UpdateXkBbs
  //@description: 更新文章
  //@param: e *model.ExaCustomer
  //@return: err error
  
  func (cbbs *XkBssService) UpdateXkBbs(xkBbs *bbs.XkBbs) (err error) {
  	err = global.GVA_DB.Save(xkBbs).Error
  	return err
  }
  
  // @author: feige
  // @function: DeleteXkBbs
  // @description: 删除帖子
  // @param: e model.DeleteXkBbs
  // @return: err error
  func (cbbs *XkBssService) DeleteXkBbs(xkBbs *bbs.XkBbs) (err error) {
  	err = global.GVA_DB.Delete(&xkBbs).Error
  	return err
  }
  
  // @author: feige
  // @function: GetXkBbs
  // @description: 根据ID获取帖子信息
  // @param: id uint
  // @return: xkBbs *bbs.XkBbs, err error
  func (cbbs *XkBssService) GetXkBbs(id uint) (xkBbs *bbs.XkBbs, err error) {
  	err = global.GVA_DB.Where("id = ?", id).First(&xkBbs).Error
  	return
  }
  
  //@author: [piexlmax](https://github.com/piexlmax)
  //@function: GetCustomerInfoList
  //@description: 分页获取客户列表
  //@param: sysUserAuthorityID string, info request.PageInfo
  //@return: list interface{}, total int64, err error
  
  func (cbbs *XkBssService) LoadXkBbsPage(info request.PageInfo) (list interface{}, total int64, err error) {
  	limit := info.PageSize
  	offset := info.PageSize * (info.Page - 1)
  
  	db := global.GVA_DB.Model(&bbs.XkBbs{})
  
  	var XkBbsList []bbs.XkBbs
  	err = db.Count(&total).Error
  	if err != nil {
  		return XkBbsList, total, err
  	} else {
  		err = db.Limit(limit).Offset(offset).Find(&XkBbsList).Error
  	}
  	return XkBbsList, total, err
  }
  
  ```

  

  ### router

  在api目录去定义路由即可

  定义一个查询路由处理方法

  ```go
  package bbsapi
  
  import (
  	"github.com/flipped-aurora/gin-vue-admin/server/global"
  	"github.com/flipped-aurora/gin-vue-admin/server/model/bbs"
  	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
  	"github.com/gin-gonic/gin"
  	"go.uber.org/zap"
  	"strconv"
  )
  
  // 定义api接口
  type XkBbsApi struct{}
  
  // GetXkBbs
  //	@Tags		GetXkBbs
  //	@Summary	根据ID查询帖子明细
  //	@Security	ApiKeyAuth
  //	@accept		application/json
  //	@Produce	application/json
  //	@Param		data	query		bbs.GetXkBbs													true	"客户ID"
  //	@Success	200		{object}	response.Response{data=exampleRes.ExaCustomerResponse,msg=string}	"获取单一客户信息,返回包括客户详情"
  //	@Router		/bbs/get?id=1 [get]
  func (e *XkBbsApi) GetXkBbs(c *gin.Context) {
  	var xkBbs bbs.XkBbs
  	// 绑定参数
  	err := c.ShouldBindQuery(&xkBbs)
  	// 如果参数没有直接报错
  	if err != nil {
  		response.FailWithMessage(err.Error(), c)
  		return
  	}
  
  	data, err := bbsService.GetXkBbs(xkBbs.ID)
  	if err != nil {
  		global.GVA_LOG.Error("获取失败!", zap.Error(err))
  		response.FailWithMessage("获取失败", c)
  		return
  	}
  
  	response.OkWithDetailed(data, "获取成功", c)
  }
  
  func (e *XkBbsApi) GetXkBbsDetail(c *gin.Context) {
  	// 绑定参数用来获取/:id这个方式
  	id := c.Param("id")
  	// 这个是用来获取?age=123
  	//age := c.Query("age")
  	parseUint, _ := strconv.ParseUint(id, 10, 64)
  	data, err := bbsService.GetXkBbs(uint(parseUint))
  	if err != nil {
  		global.GVA_LOG.Error("获取失败!", zap.Error(err))
  		response.FailWithMessage("获取失败", c)
  		return
  	}
  	response.OkWithDetailed(data, "获取成功", c)
  }
  
  ```

  ```go
  package bbsapi
  
  import (
  	"github.com/flipped-aurora/gin-vue-admin/server/service"
  )
  
  type ApiGroup struct {
  	XkBbsApi
  }
  
  var (
  	bbsService = service.ServiceGroupApp.XkBbsServiceGroup.XkBbsService
  )
  
  ```

  

  给查询路由的方法，配置一个确切的路由地址。

  ```go
  package bbsrouter
  
  import (
  	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
  	"github.com/gin-gonic/gin"
  )
  
  type XkBbsRouter struct{}
  
  func (e *XkBbsRouter) InitXkBbsRouter(Router *gin.RouterGroup) {
  	xkBbsRouterWithoutRecord := Router.Group("bbs")
  	xkBbsApi := v1.ApiGroupApp.BbsApiGroup.XkBbsApi
  	{
  		//对外暴露的接口 http://localhost:8888/bbs/get?id=123
  		xkBbsRouterWithoutRecord.GET("get", xkBbsApi.GetXkBbs)                 // 获取单一客户信息
  		xkBbsRouterWithoutRecord.GET("getdetail/:id", xkBbsApi.GetXkBbsDetail) // 获取单一客户信息
  	}
  }
  
  ```

  

  ### 注册router

  在initilize中定义路由

  ```go
  package initialize
  
  import (
  	bbsrouter "github.com/flipped-aurora/gin-vue-admin/server/router/bbs"
  	"net/http"
  
  	"github.com/flipped-aurora/gin-vue-admin/server/docs"
  	"github.com/flipped-aurora/gin-vue-admin/server/global"
  	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
  	"github.com/flipped-aurora/gin-vue-admin/server/router"
  	"github.com/gin-gonic/gin"
  	ginSwagger "github.com/swaggo/gin-swagger"
  	"github.com/swaggo/gin-swagger/swaggerFiles"
  )
  
  // 初始化总路由
  
  func Routers() *gin.Engine {
  	Router := gin.Default()
  	InstallPlugin(Router) // 安装插件
  
  	xkBbsRouter := new(bbsrouter.XkBbsRouter)
  
  	// 方便统一添加路由组前缀 多服务器上线使用
  	PublicGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
  	{
  		// 健康监测
  		PublicGroup.GET("/health", func(c *gin.Context) {
  			c.JSON(http.StatusOK, "ok")
  		})
  	}
  	{
  		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
  		systemRouter.InitInitRouter(PublicGroup) // 自动初始化相关
  	}
  	{
  		xkBbsRouter.InitXkBbsRouter(PublicGroup)
  	}
  
  	
  	global.GVA_LOG.Info("router register success")
  	return Router
  }
  
  ```

  启动main函数，

  访问 http://localhost:8888/bbs/get?id=1

  访问 http://localhost:8888/bbs/getdetail/123?age=123

  

  

  

  - swagger生成路由接口
  - 前端开始调用接口
  - 开始编写对应菜单
  - 开始定义模块对应 spa页面或者组件
  - 然后完成curd操作

  

  





# GVA关于Model 和json tag和form tag的关系

```
// ID 只是给你go程序去使用和调用---服务端的东西，那为什么要大写，因为是私有变量，大写是公开

// 请求入参---映射到model属性--form
// 就是因为model字段都是大写，就造成开发时候传递参数也必须是大小，就烦躁,如何解决这个问题，使用form tag

// 响应返回---model---json
// 因为返回时候，如果使用model的大写属性返回，会很奇怪的，所以你就必须使用json的tag来指定的返回的名字
```





# GVA开发流程

- 写接口
  - model
    - xk_bbs
  - service
    - 对这个表进行增删改查
    - 保存
      - CreateModelName
      - SaveModelName
    - 修改
      - UpdateModelName
      - EditModelName
      - ModifyModelName
    - 删除
      - DeleteModelName
      - BatchDeleteModelName
    - 查询
      - FindModelNames
      - SearchModelNames
      - QueryModelNames
      - GetModelName(获取当个)
    - 其它
      - countModelName
  - router
  - 暴露接口swagger/apiforx
  - 调用接口







## 添加的思考

- model

- service

  ```go
  // @author: feige
  // @function: CreateXkBbs
  // @description: 创建文章
  // @param: e bbs.XkBbs
  // @return: err error
  func (cbbs *XkBbsService) CreateXkBbs(xkBbs *bbs.XkBbs) (err error) {
  	// 1： 获取数据的连接对象
  	err = global.GVA_DB.Create(xkBbs).Error
  	return err
  }
  ```

  目标是什么：把帖子数据数据保存到帖子数据库表中。

  保存必须载体：结构体

  ```go
  type XkBbs struct {
  	global.GVA_MODEL
  	Title        string `json:"title" gorm:"not null;index;comment:标题"`
  	Img          string `json:"img"  gorm:"not null;default:'';comment:封面图"`
  	Description  string `json:"description" gorm:"not null;default:'';comment:描述"`
  	Content      string `json:"content" gorm:"not null;default:'';comment:文章内容--MD格式"`
  	HtmlContent  string `json:"htmlContent" gorm:"not null;default:'';comment:文章内容--MD格式"`
  	CategoryId   uint   `json:"categoryId" gorm:"not null;default:0;comment:文章分类ID"`
  	CategoryName string `json:"categoryName" gorm:"not null;default:'';comment:文章分类名称"`
  	ViewCount    int8   `json:"viewCount" gorm:"not null;default:0;comment:文章阅读数"`
  	Comments     int8   `json:"comments" gorm:"not null;default:0;comment:评论数"`
  	CommentsOpen int8   `json:"commentsOpen" gorm:"not null;default:1;comment:是否允许评论 0 不允许  1 允许"`
  	Status       int8   `json:"status" gorm:"not null;default:1;comment:0 未发布 1 发布"`
  	IsDelete     int8   `json:"isDelete" gorm:"not null;default:0;comment:0 未删除 1 删除"`
  	UserId       uint   `json:"userId" gorm:"not null;comment:文章作者ID"`
  	Username     string `json:"username"  gorm:"not null;default:'';comment:文章发布者用户名"`
  	Avatar       string `json:"avatar"  gorm:"not null;default:'';comment:文章发布者头像"`
  }
  
  ```

  测试例子

  ```go
  func main(){
      xkBbs := new(bbs.XkBbs)
      xkBbs.Title = "我是一个帖子"
      xkBbs.CategoryId = 1
      
      // 创建XkbbsService
      xkBbsService := new(XkBbsService)
      
      xkBbsService.CreateXkBbs(xkBbs);
  ```

  如果执行CreateXkBbs方法的时候，核心代码在执行什么呢？

  ```go
  err = global.GVA_DB.Create(xkBbs).Error
  ```

  这个时候gorm框架发现你执行的是Create方法，就开始使用反射把xkBbs进行解析。然后开始读取到xkBbs结构体的属性gorm的tag如果没有就使用它属性名改成驼峰小写，开始组织insert语句。

  ```sh
  insert into xk_bbs(title,categoyr_id,xxxxxx)values(xkBbs.Title,xkBbs.categoryPid,xxxxxx);
  ```

  底层会使用事务进行sql语句执行，把数据提交到表中，完成最后的操作。然后释放资源、



## 完成添加的接口定义和测试

- 保存的请求方式一般都是使用：put或者post请求
- post如何测试
  - 你可以自己定义一个表单（这个就太麻烦了）不可取
  - postman/apifox/swagger
  - http文件
  - 命令



# 中间表的设置缘由

原来的a表和b表互相之间存在1对多的关系的时候，就可以用一个中间表来过度一下的哦。

# GVA实战帖子模块增删改查

## 1： 定义模型

```go
// 自动生成模板XkVideo
package video

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// XkVideo 结构体
type XkVideo struct {
      global.GVA_MODEL
      Title  string `json:"title" form:"title" gorm:"column:title;comment:标题;size:100;"`
      Cid  *int `json:"cid" form:"cid" gorm:"column:cid;comment:分类ID;"`
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName XkVideo 表名
func (XkVideo) TableName() string {
  return "xk_video"
}


```

## 2: 定义模型参数

```go
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/video"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type XkVideoSearch struct{
    video.XkVideo
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}

```

## 3: 自动构建生成表

```go
package initialize

import (
	"os"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"

	"github.com/flipped-aurora/gin-vue-admin/server/model/video"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	case "oracle":
		return GormOracle()
	case "mssql":
		return GormMssql()
	default:
		return GormMysql()
	}
}

func RegisterTables() {
	db := global.GVA_DB
	err := db.AutoMigrate(

		system.SysApi{},
		system.SysUser{},
		system.SysBaseMenu{},
		system.JwtBlacklist{},
		system.SysAuthority{},
		system.SysDictionary{},
		system.SysOperationRecord{},
		system.SysAutoCodeHistory{},
		system.SysDictionaryDetail{},
		system.SysBaseMenuParameter{},
		system.SysBaseMenuBtn{},
		system.SysAuthorityBtn{},
		system.SysAutoCode{},
		system.SysChatGptOption{},

		example.ExaFile{},
		example.ExaCustomer{},
		example.ExaFileChunk{}, 
		// 自动生成表
		video.XkVideo{},
	)
	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}

```

## 4:  定义service

```go
package video

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/video"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    videoReq "github.com/flipped-aurora/gin-vue-admin/server/model/video/request"
    "gorm.io/gorm"
)

type XkVideoService struct {
}

// CreateXkVideo 创建XkVideo记录
// Author [piexlmax](https://github.com/piexlmax)
func (xkVideoService *XkVideoService) CreateXkVideo(xkVideo *video.XkVideo) (err error) {
	err = global.GVA_DB.Create(xkVideo).Error
	return err
}

// DeleteXkVideo 删除XkVideo记录
// Author [piexlmax](https://github.com/piexlmax)
func (xkVideoService *XkVideoService)DeleteXkVideo(xkVideo video.XkVideo) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&video.XkVideo{}).Where("id = ?", xkVideo.ID).Update("deleted_by", xkVideo.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&xkVideo).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteXkVideoByIds 批量删除XkVideo记录
// Author [piexlmax](https://github.com/piexlmax)
func (xkVideoService *XkVideoService)DeleteXkVideoByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&video.XkVideo{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&video.XkVideo{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateXkVideo 更新XkVideo记录
// Author [piexlmax](https://github.com/piexlmax)
func (xkVideoService *XkVideoService)UpdateXkVideo(xkVideo video.XkVideo) (err error) {
	err = global.GVA_DB.Save(&xkVideo).Error
	return err
}

// GetXkVideo 根据id获取XkVideo记录
// Author [piexlmax](https://github.com/piexlmax)
func (xkVideoService *XkVideoService)GetXkVideo(id uint) (xkVideo video.XkVideo, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&xkVideo).Error
	return
}

// GetXkVideoInfoList 分页获取XkVideo记录
// Author [piexlmax](https://github.com/piexlmax)
func (xkVideoService *XkVideoService)GetXkVideoInfoList(info videoReq.XkVideoSearch) (list []video.XkVideo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&video.XkVideo{})
    var xkVideos []video.XkVideo
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["title"] = true
         	orderMap["cid"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	err = db.Limit(limit).Offset(offset).Find(&xkVideos).Error
	return  xkVideos, total, err
}

```

```go
package video

type ServiceGroup struct {
	XkVideoService
}
```

聚合Service

```go
package service

import (
	bbsservice "github.com/flipped-aurora/gin-vue-admin/server/service/bbs"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/video"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	XkBbsServiceGroup   bbsservice.ServiceGroup
	VideoServiceGroup   video.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)

```



## 5、路由定义

```go
package video

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type XkVideoRouter struct {
}

// InitXkVideoRouter 初始化 XkVideo 路由信息
func (s *XkVideoRouter) InitXkVideoRouter(Router *gin.RouterGroup) {
	xkVideoRouter := Router.Group("xkVideo").Use(middleware.OperationRecord())
	xkVideoRouterWithoutRecord := Router.Group("xkVideo")
	var xkVideoApi = v1.ApiGroupApp.VideoApiGroup.XkVideoApi
	{
		xkVideoRouter.POST("createXkVideo", xkVideoApi.CreateXkVideo)   // 新建XkVideo
		xkVideoRouter.DELETE("deleteXkVideo", xkVideoApi.DeleteXkVideo) // 删除XkVideo
		xkVideoRouter.DELETE("deleteXkVideoByIds", xkVideoApi.DeleteXkVideoByIds) // 批量删除XkVideo
		xkVideoRouter.PUT("updateXkVideo", xkVideoApi.UpdateXkVideo)    // 更新XkVideo
	}
	{
		xkVideoRouterWithoutRecord.GET("findXkVideo", xkVideoApi.FindXkVideo)        // 根据ID获取XkVideo
		xkVideoRouterWithoutRecord.GET("getXkVideoList", xkVideoApi.GetXkVideoList)  // 获取XkVideo列表
	}
}

```

路由管理

```go
package video

type RouterGroup struct {
	XkVideoRouter
}

```

路由

```go
package router

import (
	bbsrouter "github.com/flipped-aurora/gin-vue-admin/server/router/bbs"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/video"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	XkBbs   bbsrouter.RouterGroup
	Video   video.RouterGroup
}

var RouterGroupApp = new(RouterGroup)

```

## 6 ： 路由初始化

```go
package initialize

import (
	"net/http"

	"github.com/flipped-aurora/gin-vue-admin/server/docs"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.Default()
	InstallPlugin(Router)	// 安装插件

	systemRouter := router.RouterGroupApp.System
	exampleRouter := router.RouterGroupApp.Example
	xkBbsRouter := router.RouterGroupApp.XkBbs

	// 如果想要不使用nginx代理前端网页，可以修改 web/.env.production 下的
	// VUE_APP_BASE_API = /
	// VUE_APP_BASE_PATH = http://localhost
	// 然后执行打包命令 npm run build。在打开下面4行注释
	// Router.LoadHTMLGlob("./dist/*.html") // npm打包成dist的路径
	// Router.Static("/favicon.ico", "./dist/favicon.ico")
	// Router.Static("/static", "./dist/assets")   // dist里面的静态资源
	// Router.StaticFile("/", "./dist/index.html") // 前端网页入口页面

	Router.StaticFS(global.GVA_CONFIG.Local.StorePath, http.Dir(global.GVA_CONFIG.Local.StorePath))	// 为用户头像和文件提供静态地址
	// Router.Use(middleware.LoadTls())  // 如果需要使用https 请打开此中间件 然后前往 core/server.go 将启动模式 更变为 Router.RunTLS("端口","你的cre/pem文件","你的key文件")
	// 跨域，如需跨域可以打开下面的注释
	// Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	// Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	//global.GVA_LOG.Info("use middleware cors")
	docs.SwaggerInfo.BasePath = global.GVA_CONFIG.System.RouterPrefix
	Router.GET(global.GVA_CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用

	PublicGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup)	// 注册基础功能路由 不做鉴权
		systemRouter.InitInitRouter(PublicGroup)	// 自动初始化相关
	}
	{
		xkBbsRouter.InitXkBbsRouter(PublicGroup)
	}

	PrivateGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		systemRouter.InitApiRouter(PrivateGroup)			// 注册功能api路由
		systemRouter.InitJwtRouter(PrivateGroup)			// jwt相关路由
		systemRouter.InitUserRouter(PrivateGroup)			// 注册用户路由
		systemRouter.InitMenuRouter(PrivateGroup)			// 注册menu路由
		systemRouter.InitSystemRouter(PrivateGroup)			// system相关路由
		systemRouter.InitCasbinRouter(PrivateGroup)			// 权限相关路由
		systemRouter.InitAutoCodeRouter(PrivateGroup)			// 创建自动化代码
		systemRouter.InitAuthorityRouter(PrivateGroup)			// 注册角色路由
		systemRouter.InitSysDictionaryRouter(PrivateGroup)		// 字典管理
		systemRouter.InitAutoCodeHistoryRouter(PrivateGroup)		// 自动化代码历史
		systemRouter.InitSysOperationRecordRouter(PrivateGroup)		// 操作记录
		systemRouter.InitSysDictionaryDetailRouter(PrivateGroup)	// 字典详情管理
		systemRouter.InitAuthorityBtnRouterRouter(PrivateGroup)		// 字典详情管理
		systemRouter.InitChatGptRouter(PrivateGroup)			// chatGpt接口
		exampleRouter.InitCustomerRouter(PrivateGroup)			// 客户路由
		exampleRouter.InitFileUploadAndDownloadRouter(PrivateGroup)	// 文件上传下载功能路由
	}
	{
		videoRouter := router.RouterGroupApp.Video
		videoRouter.InitXkVideoRouter(PrivateGroup)
	}
	global.GVA_LOG.Info("router register success")
	return Router
}

```

## 7：swagger暴露接口调用

### 安装 swagger[#](https://www.gin-vue-admin.com/guide/start-quickly/swagger.html#_1-安装-swagger)

- 可以翻墙



```
go install github.com/swaggo/swag/cmd/swag@version
```

- 无法翻墙 由于国内没法安装 [go.org/x](http://go.org/x) 包下面的东西，推荐使用 [goproxy.cn](https://goproxy.cn/zh/) 或者 [goproxy.cn/](https://goproxy.cn/)



```
# Go 版本 是 1.16 ~ 最新版 可以忽略以下步骤一
# 步骤一、启用 Go Modules 功能
go env -w GO111MODULE=on 
# 步骤二、配置 GOPROXY 环境变量
go env -w GOPROXY=https://goproxy.cn,direct

# 使用如下命令下载swag
go install github.com/swaggo/swag/cmd/swag@version
```

###  生成API文档[#](https://www.gin-vue-admin.com/guide/start-quickly/swagger.html#_2-生成api文档)

```
cd server
swag init
```

执行上面的命令后，server目录下会出现docs文件夹，打开浏览器复制如下地址，即可查看swagger文档

```
http://localhost:8888/swagger/index.html
```

### 如果遇到安装报错可以执行下面的命令

```sh
## 如果不行，获取更新包：
go get github.com/swaggo/echo-swagger

##以上还是不行，在更新
go get github.com/alecthomas/template
```

