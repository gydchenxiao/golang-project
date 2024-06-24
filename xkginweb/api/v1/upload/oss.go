package upload

import (
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path"
	"strings"
	"time"
	"xkginweb/commons/response"
	"xkginweb/utils"
)

type OSSUploadApi struct{}

var (
	// oss服务站点
	Endpoint = "oss-cn-guangzhou.aliyuncs.com"
	// 安全密钥信息
	AccessKeyId = "LTAI5tRcQjUkHoansqTwT8Q9"
	// 私钥，注意这个不能被别人知道，否则别人就可以通过代码oss中文件信息全部下载和读取出来
	AccessKeySecret = "YrWMgOIX3zG3qNsoBYq29tTP3QeuOq"
)

func (ossApi *OSSUploadApi) UploadFile(c *gin.Context) {
	// 获取文件上传传递过来的参数
	dir := c.PostForm("dir")
	//xxx := c.PostForm("xxxx")
	// 单文件
	file, _ := c.FormFile("file")
	// 读取文件后缀
	ext := path.Ext(file.Filename)
	// 读取文件名并加密
	name := strings.TrimSuffix(file.Filename, ext)
	name = utils.MD5V([]byte(name))
	format := time.Now().Format("20060102150405")
	// 拼接新文件名
	filename := name + format + ext

	log.Println("文件上传以后得名字是：", filename)
	ymspath := time.Now().Format("2006/01/02")
	var datePath string
	if dir != "" {
		datePath = dir + "/" + ymspath
	} else {
		datePath = ymspath
	}
	// 拼接路径和文件名
	relativePath := datePath + "/" + filename

	bucketname := "ksdcourse"
	// oss创建bucketname
	ossApi.CreateBuketName(bucketname)
	// oss文件上传开始
	client := ossApi.InitOssClient()
	// 获取到bucket对象,开始往bucket中添加文件信息
	bucket, _ := client.Bucket(bucketname)
	// 将用户选择的文件进行io流处理
	open, _ := file.Open()
	// 开始进行oss的文件上传
	err := bucket.PutObject(relativePath, open)
	if err != nil {
		response.FailWithMessage("上传成功", c)
		// 这里最好阻止一下
		return
	}

	// 定义一个map方法
	m := map[string]any{}
	m["relativeurl"] = relativePath
	m["url"] = "https://" + bucketname + "." + Endpoint + "/" + relativePath
	m["size"] = file.Size
	m["filename"] = file.Filename
	m["newfilename"] = filename
	m["ext"] = ext

	response.Ok(m, c)
}

// 1： 创建bucketname
func (ossApi *OSSUploadApi) InitOssClient() *oss.Client {
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	client, err := oss.New(Endpoint, AccessKeyId, AccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	return client
}

// 创建bucketname
func (ossApi *OSSUploadApi) CreateBuketName(bucketName string) (bool, error) {
	// 2: 创建bucketname
	exist, _ := ossApi.InitOssClient().IsBucketExist(bucketName)
	if !exist {
		err := ossApi.InitOssClient().CreateBucket(bucketName)
		if err != nil {
			return false, errors.New("创建bucketname【" + bucketName + "】失败了")
		}
		return true, nil
	}
	return false, errors.New("【" + bucketName + "】已经存在了")
}
