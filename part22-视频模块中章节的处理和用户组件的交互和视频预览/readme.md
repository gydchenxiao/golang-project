## 1. element分页组件发两次请求的问题



`@/api/user.js`

```js
// 查询所有的作者信息
export const findUserAuthors = (data)=>{
  return service({
    url: '/user/find/authors',
    method: 'post',
    data
  })
}
```



vue 页面组件

```vue
<template>
  <!---添加和编辑视频 -->
  <el-dialog v-model="dialogVisible" title="添加课程" top="10vh" width="640px" :before-close="handleClose">
    <div>
      <el-input v-model="pageInfo.keyword" @blur="handleSearch"></el-input>
      <el-table v-loading="tableLoading" :data="tableData" @selection-change="handleSelectionChange"
        height="calc(100vh - 280px)" style="width: 100%">
        <el-table-column label="ID" prop="id" align="center"></el-table-column>
        <el-table-column label="作者">
          <template #default="scope">
            <div style="display: flex;align-items: center;">
              <el-avatar :src="scope.row.avatar" />
              <span style="margin-left: 10px;">{{ scope.row.author_name }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column label="昵称" prop="nickname"></el-table-column>
        <el-table-column label="账号" prop="account"></el-table-column>
        <el-table-column width="200" label="操作">
          <template #default="scope">
            <el-button size="small" type="primary" @click="handleSelect(scope.$index)"><el-icon>
                <Edit />
              </el-icon>选择</el-button>
          </template>
        </el-table-column>
      </el-table>

      {{ pageInfo }}
      <div class="pagination" style="display:flex;justify-content: center;">
        <el-pagination :current-page="pageInfo.pageNum" :page-size="pageInfo.pageSize"
          :page-sizes="[2, 10, 20, 30, 50, 80, 100]" layout="total, sizes, prev, pager, next, jumper"
          :total="pageInfo.total" @size-change="handleSizeChange" @current-change="handleCurrentChange" />
      </div>
    </div>
  </el-dialog>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { findUserAuthors } from '@/api/user.js'
const pageInfo = ref({
  pageNum: 1,
  pageSize: 2,
  total: 0,
  keyword: ""
})
const dialogVisible = ref(true)
const tableLoading = ref(false)
const tableData = ref([])
const emits = defineEmits(["select"])

// 查询用户作者信息
const handleLoadData = async () => {
  tableLoading.value = true
  // console.log(pageInfo)
  const resp = await findUserAuthors(pageInfo.value)
  console.log('resp', resp)
  tableData.value = resp.data.list
  pageInfo.value.pageNum = resp.data.pageNum
  pageInfo.value.pageSize = resp.data.pageSize
  if (resp.data.total != 0)
    pageInfo.value.total = resp.data.total
  tableLoading.value = false
}

// 搜索
const handleSearch = () => {
  pageInfo.value.pageNum = 1
  handleLoadData()
}

// 分页改变
const handleCurrentChange = (pageNum) => {
  // alert(pageNum) // 上面使用的时候需要判断 resp.data.total 是否为零的哦。但是其他的页面好像是没有判断，不管了，碰到就先解决了
  pageInfo.value.pageNum = pageNum
  handleLoadData();
}
// 改变分页大小
const handleSizeChange = (pageSize) => {
  pageInfo.value.pageSize = pageSize
  handleLoadData();
}

// 选择用户
const handleSelect = (index) => {
  emits("select", tableData.value[index])
  handleClose();
}

// 打开用户弹窗组件
const handleOpen = () => {
  dialogVisible.value = true;
  // 选择用户按钮的才来加载 / 点击分页的时候
  handleLoadData()
}
// 关闭用户弹窗组件
const handleClose = () => {
  dialogVisible.value = false;
}

// 子组件默认包含是私有
defineExpose({
  handleOpen,
  handleClose
})

</script>

<style lang="scss" scoped>
.imgbox-up {
  width: 200px;
  height: 140px;
  margin-top: 10px;
  background: #fafafa;
  border: 1px solid #eee;
  margin-top: 10px;
  display: flex;
  margin-right: 10px;
  cursor: pointer;
  flex-direction: column;
  align-items: center;
  font-size: 32px;
  justify-content: center;
  color: #eee;

  .info {
    font-size: 12px;
    color: #999;
  }
}

.imgbox-up:hover {
  background: #ccc;
}

.ksd-mx-1 {
  margin-right: 5px;
}

.ksd-taglist {
  position: absolute;
  right: 0;
  top: 35px
}
</style>>
```

**首先分页组件的页数是由total和size决定**。

通过后端返回数据进行赋值，比如现在pageSize是2,total是4。
那么现在的页数是两页。

而切换到第二页的时候发送两次请求，并且页数切换到了1。

因为total或pageSize的变化而导致页数变成了一页从而导致这种情况发生，组件给你切换到了第一页然后再次触发了change事件，而发送请求在change中，所以出现了发送两次请求。

我遇到的情况是，pageNum为1时后端传的total正常，**而pageNum切换到2时后端返回total为0（用Apifox测试一下接口）**，组件给你变成1页并且切到1页，啥也不说了，先打一顿后端。



## 2. 但是其他的页面使用element-plus分页组件时没有这种情况

碰到一个先解决一个