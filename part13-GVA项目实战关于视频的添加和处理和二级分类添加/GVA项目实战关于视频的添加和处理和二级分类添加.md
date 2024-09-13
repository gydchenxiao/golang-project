## 1. 关于视频分类页面二级结构的功能实现

我们的视频页面是一个 el-table 组件展示的二级树形界面，控制二级元素隐藏(或者设置为不可选中)添加子分类的按钮来实现只有二级结构的功能。

### 1.1功能一：添加一级子分类

![image-20240304203006041](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20240304203006041.png)

功能要求：

1. 标题设置为添加父分类。

   在 el-dialog 标签上设置 :title="categoryTitle"（title属性），并设置 var categoryTitle = ref("添加父分类")，在后面不同的场景下变化即可。注意 ref 类型的数据使用的时候得 .value 的哦。

2. 可见父分类的展示为跟分类，实际就是0 ---> 在查询所有的主分类categoryList最前面unshift一个根分类即可

   ```js
   categoryList.value.unshift({ id: 0, categoryName: "根分类" })
   ```

2. 分类名称是必须要填写的，得通过做校验的方式来控制。

   ```js
   // 调用校验规则
   const rules = reactive({
     parentId: [
       { required: true, message: '请选择分类', trigger: 'blur' }
     ],
     categoryName: [
       { required: true, message: '请输入分类名称', trigger: 'blur' }
     ]
   })
   ```

3. 分类描述是可写可不写的。

4. 剩下的分类排序(注意数字不能使得被转化成string类型了哦，以及数据库中设置的数据范围是否足够的问题)，发布状态和删除状态都是有默认值的。也是可以手动填写的。



### 1.2功能二：添加子分类

可以看到下面父分类是不可选的，在当前一级元素(父分类)下添加子分类所以，新添加的子分类的父分类就是当前元素的ID。其他的都是和上面的内容一致的。

![image-20240304210044335](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20240304210044335.png)



### 1.3功能三：编辑（子分类/父分类）

实现下面两个的功能的时候：

1. 注意编辑父分类的时候是不能修改父分类的；
2. 编辑子分类的时候，编辑的父分类是可以移动当前元素在其他元素下面的，注意下面的可选项都是一级元素(父分类的哦，也是为了不能移动到二级元素下面打破只有二级结构的哦)。

#### 1.3.1编辑父分类

![image-20240304204837142](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20240304204837142.png)



#### 1.3.2编辑子分类

![image-20240304205843892](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20240304205843892.png)



## 2. 去掉vscode中vue文件中的红色警告线

![image-20240304093303326](D:\golang-project\golang-project\src\gin-vue-admin-readme\part13-GVA项目实战关于视频的添加和处理和二级分类添加\image-20240304093303326.png)



## 3. 关于上面的三个功能使用的组件

上面三个功能的实现都是使用的同一个组件在 gin-vue-admin\web\src\view\vido\category.vue 视频分类页面中使用的哦。

gin-vue-admin\web\src\view\vido\category.vue：

