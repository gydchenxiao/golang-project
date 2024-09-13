# 一、自定义的方式实现后端的鉴权处理



## 01、上节课为什么debug不生效

造成的原因是：版本升级带来隐患。

go 1.19.2—-泛型（泛型约束），对你的指定模板类型进行约束.

**什么是泛型约束**

```go
type BaseService[D any, T any] struct{}
```



## 02、解决一个bug问题，关于骨架屏幕不退的问题

把原来的状态管理骨架屏幕的代码移植到App.vue 中

```js
import {useSkeletonStore} from '@/stores/skeleton.js'
const skeletonStore = useSkeletonStore()


onMounted(() => {
  setTimeout(() => {
      skeletonStore.skLoading = false;
  }, 600)
})
```



# 二、权限分配

核心：所谓权限控制：其实就于未来不同角色可以访问到不同权限（具体就是你在router定义的每个接口的调用访问权限）。

- role 1 — /api/sys/user/load —-A1001—user:load—有记录——又权限可以访问
- role 1 — /api/sys/user/load —-A1001—user:load —无记录——权限不足

## 关于权限API的分配和管理

这部分逻辑和菜单是一模一样的

### 1：创建表

```sql
CREATE TABLE `sys_apis` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `path` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'api路径',
  `description` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'api中文描述',
  `parent_id` bigint(20) unsigned DEFAULT NULL COMMENT '隶属于菜单的api',
  `method` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT 'POST' COMMENT '方法',
  `is_deleted` bigint(20) unsigned DEFAULT '0',
  `title` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'api路径名称',
  `code` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '权限代号',
  `sort` bigint(20) DEFAULT NULL COMMENT '排序标记',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `path` (`path`),
  UNIQUE KEY `code` (`code`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='权限表';
```

### 2： 结构体

```go
package sys

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type SysApis struct {
	ID          uint                  `gorm:"primarykey;comment:主键ID" json:"id" structs:"-"` // 主键ID
	CreatedAt   time.Time             `gorm:"type:datetime(0);autoCreateTime;comment:创建时间" json:"createdAt" structs:"-"`
	UpdatedAt   time.Time             `gorm:"type:datetime(0);autoUpdateTime;comment:更新时间" json:"updatedAt" structs:"-"`
	IsDeleted   soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt;default:0" json:"isDeleted" structs:"is_deleted"`
	Title       string                `json:"title" gorm:"comment:api路径名称"`          // api路径
	Path        string                `json:"path" gorm:"comment:api路径"`             // api路径
	Description string                `json:"description" gorm:"comment:api中文描述"`    // api中文描述
	ParentId    uint                  `json:"parentId" gorm:"comment:隶属于菜单的api"`     // api组
	Method      string                `json:"method" gorm:"default:POST;comment:方法"` // 方法:创建POST(默认)|查看GET|更新PUT|删除DELETE
	Code        string                `json:"code" gorm:"comment:权限代号"`              // 方法:创建POST(默认)|查看GET|更新PUT|删除DELETE
	// 忽略该字段，- 表示无读写，-:migration 表示无迁移权限，-:all 表示无读写迁移权限
	Children []*SysApis `gorm:"-" json:"children"`
}

func (s *SysApis) TableName() string {
	return "sys_apis"
}

```

### 3: service

```go
package sys

import (
	"xkginweb/global"
	"xkginweb/model/entity/sys"
	"xkginweb/service/commons"
)

// 对用户表的数据层处理
type SysApisService struct {
	commons.BaseService[uint, sys.SysApis]
}

// 添加
func (service *SysApisService) SaveSysApis(sysApis *sys.SysApis) (err error) {
	err = global.KSD_DB.Create(sysApis).Error
	return err
}

// 修改
func (service *SysApisService) UpdateSysApis(sysApis *sys.SysApis) (err error) {
	err = global.KSD_DB.Unscoped().Model(sysApis).Updates(sysApis).Error
	return err
}

