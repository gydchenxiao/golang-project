# Go 整合zap和日志文件分割处理



## 01、下载和安装

```go
# zap核心组件
go get -u go.uber.org/zap
# 日志文件分割,日志文件的保留
go get gopkg.in/natefinch/lumberjack.v2
```

## 02、日志文件对象的初始化

### 1： 定义全局的日志对象

这个日志对象就未来在你代码去使用，但是前提必须要初始化，如何完成初始化呢，在global.go中定义个日志对象即可。然后吧日志进行初始化。找到项目的global.go文件修改如下：

```go
package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sync"
	"xkginweb/commons/parse"
)

var (
	Log        *zap.Logger // -------------------------- 新增代码
	SugarLog   *zap.SugaredLogger // -------------------------- 新增代码
	Lock       sync.RWMutex
	Yaml       map[string]interface{}
	Config     *parse.Config
	KSD_DB     *gorm.DB
	BlackCache local_cache.Cache
	REDIS      *redis.Client
)

```

### 2：在 [initilization](C:\Users\zxc\go\xkginweb\initilization)*中定义*init_zaplog.go的文件来初始化日志对象信息

```go
package initilization

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
	"xkginweb/global"
)

func InitLogger(mode string) {
	var (
		allCore []zapcore.Core
		core    zapcore.Core
	)
	encoder := getEncoder()
	writeSyncerInfo := getLumberJackWriterInfo()
	writeSyncerError := getLumberJackWriterError()
	// 日志是输出终端
	if mode == "debug" || mode == "info" {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		allCore = append(allCore, zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel))
	}

	if mode == "error" {
		allCore = append(allCore, zapcore.NewCore(encoder, writeSyncerError, zapcore.ErrorLevel))
	}

	if mode == "info" {
		allCore = append(allCore, zapcore.NewCore(encoder, writeSyncerInfo, zapcore.InfoLevel))
	}

	core = zapcore.NewTee(allCore...)
	logger := zap.New(core, zap.AddCaller())
	defer logger.Sync()
	global.Log = logger
	global.SugarLog = logger.Sugar()
}

func getLumberJackWriterError() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./zap_error.log", // 日志文件位置
		MaxSize:    5,                 // 进行切割之前，日志文件最大值(单位：MB)，默认100MB
		MaxBackups: 5,                 // 保留旧文件的最大个数
		MaxAge:     1,                 // 保留旧文件的最大天数
		Compress:   false,             // 是否压缩/归档旧文件
	}

	// 输入文件和控制台
	//return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	// 只输出文件
	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberJackLogger))
}

func getLumberJackWriterInfo() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./zap_info.log", // 日志文件位置
		MaxSize:    5,                // 进行切割之前，日志文件最大值(单位：MB)，默认100MB
		MaxBackups: 5,                // 保留旧文件的最大个数
		MaxAge:     1,                // 保留旧文件的最大天数
		Compress:   false,            // 是否压缩/归档旧文件
	}

	// 输入文件和控制台
	//return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	// 只输出文件
	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberJackLogger))
}

// json的方式输出
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = customTimeEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// 空格的方式输出
//func getEncoder() zapcore.Encoder {
//	encoderConfig := zap.NewProductionEncoderConfig()
//	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
//	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
//	return zapcore.NewConsoleEncoder(encoderConfig)
//}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

```

### 3：在main.go中然后初始化 InitLogger的方法即可。

```go
package main

import (
	"xkginweb/initilization"
)

func main() {
	// 解析配置文件
	initilization.InitViper()
    // 初始化日志 开发的时候建议设置成：debug ，发布的时候建议设置成：info/error
	initilization.InitLogger("error")
	// 初始化中间 redis/mysql/mongodb
	initilization.InitMySQL()
	// 初始化缓存
	initilization.InitRedis()
	// 定时器
	// 并发问题解决方案
	// 异步编程
	// 初始化路由
	initilization.RunServer()
}

```

### 4：使用

#### sugar用法

