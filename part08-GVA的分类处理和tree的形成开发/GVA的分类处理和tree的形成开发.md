## 1. 设置后端的model，service，router，api

### 1.model/videocategory/xk_video_category.go

后端数据库数据模型的设置

```go
package videocategory

import "time"

type XkVideoCategory struct {
	ID           uint      `gorm:"primarykey;comment:主键ID" json:"id" form:"id"`
	CategoryName string    `json:"categoryName" gorm:"not null;default:'';comment:分类名称"`
	Description  string    `json:"description" gorm:"not null;default:'';comment:分类描述"`
	CreateTime   time.Time `gorm:"type:datetime(0);comment:创建时间" json:"createTime"`
	UpdatedTime  time.Time `gorm:"type:datetime(0);comment:更新时间" json:"updatedTime"`
	ParentId     uint      `json:"parentId" gorm:"not null;default:0;comment:分类的主ID"`
	Status       int8      `json:"status" gorm:"not null;default:1;comment:0 未发布 1 发布"`
	Sorted       int8      `json:"sorted" gorm:"not null;default:1;comment:0 排序"`
	IsDelete     int8      `json:"isDelete" gorm:"not null;default:0;comment:0 未删除 1 删除"`
	// 忽略该字段，- 表示无读写，-:migration 表示无迁移权限，-:all 表示无读写迁移权限
	Children []*XkVideoCategory `gorm:"-" json:"childrens"`
}

// 建设表
// 1: 使用合理命令
// 2: 数据类型也要合理，长度
// 3: 尽量不要使用null，
// 4: 而且状态列一定给默认值
// 5 : 一定要写注释

// 指定当前model的数据库生成的表明
func (XkVideoCategory) TableName() string {
	// 这里的命名不能-只能用_下划线
	return "xk_video_category"
}

// 1. 关于上面自己设置的主键 Vid 的 form 属性没有设置的问题：
// 那么浏览器输入 http://127.0.0.1:8888/videocategory/get?vid=2 的时候，我们后台获取的 vid 默认是为 0 的哦，自然就 "record not found"
// 加上 form:"vid" 之后是没有这个问题的

// 2. 关于设置的数据库字段 ParentId 系统在数据库默认设置为了 parent_id。自己如果有设置为 json:"parentId" 是没有反应的

```

并设置进去 initialize/gorm.go 内，自动在数据库形成数据表`videocategory.XkVideoCategory{}`

并建立一个 request 文件夹，写入对应的文件



### 2.service/videocategory/xk_video_category_service.go

并写好对应的 enter.go 文件