// 按照map的方式更新
func (service *SysApisService) UpdateSysApisMap(sysApis *sys.SysApis, mapFileds *map[string]any) (err error) {
	err = global.KSD_DB.Unscoped().Model(sysApis).Updates(mapFileds).Error
	return err
}

// 删除
func (service *SysApisService) DelSysApisById(id uint) (err error) {
	var sysApis sys.SysApis
	err = global.KSD_DB.Where("id = ?", id).Delete(&sysApis).Error
	return err
}

// 批量删除
func (service *SysApisService) DeleteSysApissByIds(sysApiss []sys.SysApis) (err error) {
	err = global.KSD_DB.Delete(&sysApiss).Error
	return err
}

// 根据id查询信息
func (service *SysApisService) GetSysApisByID(id uint) (sysApiss *sys.SysApis, err error) {
	err = global.KSD_DB.Unscoped().Omit("created_at", "updated_at").Where("id = ?", id).First(&sysApiss).Error
	return
}

func (service *SysApisService) FinApiss(keyword string) (sysApis []*sys.SysApis, err error) {
	db := global.KSD_DB.Unscoped().Order("sort asc")
	if len(keyword) > 0 {
		db.Where("title like ?", "%"+keyword+"%")
	}
	err = db.Find(&sysApis).Error
	return sysApis, err
}

/**
*   开始把数据进行编排--递归
*   Tree(all,0)
 */
func (service *SysApisService) Tree(allSysApis []*sys.SysApis, parentId uint) []*sys.SysApis {
	var nodes []*sys.SysApis
	for _, dbApis := range allSysApis {
		if dbApis.ParentId == parentId {
			childrensApis := service.Tree(allSysApis, dbApis.ID)
			if len(childrensApis) > 0 {
				dbApis.Children = append(dbApis.Children, childrensApis...)
			}
			nodes = append(nodes, dbApis)
		}
	}
	return nodes
}

/*
*
数据复制
*/
func (service *SysApisService) CopyData(id uint) (dbData *sys.SysApis, err error) {
	// 2: 查询id数据
	sysApisData, err := service.GetByID(id)
	if err != nil {
		return nil, err
	}
	// 3: 开始复制
	sysApisData.ID = 0
	sysApisData.Path = ""
	sysApisData.Code = ""
	// 4: 保存入库
	data, err := service.Save(sysApisData)

	return data, err
}

```

### 4: 定义接口

```go
package sys

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"xkginweb/commons/response"
	"xkginweb/global"
	"xkginweb/model/entity/sys"
)

type SysApisApi struct {
	global.BaseApi
}

// 拷贝
func (api *SysApisApi) CopyData(c *gin.Context) {
	// 1: 获取id数据 注意定义李媛媛的/:id
	id := c.Param("id")
	data, _ := sysApisService.CopyData(api.StringToUnit(id))
	response.Ok(data, c)
}

