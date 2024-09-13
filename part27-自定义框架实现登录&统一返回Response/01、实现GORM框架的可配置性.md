# 实现GORM框架的可配置性



## 01、可配置的数据库处理

1： 在yaml配置数据库信息

```yaml
# 数据库配置
# "root:mkxiaoer@tcp(127.0.0.1:3306)/ksd-social-db?charset=utf8&parseTime=True&loc=Local", // DSN data source name
database:
  mysql:
    host: 127.0.0.1
    port: 3306
    dbname: ksd-social-db
    username: root
    password: mkxiaoer
    config: charset=utf8&parseTime=True&loc=Local


```

2: 定义yaml的解析结构体

```go
package parse

//database:--------------------------------struct
//	mysql: --------------------------------struct
//		host: 127.0.0.1 -------------------field
//		port: 3306 -------------------field
//		dbname: ksd-social-db -------------------field
//		username: root -------------------field
//		password: mkxiaoer -------------------field
//		config: charset=utf8&parseTime=True&loc=Local -------------------field

type Database struct {
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}

type Mysql struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Dbname   string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Config   string `mapstructure:"config" json:"config" yaml:"config"`
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}

```

3: 在Config.go进行定义

```go
package parse

// 配置接下入口
type Config struct {
	Database Database `mapstructure:"database" json:"database" yaml:"database"`
}

```

因为commons包下的InitViper() 方法是把Config进行映射和绑定的。也就是说commons.Config其实就是整个YAML文件本身。里面定义属性（大部分都是结构体），其实也代表是某些业务配置的第一级、

## 02、使用自己的框架来实现登录接口 

1:  定义一个登录路由处理方法

```go
package login

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 登录业务
type LoginApi struct{}

// 登录的接口处理
func (api *LoginApi) ToLogined(c *gin.Context) {
	c.JSON(http.StatusOK, "我是一个登录")
}

```

2: 定义具体的登录路由

```go
package login

import (
	"github.com/gin-gonic/gin"
	"xkginweb/api/v1/login"
)

// 登录路由
type LoginRouter struct{}

func (router *LoginRouter) InitLoginRouter(Router *gin.Engine) {
	loginApi := login.LoginApi{}
	// 单个定义
	//Router.GET("/login/toLogin", loginApi.ToLogined)
	//Router.GET("/login/toReg", loginApi.ToLogined)
	//Router.GET("/login/forget", loginApi.ToLogined)
	// 用组定义--（推荐）
	loginRouter := Router.Group("/login")
	{
		loginRouter.GET("/toLogin", loginApi.ToLogined)
		loginRouter.GET("/toReg", loginApi.ToLogined)
		loginRouter.GET("/forget", loginApi.ToLogined)
	}
}

```

3: 注册路由

```go
package initilization

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"xkginweb/global"
	"xkginweb/router"
	"xkginweb/router/login"
)

func InitGinRouter() *gin.Engine {
	// 创建gin服务
	ginServer := gin.Default()
	// 提供服务组
	courseRouter := router.RouterWebGroupApp.Course.CourseRouter
	videoRouter := router.RouterWebGroupApp.Video.VideoRouter
	loginRouter := login.LoginRouter{}
	// 接口隔离，比如登录，健康检查都不需要拦截和做任何的处理
	loginRouter.InitLoginRouter(ginServer)
	// 业务模块接口，
	publicGroup := ginServer.Group("/api")
	{
		videoRouter.InitVideoRouter(publicGroup)
		courseRouter.InitCourseRouter(publicGroup)
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
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)

	s2 := s.ListenAndServe().Error()
	fmt.Println("服务启动完毕 ", s2)
}

```

4: 访问 http://localhost:9899/login/toLogin



## 03、快速的整合gorm和db实现登录

1: 准备一个用户表

```go
package user

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID           uint `gorm:"primaryKey"`
	Name         string
	Email        *string
	Age          uint8
	Birthday     time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time      `gorm:"autoUpdateTime"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