```go
global.SugarLog.Infow("failed to fetch URL",
  // Structured context as loosely typed key-value pairs.
  "url", url,
  "attempt", 3,
  "backoff", time.Second,
)
global.SugarLog.Infof("Failed to fetch URL: %s", url)
```

#### 非sugar
```go
global.Log.Info("failed to fetch URL",
  // Structured context as strongly typed Field values.
  zap.String("url", url),
  zap.Int("attempt", 3),
  zap.Duration("backoff", time.Second),
)
```



## 日志对象是如何初始化呢？

### 日志格式输出的风格

- 默认情况是用空格来定义的，如下：

  ```go
  
  	// 日志级别是：debug 或者 info .那么就默认encoder输出格式化改成正常输出
  	if mode == "debug" || mode == "info" {
  		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
  		allCore = append(allCore, zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel))
  	}
  ```

  上面代码告诉我们。在debug和info级别下。使用 `NewConsoleEncoder` 来进行对日志格式进行输出到控制台，而这种日志格式如下：

  ```go
  2023-08-07T20:33:02.366+0800    DEBUG   initilization/init_gorm.go:53   数据库连接成功。开始运行
  ```

- json格式输出

  ```go
  // 如果error错误级别
  if mode == "error" {
      allCore = append(allCore, zapcore.NewCore(encoder, writeSyncerError, zapcore.ErrorLevel))
  }
  
  // 如果是info级别，也写入日志文件
  if mode == "info" {
      allCore = append(allCore, zapcore.NewCore(encoder, writeSyncerInfo, zapcore.InfoLevel))
  }
  ```

  上面的代码告诉，如果在error或者info的日志级别下，会使用  `encoder == getEncoder()` , 代码如下：

  ```go
  // json的方式输出
  func getEncoder() zapcore.Encoder {
  	encoderConfig := zap.NewProductionEncoderConfig()
  	encoderConfig.EncodeTime = customTimeEncoder
  	return zapcore.NewJSONEncoder(encoderConfig)
  }
  ```

  上面代码就告诉，如果error和info的日志级别，采用的是json的方式来格式化你的日志内容。同时把格式化好json日志内容写入日志文件中

  debug(在开发阶段设置，因为你在开发阶段都已经错误或者问题都解决了才上生存。) > info > warn>error >Fatal

- 那么在开发中我们一般使用什么？info / error 

  





## 日志级别

debug(在开发阶段设置，因为你在开发阶段都已经错误或者问题都解决了才上生存。) > info > warn>error >Fatal

## 日志如何写到日志文件

日志写入到文件使用： `gopkg.in/natefinch/lumberjack.v2`

```go

// 错误日志写入到文件为止
func getLumberJackWriterError() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./kva_error.log", // 日志文件位置
		MaxSize:    5,                 // 进行切割之前，日志文件最大值(单位：MB)，默认100MB
		MaxBackups: 5,                 // 保留旧文件的最大个数
		MaxAge:     1,                 // 保留旧文件的最大天数
		Compress:   false,             // 是否压缩/归档旧文件
	}

	// 输入文件和控制台
	//return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	// 只输出文件
	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberJackLogger))
}

// info日志文件的指定
func getLumberJackWriterInfo() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./kva_info.log", // 日志文件位置
		MaxSize:    5,                // 进行切割之前，日志文件最大值(单位：MB)，默认100MB
		MaxBackups: 5,                // 保留旧文件的最大个数
		MaxAge:     1,                // 保留旧文件的最大天数
		Compress:   false,            // 是否压缩/归档旧文件
	}

	// 输入文件和控制台
	//return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	// 只输出文件
	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(lumberJackLogger))
}

```





## 开始思考，如果扩展日志字段呢？

```
global.Log.Error("我是是一个error的日志级别")
global.Log.Fatal("我是fatal日志级别")
```

写入日志文件