// 保存
func (api *SysApisApi) SaveData(c *gin.Context) {
	// 1: 第一件事情就准备数据的载体
	var sysApis sys.SysApis
	err := c.ShouldBindJSON(&sysApis)
	if err != nil {
		// 如果参数注入失败或者出错就返回接口调用这。出错了.
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 创建实例，保存帖子
	err = sysApisService.SaveSysApis(&sysApis)
	// 如果保存失败。就返回创建失败的提升
	if err != nil {
		response.FailWithMessage("创建失败", c)
		return
	}
	// 如果保存成功，就返回创建创建成功
	response.Ok("创建成功", c)
}

// 状态修改
func (api *SysApisApi) UpdateStatus(c *gin.Context) {
	type Params struct {
		Id    uint   `json:"id"`
		Filed string `json:"field"`
		Value any    `json:"value"`
	}
	var params Params
	err := c.ShouldBindJSON(&params)
	if err != nil {
		// 如果参数注入失败或者出错就返回接口调用这。出错了.
		response.FailWithMessage(err.Error(), c)
		return
	}

	flag, _ := sysApisService.UnUpdateStatus(params.Id, params.Filed, params.Value)
	// 如果保存失败。就返回创建失败的提升
	if !flag {
		response.FailWithMessage("更新失败", c)
		return
	}
	// 如果保存成功，就返回创建创建成功
	response.Ok("更新成功", c)
}

// 编辑修改
func (api *SysApisApi) UpdateById(c *gin.Context) {
	// 1: 第一件事情就准备数据的载体
	var sysApis sys.SysApis
	err := c.ShouldBindJSON(&sysApis)
	if err != nil {
		// 如果参数注入失败或者出错就返回接口调用这。出错了.
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 结构体转化成map呢？
	m := structs.Map(sysApis)
	m["is_deleted"] = sysApis.IsDeleted
	err = sysApisService.UpdateSysApisMap(&sysApis, &m)
	// 如果保存失败。就返回创建失败的提升
	if err != nil {
		fmt.Println(err)
		response.FailWithMessage("更新失败", c)
		return
	}
	// 如果保存成功，就返回创建创建成功
	response.Ok("更新成功", c)
}

// 根据id删除
func (api *SysApisApi) DeleteById(c *gin.Context) {
	// 绑定参数用来获取/:id这个方式
	id := c.Param("id")
	// 开始执行
	parseUint, _ := strconv.ParseUint(id, 10, 64)
	err := sysApisService.DelSysApisById(uint(parseUint))
	if err != nil {
		response.FailWithMessage("删除失败", c)
		return
	}
	response.Ok("ok", c)
}

// 根据id查询信息
func (api *SysApisApi) GetById(c *gin.Context) {
	// 根据id查询方法
	id := c.Param("id")
	// 根据id查询方法
	parseUint, _ := strconv.ParseUint(id, 10, 64)
	sysUser, err := sysApisService.GetSysApisByID(uint(parseUint))
	if err != nil {
		global.SugarLog.Errorf("查询用户: %s 失败", id)
		response.FailWithMessage("查询用户失败", c)
		return
	}

	response.Ok(sysUser, c)
}

// 批量删除
func (api *SysApisApi) DeleteByIds(c *gin.Context) {
	// 绑定参数用来获取/:id这个方式
	ids := c.Query("ids")
	idstrings := strings.Split(ids, ",")
	var sysApis []sys.SysApis
	for _, id := range idstrings {
		parseUint, _ := strconv.ParseUint(id, 10, 64)
		sysApi := sys.SysApis{}
		sysApi.ID = uint(parseUint)
		sysApis = append(sysApis, sysApi)
	}

	err := sysApisService.DeleteSysApissByIds(sysApis)
	if err != nil {
		response.FailWithMessage("删除失败", c)
		return
	}
	response.Ok("ok", c)
}

// 查询权限信息
func (api *SysApisApi) FindApisTree(c *gin.Context) {
	keyword := c.Query("keyword")
	sysApis, err := sysApisService.FinApiss(keyword)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		return
	}
	response.Ok(sysApisService.Tree(sysApis, 0), c)
}

```

### 5: 定义具体的接口路由

```go
package sys

import (
	"github.com/gin-gonic/gin"
	v1 "xkginweb/api/v1"
)

// 登录路由
type SysApisRouter struct{}

func (r *SysApisRouter) InitSysApisRouter(Router *gin.RouterGroup) {
	sysApisApi := v1.WebApiGroupApp.Sys.SysApisApi
	// 用组定义--（推荐）
	router := Router.Group("/sys")
	{
		// 获取菜单列表
		router.POST("/apis/tree", sysApisApi.FindApisTree)
		// 保存
		router.POST("/apis/save", sysApisApi.SaveData)
		// 复制数据
		router.POST("/apis/copy/:id", sysApisApi.CopyData)
		// 修改
		router.POST("/apis/update", sysApisApi.UpdateById)
		// 启用和未启用 （控制启用，发布，删除）
		router.POST("/apis/update/status", sysApisApi.UpdateStatus)
		// 删除单个 :id 获取参数的时候id := c.Param("id")，传递的时候/sys/user/del/100
		router.POST("/apis/del/:id", sysApisApi.DeleteById)
		// 删除多个  获取参数的时候ids := c.Query("ids")，传递的时候/sys/user/dels?ids=1,2,3,4
		router.POST("/apis/dels", sysApisApi.DeleteByIds)
		// 查询明细 /user/get/1/xxx
		router.POST("/apis/get/:id", sysApisApi.GetById)
	}
}

```

### 6： 然后在 [initilization](C:\Users\zxc\go\xkginweb\initilization)的init-router.go进行注册

```go
package initilization

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
	"xkginweb/commons/filter"
	"xkginweb/commons/middle"
	"xkginweb/global"
	"xkginweb/router"
)