```go
package videoscategoryservice

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/videocategory"
)

// 定义bbs的service提供xkbbs的数据curd的操作

type XkVideoCategoryService struct{}

// @author: feige
// @function: CreateXkBbs
// @description: 创建文章
// @param: e videocategory.VideoCategory
// @return: err error
func (cvideocategory *XkVideoCategoryService) CreateXkVideoCategory(xkVideoCategory *videocategory.XkVideoCategory) (err error) {
	// 1： 获取数据的连接对象
	err = global.GVA_DB.Create(xkVideoCategory).Error
	return err
}

// @author: feige
// @function: UpdateXkBbs
// @description: 更新文章
// @param: e *model.ExaCustomer
// @return: err error
func (cvideocategory *XkVideoCategoryService) UpdateXkVideoCategory(xkVideoCategory *videocategory.XkVideoCategory) (err error) {
	//1. 下面的 Save 方法是把所有的属性都给 update 了一遍的哦（具体 sql 语句可以通过设置 config.yaml 文件中的 mysql 中的 log-mode 为 debug 模式，log-zap 为 true 来看控制台输出的 sql 语句）
	//err = global.GVA_DB.Save(xkBbs).Error
	//return err

	//2. 下面的 uodate 方法是把修改的属性进行 update，没有修改的属性数不会 update 的（推荐的写法）
	err = global.GVA_DB.Model(xkVideoCategory).Updates(xkVideoCategory).Error
	return err
}

// @author: feige
// @function: DeleteXkBbs
// @description: 删除帖子
// @param: e model.DeleteXkBbs
// @return: err error
func (cvideocategory *XkVideoCategoryService) DeleteXkVideoCategory(xkVideoCategory *videocategory.XkVideoCategory) (err error) {
	err = global.GVA_DB.Delete(&xkVideoCategory).Error
	return err
}

// @author: feige
// @function: DeleteXkBbsById
// @description: 根据ID删除帖子
// @param: e model.DeleteXkBbsById
// @return: err error
func (cvideocategory *XkVideoCategoryService) DeleteXkVideoCategoryById(id uint) (err error) {
	var xkVideoCategory videocategory.XkVideoCategory
	// 注意 xk_video_category 数据表中的主键是 id 不是 id 字段
	err = global.GVA_DB.Where("id = ?", id).Delete(&xkVideoCategory).Error
	return err
}

// @author: feige
// @function: GetXkBbs
// @description: 根据ID获取帖子信息
// @param: id uint
// @return: xkVideoCategory*videocategory.VideoCategory, err error
func (cvideocategory *XkVideoCategoryService) GetXkVideoCategory(id uint) (xkVideoCategory *videocategory.XkVideoCategory, err error) {
	// 注意 xk_video_category 数据表中的主键是 id 不是 id 字段
	err = global.GVA_DB.Where("id = ?", id).First(&xkVideoCategory).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetCustomerInfoList
//@description: 分页获取客户列表
//@param: sysUserAuthorityID string, info request.PageInfo
//@return: list interface{}, total int64, err error

func (cvideocategory *XkVideoCategoryService) LoadXkVideoCategoryPage(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB.Model(&videocategory.XkVideoCategory{})

	var XkBbsList []videocategory.XkVideoCategory
	err = db.Count(&total).Error
	if err != nil {
		return XkBbsList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Find(&XkBbsList).Error
	}
	return XkBbsList, total, err
}

// 查询视频的分类信息
func (xkcategory *XkVideoCategoryService) FindCategories() (categories []videocategory.XkVideoCategory, err error) {
	err = global.GVA_DB.Where("status = 1 and is_delete = 0 and parent_id = 0").Find(&categories).Error
	if categories != nil && len(categories) > 0 {
		for index, category := range categories {
			childrens, _ := xkcategory.findCategoresChildren(category.ID)
			if childrens != nil {
				category.Children = childrens
			}
			// 记得放回去
			categories[index] = category
		}
	}
	return categories, err
}

// 查询视频的分类信息
func (xkcategory *XkVideoCategoryService) findCategoresChildren(parentId uint) (categories []*videocategory.XkVideoCategory, err error) {
	err = global.GVA_DB.Where("status = 1 and is_delete = 0 and parent_id = ?", parentId).Find(&categories).Error
	return categories, err
}

// 查询视频的分类信息
func (xkcategory *XkVideoCategoryService) dataTransfer() {
	//?
}

```



### 3.router/videocategory/xk_video_category_router.go

并写好对应的 enter.go 文件

并设置进去 initialize/router.go 内，注入路由`xkVideoCategoryRouter.InitXkVideoCategoryRouter(PublicGroup)`

```go
package videocategoryrouter

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type XkVideoCategoryRouter struct{}

func (e *XkVideoCategoryRouter) InitXkVideoCategoryRouter(Router *gin.RouterGroup) {
	xkVideoCategoryApi := v1.ApiGroupApp.VideoCategoryApiGroup.XkVideoCategoryApi

	// 这个路由多了一个对put和post请求的中间件处理，而这个中间件做了一些对post和put的参数的处理和一些公共信息的处理
	//xkVideoCategoryCustomerRouterWithoutRecord := Router.Group("videocategory").Use(middleware.OperationRecord())
	xkVideoCategoryCustomerRouterWithoutRecord := Router.Group("videocategory") // 上面的一行代码街上了中间件的处理，这里我们是不加上的，因为加上了可能是会有问题的
	{
		xkVideoCategoryCustomerRouterWithoutRecord.POST("save", xkVideoCategoryApi.CreateXkVideoCategory)   // 添加
		xkVideoCategoryCustomerRouterWithoutRecord.POST("update", xkVideoCategoryApi.UpdateXkVideoCategory) // 更新
		xkVideoCategoryCustomerRouterWithoutRecord.DELETE("delete/:id", xkVideoCategoryApi.DeleteById)      // 删除
		xkVideoCategoryCustomerRouterWithoutRecord.POST("page", xkVideoCategoryApi.LoadXkVideoCategoryPage) // 分页查询
	}

	// 这个路由是没有中间件的路由
	xkVideoCategoryRouterWithoutRecord := Router.Group("videocategory")
	{
		//对外暴露的接口 http://localhost:8888/videocategory/get?id=123
		xkVideoCategoryRouterWithoutRecord.GET("get", xkVideoCategoryApi.GetXkVideoCategory) // 查询

		xkVideoCategoryRouterWithoutRecord.GET("getdetail/:id", xkVideoCategoryApi.GetXkVideoCategoryDetail) // 获取单一客户信息

		xkVideoCategoryRouterWithoutRecord.GET("find", xkVideoCategoryApi.FindCategories) // find
	}
}

```

