package course

import (
	"github.com/gin-gonic/gin"
	"xkginweb/commons/response"
)

type CourseController struct{}

// 获取明细
func (courseController *CourseController) GetByID(c *gin.Context) {
	// 绑定参数用来获取/:id这个方式
	//id := c.Param("id")
	// 绑定参数 ?ids=1111

	response.Ok("success", c)
	return
}

// 查询video
func (courseController *CourseController) FindVideos(c *gin.Context) {
	// service new
	// model new

	response.Ok("success", c)
	return
}

/*

下面是中间件的一种原生写法，因为每个接口都要使用下面的这种方式的话，那么重复的代码是很多的，所以我们使用了一个方法封装成中间件了，统一在初始化路由的时候加上jwt中间件处理

// 查询video
func (courseController *CourseController) FindVideos(c *gin.Context) {
	// service new
	// model new

	// 获取参数中的 token
	//token := c.Param("token") // error
	token, flag := c.GetQuery("token") // 是用的 GetQuery() 这个方法的哦
	if flag == false || token == "" {
		response.Fail(60002, "token 参数获取失败", c)
		return
	}

	// 1. 初始化 jwt 对象
	myJwt := jwtgo.NewJWT()
	// 2. myJwt 解析上面的 token 是否正确
	parserToken, err2 := myJwt.ParserToken(token)
	if err2 != nil || parserToken == nil {
		response.Fail(60002, "token 解析失败", c)
		return
	}

	response.Ok("token 解析成功", c)
	return
}
*/