func InitGinRouter() *gin.Engine {
	// 打印gin的时候日志是否用颜色标出
	//gin.ForceConsoleColor()
	//gin.DisableConsoleColor()
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 创建gin服务
	ginServer := gin.Default()
	// 提供服务组
	courseRouter := router.RouterWebGroupApp.Course.CourseRouter

	videoRouter := router.RouterWebGroupApp.Video.XkVideoRouter

	userStateRouter := router.RouterWebGroupApp.State.UserStateRouter

	bbsRouter := router.RouterWebGroupApp.BBs.XkBbsRouter
	bbsCategoryRouter := router.RouterWebGroupApp.BBs.BBSCategoryRouter

	loginRouter := router.RouterWebGroupApp.Login.LoginRouter
	logoutRouter := router.RouterWebGroupApp.Login.LogoutRouter
	codeRouter := router.RouterWebGroupApp.Code.CodeRouter

	sysMenusRouter := router.RouterWebGroupApp.Sys.SysMenusRouter
	sysApisRouter := router.RouterWebGroupApp.Sys.SysApisRouter // -----------------新增代码
	sysUserRouter := router.RouterWebGroupApp.Sys.SysUsersRouter
	sysRolesRouter := router.RouterWebGroupApp.Sys.SysRolesRouter
	sysUserRolesRouter := router.RouterWebGroupApp.Sys.SysUserRolesRouter
	sysRoleMenusRouter := router.RouterWebGroupApp.Sys.SysRoleMenusRouter
	sysRoleApisRouter := router.RouterWebGroupApp.Sys.SysRoleApisRouter

	// 解决接口的跨域问题
	ginServer.Use(filter.Cors())
	// 接口隔离，比如登录，健康检查都不需要拦截和做任何的处理
	// 业务模块接口，
	privateGroup := ginServer.Group("/api")
	// 无需jwt拦截
	{
		loginRouter.InitLoginRouter(privateGroup)
		codeRouter.InitCodeRouter(privateGroup)
	}
	// 会被jwt拦截
	privateGroup.Use(middle.JWTAuth()).Use(middle.RBAC())
	{
		logoutRouter.InitLogoutRouter(privateGroup)
		videoRouter.InitXkVideoRouter(privateGroup)
		courseRouter.InitCourseRouter(privateGroup)
		userStateRouter.InitUserStateRouter(privateGroup)
		bbsRouter.InitXkBbsRouter(privateGroup)
		bbsCategoryRouter.InitBBSCategoryRouter(privateGroup)
		sysMenusRouter.InitSysMenusRouter(privateGroup)
		sysUserRouter.InitSysUsersRouter(privateGroup)
		sysRolesRouter.InitSysRoleRouter(privateGroup)
		sysApisRouter.InitSysApisRouter(privateGroup)// ---------------------------------新增代码
		sysUserRolesRouter.InitSysUserRolesRouter(privateGroup)
		sysRoleMenusRouter.InitSysRoleMenusRouter(privateGroup)
		sysRoleApisRouter.InitSysRoleApisRouter(privateGroup)
	}

	fmt.Println("router register success")
	return ginServer
}

