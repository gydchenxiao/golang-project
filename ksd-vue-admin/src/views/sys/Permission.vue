<template>
  <div class="kva-container">
    <div class="kva-contentbox">
      <home-page-header>
        <div class="kva-form-search">
          <el-form :inline="true" ref="searchForm" :model="queryParams">
            <el-form-item>
              <el-button type="primary" icon="Plus" @click="handleAdd">添加权限</el-button>
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
        <el-table :data="tableData" style="width: 100%; margin-bottom: 20px" row-key="id" border stripe
          :height="settings.tableHeight()">
          <el-table-column fixed prop="id" label="ID" align="center" width="80" />
          <el-table-column fixed prop="parentId" label="父ID" align="center" width="80" />
          <el-table-column prop="title" label="展示名字" align="center">
            <template #default="{ row }">
              <el-input v-model="row.title" style="text-align:center" @change="handleChange(row, 'title')"></el-input>
            </template>
          </el-table-column>
          <el-table-column prop="code" label="编号" align="center">
            <template #default="{ row }">
              <el-input v-model="row.code" style="text-align:center" @change="handleChange(row, 'code')"></el-input>
            </template>
          </el-table-column>
          <el-table-column prop="code" label="访问路径" align="center">
            <template #default="{ row }">
              <el-input v-model="row.path" style="text-align:center" @change="handleChange(row, 'path')"></el-input>
            </template>
          </el-table-column>
          <el-table-column prop="code" label="请求方式" align="center">
            <template #default="{ row }">
              <el-select v-model="row.method" style="width: 100%;" @change="handleChange(row, 'method')"
                placeholder="请选择请求方式">
                <el-option v-for="item in [{ id: 'GET', title: 'GET' }, { id: 'POST', title: 'POST' }]" :key="item.id"
                  :label="item.title" :value="item.id" />
              </el-select>
            </template>
          </el-table-column>
          <el-table-column label="是否删除" align="center" width="180">
            <template #default="{ row }">
              <el-switch v-model="row.isDeleted" @change="handleChange(row, 'isDeleted')" active-color="#ff0000"
                active-text="已删除" inactive-text="未删除" :active-value="1" :inactive-value="0" />
            </template>
          </el-table-column>
          <!-- <el-table-column label="创建时间" align="center" width="160">
            <template #default="scope">
              {{ formatTimeToStr(scope.row.createdAt, "yyyy/MM/dd hh:mm:ss") }}
            </template>
          </el-table-column>
          <el-table-column label="更新时间" align="center" width="160">
            <template #default="scope">
              {{ formatTimeToStr(scope.row.updatedAt, "yyyy/MM/dd hh:mm:ss") }}
            </template>
          </el-table-column> -->
          <el-table-column fixed="right" align="left" label="操作" width="350">
            <template #default="{ row, $index }">
              <el-button icon="edit" @click="handleAddChild(row)" v-if="row.parentId == 0"
                type="primary">添加子权限</el-button>
              <el-button icon="edit" @click="handleEdit(row)" type="primary">编辑</el-button>
              <el-button icon="Tickets" @click="handleCopy(row)" type="success">复制</el-button>
              <el-button icon="remove" @click="handleRemove(row)" type="danger">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
      </home-page-header>
    </div>

    <!--添加和修改菜单-->
    <add-sys-apis ref="addRef" @load="handleLoadData"></add-sys-apis>
  </div>
</template>

<script  setup>
import { C2B, B2C } from '@/utils/wordtransfer'
import KVA from '@/utils/kva.js'
import settings from '@/settings';
import { formatTimeToStr } from '@/utils/date'
import AddSysApis from '@/views/sys/components/AddSysApis.vue'
import { LoadTreeData, UpdateStatus, CopyData, DelById } from '@/api/sysapis.js';
import { reactive } from 'vue';
import { useUserStore } from '@/stores/user.js'
const userStore = useUserStore()
const addRef = ref(null);

// 搜索属性定义
let queryParams = reactive({
  keyword: ""
})

// 数据容器
const tableData = ref([])
const searchForm = ref(null)
// 搜索
const handleSearch = () => {
  handleLoadData()
}

// 查询列表
const handleLoadData = async () => {
  const resp = await LoadTreeData(queryParams)
  tableData.value = resp.data.data
}

// 添加
const handleAdd = () => {
  addRef.value.handleOpen('', 'save', tableData.value?.length)
}

// 编辑
const handleEdit = async (row) => {
  // 在打开,再查询，
  addRef.value.handleOpen(row.id, 'edit', tableData.value?.length)
}

// 添加子菜单
const handleAddChild = (row) => {
  addRef.value.handleOpen(row, 'child', row.children?.length)
}

// 改变序号 sorted,标题 title、启用 status,isDeleted
const handleChange = async (row, field) => {
  var value = row[field];//row.isDeleted=0 
  var params = { id: row.id, field: field, value: value };
  await UpdateStatus(params);
  KVA.notifySuccess("更新成功")
  if (field == "sort") {
    tableData.value.sort((a, b) => a.sort - b.sort);
  }
}



// 物理删除
const handleRemove = async (row) => {
  try {
    await KVA.confirm("警告", "你确定要抛弃我么？", { icon: "error" })
    await DelById(row.id)
    KVA.notifySuccess("操作成功")
    // userStore.handlePianaRole(0, "")
    handleLoadData()
  } catch (e) {
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
onMounted(() => {
  handleLoadData()

  // console.log("C2B", C2B("isDeletedNum"))
  // console.log("B2C", B2C("is_deleted_num"))
})


</script>


