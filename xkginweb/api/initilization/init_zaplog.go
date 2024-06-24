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

	//--------------------------------------------控制台输出（比如：cmd/golang）
	// 日志输出模板格式化----json/普通
	encoder := getEncoder()

	//--------------------------------------------日志文件输出
	// 日志文件的输出info输出---info日志级别低都写入到info.log文件中
	writeSyncerInfo := getLumberJackWriterInfo()
	// 日志文件的输出error输出----如果比error低就全部写入这里
	writeSyncerError := getLumberJackWriterError()

	// 日志级别是：debug 或者 info .那么就默认encoder输出格式化改成正常输出
	if mode == "debug" || mode == "info" || mode == "error" {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		allCore = append(allCore, zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel))
	}

	// 如果error错误级别
	if mode == "error" {
		allCore = append(allCore, zapcore.NewCore(encoder, writeSyncerError, zapcore.ErrorLevel))
	}

	// 如果是info级别，也写入日志文件
	if mode == "info" {
		allCore = append(allCore, zapcore.NewCore(encoder, writeSyncerInfo, zapcore.InfoLevel))
	}

	// 初始化zap
	core = zapcore.NewTee(allCore...)
	logger := zap.New(core, zap.AddCaller())
	defer logger.Sync()
	// 全局日志对象初始化
	global.Log = logger
	global.SugarLog = logger.Sugar()
}

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
