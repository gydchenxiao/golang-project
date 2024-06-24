package code

//
//import (
//	"github.com/gin-gonic/gin"
//	"github.com/mojocn/base64Captcha"
//	"xkginweb/commons/response"
//)
//
//// 官网：https://github.com/mojocn/base64Captcha
//// go get github.com/mojocn/base64Captcha
//// Captcha 图形验证码
//
//var store = base64Captcha.DefaultMemStore
//
//type CodeController struct{}
//
//// /获取验证码
//func (api *CodeController) Captcha(c *gin.Context) {
//	/*
//		driverString := base64Captcha.DriverString{
//			Height:          30,
//			Width:           60,
//			NoiseCount:      0,
//			ShowLineOptions: 2 | 2,
//			Length:          4,
//			Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
//			BgColor: &color.RGBA{
//				R: 3,
//				G: 102,
//				B: 214,
//				A: 125,
//			},
//			Fonts: []string{"wqy-microhei.ttc"},
//		}
//
//		driver := driverString.ConvertFonts()
//	*/
//
//	driver := base64Captcha.DefaultDriverDigit
//	// 验证码的长度
//	driver.Length = 4
//	// 生成验证码
//	captcha := base64Captcha.NewCaptcha(driver, store)
//	// 验证码的id和验证码的图片信息以及错误信息
//	id, b64s, err := captcha.Generate()
//	// 开始数据返回
//	if err != nil {
//		// 开始返回封装数据返回
//		response.Fail(601, "验证码生成失败", c)
//	} else {
//		// 开始返回封装数据返回
//		response.Ok(map[string]any{"id": id, "img": b64s}, c)
//	}
//}
//
//// 验证输入验证码
//func (api *CodeController) VerfifyCaptcha(c *gin.Context) {
//	type CaptchaResult struct {
//		Id   string `json:"captchaId"`
//		Code string `json:"code"`
//	}
//
//	captchaResult := CaptchaResult{}
//	err2 := c.ShouldBindJSON(&captchaResult)
//	driver := base64Captcha.DefaultDriverDigit
//	captcha := base64Captcha.NewCaptcha(driver, store)
//
//	func (c *Captcha) Verify(id, answer string, clear bool) (match bool) {
//		match = c.Store.Get(id, clear) == answer
//		return
//	}
//
//	// 验证码的长度
//	driver.Length = 4
//	// 生成验证码
//
//	// 验证码的id和验证码的图片信息以及错误信息
//	id, b64s, err := captcha.Generate()
//	// 开始数据返回
//	if err != nil {
//		// 开始返回封装数据返回
//		response.Fail(601, "验证码生成失败", c)
//	} else {
//		// 开始返回封装数据返回
//		response.Ok(map[string]any{"id": id, "img": b64s}, c)
//	}
//}