```
{"level":"error","ts":"2023-08-07 20:57:50.955","caller":"xkginweb/main.go:19","msg":"我是是一个error的日志级别"}
{"level":"fatal","ts":"2023-08-07 20:57:50.968","caller":"xkginweb/main.go:20","msg":"我是fatal日志级别"}
{"level":"error","ts":"2023-08-07 20:59:15.305","caller":"xkginweb/main.go:16","msg":"我是是一个error的日志级别"}
{"level":"fatal","ts":"2023-08-07 20:59:15.321","caller":"xkginweb/main.go:20","msg":"我是fatal日志级别"}

```

默认情况下：

- level : 日志基本
- ts : 时间
- caller : 文件和行
- msg : 日志信息

如果扩展字段

```go
global.Log.Error("SugarLog 我是一个错误日志", zap.String("ip", "127.0.0.5"))
global.SugarLog.Errorw("SugarLog 我是一个错误日志", "ip", "127.0.0.1", "port", 8888, "url", "http://www.baidu.com")
global.SugarLog.Errorf("SugarLog 你访问的地址是：%s, 端口是：%d", "127.0.0.1", 8088)
```

## 日志又如何写到kafka中



```go
package initilization

import (
	"fmt"
	"github.com/IBM/sarama"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var (
	logger *zap.Logger
	sugar  *zap.SugaredLogger
)

type LogKafka struct {
	Topic     string
	Producer  sarama.SyncProducer
	Partition int32
}

func (lk *LogKafka) Write(p []byte) (n int, err error) {
	// 构建消息
	msg := &sarama.ProducerMessage{
		Topic:     lk.Topic,
		Value:     sarama.ByteEncoder(p),
		Partition: lk.Partition,
	}
	// 发现消息
	_, _, err = lk.Producer.SendMessage(msg)
	if err != nil {
		return
	}

	return
}

func main() {
	// mode == debug 日志console输出，其他不输出；kafkaSwitch == false 默认输出到文件，kafkaSwitch == true 输出到kafka
	InitLoggerKafka("debug", true)
	// 输出日志
	sugar.Debugf("查询用户信息开始 id:%d", 1)
	sugar.Infof("查询用户信息成功 name:%s age:%d", "zhangSan", 20)
	sugar.Errorf("查询用户信息失败 error:%v", "未该查询到该用户信息")

	time.Sleep(time.Second * 1)
}

func InitLoggerKafka(mode string, kafkaSwitch bool) {
	var (
		err     error
		allCore []zapcore.Core
		core    zapcore.Core
	)

	// 日志是输出终端
	if mode == "debug" {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		allCore = append(allCore, zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel))
	}

	if kafkaSwitch { // 日志输出kafka
		// kafka配置
		config := sarama.NewConfig()                     // 设置日志输入到Kafka的配置
		config.Producer.RequiredAcks = sarama.WaitForAll // 等待服务器所有副本都保存成功后的响应
		//config.Producer.Partitioner = sarama.NewRandomPartitioner // 随机的分区类型
		config.Producer.Return.Successes = true // 是否等待成功后的响应,只有上面的RequiredAcks设置不是NoReponse这里才有用.
		config.Producer.Return.Errors = true    // 是否等待失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.

		// kafka连接
		var kl LogKafka
		kl.Topic = "LogTopic" // Topic(话题)：Kafka中用于区分不同类别信息的类别名称。由producer指定
		kl.Partition = 1      // Partition(分区)：Topic物理上的分组，一个topic可以分为多个partition，每个partition是一个有序的队列。partition中的每条消息都会被分配一个有序的id（offset）
		kl.Producer, err = sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
		if err != nil {
			panic(fmt.Sprintf("connect kafka failed: %+v\n", err))
		}
		encoder := getEncoderKafka()
		writeSyncer := zapcore.AddSync(&kl)
		allCore = append(allCore, zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel))
	} else { // 日志输出file
		encoder := getEncoder()
		writeSyncer := getLumberJackWriter()
		allCore = append(allCore, zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel))
	}

	core = zapcore.NewTee(allCore...)
	logger = zap.New(core, zap.AddCaller())
	defer logger.Sync()
	sugar = logger.Sugar()
}

func getLumberJackWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log", // 日志文件位置
		MaxSize:    1,            // 进行切割之前，日志文件最大值(单位：MB)，默认100MB
		MaxBackups: 5,            // 保留旧文件的最大个数
		MaxAge:     1,            // 保留旧文件的最大天数
		Compress:   false,        // 是否压缩/归档旧文件
	}

	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
}

func getEncoderKafka() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = customTimeEncoderKafka
	return zapcore.NewJSONEncoder(encoderConfig)
}

func customTimeEncoderKafka(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

```



