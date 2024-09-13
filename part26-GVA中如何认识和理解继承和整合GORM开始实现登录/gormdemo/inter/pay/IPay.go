package pay

type IPay interface {
	// 支付的方法---v1 v2 v3
	Alipay() string
	// 支付的方法---v1 v2 v3
	WeixinPay() string
}

// interface 接口的作用：1.      2. 升级隔离
