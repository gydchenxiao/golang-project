package initilization

import "go.uber.org/zap"

// import (
//
//	"fmt"
//	"github.com/IBM/sarama"
//	"go.uber.org/zap"
//	"go.uber.org/zap/zapcore"
//	"gopkg.in/natefinch/lumberjack.v2"
//	"os"
//	"time"
//
// )
var (
	logger *zap.Logger
	sugar  *zap.SugaredLogger
)

//
//type LogKafka struct {
//	Topic     string
//	Producer  sarama.SyncProducer
//	Partition int32
//}
//
//func (lk *LogKafka) Write(p []byte) (n int, err error) {
//	// 构建消息
//	msg := &sarama.ProducerMessage{
//		Topic:     lk.Topic,
//		Value:     sarama.ByteEncoder(p),
//		Partition: lk.Partition,
//	}
//	// 发现消息
//	_, _, err = lk.Producer.SendMessage(msg)
//	if err != nil {
//		return
//	}
//
//	return
//}
//
//func main() {
//	// mode == debug 日志console输出，其他不输出；kafkaSwitch == false 默认输出到文件，kafkaSwitch == true 输出到kafka
//	InitLoggerKafka("debug", true)
//	// 输出日志
//	sugar.Debugf("查询用户信息开始 id:%d", 1)
//	sugar.Infof("查询用户信息成功 name:%s age:%d", "zhangSan", 20)
//	sugar.Errorf("查询用户信息失败 error:%v", "未该查询到该用户信息")
//
//	time.Sleep(time.Second * 1)
//}
//
//func InitLoggerKafka(mode string, kafkaSwitch bool) {
//	var (
//		err     error
//		allCore []zapcore.Core
//		core    zapcore.Core
//	)
//
//	// 日志是输出终端
//	if mode == "debug" {
//		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
//		allCore = append(allCore, zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel))
//	}
//
//	if kafkaSwitch { // 日志输出kafka
//		// kafka配置
//		config := sarama.NewConfig()                     // 设置日志输入到Kafka的配置
//		config.Producer.RequiredAcks = sarama.WaitForAll // 等待服务器所有副本都保存成功后的响应
//		//config.Producer.Partitioner = sarama.NewRandomPartitioner // 随机的分区类型
//		config.Producer.Return.Successes = true // 是否等待成功后的响应,只有上面的RequiredAcks设置不是NoReponse这里才有用.
//		config.Producer.Return.Errors = true    // 是否等待失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
//
//		// kafka连接
//		var kl LogKafka
//		kl.Topic = "LogTopic" // Topic(话题)：Kafka中用于区分不同类别信息的类别名称。由producer指定
//		kl.Partition = 1      // Partition(分区)：Topic物理上的分组，一个topic可以分为多个partition，每个partition是一个有序的队列。partition中的每条消息都会被分配一个有序的id（offset）
//		kl.Producer, err = sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
//		if err != nil {
//			panic(fmt.Sprintf("connect kafka failed: %+v\n", err))
//		}
//		encoder := getEncoderKafka()
//		writeSyncer := zapcore.AddSync(&kl)
//		allCore = append(allCore, zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel))
//	} else { // 日志输出file
//		encoder := getEncoder()
//		writeSyncer := getLumberJackWriter()
//		allCore = append(allCore, zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel))
//	}
//
//	core = zapcore.NewTee(allCore...)
//	logger = zap.New(core, zap.AddCaller())
//	defer logger.Sync()
//	sugar = logger.Sugar()
//}
//
//func getLumberJackWriter() zapcore.WriteSyncer {
//	lumberJackLogger := &lumberjack.Logger{
//		Filename:   "./test.log", // 日志文件位置
//		MaxSize:    1,            // 进行切割之前，日志文件最大值(单位：MB)，默认100MB
//		MaxBackups: 5,            // 保留旧文件的最大个数
//		MaxAge:     1,            // 保留旧文件的最大天数
//		Compress:   false,        // 是否压缩/归档旧文件
//	}
//
//	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
//}
//
//func getEncoderKafka() zapcore.Encoder {
//	encoderConfig := zap.NewProductionEncoderConfig()
//	encoderConfig.EncodeTime = customTimeEncoderKafka
//	return zapcore.NewJSONEncoder(encoderConfig)
//}
//
//func customTimeEncoderKafka(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
//	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
//}