// 覆盖生成表
func (User) TableName() string {
	return "xk_test_user"
}

```

2: 定义service

```go
package user

import (
	"xkginweb/global"
	"xkginweb/model/user"
)

// 对用户表的数据层处理
type UserService struct{}

// nil 是go空值处理，必须是指针类型
func (service *UserService) getUserByAccount(account string) (user *user.User, err error) {
	// 根据account进行查询
	err = global.KSD_DB.Where("account = ?", account).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

```

3: 其实就调用service方法

```go
package login

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xkginweb/service/user"
)

// 登录业务
type LoginApi struct{}

// 登录的接口处理
func (api *LoginApi) ToLogined(c *gin.Context) {

	// 1：获取用户在页面上输入的账号和密码开始和数据库里数据进行校验
	userService := user.UserService{}

	// 模拟用户前端输入的账号和密码
	inputAccount := "feige1"
	inputPassword := "123456"

	dbUser, err := userService.GetUserByAccount(inputAccount)
	if err != nil {
		c.JSON(http.StatusOK, "你输入的账号和密码有误!!!")
		return
	}

	// 这个时候就判断用户输入密码和数据库的密码是否一致
	if dbUser != nil && dbUser.Password == inputPassword {
		c.JSON(http.StatusOK, dbUser)
	} else {
		c.JSON(http.StatusOK, "你输入的账号和密码有误!!!")
	}
}

