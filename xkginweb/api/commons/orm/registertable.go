package orm

import (
	"xkginweb/global"
	"xkginweb/model/entity/bbs"
	"xkginweb/model/entity/jwt"
	"xkginweb/model/entity/sys"
	"xkginweb/model/entity/user"
	"xkginweb/model/entity/video"
)

func RegisterTable() {
	db := global.KSD_DB
	// 注册和声明model
	db.AutoMigrate(user.XkUser{})
	db.AutoMigrate(user.XkUserAuthor{})
	// 系统用户，角色，权限表
	db.AutoMigrate(sys.SysApis{})
	db.AutoMigrate(sys.SysMenus{})
	db.AutoMigrate(sys.SysRoleApis{})
	db.AutoMigrate(sys.SysRoleMenus{})
	db.AutoMigrate(sys.SysRoles{})
	db.AutoMigrate(sys.SysUserRoles{})
	db.AutoMigrate(sys.SysUser{})
	// 视频表
	db.AutoMigrate(video.XkVideo{})
	db.AutoMigrate(video.XkVideoCategory{})
	db.AutoMigrate(video.XkVideoChapterLesson{})
	// 社区
	db.AutoMigrate(bbs.XkBbs{})
	db.AutoMigrate(bbs.BbsCategory{})

	// 声明一下jwt模型
	db.AutoMigrate(jwt.JwtBlacklist{})
}
