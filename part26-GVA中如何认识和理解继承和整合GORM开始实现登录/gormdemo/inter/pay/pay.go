package pay

// 支付结构体
type Pay struct {
	Id int
}

// 支付的方法---v1 v2 v3
func (p Pay) Alipay() string {
	return "alipay..."
}

// 支付的方法---v1 v2 v3
func (p Pay) WeixinPay() string {
	return "weixipay..."
}
