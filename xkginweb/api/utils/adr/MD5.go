package adr

import (
	"crypto/md5"
	"encoding/hex"
	"xkginweb/utils"
)

// 参数：需要加密的字符串
func getMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// md5加密
func Md5(str string) string {
	return getMd5(getMd5(utils.PWD_SALT + str + utils.PWD_SALT))
}

// md5 加盐
func Md5Slat(str string, slat string) string {
	return getMd5(getMd5(utils.PWD_SALT + str + slat))
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: MD5V
//@description: md5加密
//@param: str []byte
//@return: string

func MD5V(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}
