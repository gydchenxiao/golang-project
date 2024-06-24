<template>
  <el-dialog :close-on-click-modal="false" v-model="dialogFormVisible" :title="title" width="640px">
    <el-form ref="userPwdFormRef" :model="form" style="padding: 30px" :rules="rules">
      <el-form-item label="密码" label-width="100" required prop="password">
        <el-input v-model="form.password" autocomplete="off" maxlength="16" type="password" placeholder="请输入密码" />
      </el-form-item>
      <el-form-item label="确认密码" label-width="100" required prop="confirmPassword">
        <el-input v-model="form.confirmPassword" maxlength="16" autocomplete="off" placeholder="请输入确认密码"
          type="password" />
      </el-form-item>
      <el-form-item label="" label-width="100">
        <p style="color:#ff0000;font-size:12px;">默认密码是：123456</p>
      </el-form-item>
    </el-form>

    <template #footer>
      <div class="dialog-footer" style="text-align: center">
        <el-button @click="dialogFormVisible = false">取消</el-button>
        <el-button type="primary" @click="handleResetPwd">
          修改密码
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup>
import { ResetPassword } from '@/api/sysusers.js'
import KVA from '@/utils/kva.js'


// （后端）管理员修改用户数据的载体
// type UserPwdContext struct {
// 	UserId          uint   `validate:"required|gt:0" json:"userId"`                         // 修改那个用户的id
// 	Password        string `validate:"required|minLen:6|maxLen:16" json:"password" `        // 密码
// 	ConfirmPassword string `validate:"required|minLen:6|maxLen:16" json:"confirmPassword" ` // 确认密码
// }
const form = reactive({
  userId: 0,
  password: "123456",
  confirmPassword: "123456"
})

const dialogFormVisible = ref(false)

const user = ref(0)
const title = ref("修改密码")
const userPwdFormRef = ref({})
const rules = reactive({
  "password": [{ required: true, message: '请输入密码', trigger: 'blur' }],
  "confirmPassword": [{ required: true, message: '请输入确认密码', trigger: 'blur' }]
})

// 打开修改密码的弹窗
const handleOpen = (row) => {
  // 接受你要修改的用户
  user.value = row
  form.userId = row.id
  title.value = "你正在修改用户是【" + row.username + "】"
  dialogFormVisible.value = true;
}

// 修改密码
const handleResetPwd = () => {
  if (form.password != form.confirmPassword) {
    KVA.error("两次密码输入不一致...")
    return;
  }
  userPwdFormRef.value.validate(async (valid, fields) => {
    if (valid) {
      await ResetPassword(form)
      KVA.notifySuccess("修改用户【" + user.value.username + "】密码成功!")
      dialogFormVisible.value = false
    }
  })
}


defineExpose({
  handleOpen
})
</script>