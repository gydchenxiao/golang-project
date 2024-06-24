<template>
  <el-dialog :close-on-click-modal="false" @close="handleClose" v-model="dialogFormVisible" :title="title" width="960px">
    <el-form v-loading="formLoading" ref="formRef" label-width="150" :model="form" style="padding: 30px" :rules="rules">
      <el-form-item label="父级权限：" required prop="parentId">
        <el-select v-model="form.parentId" style="width: 100%;" placeholder="请选择菜单" v-if="form.parentId > 0 || form.id">
          <el-option v-for="item in parentMenus" :key="item.id" :label="item.title" :value="item.id" />
        </el-select>
        <el-input v-else disabled value="根目录" />
      </el-form-item>
      <el-form-item label="API标题：" required prop="title">
        <el-input v-model="form.title" maxlength="60" placeholder="菜单标题" />
      </el-form-item>
      <el-form-item label="API的路径：" required prop="path">
        <el-input v-model="form.path" maxlength="60" placeholder="API的路径" />
      </el-form-item>
      <el-form-item label="API的编号" required prop="path">
        <el-input v-model="form.code" maxlength="60" placeholder="API的编号" />
      </el-form-item>
      <el-form-item label="API的方法：" required prop="method">
        <el-radio-group v-model="form.method">
          <el-radio label="GET">GET</el-radio>
          <el-radio label="POST">POST</el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="是否删除：" required prop="isDeleted">
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
import { SaveData, GetById, UpdateData, LoadRootData } from '@/api/sysapis.js'
import KVA from '@/utils/kva.js'
import debounce from '@/utils/debounce'

const emits = defineEmits(["load"])
// 控制弹窗得显示
const dialogFormVisible = ref(false)
// 这个数据模型
let form = ref({
  path: "",
  code: 0,
  parentId: 0,
  method: "POST",
  isDeleted: 0,
  title: ""
})


// 弹出标题
const title = ref("保存API")
const formLoading = ref(false)
const suLoading = ref(false)
const parentMenus = ref([{ id: 0, title: "根目录", parentId: 0 }])
// 表单ref, 用于提交最终得验证处理
const formRef = ref({})
// 表单验证谷子额
const rules = reactive({
  "path": [{ required: true, message: '请输入API路径', trigger: 'blur' }],
  "title": [{ required: true, message: '请输入API的名称', trigger: 'blur' }],
  "menuId": [{ required: true, message: '请选择菜单', trigger: 'blur' }]
})

// 打开弹窗
const handleOpen = async (opid, flag, sort) => {
  console.log('flag', flag)

  // 每次打开头重置
  dialogFormVisible.value = true;
  // 加载父权限
  await handleLoadRoot();
  if (opid && flag === "edit") {
    title.value = "编辑权限"
    formLoading.value = true;
    const resp = await GetById(opid)
    formLoading.value = false;
    // 这里地方要注意，一定要用reactive进行包裹处理
    form.value = resp.data.data
    title.value = "编辑权限【" + form.value.title + "】"
  } else if (flag === 'child') {
    title.value = "添加子权限"
    form.value.parentId = opid.id
    form.value.path = opid.path + "/"
  }
}


// 保存用户
const handleSubmit = () => {
  console.log('handleSubmit')

  suLoading.value = true
  debounce(() => {
    formRef.value.validate(async (valid, fields) => {
      if (valid) {
        try {
          form.value.id ? await UpdateData(form.value) : await SaveData(form.value)
          suLoading.value = false
          // 提示保存成功
          KVA.notifySuccess(form.value.id ? "更新成功" : "保存成功")
          // 关闭弹窗 
          handleClose()
          // 刷新方法
          emits("load")
        } catch (ex) {
          suLoading.value = false
        }
      } else {
        suLoading.value = false
      }
    })
  }, 1000)
}

const handleClose = () => {
  delete form.value.id
  parentMenus.value = [{ id: 0, title: "根目录", parentId: 0 }]
  formRef.value.resetFields()
  dialogFormVisible.value = false
}

// 开始查询父菜单
const handleLoadRoot = async () => {
  const resp = await LoadRootData()
  parentMenus.value = parentMenus.value.concat(resp.data.data)
}

// 把弹出打开得方法暴露给父组件进行调用
defineExpose({
  handleOpen
})
</script>