```

开始模拟用户是否登录成功。

- 正常的的情况返回的，status=200 ,ok。 说明没有问题，返回整个用户对象
- 如果输入错误账号和密码，那就直接提示：你输入的账号和密码有误!!!"  

你发现一个问题，这个就无论你返回的正确的信息还是错误的信息。你都必须是：http.StatusOK  。那么这个状态到底是什么含义呢？能u自定义呢？答案是不能。但是这个状态有什么用。可以看到整个请求和响应的所以的执行过程，我们可以通过http.Status这个状态可以分析请求和响应的错误到底是什么？



## 04、关于http.Status

所有的web框架，为了监听请求和响应整个执行过程，定义了一系列的状态。然后给开发者使用。不同状态，浏览器进行处理的结果和响应方式可能都会不一样。

### 1、 HTTP Status Code 1xx 请求信息

这一组[状态码](https://so.csdn.net/so/search?q=状态码&spm=1001.2101.3001.7020)表明这是一个临时性响应。此响应仅由状态行和可选的HTTP头组成，以一个空行结尾。由于HTTP／1.0未定义任何1xx状态码，所以不要向HTTP／1.0客户端发送1xx响应。

| Http状态码    | Http Status Code                                             | Http状态码含义中文说明 |
| :------------ | :----------------------------------------------------------- | :--------------------- |
| ***\*100\**** | [100 Continue](https://seo.juziseo.com/doc/http_code/100)    | 请继续请求             |
| ***\*101\**** | [101 Switching Protocols](https://seo.juziseo.com/doc/http_code/101) | 请切换协议             |
| ***\*102\**** | [102 Processing](https://seo.juziseo.com/doc/http_code/102)  | 将继续执行请求         |



### 2、 HTTP Status Code 2xx 成功状态

这一组状态码表明客户端的请求已经被服务器端成功接收并正确解析。

| Http状态码    | Http Status Code                                             | Http状态码含义中文说明                  |
| :------------ | :----------------------------------------------------------- | :-------------------------------------- |
| ***\*200\**** | [200 OK](https://seo.juziseo.com/doc/http_code/200)          | 请求成功                                |
| ***\*201\**** | [201 Created](https://seo.juziseo.com/doc/http_code/201)     | 请求已被接受，等待资源响应              |
| ***\*202\**** | [202 Accepted](https://seo.juziseo.com/doc/http_code/202)    | 请求已被接受，但尚未处理                |
| ***\*203\**** | [203 Non-Authoritative Information](https://seo.juziseo.com/doc/http_code/203) | 请求已成功处理，结果来自第三方拷贝      |
| ***\*204\**** | [204 No Content](https://seo.juziseo.com/doc/http_code/204)  | 请求已成功处理，但无返回内容            |
| ***\*205\**** | [205 Reset Content](https://seo.juziseo.com/doc/http_code/205) | 请求已成功处理，但需重置内容            |
| ***\*206\**** | [206 Partial Content](https://seo.juziseo.com/doc/http_code/206) | 请求已成功处理，但仅返回了部分内容      |
| ***\*207\**** | [207 Multi-Status](https://seo.juziseo.com/doc/http_code/207) | 请求已成功处理，返回了多个状态的XML消息 |
| ***\*208\**** | [208 Already Reported](https://seo.juziseo.com/doc/http_code/208) | 响应已发送                              |
| ***\*226\**** | [226 IM Used](https://seo.juziseo.com/doc/http_code/226)     | 已完成响应                              |

### 3、 HTTP Status Code 3xx 重定向状态

这一组状态码表示客户端需要采取更进一步的行动来完成请求。通常，这些状态码用来重定向，后续的请求地址（重定向目标）在本次响应的Location域中指明。

| Http状态码    | Http Status Code                                             | Http状态码含义中文说明         |
| :------------ | :----------------------------------------------------------- | :----------------------------- |
| ***\*300\**** | [300 Multiple Choices](https://seo.juziseo.com/doc/http_code/300) | 返回多条重定向供选择           |
| ***\*301\**** | [301 Moved Permanently](https://seo.juziseo.com/doc/http_code/301) | 永久重定向                     |
| ***\*302\**** | [302 Found](https://seo.juziseo.com/doc/http_code/302)       | 临时重定向                     |
| ***\*303\**** | [303 See Other](https://seo.juziseo.com/doc/http_code/303)   | 当前请求的资源在其它地址       |
| ***\*304\**** | [304 Not Modified](https://seo.juziseo.com/doc/http_code/304) | 请求资源与本地缓存相同，未修改 |
| ***\*305\**** | [305 Use Proxy](https://seo.juziseo.com/doc/http_code/305)   | 必须通过代理访问               |
| ***\*306\**** | [306 (已废弃)Switch Proxy](https://seo.juziseo.com/doc/http_code/306) | (已废弃)请切换代理             |
| ***\*307\**** | [307 Temporary Redirect](https://seo.juziseo.com/doc/http_code/307) | 临时重定向，同302              |
| ***\*308\**** | [308 Permanent Redirect](https://seo.juziseo.com/doc/http_code/308) | 永久重定向，且禁止改变http方法 |

### 4、 HTTP Status Code 4xx 客户端错误

这一组状态码表示客户端的请求存在错误，导致服务器无法处理。除非响应的是一个HEAD请求，否则服务器就应该返回一个解释当前错误状况的实体，以及这是临时的还是永久性的状况。这些状态码适用于任何请求方法。浏览器应当向用户显示任何包含在此类错误响应中的实体内容。

| Http状态码    | Http Status Code                                             | Http状态码含义中文说明               |
| :------------ | :----------------------------------------------------------- | :----------------------------------- |
| ***\*400\**** | [400 Bad Request](https://seo.juziseo.com/doc/http_code/400) | 请求错误，通常是访问的域名未绑定引起 |
| ***\*401\**** | [401 Unauthorized](https://seo.juziseo.com/doc/http_code/401) | 需要身份认证验证                     |
| ***\*402\**** | [402 Payment Required](https://seo.juziseo.com/doc/http_code/402) | -                                    |
| ***\*403\**** | [403 Forbidden](https://seo.juziseo.com/doc/http_code/403)   | 禁止访问                             |
| ***\*404\**** | [404 Not Found](https://seo.juziseo.com/doc/http_code/404)   | 请求的内容未找到或已删除             |
| ***\*405\**** | [405 Method Not Allowed](https://seo.juziseo.com/doc/http_code/405) | 不允许的请求方法                     |
| ***\*406\**** | [406 Not Acceptable](https://seo.juziseo.com/doc/http_code/406) | 无法响应，因资源无法满足客户端条件   |
| ***\*407\**** | [407 Proxy Authentication Required](https://seo.juziseo.com/doc/http_code/407) | 要求通过代理的身份认证               |
| ***\*408\**** | [408 Request Timeout](https://seo.juziseo.com/doc/http_code/408) | 请求超时                             |
| ***\*409\**** | [409 Conflict](https://seo.juziseo.com/doc/http_code/409)    | 存在冲突                             |
| ***\*410\**** | [410 Gone](https://seo.juziseo.com/doc/http_code/410)        | 资源已经不存在(过去存在)             |
| ***\*411\**** | [411 Length Required](https://seo.juziseo.com/doc/http_code/411) | 无法处理该请求                       |
| ***\*412\**** | [412 Precondition Failed](https://seo.juziseo.com/doc/http_code/412) | 请求条件错误                         |
| ***\*413\**** | [413 Payload Too Large](https://seo.juziseo.com/doc/http_code/413) | 请求的实体过大                       |
| ***\*414\**** | [414 Request-URI Too Long](https://seo.juziseo.com/doc/http_code/414) | 请求的URI过长                        |
| ***\*415\**** | [415 Unsupported Media Type](https://seo.juziseo.com/doc/http_code/415) | 无法处理的媒体格式                   |
| ***\*416\**** | [416 Range Not Satisfiable](https://seo.juziseo.com/doc/http_code/416) | 请求的范围无效                       |
| ***\*417\**** | [417 Expectation Failed](https://seo.juziseo.com/doc/http_code/417) | 无法满足的Expect                     |
| ***\*418\**** | [418 I'm a teapot](https://seo.juziseo.com/doc/http_code/418) | 愚人节笑话                           |
| ***\*421\**** | [421 There are too many connections from your internet address](https://seo.juziseo.com/doc/http_code/421) | 连接数超限                           |
| ***\*422\**** | [422 Unprocessable Entity](https://seo.juziseo.com/doc/http_code/422) | 请求的语义错误                       |
| ***\*423\**** | [423 Locked](https://seo.juziseo.com/doc/http_code/423)      | 当前资源被锁定                       |
| ***\*424\**** | [424 Failed Dependency](https://seo.juziseo.com/doc/http_code/424) | 当前请求失败                         |
| ***\*425\**** | [425 Unordered Collection](https://seo.juziseo.com/doc/http_code/425) | 未知                                 |
| ***\*426\**** | [426 Upgrade Required](https://seo.juziseo.com/doc/http_code/426) | 请切换到TLS/1.0                      |
| ***\*428\**** | [428 Precondition Required](https://seo.juziseo.com/doc/http_code/428) | 请求未带条件                         |
| ***\*429\**** | [429 Too Many Requests](https://seo.juziseo.com/doc/http_code/429) | 并发请求过多                         |
| ***\*431\**** | [431 Request Header Fields Too Large](https://seo.juziseo.com/doc/http_code/431) | 请求头过大                           |
| ***\*449\**** | [449 Retry With](https://seo.juziseo.com/doc/http_code/449)  | 请重试                               |
| ***\*451\**** | [451 Unavailable For Legal Reasons](https://seo.juziseo.com/doc/http_code/451) | 访问被拒绝（法律的要求）             |
| ***\*499\**** | [499 Client Closed Request](https://seo.juziseo.com/doc/http_code/499) | 客户端主动关闭了连接                 |

### 5、 HTTP Status Code 5xx 服务器错误状态

这一组状态码说明服务器在处理请求的过程中有错误或者异常状态发生，也有可能是服务器意识到以当前的软硬件资源无法完成对请求的处理。除非这是一个HEAD请求，否则服务器应当包含一个解释当前错误状态以及这个状况是临时的还是永久的解释信息实体。浏览器应当向用户展示任何在当前响应中被包含的实体。

| Http状态码    | Http Status Code                                             | Http状态码含义中文说明   |
| :------------ | :----------------------------------------------------------- | :----------------------- |
| ***\*500\**** | [500 Internal Server Error](https://seo.juziseo.com/doc/http_code/500) | 服务器端程序错误         |
| ***\*501\**** | [501 Not Implemented](https://seo.juziseo.com/doc/http_code/501) | 服务器不支持的请求方法   |
| ***\*502\**** | [502 Bad Gateway](https://seo.juziseo.com/doc/http_code/502) | 网关无响应               |
| ***\*503\**** | [503 Service Unavailable](https://seo.juziseo.com/doc/http_code/503) | 服务器端临时错误         |
| ***\*504\**** | [504 Gateway Timeout](https://seo.juziseo.com/doc/http_code/504) | 网关超时                 |
| ***\*505\**** | [505 HTTP Version Not Supported](https://seo.juziseo.com/doc/http_code/505) | 服务器不支持的HTTP版本   |
| ***\*506\**** | [506 Variant Also Negotiates](https://seo.juziseo.com/doc/http_code/506) | 服务器内部配置错误       |
| ***\*507\**** | [507 Insufficient Storage](https://seo.juziseo.com/doc/http_code/507) | 服务器无法存储请求       |
| ***\*508\**** | [508 Loop Detected](https://seo.juziseo.com/doc/http_code/508) | 服务器因死循环而终止操作 |
| ***\*509\**** | [509 Bandwidth Limit Exceeded](https://seo.juziseo.com/doc/http_code/509) | 服务器带宽限制           |
| ***\*510\**** | [510 Not Extended](https://seo.juziseo.com/doc/http_code/510) | 获取资源策略未被满足     |
| ***\*511\**** | [511 Network Authentication Required](https://seo.juziseo.com/doc/http_code/511) | 需验证以许可连接         |
| ***\*599\**** | [599 Network Connect Timeout Error](https://seo.juziseo.com/doc/http_code/599) | 网络连接超时             |

上面的这些状态全部都是描述请求和响应的整个过程的状态。不包含业务的状态。我举例例子

```go
c.JSON(http.StatusOK, "你输入的账号和密码有误!!!")
```

在开发中一个接口：成功只有一种情况，但是失败和错误就有N种情况。那么这些N情况的返回到底选择什么样子状态就变得非常的重要。你接下来就困难选择症（你不知道到底选择什么？），而且你在这些状态找不到适合的，所以业务的状态处理不应该选择web框架中提供的。因为这些状态根本就不是让你来做业务错误的状态监控，别人专门去web请求和响应的状态的监听。

那么怎么处理。通过自己的方式来定义状态。但是所有的返回都用：http.StatusOK

- 只要请求和响应是正常的，无论正确和错误，我们都用http.StatusOK来返回，但是区分和界定用自己定义的状态来定义业务、

这也就是为什么我们要做自定义返回的意义和价值了.

## 05、开始学会使用map来封装返回

```go
package login

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xkginweb/service/user"
)