```vue
<template>
  <div style="padding: 20px;background: #fff;">
    <!-- 头部的添加，删除，搜索区域 -->
    <div class="ksd-search-container">
      <el-button type="success" icon="Plus" @click="handleOpenAdd">添加一级分类</el-button>
      <el-button type="danger" icon="Delete" @click="handleBatchDel">批量删除</el-button>

      <el-select v-model="searchForm.status" @change="handleSearch" @clear="searchForm.status = -1"
        style="margin-left: 10px;" clearable class="m-2" placeholder="请选择帖子状态" size="large">
        <el-option v-for="item in statusOptions" :key="item.value" :label="item.label" :value="item.value" />
      </el-select>

      <el-input v-model="searchForm.keyword" clearable class="w-50" size="small"
        :input-style="{ 'width': '240px', 'margin': '0 10px' }" placeholder="请输入分类名称" :suffix-icon="Search" />

      <el-button type="success" style="margin-left: 40px;" @click="handleSearch" icon="Search">搜索</el-button>

    </div>

    <!-- 数据的内容展示区域 -->
    <el-table :data="categoryList" style="margin-bottom: 20px" row-key="id" border :default-expand-all="false">
      <el-table-column prop="id" label="ID" width="主键" align="center" />
      <el-table-column prop="categoryName" label="分类名称" align="center" />
      <el-table-column label="状态" align="center">
        <template #default="scope">
          <el-switch v-model="scope.row.status" size="large" :active-value="1" :inactive-value="0"
            @change="handleChangeStatus(scope.row, 'status')" />
        </template>
      </el-table-column>
      <el-table-column label="删除状态" align="center">
        <template #default="scope">
          <el-switch v-model="scope.row.isDelete" size="large" :active-value="1" :inactive-value="0"
            @change="handleChangeStatus(scope.row, 'is_delete')" />
        </template>
      </el-table-column>
      <el-table-column label="排序" align="center">
        <template #default="scope">
          <el-input v-model="scope.row.sorted"></el-input>
        </template>
      </el-table-column>
      <el-table-column label="创建时间" align="center">
        <template #default="scope">
          {{ formatTimeToStr(scope.row.createTime, "yyyy/MM/dd hh:mm:ss") }}
        </template>
      </el-table-column>
      <el-table-column label="更新时间" align="center">
        <template #default="scope">
          {{ formatTimeToStr(scope.row.updateTime, "yyyy/MM/dd hh:mm:ss") }}
        </template>
      </el-table-column>
      <el-table-column label="操作" align="right" width="300px">
        <template #default="scope">
          <!-- （子是一个二级结构）:disabled="scope.row.parentId != 0" 的作用：如果是二级标签的话，就不提供添加子分类的功能，就把这个 el-button 给不让选中了。但是我们使用的是 :style 隐藏的方式 -->
          <el-button size="small" :style="{ display: scope.row.parentId != 0 ? 'none' : '' }" type="primary"
            @click="handleAddCategory(scope.row)"><el-icon>
              <Plus />
            </el-icon>添加子分类</el-button>
          <el-button size="small" type="primary" @click="handleEdit(scope.row, scope.row.parentId == 0)"><el-icon>
              <Edit />
            </el-icon>编辑</el-button>
          <el-button size="small" type="danger" @click="handleDelete(scope.$index, scope.row)"
            icon="Delete">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <!--添加分类的组件-->
    <add-category ref="addCategoryRef" @success="handleReload"></add-category>
  </div>
</template>

<script setup>
import { onMounted, reactive, ref } from "vue"
import { loadCategory } from "@/api/videocategory.js"
import { formatTimeToStr } from '@/utils/date'
import AddCategory from '@/view/vido/components/AddCategory.vue'
import Ksd from '@/utils/index.js'

// 1. 子组件和父组件的通信处理
// ref="addCategoryRef"：获取添加子组件的对象（用来调用子组件的函数或者使用子组件的数据）
const addCategoryRef = ref(null);
// 打开添加一级分类
const handleOpenAdd = () => {
  addCategoryRef.value.handleOpen();
}
// 打开添加子分类
const handleAddCategory = (row) => {
  addCategoryRef.value.handleAddOpen(row, true);
}
// 编辑就打开
const handleEdit = (row, flag) => {
  addCategoryRef.value.handleOpen(row, flag);
}
// @success="handleReload"：添加和编辑成功的刷新方法（子组件用来调用父组件的函数）
const handleReload = () => {
  handleLoadData()
}

const handleBatchDel = () => {

}

// 2. 数据加载
// 数据的管理和搜索查询
const searchForm = reactive({
  keyword: "",
  status: "",
})

// 校验使用的
const rules = reactive({
  title: [
    { required: true, message: '请输入分类名称', trigger: 'blur' }
  ]
})

// 定义接受分类数据的容器--响应式
var categoryList = ref([])

// 定义查询分类数据的方法
const handleLoadData = async () => {
  const res = await loadCategory()
  categoryList.value = res.data
}

// 3. 生命周期执行函数
onMounted(() => {
  handleLoadData()
})
</script>

<style scoped>
.ksd-search-container {
  margin-bottom: 10px;
  background: #f7fbff;
  padding: 6px;
  display: flex;
}

.custom-tree-node {
  display: flex;
  justify-content: space-between;
  width: 100%;
}
</style>
```



