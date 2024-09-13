package api

import (
	"gormdemo/global"
	"gormdemo/inter/pay"
)

type BbsApi struct {
}

// 保存文章-----http://localhost:8888/buy/news/1
func (api BbsApi) BuyNews() {
	pv1 := global.IOC["payv1"]
	flag := 1

	// 使用接口来完成
	realPv1 := pay.IPay(pv1)
	if flag == 1 {
		realPv1.Alipay()
	} else {
		realPv1.WeixinPay()
	}
}