func RunServer() {

	// 初始化路由
	Router := InitGinRouter()
	// 为用户头像和文件提供静态地址
	Router.StaticFS("/static", http.Dir("/static"))
	address := fmt.Sprintf(":%d", global.Yaml["server.port"])
	// 启动HTTP服务,courseController
	s := initServer(address, Router)
	global.Log.Debug("服务启动成功：端口是：", zap.String("port", address))
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)

	s2 := s.ListenAndServe().Error()
	global.Log.Info("服务启动完毕", zap.Any("s2", s2))
}

```

### 7： 前端定义api接口

```go
import request from '@/request/index.js'
import { C2B  } from '../utils/wordtransfer'

/**
 * 查询权限列表并分页
 */
export const LoadTreeData = (data)=>{
   return request.post(`/sys/apis/tree`,data)
}

/**
 * 根据id查询权限信息
 */
export const GetById = ( id )=>{
   return request.post(`/sys/apis/get/${id}`)
}

/**
 * 保存权限
 */
export const SaveData = ( data )=>{
   return request.post(`/sys/apis/save`,data)
}

/**
 * 更新权限信息
 */
export const UpdateData = ( data )=>{
   return request.post(`/sys/apis/update`,data)
}


/**
 * 根据id删除权限信息
 */
export const DelById = ( id )=>{
   return request.post(`/sys/apis/del/${id}`)
}

/**
 * 根据ids批量删除权限信息
 */
export const DelByIds = ( ids )=>{
   return request.post(`/sys/apis/dels?ids=${ids}`)
}

/**
 * 权限启用和未启用
 */
export const UpdateStatus = ( data )=>{
   data.field = C2B(data.field)
   return request.post(`/sys/apis/update/status`,data)
}


/**
 * 复制数据
 */
export const CopyData = ( id )=>{
   return request.post(`/sys/apis/copy/${id}`,{})
}

```

### 8: 开始对接页面 views/sys/Permission.vue

```vue
<template>
  <div class="kva-container">
    <div class="kva-contentbox">
      <home-page-header>
        <div class="kva-form-search">
          <el-form :inline="true" ref="searchForm" :model="queryParams">
            <el-form-item>
              <el-button type="primary"  icon="Plus" @click="handleAdd">添加权限</el-button>
            </el-form-item>
            <el-form-item label="关键词：">
              <el-input v-model="queryParams.keyword" placeholder="请输入菜单名称..." maxlength="10" clearable />
            </el-form-item>
            <el-form-item>
              <el-button type="primary" icon="Search" @click.prevent="handleSearch">搜索</el-button>
              <el-button type="danger" icon="Refresh" @click.prevent="handleReset">重置</el-button>
            </el-form-item>
          </el-form>
        </div>
        <!-- default-expand-all -->
        <el-table
          :data="tableData"
          style="width: 100%; margin-bottom: 20px"
          row-key="id"
          border
          stripe
          :height="settings.tableHeight()"
        >
          <el-table-column fixed prop="id" label="ID" align="center" width="80"  />
          <el-table-column fixed prop="parentId" label="父ID" align="center" width="80" />
          <el-table-column prop="title" label="展示名字" align="center" >
            <template #default="{row}">
                <el-input v-model="row.title" style="text-align:center" @change="handleChange(row,'title')"></el-input>
            </template>
          </el-table-column>
          <el-table-column prop="code" label="编号" align="center" >
            <template #default="{row}">
                <el-input v-model="row.code" style="text-align:center" @change="handleChange(row,'code')"></el-input>
            </template>
          </el-table-column>
          <el-table-column prop="code" label="访问路径" align="center" >
            <template #default="{row}">
                <el-input v-model="row.path" style="text-align:center" @change="handleChange(row,'path')"></el-input>
            </template>
          </el-table-column>
          <el-table-column prop="sort" label="排序"  align="center" width="180">
            <template #default="{row}">
                <el-input-number v-model="row.sort" @change="handleChange(row,'sort')"></el-input-number>
            </template>
          </el-table-column>
          <el-table-column label="是否删除" align="center" width="180">
            <template #default="{row}">
              <el-switch 
              v-model="row.isDeleted" 
              @change="handleChange(row,'isDeleted')" 
              active-color="#ff0000"
               active-text="已删除" 
               inactive-text="未删除" 
               :active-value="1" 
               :inactive-value="0"/>
            </template>
          </el-table-column>
          <el-table-column label="创建时间" align="center" width="160">
            <template #default="scope">
              {{ formatTimeToStr(scope.row.createdAt,"yyyy/MM/dd hh:mm:ss") }}
            </template>
          </el-table-column>
          <el-table-column label="更新时间" align="center" width="160">
            <template #default="scope">
              {{ formatTimeToStr(scope.row.updatedAt,"yyyy/MM/dd hh:mm:ss") }}
            </template>
          </el-table-column>
          <el-table-column fixed="right" align="center" label="操作" width="350">
            <template #default="{row,$index}">
              <el-button text icon="edit" @click="handleEdit(row)"  type="primary">编辑</el-button>
              <el-button text icon="Tickets" @click="handleCopy(row)"  type="success">复制</el-button>
              <el-button text icon="remove" @click="handleRemove(row)"  type="danger">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
      </home-page-header>
    </div>
    <!--添加和修改菜单-->
    <add-sys-apis ref="addRef"  @load="handleLoadData"></add-sys-apis>
  </div>