### 4.api/v1/videocategory/xk_videocategory_controller.go

并写好对应的 enter.go 文件

```go
package videocategoryapi

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/videocategory"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// 定义api接口
type XkVideoCategoryApi struct{}

// 添加数据到数据库
func (e *XkVideoCategoryApi) CreateXkVideoCategory(c *gin.Context) {
	// 1: 第一件事情就准备数据的载体
	var xkVideoCategory videocategory.XkVideoCategory
	err := c.ShouldBindJSON(&xkVideoCategory)
	if err != nil {
		// 如果参数注入失败或者出错就返回接口调用这。出错了.
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 创建实例，保存帖子
	//xkVideoCategoryService := new(service.videocategory.XkVideoCategoryService) // error：server/service 拿不到 server/service/videocategory 中的内容
	//xkVideoCategoryService := new(videocategory2.XkVideoCategoryService) // 得 import server/service/videocategory
	err = xkVideoCategoryService.CreateXkVideoCategory(&xkVideoCategory)
	// 如果保存失败。就返回创建失败的提升
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	// 如果保存成功，就返回创建创建成功
	response.OkWithMessage("创建成功", c)
}

// 更新修改数据库中的数据
func (e *XkVideoCategoryApi) UpdateXkVideoCategory(c *gin.Context) {
	// 1: 第一件事情就准备数据的载体
	var xkVideoCategory videocategory.XkVideoCategory
	err := c.ShouldBindJSON(&xkVideoCategory)
	if err != nil {
		// 如果参数注入失败或者出错就返回接口调用这。出错了.
		response.FailWithMessage(err.Error(), c)
		return
	}

	//xkVideoCategoryService := new(videocategory2.XkVideoCategoryService) // 得 import server/service/videocategory
	err = xkVideoCategoryService.UpdateXkVideoCategory(&xkVideoCategory)
	// 如果保存失败。就返回创建失败的提升
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	// 如果保存成功，就返回创建创建成功
	response.OkWithMessage("更新成功", c)
}

// GetXkVideoCategory
//
//	@Tags		GetXkVideoCategory
//	@Summary	根据ID查询帖子明细
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Param		data	query		videocategory.GetXkVideoCategory													true	"客户ID"
//	@Success	200		{object}	response.Response{data=exampleRes.ExaCustomerResponse,msg=string}	"获取单一客户信息,返回包括客户详情"
//	@Router		/videocategory/get?id=1 [get]
//
// 查询数据库中的数据
func (e *XkVideoCategoryApi) GetXkVideoCategory(c *gin.Context) {
	var xkVideoCategory videocategory.XkVideoCategory
	// 绑定参数
	err := c.ShouldBindQuery(&xkVideoCategory)
	// 如果参数没有直接报错
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	data, err := xkVideoCategoryService.GetXkVideoCategory(xkVideoCategory.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	response.OkWithDetailed(data, "获取成功", c)
}

// 查询数据库中的数据
func (e *XkVideoCategoryApi) GetXkVideoCategoryDetail(c *gin.Context) {
	// 绑定参数用来获取/:id这个方式
	id := c.Param("id")

	// 这个是用来获取?age=123
	//age := c.Query("age")

	parseUint, _ := strconv.ParseUint(id, 10, 64)
	data, err := xkVideoCategoryService.GetXkVideoCategory(uint(parseUint))
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(data, "获取成功", c)
}

// http://localhost:8888/videocategory/delete/:id
func (e *XkVideoCategoryApi) DeleteById(c *gin.Context) {
	// 绑定参数用来获取/:id这个方式
	id := c.Param("id")
	// 开始执行
	parseUint, _ := strconv.ParseUint(id, 10, 64)

	//xkVideoCategoryService := new(videocategory2.XkVideoCategoryService) // 得 import server/service/videocategory
	err := xkVideoCategoryService.DeleteXkVideoCategoryById(uint(parseUint))
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithDetailed("ok", "获取成功", c)
}

// LoadXkVideoCategoryPage
func (e *XkVideoCategoryApi) LoadXkVideoCategoryPage(c *gin.Context) {
	// 创建一个分页对象
	var pageInfo request.PageInfo
	// 把前端json的参数传入给PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	//xkVideoCategoryService := new(videocategory2.XkVideoCategoryService) // 得 import server/service/videocategory
	xkVideoCategoryPage, total, err := xkVideoCategoryService.LoadXkVideoCategoryPage(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     xkVideoCategoryPage,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// 查询视频分类信息
func (e *XkVideoCategoryApi) FindCategories(c *gin.Context) {
	catgories, err := xkVideoCategoryService.FindCategories()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(catgories, "获取成功", c)
}

```





