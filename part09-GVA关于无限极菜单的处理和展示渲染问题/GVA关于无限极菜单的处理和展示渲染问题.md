## 1. 关于前端用树形结构展示数据el-tree组件

### 1.1view/vido/category.vue

注意组件属性的使用，以及用户体验

```vue
<template>
  <div style="padding: 20px;background: #fff;">
    <!-- check-strictly -->
    <!-- :default-checked-keys="checkedIds" -->
    <el-tree :data="categoryList" show-checkbox node-key="id" default-expand-all draggable ref="categoryTreeRef"
      @node-click="handleClick" :expand-on-click-node="false">
      <template #default="{ node, data }">
        <div class="custom-tree-node">
          <span>{{ data.categoryName }}</span>
          <span>
            <a @click="handleEdit(node, data)"> 修改 </a>
            <a style="margin-left: 8px" @click="handleDel(node, data)"> 删除 </a>
          </span>
        </div>
      </template>
    </el-tree>
    <el-button @click="handleSubmit">确定</el-button>
  </div>
</template>

<script setup>
import { onMounted, ref } from "vue"
import { loadCategory } from "@/api/videocategory.js"

// 定义接受分类数据的容器--响应式
var categoryList = ref([])
const checkedIds = ref([])
const categoryTreeRef = ref({})

// 定义查询分类数据的方法
const handleLoadData = async () => {
  const res = await loadCategory()
  categoryList.value = res.data
  // 把所有的节点都选中
  checkedIds.value = res.data.map(data => data.id)
  // 使用method的方式来选择节点
  categoryTreeRef.value.setCheckedKeys(checkedIds.value)
}

// 提交方法
const handleSubmit = () => {
  const allSelectIds = categoryTreeRef.value.getCheckedKeys()
  console.log('allSelectIds', allSelectIds)
}

//删除
const handleDel = (node, data) => {
  console.log("你要删除的id是", node, data)
  categoryTreeRef.value.remove(node, data)
}

//编辑
const handleEdit = (node, data) => {
  console.log("编辑的信息是", node, data)
  data.categoryName = "javaxxxxxx"
}

// 点击触发 
const handleClick = (data, node) => {
  console.log(node, data)
}
// 生命周期执行函数
onMounted(() => {
  handleLoadData()
})
</script>

<style scoped>
.custom-tree-node {
  display: flex;
  justify-content: space-between;
  width: 100%;
}
</style>
```



### 1.2api/videocategory.js

异步请求后端的数据

```js
import service from '@/utils/request'

// export const loadCategory = (params) => {
//   return service({
//     url: '/videocategory/find',
//     method: 'get',
//     params
//   })
// }
// 下面的写法也是可以的哦
export const loadCategory = () => {
  return service({
    url: '/videocategory/find',
    method: 'get',
  })
}
```





## 2. 后端的处理

### 2.1model/videocategory/xk_video_categorys_model.go

因为一个节点下面可能会有多个节点，所以数据model中设置了一个 childrens 字段

```go
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
	Children []*XkVideoCategory `gorm:"-" json:"children"`
}
```

注意上面的 json:"children" 还是 json:"childrens" 的区别？
如果前端使用的是 tree 树形表结构展示的数据，那么后端可以 json:"childrens"，前端设置 :tree-props="{ children: 'childrens' }"。
如果后端是 json:"children" 前端就不用设置了。
不过 el-table 有属性 :tree-props="{ children: 'childrens' }"， el-tree 没有属性 :tree-props="{ children: 'childrens' }"。



### 2.2service/videocategory/xk_video_category_service.go

提供访问数据库的gorm语句，重要的就是多级结构的递归查询。