</template>

<script  setup>
import { C2B,B2C } from '@/utils/wordtransfer'
import KVA from '@/utils/kva.js'
import settings from '@/settings';
import { formatTimeToStr } from '@/utils/date'
import AddSysApis from '@/views/sys/components/AddSysApis.vue'
import { LoadTreeData,UpdateStatus,CopyData,DelById } from '@/api/sysapis.js';
import { reactive } from 'vue';
import { useUserStore } from '@/stores/user.js'
const userStore = useUserStore()
const addRef = ref(null);

// 搜索属性定义
let queryParams = reactive({
  keyword:""
})

// 数据容器
const tableData = ref([]) 
const searchForm = ref(null)
// 搜索
const handleSearch = ()=> {
  handleLoadData()
}

// 查询列表
const handleLoadData = async ()=>{
  const resp = await LoadTreeData(queryParams)
  tableData.value = resp.data
}

// 添加
const handleAdd = ()=>{
  addRef.value.handleOpen('','save',tableData.value?.length)
}

// 编辑
const handleEdit =  async (row) => {
  // 在打开,再查询，
  addRef.value.handleOpen(row.id,'edit',tableData.value?.length)
}

// 添加子菜单
const handleAddChild = (row) => {
  addRef.value.handleOpen(row,'child',row.children?.length)
}

// 改变序号 sorted,标题 title、启用 status,isDeleted
const handleChange = async (row,field) =>{
  var value = row[field];//row.isDeleted=0 
  var params = {id:row.id,field:field,value:value};
  await UpdateStatus(params); 
  KVA.notifySuccess("更新成功")
  if(field=="sort"){
    tableData.value.sort((a,b)=>a.sort-b.sort);
  }
}



// 物理删除
const handleRemove =  async (row) => {
  try{
    await KVA.confirm("警告","你确定要抛弃我么？",{icon:"error"})
    await DelById(row.id)
    KVA.notifySuccess("操作成功")
    userStore.handlePianaRole(0,"")
    handleLoadData()
  }catch(e){
    KVA.notifyError("操作失败")
  }
}

// 重置搜索表单
const handleReset = () => {
  queryParams.keyword = ""
  searchForm.value.resetFields()
  handleLoadData()
}



// 复制
const handleCopy = async (row) => {
  await CopyData(row.id);
  KVA.notifySuccess("复制成功")
  handleLoadData()
}

// 生命周期加载
onMounted(()=>{
  handleLoadData()

  console.log("C2B",C2B("isDeletedNum"))
  console.log("B2C",B2C("is_deleted_num"))
})


</script>

```

具体后续就不展开了

