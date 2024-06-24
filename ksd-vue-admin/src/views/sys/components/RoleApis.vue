<template>
    <el-drawer v-model="roleDrawer" size="40%" :with-header="false" :direction="rtl" :before-close="handleClose">
        <header>
            <h3>当前授权api的角色是：【{{ title }}】</h3>
        </header>
        <div style="display: flex;justify-content: space-between;margin-top: 20px;">
            <el-input style="width: 60%;" v-model="keyword" placeholder="筛选" maxlength="20"></el-input>
            <el-button type="primary" @click="handleSubmitMenusToRole" icon="check"
                style="padding:8px 16px!important;">确定</el-button>
        </div>
        <div class="treebox">
            <el-tree ref="treeRef" :data="menuTreeData" @check-change="handleSelectNodes" show-checkbox node-key="id"
                default-expand-all highlight-current :filter-node-method="filterNode" :default-expanded-keys="[2, 3]"
                :default-checked-keys="[5]" :props="defaultProps" />
        </div>
    </el-drawer>
</template>


<script setup>
// 引入菜单api
import { LoadTreeData } from '@/api/sysapis.js'
import KVA from '@/utils/kva.js'
import { SelectRoleApis, SaveRoleApis } from '@/api/sysroles.js'

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
    // 可以看到有两层 data 的哦，当然在响应拦截那里可以设置返回 resp.data，那么在这里就不用两层.data了
    // console.log('resp', resp) 
    menuTreeData.value = resp.data.data
}

// 关闭抽屉
const handleClose = (done) => {
    ElMessageBox.confirm('你确定要放弃该操作吗？').then(() => {
        menuTreeData.value = []
        treeRef.value.setCheckedKeys([])
        done()
    })
}

// 打开抽屉
const handleOpen = async (role) => {
    title.value = `你当前操作的角色是：【${role.roleName}】`
    currentRole.value = role
    roleDrawer.value = true
    // 使用加载--立即查询所有菜单信息
    await hanelLoadTreeData()
    // 查询当前角色已经授权的菜单
    const resp = await SelectRoleApis(role.id)
    if (resp.data.data && resp.data.data?.length > 0) {
        var menuIdArr = resp.data.data.map(c => c.id)
        // 然后把授权的菜单，全部选中
        treeRef.value.setCheckedKeys(menuIdArr)
    } else {
        treeRef.value.setCheckedKeys([])
    }
}

// 点击菜单获取节点信息
const handleSelectNodes = () => {
    selectNodes.value = treeRef.value.getCheckedKeys();
}

// 开始给角色授权菜单
const handleSubmitMenusToRole = async () => {
    // 获取所有选中的节点
    let selectNodesIds = treeRef.value.getCheckedKeys();
    if (selectNodesIds && selectNodesIds.length == 0) {
        KVA.notifyError("请选择一个项进行操作!!!")
        return;
    }
    // 开始获取
    var menuIds = selectNodesIds.join(",")
    var roleId = currentRole.value.id;
    // 发起异步请求把role和menuids传递给服务器，开始绑定
    const resp = await SaveRoleApis({ "roleId": roleId, "apiIds": menuIds })
    // console.log('resp', resp)
    if (resp.data.code = 20000) { // code 返回码是后端定义的
        KVA.notifySuccess("授权成功!!!")
        roleDrawer.value = false // 关闭抽屉
    } else if (resp.data.code == 40001) {
        KVA.notifyError("授权失败!!!")
    }
}

// 过滤节点
const filterNode = (value, data) => {
    if (!value) return true
    return data.title.includes(value)
}

// 监听
watch(keyword, (val) => {
    treeRef.value?.filter(val)
})


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