## 2. 设置前端的vue页面和js函数

### 1.mange.vue 视频列表页面

使用的是 element-plus 中的 table 组件类型中的 树形数据与懒加载（tree 形列表）

1. 注意el-table 组件中的 `:tree-props="{ children: 'childrens' }"` 属性的使用意义

   因为后端 model 设置的时候我们的 `Children []*XkVideoCategory gorm:"-" json:"childrens"` 设置的 json 为 childrens 注意有一个 s，但是官方的是没有 s 的，所以 :tree-props 属性的意义是把数据库中的 childrens 替代官方组件中的children。

   我们也是可以修改后端设置 model 时改为 json:children 来解决的。但是可能会因为后端可能是外包出去的内容，是得我们不方便修改后端的内容，所以我们可以在前端组件设置一下。

2. 其次就是配合 videocategory.js 调试前端页面

```js
<template>
    <div>
        <el-table :data="videoCategoryList" row-key="id" border lazy :load="load" :tree-props="{ children: 'childrens' }">
            <el-table-column prop="description" label="Description" width="150" />
            <el-table-column prop="categoryName" label="CategoryName" width="150" />
            <el-table-column prop="createTime" label="CreateTime" width="210" />
            <el-table-column prop="updatedTime" label="UpdatedTime" width="210" />
            <el-table-column label="Status" width="100">
                <template #default="scope">
                    <span v-if="scope.row.status == 1" style="color:green">已发布</span>
                    <span v-if="scope.row.status == 0" style="color:red">未发布</span>
                </template>
            </el-table-column>
            <el-table-column label="Sorted">
                <template #default="scope">
                    <el-input v-model="scope.row.sorted"></el-input>
                </template>
            </el-table-column>
            <el-table-column label="操作">
                <template #default="scope">
                    <el-button size="small" type="primary" @click="handleEdit(scope.$index, scope.row)"><el-icon>
                            <Edit />
                        </el-icon>编辑</el-button>
                    <el-button size="small" type="danger" @click="handleDelete(scope.$index, scope.row)"
                        icon="Delete">删除</el-button>
                </template>
            </el-table-column>
        </el-table>

        <!-- 分类添加和编辑 :close-on-press-escape="false" :close-on-click-modal="false" -->
        <el-dialog v-model="dialogVisible" title="添加分类" width="600px" :before-close="handleClose">
            <el-form ref="categoryFormRef" :model="category" :rules="rules" label-width="120px" class="demo-ruleForm"
                :size="formSize" status-icon>
                <el-form-item label="分类名称" prop="categoryName">
                    <el-input v-model="category.categoryName" placeholder="请输入分类名称" maxlength="50" />
                </el-form-item>
                <el-form-item label="分类描述" prop="description">
                    <el-input v-model="category.description" type="textarea" placeholder="请输入分类描述" maxlength="50" />
                </el-form-item>
                <el-form-item label="分类排序" prop="sorted">
                    <el-input v-model="category.sorted" type="number" placeholder="请输入分类排序" maxlength="11" />
                </el-form-item>
                <el-form-item label="发布状态" prop="status">
                    <el-radio-group v-model="category.status" class="ml-4">
                        <el-radio :label="1">发布</el-radio>
                        <el-radio :label="0">未发布</el-radio>
                    </el-radio-group>
                </el-form-item>
                <el-form-item label="删除状态" prop="isDelete">
                    <el-radio-group v-model="category.isDelete" class="ml-4">
                        <el-radio :label="1">已删除</el-radio>
                        <el-radio :label="0">未删除</el-radio>
                    </el-radio-group>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="handleSubmit"><el-icon>
                            <Check />
                        </el-icon>保存</el-button>
                    <el-button @click="handleCloseDialog"><el-icon>
                            <Close />
                        </el-icon>关闭</el-button>
                </el-form-item>
            </el-form>
        </el-dialog>
    </div>
</template>

<script setup>
import { loadVideoCategoryData, findVideoCategory, saveVideoCategory, delVideoCategory, updateVideoCategory } from '@/api/videomange.js'
import { onMounted, reactive, ref } from 'vue'
import Ksd from '@/utils/index.js'
import { ElMessageBox } from 'element-plus'

let category = reactive({
    categoryName: '',
    description: '',
    parentId: 0,
    sorted: 1,
    status: 1,
    isDelete: 0,
})
const videoCategoryList = ref([]) // video 分类列表，注意使用的时候需要 .value 的哦（使用 reactive 数据类型好型是有一些问题的哦）

const dialogVisible = ref(false)
const categoryFormRef = ref(null)
// 点击添加和编辑对话框
const handleEdit = (index, row) => {
    // 这里有一个疑问？
    // category = row // error: 交互的时候有问题的，因为在点击编辑的时候，假如修改 categoryName, 但是我们关闭不提交修改的时候，我们得刷新一下页面才能看到原来的 categoryName 的数据
    // 可以使用下面的两种方式解决的哦
    // category = { ...row }
    Object.assign(category, row)
    // 数据回填form
    dialogVisible.value = true
}
// 交换关闭
const handleClose = () => {
    ElMessageBox.confirm(
        '你确定要离开吗？',
        '提示',
        {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning',
        }
    ).then(() => {
        handleCloseDialog()
    })
}
// 关闭添加和编辑对话框
const handleCloseDialog = () => {
    dialogVisible.value = false
}
// 保存和编辑
const handleSubmit = async () => {
    await categoryFormRef.value.validate(async (valid, fields) => {
        if (valid) {
            let res
            let message = '添加成功'
            category.sorted = parseInt(category.sorted)
            if (category.id) {
                message = '编辑成功'
                res = await updateVideoCategory(category)
            } else {
                res = await saveVideoCategory(category)
            }

            if (res.code === 0) {
                // 关闭当前添加窗口
                handleCloseDialog()
                // 重置model对象
                categoryFormRef.value.resetFields()
                categoryFormRef.value.clearValidate()
                // 添加刷新当前页
                handleLoadData()
                Ksd.success(message)
            }
        }
    })
}

// 交换关闭
const handleDelete = async (index, row) => {
    const result = await Ksd.confirm('提示', '你确定要抛弃我吗？', { cbtn: '容我想想', sbtn: '直接删了', type: 'error' })
    if (result === 'confirm') {
        const res = await delVideoCategory(row.id)
        if (res.data === 'ok') {
            Ksd.success('删除成功!!!')
            // 第一种方案：单删行，不刷新 (删除最后一条要记得删除，可能还要同步总数)
            // pageInfo.resultList.splice(index, 1)
            // pageInfo.total --
            // if( pageInfo.resultList.length === 0){
            //  handleLoadData()
            // }
            // 第二种方案：直接调用查询方法(推荐)
            handleLoadData()
        }
    }
}

// 这里是查询-------------------------------------------------

// 1: 设置请求分页的参数
const pageInfo = reactive({
    page: 1,
    pageSize: 10,
    total: 0,
    status: -1,
    resultList: [],
    keyword: ''
})
// 获取删除的ids
const ids = ref([])
const loading = ref(false)

// 2: 数据加载
const handleLoadData = async () => {
    loading.value = true
    const res = await loadVideoCategoryData(pageInfo)
    loading.value = false
    pageInfo.total = res.data.total
    pageInfo.resultList = res.data.list
}

// 查询所有有效的分类信息
const handleLoadCategoryData = async () => {
    const resposne = await findVideoCategory()
    videoCategoryList.value = resposne.data
}

// 3: 生命周期初始化数据加载
onMounted(() => {
    handleLoadData()
    handleLoadCategoryData()
})
</script >
```



### 3.videocategory.js 请求视频列表数据的函数

注意 /videocategory/find 在后端的具体方法的实现（递归内容）

```js
import service from '@/utils/request'

// 获取 video 一级分类
// export const findVideoCategory = (params) => {
//   return service({
//     url: '/videocategory/find',
//     method: 'get',
//     params
//   })
// }
// 下面的写法也是可以的哦
export const findVideoCategory = () => {
  return service({
    url: '/videocategory/find',
    method: 'get',
  })
}

export const loadVideoCategoryData = (data) => {
  return service({
    url: '/videocategory/page',
    method: 'post',
    data,
  })
}

export const saveVideoCategory = (data) => {
  return service({
    url: '/videocategory/save',
    method: 'post',
    data,
  })
}

export const updateVideoCategory = (data) => {
  return service({
    url: '/videocategory/update',
    method: 'post',
    data,
  })
}

// 根据id删除video
export const delVideoCategory = (id) => {
  return service({
    url: `/videocategory/delete/${id}`,
    method: 'delete',
  })
}
```

