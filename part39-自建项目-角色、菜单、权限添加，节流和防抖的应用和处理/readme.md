# gorm框架更新多列中，更新0值失败的问题

参考文献：https://gorm.io/zh_CN/docs/update.html

默认情况下：gorm框架更新结构体的时候，只能更新那些非0列。如果你更新为0的列那么久必须使用map

解决方案：

1：直接把0该换其它的状态。（一般不会使用）

2:   把结构体转化成map的方式来进行处理即可

1: 安装组件

```go
go get github.com/fatih/structs
```

2: 定义结构体

```go
package model

type SysUser struct {
	ID        uint   `json:"ID" structs:"-"`
	UUID      string `json:"uuid" structs:"-" ` // 用户UUID
	Slat      string `json:"slat" structs:"-" ` // 用户登录密码
	Enable    int    `json:"enable" structs:"enable" `
	Account   string `json:"account" structs:"account"`    // 用户登录名
	Password  string `json:"password" structs:"password" ` // 密码加盐
	Username  string `json:"username" structs:"username" ` // 用户昵称
	Avatar    string `json:"avatar" structs:"avatar" `     // 用户头像
	Phone     string `json:"phone" structs:"phone" `       // 用户手机号
	Email     string `json:"email" structs:"email" `
	IsDeleted int    `json:"email" structs:"is_deleted"`
}

```

3: 写个测试

```go
package main

import (
	"fmt"
	"github.com/fatih/structs"
	"strutstomap/model"
)

func main() {

	sysUser := model.SysUser{}
	sysUser.ID = 1
	sysUser.UUID = "1111"
	sysUser.Slat = "1111"
	sysUser.Avatar = "XXXXX"
	sysUser.Email = "xxxx@qq.com"
	sysUser.Username = "飞飞"
	sysUser.Account = "feige"
	sysUser.IsDeleted = 0

	fmt.Println("user to map：", structs.Map(sysUser))

}

```

可以看到id,uuid,slat被忽略掉了（structs:"-"`  和 gorm 框架中不写入数据库表字段的设置差不多）。而增加structs的列都会按照你指定的列名作为map的key。

1. structs.Map(sysUser): 把 sysUser struct 转成 Map 类型。

   go get github.com/fatih/structs  中的一个工具函数

2. sysUser struct 不能继承。

   ```golang
   // structs 属性的设置是为了：gorm框架更新多列中，更新0值失败的问题（因为修改了 is_deleted 软删除字段）
   type GVA_MODEL struct {
   	ID        uint      `gorm:"primarykey;comment:主键ID" json:"id" structs:"-"` // 主键ID
   	CreatedAt time.Time `gorm:"type:datetime(0);autoCreateTime;comment:创建时间" json:"createdAt" structs:"-"`
   	UpdatedAt time.Time `gorm:"type:datetime(0);autoUpdateTime;comment:更新时间" json:"updatedAt" structs:"-"`
   
   	// 注意下面的 json:"isDeleted 是 isDeleted 不是 is_deleted 的哦
   	IsDeleted soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt;default:0" json:"isDeleted" structs:"is_deleted"`
   }
   
   // structs 属性的设置是为了：gorm框架更新多列中，更新0值失败的问题（因为修改了 is_deleted 软删除字段）
   type SysRoles struct {
   	global.GVA_MODEL `structs:"-"`
   	RoleName         string `json:"roleName" gorm:"comment:角色名" structs:"role_name"`  // 角色名
   	RoleCode         string `json:"roleCode" gorm:"comment:角色代号" structs:"role_code"` // 角色代号
   }
   
   ```

   所以使用的时候得把 is_deleted 字段再加上的哦：

   ```golang
   	// 结构体转化成map呢？
   	m := structs.Map(sysRole)
   	m["is_deleted"] = sysRole.IsDeleted // global.GVA_MODEL `structs:"-"` 所以需要再加上的哦
   	err = sysRolesService.UpdateSysRolesMap(&sysRole, &m)
   ```

3. 注意：structs 后跟的是数据库表中的字段名，并且作为 Map 结构的 key 值。