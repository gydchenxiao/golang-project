<template>
  <div class="kva-container">
    <div class="kva-contentbox">
      <home-page-header>
        <div class="kva-form-search">
          <el-form :inline="true" ref="searchForm" :model="queryParams">
            <el-form-item>
              <el-button type="primary" icon="Plus" @click="handleAdd">添加根菜单</el-button>
            </el-form-item>
            <el-form-item label="关键词：">
              <el-input v-model="queryParams.keyword" @keydown.enter="handleSearch" placeholder="请输入菜单名称..."
                maxlength="10" clearable />
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
          <el-table-column fixed prop="id" label="ID" align="center" width="70" />
          <el-table-column fixed prop="parentId" label="父ID" align="center" width="50" />
          <el-table-column prop="title" label="展示名字" align="center">
            <template #default="{ row }">
              <el-input v-model="row.title" style="text-align:center" @change="handleChange(row, 'title')"></el-input>
            </template>
          </el-table-column>
          <el-table-column label="图标" align="center">
            <template #default="{ row }">
              <el-icon>
                <component :is="row.icon" />
              </el-icon>
              {{ row.icon }}
            </template>
          </el-table-column>
          <el-table-column prop="path" label="访问路径" align="center" />
          <!-- <el-table-column prop="name" label="国际化名字" align="center" /> -->
          <el-table-column prop="sort" label="排序" align="center" width="180px">
            <template #default="{ row }">
              <el-input-number v-model="row.sort" @change="handleChange(row, 'sort')"></el-input-number>
            </template>
          </el-table-column>
          <el-table-column label="是否隐藏" align="center" width="180">
            <template #default="{ row }">
              <el-switch v-model="row.hidden" @change="handleChange(row, 'hidden')" active-color="#4caf50"
                active-text="显示中" inactive-text="已隐藏" :active-value="1" :inactive-value="0" />
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
          <el-table-column fixed="right" align="left" label="操作" width="340">
            <template #default="{ row, $index }">
              <el-button icon="edit" @click="handleAddChild(row)" v-if="row.parentId == 0"
                type="primary">添加子菜单</el-button>
              <el-button icon="Tickets" @click="handleCopy(row)" type="primary">复制</el-button>
              <el-button icon="edit" @click="handleEdit(row)" type="primary">编辑</el-button>
              <el-button icon="edit" @click="handleRemove(row)" type="primary">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
      </home-page-header>
    </div>

    <!-- 菜单组件 -->
    <add-sys-menus ref="addRef" @load="handleLoadData"></add-sys-menus>
  </div>
</template>

<script  setup>
import KVA from '@/utils/kva.js'
import settings from '@/settings';
import { formatTimeToStr } from '@/utils/date'
import AddSysMenus from '@/views/sys/components/AddSysMenus.vue'
import { LoadTreeData, UpdateStatus, CopyData, DelById } from '@/api/sysmenus.js';
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
  // console.log('resp', resp)
  tableData.value = resp.data.data
}

// 添加
const handleAdd = () => {
  addRef.value.handleOpen('', 'save', tableData.value?.length)
}

// 复制
const handleCopy = async (row) => {
  await CopyData(row.id)
  KVA.notifySuccess('复制成功')
  handleLoadData()
}

// 编辑
const handleEdit = async (row) => {
  // console.log('row', row)
  // 在打开,再查询，
  addRef.value.handleOpen(row.id, 'edit', tableData.value?.length)
}

// 删除
const handleRemove = async (row) => {
  try {
    await KVA.confirm("警告", "你确定要抛弃我么？", { icon: "error" })
    await DelById(row.id)
    KVA.notifySuccess("操作成功")
    // userStore.handlePianaRole(0, "") // 不切换角色不用了
    handleLoadData()
  } catch (e) {
    KVA.notifyError("操作失败")
  }
}

// 添加子菜单
const handleAddChild = (row) => {
  // console.log('row', row)
  addRef.value.handleOpen(row, 'child', row.children?.length)
}

// 重置搜索表单
const handleReset = () => {
  queryParams.keyword = ""
  searchForm.value.resetFields()
  handleLoadData()
}

// // 删除单个
// const handleDel = async (row) => {
//   var params = {
//     id: row.id,
//     field: 'is_deleted',
//     value: row.isDeleted
//   }
//   await UpdateStatus(params)
//   KVA.notifySuccess("操作成功")
// }
// 改变序号
const handleChange = async (row, field) => {
  var value = row[field];
  var params = { id: row.id, field: field, value: value };
  await UpdateStatus(params);
  KVA.notifySuccess("更新成功")
  if (field == "sort") {
    tableData.value.sort((a, b) => a.sort - b.sort);
  }
  if (field == "hidden") {
    userStore.handlePianaRole(0, "")
  }
}

// 生命周期加载
onMounted(() => {
  handleLoadData()
})


</script>

