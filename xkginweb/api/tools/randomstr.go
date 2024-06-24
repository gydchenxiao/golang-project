package tools

import (
	"math/rand"
	"time"
)

type RandType int64

const (
	RtNum RandType = 1 << iota
	RtLowAlpha
	RtUpAlpha
	RtPunct
)

func (r RandType) String() string {
	str := ""
	switch {
	case r&RtNum == RtNum:
		str += " 数字随机码 "
		fallthrough
	case r&RtLowAlpha == RtLowAlpha:
		str += " 小写字母随机码 "
		fallthrough
	case r&RtUpAlpha == RtUpAlpha:
		str += " 大写字母随机码 "
		fallthrough
	case r&RtPunct == RtPunct:
		str += " 符号随机码 "
	default:
		str = " 未定义的类型 "

	}
	return str
}

func RandomStr(ty RandType, width int) string {
	str := ""
	switch {
	case ty&RtNum == RtNum:
		str += "0123456789"
		fallthrough
	case ty&RtLowAlpha == RtLowAlpha:
		str += "abcdefghijklmnopqrstuvwxyz"
		fallthrough
	case ty&RtUpAlpha == RtUpAlpha:
		str += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		fallthrough
	case ty&RtPunct == RtUpAlpha:
		str += ",./;'[]-=+!@#$%^&*()`~"
	default:
		str += "ABDEFGHIJKMNQRTUWY123456789abdefghijkmnqrtuwy"
	}
	rand.Seed(time.Now().Unix())
	idx := rand.Int() % len(str)
	res := ""
	for i := 0; i < width; i++ {
		res += str[idx : idx+1]
	}
	return res
}
