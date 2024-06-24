package login

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/patrickmn/go-cache"
	"strconv"
	"time"
	"xkginweb/commons/jwtgo"
	"xkginweb/commons/response"
	"xkginweb/global"
	msys "xkginweb/model/entity/sys"
	"xkginweb/utils"
)

// 登录业务
type LoginApi struct{}

// 定义验证的 store --默认是存储在 go 内存中
var store = base64Captcha.DefaultMemStore

// 登录的接口处理
func (api *LoginApi) ToLogined(c *gin.Context) {
	// 1.构造获取用户上传的账号密码和验证码
	type LoginParam struct {
		Account  string // `json:"Account"`   json 管的是返回，LoginParam 不需要返回
		Password string // `json:"Password"`
		Code     string // 用户输入的验证码

		CodeId string // 生成验证码产生的 codeId，也是得再传递回来的
	}
	param := LoginParam{}
	err2 := c.ShouldBindJSON(&param)
	if err2 != nil {
		response.Fail(60002, "参数绑定有误", c)
		return
	}

	// 2. 先验证验证码是否输入正确
	//if len(param.Code) == 0 || len(param.CodeId) == 0 { // 这里前端也是做了检测的
	//	response.Fail(60002, "输入的验证码为空", c)
	//	return
	//}
	//verify := store.Verify(param.CodeId, param.Code, true)
	//if !verify {
	//	response.Fail(60002, "你输入的验证码有误!!", c)
	//	return
	//}

	// 3. 再验证输入的数据的合理性
	inputAccount := param.Account
	inputPassword := param.Password
	if len(inputAccount) == 0 {
		response.Fail(60002, "请输入账号", c)
		return
	}
	if len(inputPassword) == 0 {
		response.Fail(60002, "请输入密码", c)
		return
	}

	// 4. 获取用户在页面上输入的账号和密码开始和数据库里数据进行校验
	// 上面的 Account 和 Password 有值（鉴权没问题），就查询数据库
	dbUser, err := sysUserService.GetUserByAccount(inputAccount)
	if err != nil {
		response.Fail(60002, "你输入的账号和密码有误", c)
		return
	}

	// 这个时候就判断用户输入密码和数据库的密码是否一致
	// inputPassword = utils.Md5(123456) = 2ec9f77f1cde809e48fabac5ec2b8888
	// dbUser.Password = 2ec9f77f1cde809e48fabac5ec2b8888
	// 这个时候就判断用户输入密码和数据库的密码是否一致
	if dbUser != nil && dbUser.Password == utils.Md5Slat(inputPassword, dbUser.Slat) { // md5 加盐加密
		// 根据用户id(dbUser.ID)查询用户的角色 ---> user ==> role
		//userRoles[] * vo.SysRolesVo
		userRoles, err := sysUserRoleService.SelectUserRoles(dbUser.ID)
		if err != nil { // 查询失败
			response.Fail(60002, "根据用户ID, 查询对应的角色失败", c)
			return
		} else if err == nil && len(userRoles) > 0 { // 查询成功 && 查询的结果大于零
			// 封装一个方法出去
			token := api.generaterToken(c, userRoles[0].RoleCode, userRoles[0].ID, dbUser)

			//length := len(sysRolesVo)
			//roleCopy := [length]map[string]any{} // 切片的长度不能定义成一个变量
			// 根据用户角色菜单信息
			roleMenus, _ := sysRoleMenusService.SelectRoleMenus(userRoles[0].ID)
			// 根据用户角色查询用户的角色的权限 ---> roleApi
			permissions, _ := sysRoleApisService.SelectRoleApis(userRoles[0].ID)

			// 这个uuid是用于挤下线使用
			uuid := utils.GetUUID()
			userIdStr := strconv.FormatUint(uint64(dbUser.ID), 10)
			global.Cache.Set("LocalCache:Login:"+userIdStr, uuid, cache.NoExpiration)

			// 3. 生成的 token 成功了，那么后面的业务接口 video/bbs ... 都是要通过 jwt 的 token 验证的哦（在初始化路由的地方统一设置jwt中间件）
			//response.Ok(map[string]interface{}{"roleCopy": roleCopy, "user": dbUser, "token": token, "roles": roles, "permissions": permissions}, c)  // 测试
			response.Ok(map[string]interface{}{"user": dbUser, "token": token, "roles": userRoles, "uuid": uuid, "roleMenus": sysMenuService.Tree(roleMenus, 0), "permissions": permissions}, c) // 前面的 interface{} 也是可以使用一个单词 any 代替的哦
			//return
		} else {
			// 登陆是没有问题的，但是 len(sysRolesVo) <= 0 了，返回空的结果
			response.Ok(map[string]any{"user": dbUser, "token": "", "roles": map[string]any{}, "roleMenus": map[string]any{}, "permissions": map[string]any{}}, c)
			//return
		}
	} else {
		// 没有查询到输入的用户 / 密码验证失败
		response.Fail(60002, "你输入的账号和密码有误", c)
		//return
	}
}

