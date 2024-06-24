package upload

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
	"xkginweb/commons/response"
	"xkginweb/global"
	"xkginweb/utils"
)

type LocalUploadApi struct{}

func (receiver *LocalUploadApi) UploadFile(c *gin.Context) {
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
	filename := name + "_" + format + ext

	log.Println("文件上传以后得名字是：", filename)
	ymspath := time.Now().Format("2006/01/02")
	var relativePath string
	if dir != "" {
		relativePath = global.Config.Local.Path + "/" + dir + "/" + ymspath
	} else {
		relativePath = global.Config.Local.Path + "/" + ymspath
	}
	// 拼接路径和文件名
	filepath := relativePath + "/" + filename
	// 创建父目录
	err := os.MkdirAll(relativePath, os.ModeDir)
	if err != nil {
		response.FailWithMessage("文件创建目录失败", c)
		return
	}
	// 上传文件至指定的完整文件路径
	c.SaveUploadedFile(file, filepath)
	// 定义一个map方法
	m := map[string]any{}
	m["relativeurl"] = filepath
	m["url"] = global.Config.Local.Fileserver + filepath
	m["size"] = file.Size
	m["filename"] = file.Filename
	m["newfilename"] = filename
	m["ext"] = ext

	response.Ok(m, c)
}

func (receiver *LocalUploadApi) UploadFileWangEditor(c *gin.Context) {
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
	filename := name + "_" + format + ext

	log.Println("文件上传以后得名字是：", filename)
	ymspath := time.Now().Format("2006/01/02")
	var relativePath string
	if dir != "" {
		relativePath = global.Config.Local.Path + "/" + dir + "/" + ymspath
	} else {
		relativePath = global.Config.Local.Path + "/" + ymspath
	}
	// 拼接路径和文件名
	filepath := relativePath + "/" + filename
	// 创建父目录
	err := os.MkdirAll(relativePath, os.ModeDir)
	if err != nil {
		response.FailWithMessage("文件创建目录失败", c)
		return
	}
	// 上传文件至指定的完整文件路径
	c.SaveUploadedFile(file, filepath)
	// 定义一个map方法
	m := map[string]any{}
	mdata := map[string]any{}
	m["errno"] = 0
	m["message"] = "success"
	mdata["url"] = global.Config.Local.Fileserver + filepath
	m["data"] = mdata
	// 文件上传成功
	c.JSON(http.StatusOK, m)
}
