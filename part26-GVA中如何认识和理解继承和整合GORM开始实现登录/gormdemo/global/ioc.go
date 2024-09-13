package global

import (
	"gormdemo/inter/pay"
)

// 模拟创建一个容器
func InitIoc() {
	// 把这些对象全部注册到
	m := map[string]pay.IPay{}
	m["payv1"] = pay.Pay{}
	m["payv2"] = pay.PayV2{}

	IOC = m
}
