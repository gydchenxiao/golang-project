package main

import (
	"fmt"
	"gormdemo/inter/demo"
)

func main() {
	// 自己东西方法自己调用
	user := demo.XkUser{}
	name := user.TableName()
	fmt.Println(name)

	// 自己东西方法自己调用
	person := demo.XkPerson{}
	name2 := person.TableName()
	fmt.Println(name2)
}
