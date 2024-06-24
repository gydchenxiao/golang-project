<template>
  <div class="kva-container">
    <div class="kva-contentbox">
      <!-- HmoePageHeader 组件 -->
      <home-page-header>
        <!-- 搜索栏 -->
        <div class="kva-form-search">
          <el-form :inline="true" :model="queryParams">
            <el-form-item>
              <el-button type="primary" v-permission="['A1001', 'A1002']" icon="Plus" @click="handleAdd">添加</el-button>
              <el-button type="danger" v-permission="'A1003'" icon="Delete" @click="handleDels">删除</el-button>
            </el-form-item>
            <el-form-item label="关键词：">
              <el-input v-model="queryParams.keyword" placeholder="请输入搜索关键词..." maxlength="10" clearable
                @keydown.enter="handleSearch" />
            </el-form-item>
            <el-form-item>
              <el-button type="primary" icon="Search" @click.prevent="handleSearch">搜索</el-button>
            </el-form-item>
          </el-form>
        </div>

        <!-- el-table 组件数据展示 -->
        <el-table :data="tableData" @selection-change="handleSelectionChange" style="width: 100%"
          :height="settings.tableHeight()">
          <el-table-column type="selection" fixed="left" width="55" />
          <el-table-column prop="id" fixed="left" label="ID" />
          <el-table-column label="昵称" fixed="left" width="150">
            <template #default="scope">
              <el-avatar :src="scope.row.avatar" size="small"></el-avatar>
              <span style="margin-left: 5px;position: relative;top:-5px;">{{ scope.row.username }}</span>
            </template>
          </el-table-column>
          <el-table-column prop="account" label="账号" width="120" />
          <!-- <el-table-column prop="phone" label="用户手机号" width="150" /> -->
          <!-- <el-table-column prop="email" label="用户邮箱" width="180" /> -->
          <el-table-column label="是否被启用" align="center" width="180">
            <template #default="scope">
              <el-switch v-model="scope.row.enable" @change="handleChangeEnable(scope.row)" active-text="启用中"
                inactive-text="禁止中" :active-value="1" :inactive-value="0" />
            </template>
          </el-table-column>
          <el-table-column label="删除状态" align="center" width="200">
            <template #default="scope">
              <el-switch v-model="scope.row.isDeleted" @change="handleDel(scope.row)" active-color="#ff0000"
                active-text="已删除" inactive-text="未删除" :active-value="1" :inactive-value="0" />
            </template>
          </el-table-column>
          <el-table-column label="授予角色" align="center" width="320">
            <template #default="{ row, $index }">
              <el-cascader style="width: 100%;" v-model="row.roleIds" :options="rolesData" @change="handleChangeRole(row)"
                :props="props" />
            </template>
          </el-table-column>
          <el-table-column label="创建时间" align="center" width="160">
            <template #default="scope">
              {{ formatTimeToStr(scope.row.createdAt, "yyyy/MM/dd hh:mm:ss") }}
            </template>
          </el-table-column>
          <el-table-column label="更新时间" align="center" width="160">
            <template #default="scope">
              {{ formatTimeToStr(scope.row.updatedAt, "yyyy/MM/dd hh:mm:ss") }}
            </template>
          </el-table-column>
          <el-table-column fixed="right" align="center" label="操作" width="240">
            <template #default="{ row, $index }">
              <el-button icon="edit" @click="handleEdit(row)" type="primary">编辑</el-button>
              <el-button icon="switch" @click="handleResetPwd(row)" type="success">重置密码</el-button>
            </template>
          </el-table-column>
        </el-table>

        <!-- 分页 -->
        <div class="kva-pagination-box">
          <el-pagination :current-page="queryParams.page" :page-size="queryParams.pageSize"
            :page-sizes="[10, 20, 30, 50, 100, 200]" small layout="total, sizes, prev, pager, next, jumper" :total="total"
            @size-change="handleSizeChange" @current-change="handleCurrentChange" />
        </div>
      </home-page-header>
    </div>
    <!--自定义修改密码得弹窗-->
    <update-pwd ref="userPwdRef"></update-pwd>
    <!--添加和修改系统用户-->
    <add-sys-user ref="addRef" @load="handleLoadData"></add-sys-user>
  </div>
</template>

