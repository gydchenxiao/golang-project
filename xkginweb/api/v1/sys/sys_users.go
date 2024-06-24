package sys

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
	"strings"
	"xkginweb/commons/response"
	"xkginweb/global"
	context2 "xkginweb/model/context"
	req "xkginweb/model/entity/comms/request"
	resp "xkginweb/model/entity/comms/response"
	"xkginweb/model/entity/sys"
	"xkginweb/utils"
	"xkginweb/utils/adr"
)

type SysUsersApi struct {
	global.BaseApi
}

// 管理员修改用户的密码
func (api *SysUsersApi) UpdatePwd(c *gin.Context) {
	// 1: 接受用户输入的密码和确认密码 和 还要当前操作的用户id
	userPwdContext := context2.UserPwdContext{}
	err := c.ShouldBindJSON(&userPwdContext)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	// 服务端开始验证
	v := validate.Struct(userPwdContext)
	if !v.Validate() {
		response.FailWithValidatorData(v, c)
		return
	}

	if userPwdContext.Password != userPwdContext.ConfirmPassword {
		response.FailWithMessage("密码或确认密码不一致!", c)
		return
	}

	// 2: 然后根据用户id查询是否存在。
	sysUser, _ := sysUserService.GetSysUserByID(userPwdContext.UserId)
	if sysUser.ID == 0 {
		// 5 : 如果不存在，操作不合法,查无此用户信息
		response.FailWithMessage("操作不合法,查无此用户信息", c)
		return
	}

	// 针对性的去更新和处理
	newSysUser := sys.SysUser{}
	newSysUser.ID = sysUser.ID
	newSysUser.Slat = utils.GetUUID()
	newSysUser.Password = adr.Md5Slat(userPwdContext.Password, newSysUser.Slat)
	// 4 : 如果存在 就把slat和 旧密码都修改掉
	sysUserService.UpdateSysUser(&newSysUser)
	//5： 如果没有数据返回，直接返回成功就用OkSuccess
	response.OkSuccess(c)
}

// 修改自己密码
func (api *SysUsersApi) UpdatePwdSelf(c *gin.Context) {
	// 1: 接受用户输入的密码和确认密码 和 还要当前操作的用户id
	sysUserPwdContext := context2.SysUserPwdContext{}
	err := c.ShouldBindJSON(&sysUserPwdContext)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	// 找个验证框架去处理会更好
	if len(sysUserPwdContext.Password) == 0 || len(sysUserPwdContext.ConfirmPassword) == 0 {
		response.FailWithMessage("密码或确认密码不能为空!", c)
		return
	}

	if sysUserPwdContext.Password != sysUserPwdContext.ConfirmPassword {
		response.FailWithMessage("密码或确认密码不一致!", c)
		return
	}

	// 2: 然后根据用户id查询是否存在。
	// 获取当前登录的用户ID = 1 ---l
	sysUser, _ := sysUserService.GetSysUserByID(api.GetLoginUserId(c))
	if sysUser.ID == 0 {
		// 5 : 如果不存在，操作不合法,查无此用户信息
		response.FailWithMessage("操作不合法,查无此用户信息", c)
		return
	}

	// 针对性的去更新和处理
	newSysUser := sys.SysUser{}
	newSysUser.ID = sysUser.ID
	newSysUser.Slat = utils.GetUUID()
	newSysUser.Password = adr.Md5Slat(sysUserPwdContext.Password, newSysUser.Slat)
	// 4 : 如果存在 就把slat和 旧密码都修改掉
	sysUserService.UpdateSysUser(&newSysUser)
	//5： 如果没有数据返回，直接返回成功就用OkSuccess
	response.OkSuccess(c)
}

// 保存
func (api *SysUsersApi) SaveData(c *gin.Context) {
	// 1: 第一件事情就准备数据的载体
	var sysUserContext context2.SysUserContext
	err := c.ShouldBindJSON(&sysUserContext)
	if err != nil {
		// 如果参数注入失败或者出错就返回接口调用这。出错了.
		response.FailWithBindParams(c)
		return
	}

	// 验证
	validation := validate.Struct(sysUserContext)
	if !validation.Validate() {
		response.FailWithValidatorData(validation, c)
		return
	}

	// 判断是否当前账号已经被添加(数据库对于 account 账号字段也是做了处理的哦，添加UNIQUE索引)
	// 数据库和后端都做一下处理是更好的哦
	exitSysUser, _ := sysUserService.GetUserByAccount(sysUserContext.Account)
	if exitSysUser != nil {
		response.FailWithMessage("用户账号已经存在!", c)
		return
	}

	// 校验----框架
	sysUser := sys.SysUser{}
	// 对象复制
	utils.CopyProperties(&sysUser, sysUserContext)
	// 设置uuid
	sysUser.UUID = utils.GetUUID()
	// 设置密码盐
	sysUser.Slat = utils.GetUUID()
	// 密码加密
	sysUser.Password = adr.Md5Slat(sysUser.Password, sysUser.Slat)
	// 创建实例，保存帖子
	err = sysUserService.SaveSysUser(&sysUser)
	// 如果保存失败。就返回创建失败的提升
	if err != nil {
		response.FailWithMessage("创建失败，造成原因可能账号重复", c)
		return
	}
	// 如果保存成功，就返回创建创建成功
	response.Ok("创建成功", c)
}