// 登录业务
type LoginApi struct{}

// 登录的接口处理
func (api *LoginApi) ToLogined(c *gin.Context) {

	// 1：获取用户在页面上输入的账号和密码开始和数据库里数据进行校验
	userService := user.UserService{}

	m := map[string]any{}

	// 模拟用户前端输入的账号和密码
	inputAccount := "feige1"
	inputPassword := "123456"

	if len(inputAccount) == 0 {
		m["code"] = 60002
		m["msg"] = "请输入账号!!!"
		c.JSON(http.StatusOK, m)
		return
	}

	if len(inputPassword) == 0 {
		m["code"] = 60002
		m["msg"] = "请输入密码!!!"
		c.JSON(http.StatusOK, "请输入密码!!!")
		return
	}

	dbUser, err := userService.GetUserByAccount(inputAccount)
	if err != nil {
		m["code"] = 60001
		m["msg"] = "你输入的账号和密码有误!!!"
		c.JSON(http.StatusOK, m)
		return
	}
	// 这个时候就判断用户输入密码和数据库的密码是否一致
	if dbUser != nil && dbUser.Password == inputPassword {
		m["code"] = 20000
		m["msg"] = dbUser
		c.JSON(http.StatusOK, dbUser)
	} else {
		m["code"] = 60001
		m["msg"] = "你输入的账号和密码有误"
		c.JSON(http.StatusOK, "你输入的账号和密码有误!!!")
	}
}

