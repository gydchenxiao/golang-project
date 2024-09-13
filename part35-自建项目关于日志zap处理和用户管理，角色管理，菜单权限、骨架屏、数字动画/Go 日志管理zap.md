# go日志框架zap

## 官网地址
https://github.com/uber-go/zap
http://www.topgoer.com/%E9%A1%B9%E7%9B%AE/log/ZapLogger.html



## 为什么要用日志

平时开发我们肯定都比较喜欢实用fmt.println输出和调试数据和查看的一些数据。确实很方便，但是这种代表对于开发者来确实是有价值的。但是运维或者未来把项目发布的生产以后。那么你面临的问题没有控制台。或者你根本就没权限去查看这些控制台的信息。如果上线报错了。你怎么快速把错误和问题定位出来。那么单纯使用fmt.println就满足不了需求了。一般说用日志框架

因为大部分日志框架都已经集成可以写入到日志文件。而且还可以根据日志级别来区分和写入不同日志文件。比如：debug —debug.log .,error —-error.log 。有没有全自动化的过程呢，有使用ELK 。 这种就实时监控和查看你的线上的日志。快速定位问题和原因。

ELK ====== ElasticSearch(数据存储)———Kafka——-Logstash(收集日志的组件) —-Kibana (复杂展示日志信息)

## 介绍
Zap是非常快的、结构化的，分日志级别的Go日志库。
它同时提供了结构化日志记录和printf风格的日志记录
它非常的快 根据Uber-go Zap的文档，它的性能比类似的结构化日志包更好——也比标准库更快。 

## 安装
```go
# zap核心组件---1m
go get -u go.uber.org/zap
# 日志文件分割,日志文件的保留
go get gopkg.in/natefinch/lumberjack.v2
```
## 日志初始化
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

	writeSyncer := getLumberJackWriter()
	// 日志是输出终端
	if mode == "debug" {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		allCore = append(allCore, zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel))
	}

	encoder := getEncoder()
	allCore = append(allCore, zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel))
	core = zapcore.NewTee(allCore...)
	logger := zap.New(core, zap.AddCaller())
	defer logger.Sync()
	global.Log = logger
	global.SugarLog = logger.Sugar()
}

func getLumberJackWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./zap.log", // 日志文件位置
		MaxSize:    5,           // 进行切割之前，日志文件最大值(单位：MB)，默认100MB
		MaxBackups: 5,           // 保留旧文件的最大个数
		MaxAge:     1,           // 保留旧文件的最大天数
		Compress:   false,       // 是否压缩/归档旧文件
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

## 使用

### sugar用法

```go
logger, _ := zap.NewProduction()
defer logger.Sync() // flushes buffer, if any
sugar := logger.Sugar()
sugar.Infow("failed to fetch URL",
  // Structured context as loosely typed key-value pairs.
  "url", url,
  "attempt", 3,
  "backoff", time.Second,
)
sugar.Infof("Failed to fetch URL: %s", url)
```

### 非sugar
```go
logger, _ := zap.NewProduction()
defer logger.Sync()
logger.Info("failed to fetch URL",
  // Structured context as strongly typed Field values.
  zap.String("url", url),
  zap.Int("attempt", 3),
  zap.Duration("backoff", time.Second),
)
```