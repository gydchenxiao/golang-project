package main

import (
	"flag"
	"xkginweb/initilization"
)

func main() {
	// 设置 配置文件env 和 服务端口server.port
	var env string
	var port int
	flag.StringVar(&env, "env", "dev", "环境标识")
	flag.IntVar(&port, "server.port", -1, "测试端口")
	flag.Parse()
	// 这里也可以使用一个结构体
	args := map[string]any{"env": env, "server.port": port}
	// 设置环境变量
	//viper.SetDefault("env", env)
	// 解析配置文件
	initilization.InitViper(args)

	// 日志管理
	// 初始化日志 开发的时候建议设置成：debug ，发布的时候建议设置成：info/error
	initilization.InitLogger("debug")

	// 初始化中间 redis/mysql/mongodb
	initilization.InitMySQL()

	// 初始化本地缓存
	initilization.InitCache()

	// jwt + login + 权限

	// 定时器

	// 并发问题解决方案

	// 异步编程

	// 初始化路由
	initilization.RunServer()
}
