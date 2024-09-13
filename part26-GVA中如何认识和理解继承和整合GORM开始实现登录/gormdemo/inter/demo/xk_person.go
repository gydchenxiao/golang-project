package demo

import "fmt"

// 定义子类去实现
type XkPerson struct {
}

func (XkPerson) TableName() string {
	fmt.Println("我是Itable的子类")
	return "xk_person_test"
}
