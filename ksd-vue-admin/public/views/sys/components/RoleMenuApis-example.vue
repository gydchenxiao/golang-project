<template>
  <el-drawer v-model="roleDrawer" size="40%" :title="title" :direction="rtl" :before-close="handleClose">
    <el-tabs type="border-card">
      <el-tab-pane label="角色菜单">
        <div style="display: flex;justify-content: space-between;">
          <el-input style="width: 60%;" v-model="keyword" placeholder="筛选" maxlength="20"></el-input>
          <el-button type="primary" @click="handleSubmitMenusToRole" icon="check"
            style="padding:8px 16px!important;">确定</el-button>
        </div>
        {{ selectNodes }}
        <div class="treebox">
          <el-tree ref="treeRef" :data="menuTreeData" @check-change="handleSelectNodes" show-checkbox node-key="id"
            :default-expanded-keys="[2, 3]" :default-checked-keys="[5]" :props="defaultProps" />
        </div>
      </el-tab-pane>

      <el-tab-pane label="角色API"></el-tab-pane>
    </el-tabs>
  </el-drawer>
</template>

<script setup>
import KVA from '@/utils/kva.js'
// 引入菜单api
import { LoadTreeData } from '@/api/sysmenus.js'
import { useUserStore } from '@/stores/user.js'

const userStore = useUserStore()

// 获取tree的实例
const treeRef = ref({})
const selectNodes = ref([])

// 抽屉控制
const roleDrawer = ref(false)
// 给那个角色授权
const currentRole = ref({})
// 标题
const title = ref("角色授权")
// 搜索菜单信息
const keyword = ref("")
// 菜单映射字段
const defaultProps = {
  children: 'children',
  label: 'title',
}
// 菜单tree数据容器
const menuTreeData = ref([])

// 加载菜单数据
const hanelLoadTreeData = async () => {
  const resp = await LoadTreeData();
  menuTreeData.value = resp.data.data
}

// 关闭抽屉
const handleClose = (done) => {
  ElMessageBox.confirm('你确定要放弃该操作吗？').then(() => {
    done()
  })
}

// 打开抽屉
const handleOpen = async (role) => {
  title.value = `你当前操作的角色是：【${role.roleName}】`
  currentRole.value = role
  roleDrawer.value = true

  // 使用加载--立即查询菜单信息
  await hanelLoadTreeData()
  // 查询当前角色已经授权的菜单

  // 然后把授权的菜单，全部选中
  treeRef.value.setCheckedKeys([
    1, 2
  ])
}

// 点击菜单获取节点信息
const handleSelectNodes = () => {
  selectNodes.value = treeRef.value.getCheckedKeys();
}

// 开始给角色授权菜单
const handleSubmitMenusToRole = () => {
  // 获取所以选中的节点
  let selectNodesList = treeRef.value.getCheckedNodes();
  // 开始获取
  if (selectNodesList && selectNodesList.length > 0) {
    var menuIds = selectNodesList.map(c => c.id).join(",")
    var roleId = currentRole.value.id;
    var roleName = currentRole.value.roleName;
    // alert(menuIds)
    // alert(roleId)
    // 发起异步请求把role和menuids传递给服务器，开始绑定

    // 如果你授权的是当前角色
    if (userStore.currRoleId == roleId) {
      userStore.handlePianaRole(roleId, roleName)
    }

    // 关闭抽屉 
    KVA.notifySuccess("操作成功")
    roleDrawer.value = false
  }
}

// 暴露方法给外部组件调用
defineExpose({
  handleOpen
})
</script>

<style scoped lang="scss">
.treebox {
  margin-top: 20px;
}
</style>