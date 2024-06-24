<template>
  <el-dialog :close-on-click-modal="false" v-model="roleDialogFormVisible" :title="title" width="640px">
    <el-form v-loading="formLoading" ref="sysRoleFormRef" :model="form" style="padding: 30px" :rules="rules">
      <el-form-item label="角色名称" label-width="100" required prop="roleName">
        <el-input v-model="form.roleName" autocomplete="off" maxlength="60" placeholder="请输入角色名称" />
      </el-form-item>
      <el-form-item label="角色代号" label-width="100" required prop="roleCode">
        <el-input v-model="form.roleCode" maxlength="60" autocomplete="off" placeholder="请输入角色代号" />
      </el-form-item>
      <el-form-item label="是否删除" label-width="100" required prop="isDeleted">
        <el-radio-group v-model="form.isDeleted">
          <el-radio :label="1">已删除</el-radio>
          <el-radio :label="0">未删除</el-radio>
        </el-radio-group>
      </el-form-item>
      <div class="dialog-footer" style="text-align: center">
        <el-button @click="handleClose" icon="remove">取消</el-button>
        <el-button type="primary" @click="handleSubmit" :loading="suLoading" icon="plus">{{
          form.id ? "更新" : "保存" }}</el-button>
      </div>
    </el-form>
  </el-dialog>
</template>

<script setup>
import { SaveData, GetById, UpdateData } from '@/api/sysroles.js'
import KVA from '@/utils/kva.js'
import debounce from '@/utils/debounce'

const emits = defineEmits(["load"])
// 控制弹窗得显示
const roleDialogFormVisible = ref(false)
// 这个数据模型
let form = reactive({
  roleName: "",
  roleCode: "",
  isDeleted: 0
})

// 弹出标题
const title = ref("保存系统角色")
const formLoading = ref(false)
const suLoading = ref(false)
// 表单ref, 用于提交最终得验证处理
const sysRoleFormRef = ref({})
// 表单验证谷子额
const rules = reactive({
  "roleName": [{ required: true, message: '请输入角色名称', trigger: 'blur' }],
  "roleCode": [{ required: true, message: '请输入角色代号', trigger: 'blur' }]
})

// 打开修改密码的弹窗
const handleOpen = async (opid) => {
  // 每次打开头重置
  roleDialogFormVisible.value = true;
  if (opid) {
    title.value = "编辑系统角色"
    formLoading.value = true;
    const resp = await GetById(opid)
    formLoading.value = false;
    // 这里地方要注意，一定要用reactive进行包裹处理
    form = reactive(resp.data.data)
    title.value = "你正在修改系统角色是【" + form.roleName + "】"
  }
}

// 保存用户
const handleSubmit = () => {
  suLoading.value = true
  debounce(() => {
    sysRoleFormRef.value.validate(async (valid, fields) => {
      if (valid) {
        try {
          form.id ? await UpdateData(form) : await SaveData(form)
          suLoading.value = false
          // 提示保存成功
          KVA.notifySuccess(form.id ? "更新成功" : "保存成功")
          // 关闭弹窗 
          handleClose()
          // 刷新方法
          emits("load")
        } catch (ex) {
          suLoading.value = false
          // 关闭弹窗 
          handleClose()
        }
      }
    })
  }, 1000)
}

const handleClose = () => {
  delete form.id
  sysRoleFormRef.value.resetFields()
  roleDialogFormVisible.value = false
}


// 把弹出打开得方法暴露给父组件进行调用
defineExpose({
  handleOpen
})
</script>