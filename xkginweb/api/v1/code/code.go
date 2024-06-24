package code

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"xkginweb/commons/code"
	"xkginweb/commons/response"
)

// 官网：https://github.com/mojocn/base64Captcha
// go get github.com/mojocn/base64Captcha
// Captcha 图形验证码

type CodeApi struct{}

var captcha = code.NewCaptcha(base64Captcha.DefaultDriverDigit, base64Captcha.DefaultMemStore)

// /获取验证码
func (api *CodeApi) Captcha(c *gin.Context) {
	// 验证码的id和验证码的图片信息以及错误信息
	id, b64s, err := captcha.Generate()
	// 开始数据返回
	if err != nil {
		// 开始返回封装数据返回
		response.Fail(601, "验证码生成失败", c)
	} else {
		// 开始返回封装数据返回
		response.Ok(map[string]any{"codeId": id, "b64sImg": b64s}, c)
	}
}

/*

// 下面的函数在登陆验证的时候是没有起到作用的
// 验证输入验证码
func (api *CodeApi) VerfifyCaptcha(c *gin.Context) {
	type CaptchaResult struct {
		Id   string `json:"id" form:"id"`
		Code string `json:"code" form:"code"`
	}

	captchaResult := CaptchaResult{}
	err2 := c.ShouldBindQuery(&captchaResult)
	if err2 != nil {

	}
	match := captcha.Verify(captchaResult.Id, captchaResult.Code, true)
	response.Ok(match, c)
}
*/
