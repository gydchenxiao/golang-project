package demo

import "fmt"

// 定义子类去实现
type XkUser struct {
}

func (XkUser) TableName() string {
	fmt.Println("我是Itable的子类")
	return "xk_user_test"
}
