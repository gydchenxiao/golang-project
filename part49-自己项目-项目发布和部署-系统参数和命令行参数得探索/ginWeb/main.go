package main

import (
	"flag"
	"fmt"
)

func main() {
	//1:go程序去读取环境变量
	//path := os.Getenv("GOPATH")
	//root := os.Getenv("GOROOT")
	//goenv := os.Getenv("goenv")
	//port := os.Getenv("port")
	//fmt.Println(path)
	//fmt.Println(root)
	//fmt.Println(goenv)
	//fmt.Println(port)

	// 设置环境变量
	//viper.SetDefault("goenv", "dev") //修改这里即可
	//initilization.InitViper()

	//2:go程序去读取命令行参数
	// --goenv=test
	var env string
	// 默认 --goenv=dev
	flag.StringVar(&env, "goenv", "dev", "环境标识")
	flag.Parse()
	fmt.Println(env)
}
