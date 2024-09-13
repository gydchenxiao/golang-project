package parse

//# nosql数据的配置
//nosql:
//	redis:
//		host: 127.0.0.1
//		port: 3306
//		password:
//		db: 0

type NoSQL struct {
	Redis Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
	ES    ES    `mapstructure:"es" json:"es" yaml:"es"`
}

// ES配置
type ES struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}

// redis配置
type Redis struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	Dbname   int    `mapstructure:"dbname" json:"dbname" yaml:"dbname"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}
