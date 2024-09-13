package main

//
//import (
//	"fmt"
//	"gookitvalidator/model"
//)
//import "github.com/gookit/validate"
//
//func main() {
//
//	// 1：定义一个结构体
//	bbs := &model.XkBbs{
//		Title: "1223423423",
//		Code:  "222",
//		Email: "sdfsdfsd@qq.com",
//	}
//
//	// 2:创建 Validation 实例
//	v := validate.Struct(bbs)
//
//	// 3: 获取验证结果
//	if v.Validate() {
//		// 你可以做你想要做的事情
//	} else {
//		//fmt.Println(v.Errors.Field("Title")["required"]) // 所有的错误消息
//		//fmt.Println(v.Errors.Field("Title")["minLen"])   // 所有的错误消息
//		//fmt.Println(v.Errors.Field("Email")["email"])    // 所有的错误消息
//		//fmt.Println(v.Errors.One())         // 返回随机一条错误消息
//		//fmt.Println(v.Errors.Field("Name")) // 返回该字段的错误消息
//		fmt.Println(v.Errors.Error())                  // 返回该字段的错误消息
//		fmt.Println(v.Errors.JSON())                   // 返回该字段的错误消息
//		fmt.Println(v.Errors.String())                 // 返回该字段的错误消息
//		fmt.Println(v.Errors.All()["title"]["minLen"]) // 返回该字段的错误消息
//	}
//}
