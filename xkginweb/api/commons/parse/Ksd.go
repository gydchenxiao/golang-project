package parse

type Ksd struct {
	Alipay Alipay `mapstructure:"alipay" json:"alipay" yaml:"alipay"`
}

type Alipay struct {
	Appid   string         `mapstructure:"appid" json:"appid" yaml:"appid"`
	Screat  string         `mapstructure:"screat" json:"screat" yaml:"screat"`
	URLS    []string       `mapstructure:"urls" json:"urls" yaml:"urls"`
	Paths   []string       `mapstructure:"paths" json:"path" yaml:"paths"`
	Routers []Router       `mapstructure:"routers" json:"routers" yaml:"routers"`
	Detail  Detail         `mapstructure:"detail" json:"detail" yaml:"detail"`
	Map     map[string]any `mapstructure:"map" json:"map" yaml:"map"`
}

type Detail struct {
	Id   int    `mapstructure:"id" json:"id" yaml:"id"`
	Name string `mapstructure:"name" json:"name" yaml:"name"`
}

type Router struct {
	Id     int    `mapstructure:"id" json:"id" yaml:"id"`
	Url    string `mapstructure:"url" json:"url" yaml:"url"`
	Filter string `mapstructure:"filter" json:"filter" yaml:"filter"`
}
