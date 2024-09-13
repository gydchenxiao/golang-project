## 1. 处理之前帖子分类页面中状态修改不成功的问题

### 1.1后端专们给这个需求做了一个 model

注意为什么要重新设置一个 model，因为前端请求的参数可能不会有数据表的所有字段的值，所以专门定制一个接口（推荐的做法），当然也是可以在更新的时候先根据 ID，查一下，在使用 utils 中的函数，把前端请求的参数覆盖刚刚查询出来的 ID 数据。

```go
package req

type StatusReq struct {
	ID    uint   `json:"id"`
	Value int8   `json:"value"` // 前端传递过来的需要更新的 status / isDelete 的值
	Field string `json:"field"`
}
```

### 1.2并且专门提供了一个 service 方法

注意下面 Update 的写法

```go
// 修改 status / isDeleted, 量身定制的更新状态的方法
func (cbbs *BBSCategoryService) UpdateBbsCategoryStatus(statusReq *req.StatusReq) (err error) {
	//
	err = global.GVA_DB.Model(new(*bbs.BBSCategory)).Where("id=?", statusReq.ID).Update(statusReq.Field, statusReq.Value).Error
	return err
}
```

### 1.3配置路由 router

```go
// 量身定制的一个更新状态的方法
		xkBbsCustomerRouterWithoutRecord.POST("updateStatus", bbsCategoryApi.UpdateBBSCategoryStatus)
```

### 1.4提供 api 方法

```go
// 量身定制的一个更新状态的方法
func (e *BBSCategoryApi) UpdateBBSCategoryStatus(c *gin.Context) {
	// 1: 第一件事情就准备数据的载体
	var statusReq req.StatusReq
	err := c.ShouldBindJSON(&statusReq) // post 请求
	if err != nil {
		// 如果参数注入失败或者出错就返回接口调用这。出错了.
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = bbsCatgoryService.UpdateBbsCategoryStatus(&statusReq) // update status 的量身定制的方法
	// 如果保存失败。就返回创建失败的提升
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	// 如果保存成功，就返回创建创建成功
	response.OkWithMessage("更新成功", c)
}
```

下面的这个方法就是先根据 ID 查询数据出来，再使用了一个 utils 中的函数把前端请求的参数值覆盖进去，使得更新的接口统一：

```go
func (e *BBSCategoryApi) UpdateBBSCategory(c *gin.Context) {
	// 1: 第一件事情就准备数据的载体
	var bbsReq req.BbsCategorySaveReq
	err := c.ShouldBindJSON(&bbsReq) // post 请求
	if err != nil {
		// 如果参数注入失败或者出错就返回接口调用这。出错了.
		response.FailWithMessage(err.Error(), c)
		return
	}
	
	// 上面加了一层 req.BbsCategorySaveReq, 因为前端不一定回返回我们设置的 model 的全部的字段的
	bbsCategory, err := bbsCatgoryService.GetBBSCategory(bbsReq.ID)
	utils.CopyProperties(bbsCategory, bbsReq)

	err = bbsCatgoryService.UpdateBBSCategory(bbsCategory)
	// 如果保存失败。就返回创建失败的提升
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	// 如果保存成功，就返回创建创建成功
	response.OkWithMessage("更新成功", c)
}
```



### 1.5前端直接使用这个接口即可

```js
// 量身定制的更新状态的函数
export const updateBbsCategoryStatus = (data) => {
  return service({
    url: '/bbscategory/updateStatus',
    method: 'post',
    data,
  })
}
```



## 2. 视频列表添加视频页面的处理

注意下面各组件之间的数据传递：v-model 配合 defineEmits 实现父组件拿到子组件的数据(也是可以通过调用子组件的函数来得到的)，ref (实现父组件调用子组件的函数)

### 2.1view/vido/mange.vue 把组件封装出去

```vue
<template>
    <div>
        <el-button type="primary" size="default" @click="handleAddVideo">添加视频</el-button>
        添加视频
        <add-video ref="addVideoRef"></add-video>
    </div>
</template>
<script setup>
import AddVideo from './components/AddVideo.vue'
import { ref } from "vue";
// 获取到子组件addvideo的对象
const addVideoRef = ref({})

// 添加视频
const handleAddVideo = () => {
    addVideoRef.value.handleOpen() // 父组件调用子组件的函数
}
</script>
```

### 2.2view/vido/components/AddVideo.vue

```vue
<template>
  <!---添加和编辑视频 -->
  <el-dialog v-model="dialogVisible" title="添加课程" top="1vh" width="90%" :before-close="handleClose">
    <el-form ref="categoryFormRef" :model="video" :rules="rules" label-width="120px" class="demo-ruleForm"
      :size="formSize" status-icon>
      <el-form-item label="课程名称" prop="title">
        <el-input v-model="video.title" placeholder="请输入课程名称" maxlength="50" />
      </el-form-item>
      <el-form-item label="课程描述" prop="description">
        <el-input v-model="video.description" type="textarea" placeholder="请输入课程描述" maxlength="50" />
      </el-form-item>
      <el-form-item label="课程排序" prop="sorted">
        <el-input v-model="video.sorted" type="number" placeholder="请输入课程排序" maxlength="11" />
      </el-form-item>
      <el-form-item label="发布状态" prop="status">
        <el-radio-group v-model="video.status" class="ml-4">
          <el-radio :label="1">发布</el-radio>
          <el-radio :label="0">未发布</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="删除状态" prop="isDelete">
        <el-radio-group v-model="video.isDelete" class="ml-4">
          <el-radio :label="1">已删除</el-radio>
          <el-radio :label="0">未删除</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="视频描述">
        <ksd-editor v-model="video.content"></ksd-editor>
      </el-form-item>
      {{ video }}
      <el-form-item>
        <el-button type="primary" @click="handleSubmit"><el-icon>
            <Check />
          </el-icon>保存</el-button>
        <el-button @click="handleClose"><el-icon>
            <Close />
          </el-icon>关闭</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>

<script setup>
import { reactive, ref } from 'vue'
import KsdEditor from '@/components/editor/KsdEditor.vue'

// 让页面子组件显示出来
const dialogVisible = ref(false)
const video = reactive({
  content: ''
})

// 打开
const handleOpen = () => {
  dialogVisible.value = true
}

// 关闭
const handleClose = () => {
  dialogVisible.value = false
}

// 子组件默认包含是私有
defineExpose({
  handleOpen,
  handleClose
})
</script>
```

### 2.3@/components/editor/KsdEditor.vue

```vue
<template>
    <div style="border: 1px solid #ccc">
        <Toolbar style="border-bottom: 1px solid #ccc" :editor="editorRef" :defaultConfig="toolbarConfig" :mode="mode" />
        <Editor style="height: 500px; overflow-y: hidden;" :defaultConfig="editorConfig" :mode="mode"
            @onChange="handleChange" @onCreated="handleCreated" />
    </div>
</template>

<script setup>
import '@wangeditor/editor/dist/css/style.css' // 引入 css
import { shallowRef } from 'vue'
import { Editor, Toolbar } from '@wangeditor/editor-for-vue'
// 编辑器实例，必须用 shallowRef
const editorRef = shallowRef()
const toolbarConfig = {}
const editorConfig = { placeholder: '请输入内容...' }

// 自定义v-model属性的时候处理
defineProps("[modelValue]")
const emits = defineEmits(["update:modelValue"])
const handleChange = (editor) => {
    emits("update:modelValue", editor.getHtml())
}


const handleCreated = (editor) => {
    editorRef.value = editor // 记录 editor 实例，重要！
}

</script>    
```