// login 成功后就下发 token（下面的方法是小写开头的，因为外面不需要使用的哦）
func (api *LoginApi) generaterToken(c *gin.Context, roleCode string, roleId uint, dbUser *msys.SysUser) string {
	// 在上面的流程都没有问题的时候，这里 server 就得把 token 给返回 client 的哦
	// 设置token续期的缓冲时间
	bf, _ := utils.ParseDuration("1d") // 缓冲时间，BufferTime
	ep, _ := utils.ParseDuration("7d") // 过期时间，ExpiresTime

	// 1: jwt生成token
	myJwt := jwtgo.NewJWT()
	// 2: 生成token
	token, err2 := myJwt.CreateToken(jwtgo.CustomClaims{
		dbUser.ID,
		dbUser.Username,
		roleCode,
		roleId,
		int64(bf / time.Second),
		jwt.StandardClaims{
			Audience:  "KSD",                                           // 受众
			Issuer:    "KSD-ADMIN",                                     // 签发者
			IssuedAt:  time.Now().Unix(),                               // 签发时间
			NotBefore: time.Now().Add(-1000).Unix(),                    // 生效时间
			ExpiresAt: time.Now().Add(ep).Add(60 * time.Second).Unix(), // 过期时间
		},
	})

	fmt.Println("当前时间是：", time.Now().Unix())
	fmt.Println("缓冲时间：", int64(bf/time.Second))
	fmt.Println("签发时间：" + time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("生效时间：" + time.Now().Add(-1000).Format("2006-01-02 15:04:05"))
	fmt.Println("过期时间：" + time.Now().Add(ep).Add(60*time.Second).Format("2006-01-02 15:04:05"))

	if err2 != nil {
		response.Fail(60002, "登录失败，token颁发不成功!", c)
	}

	return token
}

/*
// 登录的接口处理(使用上面的功能比较完善的ToLogined方法)
func (api *LoginApi) ToLogined(c *gin.Context) {
	// 1：获取用户在页面上输入的账号和密码开始和数据库里数据进行校验
	userService := user.UserService{}

	// 1: 模拟用户前端输入的账号和密码---post+query--如果参数不多可以考虑使用下面
	//axios.post("xxx?account=1&password=111")
	//inputAccount, _ := c.GetQuery("account")
	//inputPassword, _ := c.GetQuery("password")

	// 2: 模拟用户前端输入的账号和密码---post+query--如果参数多可以考虑结构体
	// //axios.post("xxx",params:{data})
	//type LoginParam struct {
	//	Account  string `form:"account"` // form 管的是传入，前端传入参数时候的得是 account，有大小写要求的哦
	//	Password string `form:"password"`
	//}
	//param := LoginParam{}
	//err2 := c.ShouldBindQuery(&param)
	//if err2 != nil {
	//	response.Fail(40001, "参数绑定有误", c)
	//}

	// 3: 模拟用户前端输入的账号和密码---post+json-
	//axios.post("xxx",{
	//    "account": "feige11111",
	//    "password": "123456"
	//})
	type LoginParam struct {
		Account  string // `json:"Account"`   json 管的是返回，LoginParam 不需要返回
		Password string // `json:"Password"`
	}
	param := LoginParam{}
	err2 := c.ShouldBindJSON(&param)
	if err2 != nil {
		response.Fail(40001, "参数绑定有误", c)
	}

	// 验证输入的数据的合理性
	inputAccount := param.Account
	inputPassword := param.Password
	if len(inputAccount) == 0 {
		response.Fail(60002, "请输入账号", c)
		return
	}
	if len(inputPassword) == 0 {
		response.Fail(60002, "请输入密码", c)
		return
	}

	// 上面的 Account 和 Password 有值（鉴权没问题），就查询数据库
	dbUser, err := userService.GetUserByAccount(inputAccount)
	if err != nil {
		response.Fail(60001, "你输入的账号和密码有误", c)
		return
	}
	// 这个时候就判断用户输入密码和数据库的密码是否一致
	if dbUser != nil && dbUser.Password == inputPassword {
		response.Ok(dbUser, c)
	} else {
		response.Fail(60001, "你输入的账号和密码有误", c)
	}
}
*/