// (系统用户启用和未启用、删除和未删除)状态修改
func (api *SysUsersApi) UpdateStatus(c *gin.Context) {
	// 先直接使用下面的结构体了，推荐后期改用context数据载体
	type Params struct {
		Id    uint   `json:"id"`
		Filed string `json:"field"`
		Value int    `json:"value"`
	}
	var params Params
	err := c.ShouldBindJSON(&params)
	if err != nil {
		// 如果参数注入失败或者出错就返回接口调用这。出错了.
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = sysUserService.UpdateStatus(params.Id, params.Filed, params.Value)
	// 如果保存失败。就返回创建失败的提升
	if err != nil {
		response.FailWithMessage("更新失败", c)
		return
	}
	// 如果保存成功，就返回创建创建成功
	response.Ok("更新成功", c)
}

/*
// 更新
func (api *SysUsersApi) UpdateById(c *gin.Context) {
	// 1: 第一件事情就准备数据的载体
	var sysUser sys.SysUser
	err := c.ShouldBindJSON(&sysUser)
	if err != nil {
		// 如果参数注入失败或者出错就返回接口调用这。出错了.
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = sysUserService.UpdateSysUser(&sysUser)
	// 如果保存失败。就返回创建失败的提升
	if err != nil {
		response.FailWithMessage("更新失败", c)
		return
	}
	// 如果保存成功，就返回创建创建成功
	response.Ok("更新成功", c)
}
*/

// 按照Map的方式更新
func (api *SysUsersApi) UpdateById(c *gin.Context) {
	// 1: 第一件事情就准备数据的载体
	var sysUser sys.SysUser
	err := c.ShouldBindJSON(&sysUser)
	if err != nil {
		// 如果参数注入失败或者出错就返回接口调用这。出错了.
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 结构体转化成map呢？
	m := structs.Map(sysUser)
	m["is_deleted"] = sysUser.IsDeleted // global.GVA_MODEL `structs:"-"` 所以需要再加上的哦
	err = sysUserService.UpdateSysUserMap(&sysUser, &m)
	// 如果保存失败。就返回创建失败的提升
	if err != nil {
		response.FailWithMessage("更新失败", c)
		return
	}
	// 如果保存成功，就返回创建创建成功
	response.Ok("更新成功", c)
}

// 根据id删除
func (api *SysUsersApi) DeleteById(c *gin.Context) {
	// 绑定参数用来获取/:id这个方式
	id := c.Param("id")
	// 开始执行
	err := sysUserService.DelSysUserById(api.StringToUnit(id))
	if err != nil {
		response.FailWithMessage("删除失败", c)
		return
	}
	response.Ok("ok", c)
}

// 批量删除
func (api *SysUsersApi) DeleteByIds(c *gin.Context) {
	// 绑定参数用来获取?ids=1,2,3,4,5,6
	ids := c.Query("ids")
	// 开始执行
	idstrings := strings.Split(ids, ",")
	var sysUsers []sys.SysUser
	for _, id := range idstrings {
		sysUser := sys.SysUser{}
		sysUser.ID = api.StringToUnit(id)
		sysUsers = append(sysUsers, sysUser)
	}
	err := sysUserService.DeleteSysUsersByIds(sysUsers)
	if err != nil {
		response.FailWithMessage("删除失败", c)
		return
	}
	response.Ok("ok", c)
}

// 根据id查询信息
func (api *SysUsersApi) GetById(c *gin.Context) {
	// 根据id查询方法
	id := c.Param("id")
	// 根据id查询方法
	sysUser, err := sysUserService.GetSysUserByID(api.StringToUnit(id))
	if err != nil {
		global.SugarLog.Errorf("查询用户: %s 失败", id)
		response.FailWithMessage("查询用户失败", c)
		return
	}

	response.Ok(sysUser, c)
}

// 分页查询信息
func (api *SysUsersApi) LoadSysUserPage(c *gin.Context) {
	// 创建一个分页对象
	var pageInfo req.PageInfo
	// 把前端json的参数传入给PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	sysUserPage, total, err := sysUserService.LoadSysUserPage(pageInfo)
	if err != nil {
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}

	response.Ok(resp.PageResult{
		List:     sysUserPage,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, c)
}
