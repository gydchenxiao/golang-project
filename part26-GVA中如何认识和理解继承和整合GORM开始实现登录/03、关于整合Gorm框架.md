# 关于项目中如何整合GORM框架



## 01、整合gorm框架

### 安装

```go
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

编写main.go来连接数据库

```go
package main

import (
"fmt"
"gorm.io/driver/mysql"
"gorm.io/gorm"
)

func main() {
db, err := gorm.Open(mysql.New(mysql.Config{
    // 登陆数据库用户名   密码   数据库名
    DSN:                       "root:020913@tcp(127.0.0.1:3306)/gva?charset=utf8&parseTime=True&loc=Local", // DSN data source name
    DefaultStringSize:         256,                                                                                     // string 类型字段的默认长度
    DisableDatetimePrecision:  true,                                                                                    // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
    DontSupportRenameIndex:    true,                                                                                    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
    DontSupportRenameColumn:   true,                                                                                    // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
    SkipInitializeWithVersion: false,                                                                                   // 根据当前 MySQL 版本自动配置
}))

if err != nil {
    fmt.Println(err)
}

fmt.Println("数据库连接成功。开始运行", db)
}

```



## 02、编写登录接口



## 03、对接第三方技术：验证码code功能
