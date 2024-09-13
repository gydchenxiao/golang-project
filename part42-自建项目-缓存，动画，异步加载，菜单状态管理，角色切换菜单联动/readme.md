# 查询角色对应的菜单

### 1:  根据用户查询角色 api/login/login.go

```go

// 登录的接口处理
func (api *LoginApi) ToLogined(c *gin.Context) {

	//  省略代码.........................
    
	// 这个时候就判断用户输入密码和数据库的密码是否一致
	// inputPassword = utils.Md5(123456) = 2ec9f77f1cde809e48fabac5ec2b8888
	// dbUser.Password = 2ec9f77f1cde809e48fabac5ec2b8888
	if dbUser != nil && dbUser.Password == adr.Md5Slat(inputPassword, dbUser.Slat) {
		token := api.generaterToken(c, dbUser)
		// 根据用户id查询用户的角色
		userRoles, _ := sysUserRolesService.SelectUserRoles(dbUser.ID)//-----新增
		// 根据用户查询菜单信息
		roleMenus, _ := sysRoleMenusService.SelectRoleMenus(userRoles[0].ID) //----新增
		// 根据用户id查询用户的角色的权限
		permissions, _ := sysRoleApisService.SelectRoleApis(userRoles[0].ID) //----新增
		
		// 查询返回
		response.Ok(map[string]any{"user": dbUser, "token": token, "roles": userRoles, "roleMenus": sysMenuService.Tree(roleMenus, 0), "permissions": permissions}, c)
	} else {
		response.Fail(60002, "你输入的账号和密码有误", c)
	}
}
```

### 2: 根据用户id查询对应的角色

```go
userRoles, _ := sysUserRolesService.SelectUserRoles(dbUser.ID)//-----新增
```

具体如下：

```go
// 查询用户授权的角色信息
func (service *SysUserRolesService) SelectUserRoles(userId uint) (sysRoles []*sys2.SysRoles, err error) {
	err = global.KSD_DB.Select("t2.*").Table("sys_user_roles t1,sys_roles t2").
		Where("t1.user_id = ? and t1.role_id = t2.id", userId).Scan(&sysRoles).Error
	return sysRoles, err
}

```

### 3:  切换角色必须要角色对应菜单也进行刷新

默认情况下，我们把用户对应的角色列表查询出来，把第一个作为默认的角色，那么就必须把默认的角色对应菜单和权限全部都全查询返回，然后页面根据服务端返回的菜单和权限进行渲染。

```go
```