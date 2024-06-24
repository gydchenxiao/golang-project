package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// 注意加密以后，那么就必须考虑到login的时候要加密比对

// 参数：需要加密的字符串
// 单独使用 MD5 加密，网上是有破解的网站的，所以后面有了 1. 加盐加密 / 2. 动态加盐加密
func getMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// md5加密
func Md5(str string) string {
	return getMd5(getMd5(PWD_SALT + str + PWD_SALT))
}

// md5 加盐加密（1. 增加黑客破解密码的难度     2. 在数据库也可以再设置一个 salt 字段是得每一个用户的 salt 都不同）
func Md5Slat(str string, slat string) string {
	return getMd5(getMd5(PWD_SALT + str + slat))
}

// @author: [piexlmax](https://github.com/piexlmax)
// @function: MD5V
// @description: md5加密
// @param: str []byte
// @return: string
func MD5V(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}