## 封装web请求日志



```go
```

## gorm数据库日志的配置

```go
package initilization

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	logger2 "gorm.io/gorm/logger"
	"log"
	"os"
	"time"
	"xkginweb/commons/orm"
	"xkginweb/global"
)

func InitMySQL() {

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

	m := global.Config.Database.Mysql
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		DontSupportRenameIndex:    true,    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,    // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,   // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		// GORM 定义了这些日志级别：Silent、Error、Warn、Info
		//Logger: logger.Default.LogMode(logger.Info),
		Logger: newLogger,
	})

	// 如果报错
	if err != nil {
		global.Log.Error("数据连接出错了", zap.String("error", err.Error()))
		panic("数据连接出错了" + err.Error()) // 把程序直接阻断，把数据连接好了在启动
	}

	global.KSD_DB = db
	// 初始化数据库表
	orm.RegisterTable()

	// 日志输出
	global.Log.Debug("数据库连接成功。开始运行", zap.Any("db", db))
}

```

## gin的日志配置

```go
package initilization

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
	"xkginweb/commons/filter"
	"xkginweb/commons/middle"
	"xkginweb/global"
	"xkginweb/router"
	"xkginweb/router/code"
	"xkginweb/router/login"
)

func InitGinRouter() *gin.Engine {
	// 打印gin的时候日志是否用颜色标出
	//gin.ForceConsoleColor()
	//gin.DisableConsoleColor()
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 创建gin服务
	ginServer := gin.Default()
	// 提供服务组
	courseRouter := router.RouterWebGroupApp.Course.CourseRouter
	videoRouter := router.RouterWebGroupApp.Video.VideoRouter
	menusRouter := router.RouterWebGroupApp.SysMenu.SysMenusRouter

	// 解决接口的跨域问题
	ginServer.Use(filter.Cors())

	loginRouter := login.LoginRouter{}
	logoutRouter := login.LogoutRouter{}
	codeRouter := code.CodeRouter{}
	// 接口隔离，比如登录，健康检查都不需要拦截和做任何的处理
	// 业务模块接口，
	privateGroup := ginServer.Group("/api")
	// 不需要拦截就放注册中间间的前面,需要拦截的就放后面
	loginRouter.InitLoginRouter(privateGroup)
	codeRouter.InitCodeRouter(privateGroup)
	// 只要接口全部使用jwt拦截
	privateGroup.Use(middle.JWTAuth())
	{
		logoutRouter.InitLogoutRouter(privateGroup)
		videoRouter.InitVideoRouter(privateGroup)
		courseRouter.InitCourseRouter(privateGroup)
		menusRouter.InitSysMenusRouter(privateGroup)
	}

	fmt.Println("router register success")
	return ginServer
}

func RunServer() {
	// 初始化路由
	Router := InitGinRouter()
	// 为用户头像和文件提供静态地址
	Router.StaticFS("/static", http.Dir("/static"))
	address := fmt.Sprintf(":%d", global.Yaml["server.port"])
	// 启动HTTP服务,courseController
	s := initServer(address, Router)
	global.Log.Debug("服务启动成功：端口是：", zap.String("port", "8088"))
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)

	s2 := s.ListenAndServe().Error()
	global.Log.Info("服务启动完毕", zap.Any("s2", s2))
}

```

这里就告诉我们一个道理。你可以自己使用zap来完成gin的日志的事情，或者完成gorm日志的事情。