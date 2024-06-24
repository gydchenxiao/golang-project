package response

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
	"net/http"
	"strconv"
)

// R Response ServerResult Result

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var (
	VALIDATOR_MAP        = map[string]string{"code": "701", "msg": "属性验证有误"}
	BINDING_PAMATERS_MAP = map[string]string{"code": "702", "msg": "参数绑定有误"}
)

func Result(code int, msg string, data any, c *gin.Context) {
	c.JSON(http.StatusOK, Response{code, msg, data})
}

var (
	CODE       = 20000
	ERROR_CODE = 40001
	MSG        = "success"
)

/*
*
成功
*/
func OkSuccess(c *gin.Context) {
	Result(CODE, MSG, nil, c)
}

/*
*
成功
*/
func Ok(data interface{}, c *gin.Context) {
	Result(CODE, MSG, data, c)
}

/*
*
失败
*/
func Fail(code int, msg string, c *gin.Context) {
	Result(code, msg, map[string]any{}, c)
}

/*
*
失败
*/
func FailWithMessage(msg string, c *gin.Context) {
	Result(ERROR_CODE, msg, map[string]any{}, c)
}

// 权限不足的返回
func FailWithPermission(c *gin.Context) {
	Result(80001, "权限不足", map[string]any{}, c)
}

/*
*
封装错误的返回
*/
func FailWithError(err error, c *gin.Context) {
	Result(ERROR_CODE, err.Error(), map[string]any{}, c)
}

func FailWithData(code int, msg string, data any, c *gin.Context) {
	Result(code, msg, data, c)
}

func FailWithValidatorData(validate *validate.Validation, c *gin.Context) {
	all := validate.Errors.All()
	one := validate.Errors.One()
	code, _ := strconv.ParseInt(VALIDATOR_MAP["code"], 10, 32)
	Result(int(code), one, all, c) // one: 获取报错信息
}

func FailWithBindParams(c *gin.Context) {
	code, _ := strconv.ParseInt(BINDING_PAMATERS_MAP["code"], 10, 32)
	Result(int(code), BINDING_PAMATERS_MAP["msg"], nil, c)
}
