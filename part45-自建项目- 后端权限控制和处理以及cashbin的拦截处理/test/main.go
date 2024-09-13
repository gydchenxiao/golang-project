package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	logger2 "gorm.io/gorm/logger"
	"log"
	"os"
	"strings"
	"time"
)

func main() {

	// 初始化gorm的日志
	newLogger := logger2.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger2.Config{
			SlowThreshold:             time.Second,  // Slow SQL threshold
			LogLevel:                  logger2.Info, // Log level
			IgnoreRecordNotFoundError: false,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,        // Don't include params in the SQL log
			Colorful:                  true,         // Disable color
		},
	)
	// Initialize a Gorm adapter and use it in a Casbin enforcer:
	// The adapter will use the MySQL database named "casbin".
	// If it doesn't exist, the adapter will create it automatically.
	// You can also use an already existing gorm instance with gormadapter.NewAdapterByDB(gormInstance)
	db, _ := gorm.Open(mysql.Open("root:mkxiaoer@(127.0.0.1:3306)/cashbin?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{
		Logger: newLogger,
	})
	a, _ := gormadapter.NewAdapterByDB(db)
	m, _ := model.NewModelFromFile("./conf/rbac_model.conf")

	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := gormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/abc", true)

	syncedCachedEnforcer, _ := casbin.NewSyncedCachedEnforcer(m, a)
	syncedCachedEnforcer.SetExpireTime(60 * 60)
	// 查询
	syncedCachedEnforcer.LoadPolicy()

	// policy_definition
	// admin	/api/sys/user/:id	POST
	// admin	/api/sys/user/list	POST

	syncedCachedEnforcer.AddRoleForUser("root", "superadmin")
	// 添加自定义函数
	syncedCachedEnforcer.AddFunction("checkSuperAdmin", func(arguments ...interface{}) (interface{}, error) {
		// 获取用户名
		username := arguments[0].(string)
		// 检查用户名的角色是否为superadmin
		return syncedCachedEnforcer.HasRoleForUser(username, "superadmin")
	})
	// 添加自定义函数
	// equals(r.sub, p.sub)
	syncedCachedEnforcer.AddFunction("equals", func(arguments ...interface{}) (interface{}, error) {
		// 获取用户名
		args1 := arguments[0].(string)
		args2 := arguments[1].(string)
		// 检查用户名的角色是否为superadmin
		return strings.EqualFold(args1, args2), nil
	})

	//比较
	enforce, _ := syncedCachedEnforcer.Enforce("admin", "/api/sys/user", "POST")
	fmt.Println(enforce)

	// 保存
	//syncedCachedEnforcer.AddPolicy("alice", "data2", "read")
	//syncedCachedEnforcer.AddPolicy("alice", "data1", "read")
	// 修改
	//old := []string{"alice", "data2", "read"}
	//new := []string{"alice", "data111", "writer"}
	//enforcer.UpdatePolicy(old, new)
	// 删除
	//syncedCachedEnforcer.RemovePolicy("alice", "data2", "read")

	//// Save the policy back to DB.
	syncedCachedEnforcer.SavePolicy()
}