```

但是map有一个非常麻烦的事情，就是key不具备面向对象。写起来很郁闷和不方便。

```go
package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// R Response ServerResult Result

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Result(code int, msg string, data any, c *gin.Context) {
	c.JSON(http.StatusOK, Response{code, msg, data})
}

var (
	CODE = 20000
	MSG  = "success"
)

/*
*
成功
*/
func Ok(c *gin.Context) {
	Result(CODE, "操作成功", map[string]interface{}{}, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(CODE, MSG, data, c)
}

/*
*
失败
*/
func Fail(code int, msg string, c *gin.Context) {
	Result(code, msg, map[string]any{}, c)
}

func FailWithData(code int, msg string, data any, c *gin.Context) {
	Result(code, msg, data, c)
}

```

## 06：如果返回的值出现的属性是大写的，说明没有配置json

```go
package user

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `gorm:"primaryKey"json:"id"`
	Name      string         `json:"name"`
	Account   string         `gorm:"unique"json:"account"`
	Password  string         `json:"password"`
	Email     *string        `json:"email"`
	Age       uint8          `json:"age"`
	Birthday  time.Time      `json:"birthday"`
	CreatedAt time.Time      `gorm:"autoUpdateTime"json:"createAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// 覆盖生成表
func (User) TableName() string {
	return "xk_admin_user"
}

```

返回的时候就变成小写，如下

```json
{
  "code": 20000,
  "msg": "success",
  "data": {
    "id": 1,
    "name": "飞哥",
    "account": "feige1",
    "password": "123456",
    "email": "123456@qq.com",
    "age": 12,
    "birthday": "2023-07-13T21:08:58+08:00",
    "createAt": "2023-07-13T21:09:05+08:00",
    "updatedAt": "2023-07-13T21:09:05+08:00",
    "DeletedAt": null
  }
}
```



## 07、接口的合理性和参数的获取和调试

1: 登录属于form提交范畴的，一般的处理方案：post请求。

```GO
package login

import (
	"github.com/gin-gonic/gin"
	"xkginweb/api/v1/login"
)

// 登录路由
type LoginRouter struct{}

func (router *LoginRouter) InitLoginRouter(Router *gin.Engine) {
	loginApi := login.LoginApi{}
	// 单个定义
	//Router.GET("/login/toLogin", loginApi.ToLogined)
	//Router.GET("/login/toReg", loginApi.ToLogined)
	//Router.GET("/login/forget", loginApi.ToLogined)
	// 用组定义--（推荐）
	loginRouter := Router.Group("/login")
	{
		loginRouter.POST("/toLogin", loginApi.ToLogined)
	}
}

```

2: 参数动态化

```go
```

一般如果你要模拟参数的动态化，要么使用：test.http文件，或者使用swagger，或者apiforx或者程序代码。显然程序代码我们还没搭建框架，自然就不能使用。我们可以考虑apifox



## 08、搭建一个自己vite的项目把登录整合进来



## 09、集成验证码功能





## 10、实现登录和验证码的校验



## 11、状态管理piana



## 12、对接jwt



## 13、动态路由gva的动态菜单



## 14、权限和角色的处理



## 15、业务开发 （章节，帖子，评论，订单）



## 16、外部配置 Nacos，日志

