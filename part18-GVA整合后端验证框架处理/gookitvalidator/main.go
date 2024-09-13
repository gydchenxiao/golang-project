package main

import (
	"fmt"
	"github.com/gookit/validate"
	"gookitvalidator/model"
	"reflect"
)

func main() {

	bbs := model.XkBbs{
		Title: "xxx33xx",
		Code:  "222",
		Email: "sdfsdfsd@.com",
	}

	toMap := structToMap(bbs)

	v := validate.New(toMap)
	// 添加规则
	v.StringRule("Title", "required|minLen:7|maxLen:15")
	v.StringRule("Email", "email")
	// 添加message
	v.AddMessages(map[string]string{
		"Title.required": "注意，请输入{field}",
		"Title.minLen":   "{field}长度必须大于等于7",
		"Title.maxLen":   "{field}长度必须小于等于15",
		"Email.email":    "{field}不合法",
	})

	v.AddTranslates(map[string]string{
		"Title": "课程标题",
		"Email": "课程邮箱",
	})

	if v.Validate() {

	} else {
		fmt.Println(v.Errors.Error())                  // 返回该字段的错误消息
		fmt.Println(v.Errors.JSON())                   // 返回该字段的错误消息
		fmt.Println(v.Errors.String())                 // 返回该字段的错误消息
		fmt.Println(v.Errors.All()["title"]["minLen"]) // 返回该字段的错误消息
	}

}

func structToMap(s interface{}) map[string]interface{} {
	m := make(map[string]interface{})

	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()

		m[field.Name] = value
	}

	return m
}
