package initilization

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"xkginweb/global"
)

func InitViper(args map[string]any) {
	//获取项目的执行路径
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	config := viper.New()
	//config.AddConfigPath(path + "/conf") //设置读取的文件路径
	//config.SetConfigName("application")  //设置读取的文件名
	//config.SetConfigType("yaml")         //设置文件的类型
	config.SetConfigFile(path + "/conf/application.yaml") // 修改不同的配置文件来设置项目的配置
	//尝试进行配置读取
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}

	config.WatchConfig()
	config.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = config.Unmarshal(&global.Config); err != nil {
			fmt.Println(err)
		}
	})

	// 这里才是把yaml配置文件解析放入到Config对象的过程---map---config
	if err = config.Unmarshal(&global.Config); err != nil {
		fmt.Println(err)
	}

	// 打印文件读取出来的内容:
	keys := config.AllKeys()
	dataMap := make(map[string]interface{})
	for _, key := range keys {
		fmt.Println("yaml存在的key是: " + key)
		dataMap[key] = config.Get(key)
	}

	// 用环境变量覆盖
	// 命令行参数覆盖 boot
	port := args["server.port"].(int)
	if port != -1 {
		dataMap["server.port"] = port
	}

	global.Yaml = dataMap
}
