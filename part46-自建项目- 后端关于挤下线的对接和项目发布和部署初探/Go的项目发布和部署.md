# Go的项目发布和部署



## 01、准备工作

- 准备一台阿里云服务器
- 开放8080的安全组

## 02、新建一个web工程

### 安装

要安装 Gin 软件包，需要先安装 Go 并设置 Go 工作区。

1.下载并安装 gin：

```sh
go get -u github.com/gin-gonic/gin
```

2.将 gin 引入到代码中：

```go
import "github.com/gin-gonic/gin"
```

### 代码

```go
package main

import "github.com/gin-gonic/gin"

func main() {
	engine := gin.Default()
	engine.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{"code": 200, "msg": "server is success"})
	})
	engine.Run(":8080")
}

```

执行

```
go run main.go
```

### 开始编译

mac  电脑执行

```go
#编译 Linux 64位可执行程序：
GOOS=linux GOARCH=amd64 go build main.go
GOOS=linux GOARCH=arm64 go build main.go
#编译Windows  64位可执行程序：
GOOS=windows GOARCH=amd64 go build main.go
GOOS=windows GOARCH=arm64 go build main.go
#编译 MacOS 64位可执行程序
GOOS=darwin GOARCH=amd64 go build main.go
GOOS=darwin GOARCH=arm64 go build main.go
```

windows执行如下：如果报错执行下面

```go
set CGO_ENABLED=0
set GOARCH=amd64
set GOOS=windows
```

```go
go build main.go
```

Linux

```sh
chmod +x main
./main
#或者
nohup ./main &
```

windows

直接双击打开main.exe即可。
