package demo

// 定义接口
type ITable interface {
	// 定义方法标准
	TableName() string
}

// 接口：
// 1：标准规范 ----可以实现多态
// 2：升级隔离的作用