<script  setup>
import { formatTimeToStr } from '@/utils/date'
import KVA from '@/utils/kva.js'
import settings from '@/settings';
import { LoadData, UpdateStatusSysUser, SaveUserRole } from '@/api/sysusers.js'
import { FindData } from '@/api/sysroles.js'
import UpdatePwd from '@/views/sys/components/UpdatePwd.vue'
import AddSysUser from '@/views/sys/components/AddSysUser.vue'

// 搜索属性定义
const queryParams = reactive({
  page: 1,
  pageSize: 10,
  keyword: ''
})
// 数据容器
const tableData = ref([])
const total = ref(0)

// 搜索
const handleSearch = () => {
  queryParams.page = 1
  queryParams.total = 0
  queryParams.pageSize = 10
  handleLoadData()
}

// 获取添加和修改系统用户的组件对象
const addRef = ref({});

// 查询用户列表信息
const handleLoadData = async () => {
  const resp = await LoadData(queryParams)
  console.log('resp', resp) // 注意：通过打印可以看到，resp 返回的数据是有两层 data 的哦
  tableData.value = resp.data.data.list.map(data => { // 加上 roleIds
    data.roleIds = data.roleIds ? data.roleIds.split(",").map(r => r * 1) : []
    return data;
  })
  console.log('tableData', tableData)
  total.value = resp.data.data.total
  queryParams.page = resp.data.data.page
}

// 批量选择
const multipleSelection = ref([])
const handleSelectionChange = (vals) => {
  multipleSelection.value = vals
}

// 添加事件
const handleAdd = () => {
  // KVA.notify("注册提示", "感谢你注册平台,<a href=''>点击此处进入查看</a>", 3000, { type: "success", position: "br" })
  // 在打开,再查询，
  addRef.value.handleOpen()
}
// 删除事件
const handleDels = () => {
  KVA.confirm("警告", "<strong>你确定要抛弃我么？</strong>", { icon: "success" }).then(() => {
    KVA.message("去请求你要删除的异步请求的方法把")
  }).catch(err => {
    KVA.error("你点击的是关闭或者取消按钮")
  })
}

// 先直接使用下面的结构体了，推荐后期改用context数据载体
// type Params struct {
// 	Id    uint   `json:"id"`
// 	Filed string `json:"field"`
// 	Value int    `json:"value"`
// }
// 启用和禁止的处理
const handleChangeEnable = async (row) => {
  // console.log('row', row)
  var params = {
    id: row.id,
    field: 'enable',
    value: row.enable
  }
  const resp = await UpdateStatusSysUser(params)
  console.log('resp', resp)
  if (resp.data.code === 20000) {
    KVA.notifySuccess("操作成功")
  }
}
// 已删除和未删除的处理
const handleDel = async (row) => {
  console.log('row', row)
  var params = {
    id: row.id,
    field: 'is_deleted',
    value: row.isDeleted
  }
  const resp = await UpdateStatusSysUser(params)
  console.log('handleDel resp', resp)
  if (resp.data.code === 20000) {
    // KVA.notifySuccess("操作成功")
  }
}

// 查询所有的角色信息
const rolesData = ref([])
const props = ref({ multiple: true, emitPath: false })
// 查询所有的角色信息
const handleFindRoles = async () => {
  const resp = await FindData();
  rolesData.value = resp.data.data.map(r => ({ label: r.roleName, value: r.id }))
}
// 授予角色
const handleChangeRole = async (row) => {
  // console.log('row', row)
  const params = { userId: row.id, roleIds: row.roleIds.join(',') }
  // console.log('params', params)
  await SaveUserRole(params) // response 拦截的时候做了处理，所以失败了，是不会做下一步的
  KVA.notifySuccess("授予角色成功!")
}


// 获取重置密码的组件对象
const userPwdRef = ref({});
// 重置密码
const handleResetPwd = (row) => {
  userPwdRef.value.handleOpen(row)
}
// 编辑
const handleEdit = async (row) => {
  // 在打开,再查询，
  addRef.value.handleOpen(row.id)
}

// 点击分页改变
const handleCurrentChange = (pageNum) => {
  queryParams.page = pageNum
  handleLoadData()
}
// 点击每页显示多少条改变
const handleSizeChange = (pageSize) => {
  queryParams.page = 1
  queryParams.pageSize = pageSize
  handleLoadData()
}

// 生命周期加载
onMounted(() => {
  handleLoadData()
  handleFindRoles()
})


</script>

