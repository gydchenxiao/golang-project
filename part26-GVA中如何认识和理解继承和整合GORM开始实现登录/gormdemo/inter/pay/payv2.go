package pay

// 支付结构体
type PayV2 struct{}

// 支付的方法---v1 v2 v3
func (p PayV2) Alipay() string {
	return "alipay v2..."
}

// 支付的方法---v1 v2 v3
func (p PayV2) WeixinPay() string {
	return "weixipay v2..."
}