gin-vue-admin\web\src\view\vido\components\AddCategory.vue:

```vue
<template>
  <!--添加主分类 -->
  <el-dialog v-model="dialogVisible" :close-on-press-escape="false" :close-on-click-modal="false" :title="categoryTitle"
    width="600px" :before-close="handleClose">
    <el-form ref="videoCategoryFormRef" :model="formData" :rules="rules" label-width="120px">
      <el-form-item label="父分类" prop="parentId">
        <!-- 如果添加的一级分类---父分类parentId=0,并且不能选中
        如果添加的子分类，====子分类的对应父分类要被选中，通过可以更换 -->
        <el-select v-model="formData.parentId" :disabled="formData.parentId == 0 || rootFlag" clearable class="m-2"
          placeholder="请选择主分类" size="large">
          <el-option v-for="item in categoryList" :key="item.id" :label="item.categoryName" :value="item.id" />
        </el-select>
      </el-form-item>
      <el-form-item label="分类名称" prop="categoryName">
        <el-input v-model="formData.categoryName" placeholder="请输入分类名称" maxlength="50" />
      </el-form-item>
      <el-form-item label="分类描述" prop="description">
        <el-input v-model="formData.description" type="textarea" placeholder="请输入分类描述" maxlength="50" />
      </el-form-item>
      <el-form-item label="分类排序" prop="sorted">
        <el-input v-model="formData.sorted" type="number" placeholder="请输入分类排序" maxlength="11" />
      </el-form-item>
      <el-form-item label="发布状态" prop="status">
        <el-radio-group v-model="formData.status" class="ml-4">
          <el-radio :label="1">发布</el-radio>
          <el-radio :label="0">未发布</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="删除状态" prop="isDelete">
        <el-radio-group v-model="formData.isDelete" class="ml-4">
          <el-radio :label="1">已删除</el-radio>
          <el-radio :label="0">未删除</el-radio>
        </el-radio-group>
      </el-form-item>
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
import { ref, reactive } from 'vue'
import { findRoot, saveVideoCategory, updateVideoCategory } from '@/api/videocategory.js'
import Ksd from '@/utils/index.js'

// 自定义事件（调用父组件的函数，实现刷新父组件/父页面的刷新）
const emits = defineEmits(["success"])

// 让页面子组件显示出来
const dialogVisible = ref(false)
const rootFlag = ref(false)
// 定义分类数据容器接收
var categoryList = ref([])
var categoryTitle = ref("添加父分类")
// 这个用于调用form表单的方法和事件，主要用于做校验和重置等相关操作
const videoCategoryFormRef = ref({});
// 定义form保存的对象数据模型-v-model
const formData = ref({
  parentId: 0,
  categoryName: '',
  description: '',
  sorted: 1,
  status: 1,
  isDelete: 0
})
// 调用校验规则
const rules = reactive({
  parentId: [
    { required: true, message: '请选择分类', trigger: 'blur' }
  ],
  categoryName: [
    { required: true, message: '请输入分类名称', trigger: 'blur' }
  ]
})

// 加载一级分类
const loadCategoriesData = async () => {
  const res = await findRoot(); // 查询所有的主分类信息
  categoryList.value = res.data;
}
// 添加父分类
const handleOpen = async (row, flag) => {
  // 加载分类
  await loadCategoriesData()
  // 添加根分类
  categoryList.value.unshift({ id: 0, categoryName: "根分类" })
  // 编辑当前的数据
  if (row) {
    // 编辑当前的数据
    formData.value = row
    // 编辑框的标题更改
    categoryTitle.value = "编辑父分类【" + row.categoryName + "】"
  } else {
    handleReset()
    formData.parentId = 0
    // 编辑框的标题更改
    categoryTitle.value = "添加父分类"
  }
  // 打开编辑框
  dialogVisible.value = true
}
const handleAddOpen = async (row, flag) => {
  // 加载分类
  await loadCategoriesData()
  rootFlag.value = flag
  // 重置成最初状态来添加子分类
  handleReset();
  // 把父亲分类选择，代表你的子分类要挂载到拿个父分类下
  formData.value.parentId = row.id;
  // 编辑框的标题更改
  categoryTitle.value = "给【" + row.categoryName + "】添加子分类"
  // 打开添加框
  dialogVisible.value = true
}

// 关闭
const handleClose = () => {
  dialogVisible.value = false
  categoryList.value = []
  handleReset();
  categoryTitle.value = "添加分类";
}

// 提交保存
const handleSubmit = async () => {
  // 进行数据校验
  await videoCategoryFormRef.value.validate(async (valid) => {
    if (valid) {
      formData.value.sorted = formData.value.sorted * 1 // sorted 在这里变成了 string 类型了，可以这样改成 number 类型
      if (formData.value.id) {
        // 执行编辑
        // console.log("formData", formData.value)
        await updateVideoCategory(formData.value)
        // 提示
        Ksd.success("更新分类【" + formData.value.categoryName + "】成功!")
      } else {
        // 执行保存
        // console.log("formData", formData.value)
        await saveVideoCategory(formData.value)
        // 提示
        Ksd.success("添加分类【" + formData.value.categoryName + "】成功!")
      }
      // 关闭方法
      handleClose();
      // 刷新父窗口的事件
      emits("success");
    }
  })
}

//重置
const handleReset = () => {
  formData.value = {
    parentId: 0,
    categoryName: '',
    description: '',
    sorted: 1,
    status: 1,
    isDelete: 0
  }
}

// 子组件默认包含是私有
defineExpose({
  handleOpen,
  handleAddOpen,
  handleClose
})
</script>

<style lang="scss" scoped></style>>
```