```go
// 1. 只能查看两级树形结构，而且写法不太好，不优雅
//// 查询视频的分类信息
//func (xkcategory *XkVideoCategoryService) FindCategories() (categories []videocategory.XkVideoCategory, err error) {
//	err = global.GVA_DB.Where("status = 1 and is_delete = 0 and parent_id = 0").Find(&categories).Error
//	if categories != nil && len(categories) > 0 {
//		for index, category := range categories {
//			childrens, _ := xkcategory.findCategoresChildren(category.ID)
//			if childrens != nil {
//				category.Children = childrens
//			}
//			// 记得放回去
//			categories[index] = category
//		}
//	}
//	return categories, err
//}
//
//// 查询视频的分类信息
//func (xkcategory *XkVideoCategoryService) findCategoresChildren(parentId uint) (categories []*videocategory.XkVideoCategory, err error) {
//	err = global.GVA_DB.Where("status = 1 and is_delete = 0 and parent_id = ?", parentId).Find(&categories).Error
//	return categories, err
//}

// 2. 两级写死的做法。只能查看两级树形结构，写法还算优雅
//func (xkcategory *XkVideoCategoryService) Tree(allDbCategoires []*videocategory.XkVideoCategory) []*videocategory.XkVideoCategory {
//	// 定义一个节点
//	//allDbCategoires =
//	//1	Java	Java	2023-06-04 20:59:19	2023-06-04 20:59:19	1	1	0	0
//	//2	Go	Go	2023-06-04 20:59:19	2023-06-04 20:59:19	1	1	0	0
//	//3	Javascript	Javascript	2023-06-04 20:59:19	2023-06-04 20:59:19	1	1	0	0
//	//4	Spring	Spring	2023-06-04 20:59:19	2023-06-04 20:59:19	1	1	0	1
//	//5	SpringBoot	SpringBoot	2023-06-04 20:59:19	2023-06-04 20:59:19	1	1	0	1
//	//6	Gin	Gin	2023-06-04 20:59:19	2023-06-04 20:59:19	1	1	0	2
//	//7	Beego	Beego	2023-06-04 20:59:19	2023-06-04 20:59:19	1	1	0	2
//	//8	XOrm	XOrm	2023-06-04 20:59:19	2023-06-04 20:59:19	1	1	0	2
//	//9	Gorm	Gorm	2023-06-04 20:59:19	2023-06-04 20:59:19	1	1	0	2
//	//10	GVA	GVA	2023-06-04 20:59:19	2023-06-04 20:59:19	1	1	0	2
//
//	//nodes
//	//1	Java	Java	2023-06-04 20:59:19	2023-06-04 20:59:19	1	1	0	0 ---parentNode-XkVideoCategory
//	//2	Go	Go	2023-06-04 20:59:19	2023-06-04 20:59:19	1	1	0	0
//	//3	Javascript	Javascript	2023-06-04 20:59:19	2023-06-04 20:59:19	1	1	0	0
//
//	var nodes []*videocategory.XkVideoCategory //---------准备空教室
//	for _, dbCategory := range allDbCategoires {
//		if dbCategory.ParentId == 0 {
//			// 这里找到所有的父类
//			nodes = append(nodes, dbCategory)
//		}
//	}
//
//	// 开始遍历父类
//	for _, dbCategory := range allDbCategoires {
//		for _, parentNode := range nodes {
//			if dbCategory.ParentId == parentNode.ID {
//				parentNode.Children = append(parentNode.Children, dbCategory)
//			}
//		}
//	}
//	return nodes
//}

// 3. 多级表也是可以的（递归）
func (xkcategory *XkVideoCategoryService) FindCategories() (categories []*videocategory.XkVideoCategory, err error) {
	err = global.GVA_DB.Where("status = 1 and is_delete = 0").Find(&categories).Error
	return categories, err
}
func (xkcategory *XkVideoCategoryService) Tree(allDbCategoires []*videocategory.XkVideoCategory, parentId uint) []*videocategory.XkVideoCategory {
	var nodes []*videocategory.XkVideoCategory //---------准备空教室
	// 开始遍历父类
	for _, dbCategory := range allDbCategoires { //1 parentId = 0 parentId=0 2 3 4 5 6 7 8 9 10
		if dbCategory.ParentId == parentId {
			dbCategory.Children = append(dbCategory.Children, xkcategory.Tree(allDbCategoires, dbCategory.ID)...)
			nodes = append(nodes, dbCategory)
		}
	}
	return nodes
}
```



### 2.3router/videocategory/xk_video_category_router.go

先用xkVideoCategoryService.FindCategories() 查出所有的数据catgories，再用xkVideoCategoryService.Tree(catgories, 0)处理刚刚查出所有的数据catgories。

```go
// 查询视频分类信息
func (e *XkVideoCategoryApi) FindCategories(c *gin.Context) {
	catgories, err := xkVideoCategoryService.FindCategories()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	//response.OkWithDetailed(catgories, "获取成功", c) // 只能查看两级树形结构
	response.OkWithDetailed(xkVideoCategoryService.Tree(catgories, 0), "获取成功", c) // 能够查看多级树形结构
}
```