### 3.1关于父子组件之间的通信如何实现

```vue
    <!--添加分类的子组件-->
    <add-category ref="addCategoryRef" @success="handleReload"></add-category>
```

#### 父元素调用子元素的函数：ref="addCategoryRef"

```vue
// ref="addCategoryRef"：获取添加子组件的对象（用来调用子组件的函数或者使用子组件的数据）
const addCategoryRef = ref(null);
```

addCategoryRef 就得到了子组件的对象。子元素再把需要给父元素使用的函数放出去：

```vue
// 子组件默认包含是私有
defineExpose({
  handleOpen,
  handleAddOpen,
  handleClose
})
```

父元素即可通过获取到的子元素的对象调用了。



#### 子元素调用父元素的函数：@success="handleReload"

父元素定义给子元素调用的函数：

```vue
// @success="handleReload"：添加和编辑成功的刷新方法（子组件用来调用父组件的函数）
const handleReload = () => {
  handleLoadData()
}

// 定义查询分类数据的方法(视频分类页面展示的所有数据 ---> 刷新页面)
const handleLoadData = async () => {
  const res = await loadCategory() 
  categoryList.value = res.data
}
```

子元素再自定义事件：

```vue
// 自定义事件（调用父组件的函数，实现刷新父组件/父页面的刷新）
const emits = defineEmits(["success"])

// 调用：即可刷新父窗口的事件（跟新了数据，就不用再刷新页面也能及时看到数据。）
emits("success");
```

