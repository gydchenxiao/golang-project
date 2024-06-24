<template>
  <el-dialog :close-on-click-modal="false" @close="handleClose" v-model="dialogFormVisible" :title="title" width="960px">
    <el-form v-loading="formLoading" ref="formRef" label-width="100" :model="form" style="padding: 30px" :rules="rules">
      <el-form-item label="父级菜单：" required prop="parentId">
        <el-select v-model="form.parentId" style="width: 100%;" placeholder="请选择菜单" v-if="form.parentId > 0 || form.id">
          <el-option v-for="item in parentMenus" :key="item.id" :label="item.title" :value="item.id" />
        </el-select>
        <el-input v-else disabled value="根目录" />
      </el-form-item>
      <el-form-item label="菜单标题：" required prop="title">
        <el-input v-model.trim="form.title" maxlength="60" placeholder="菜单标题" />
      </el-form-item>
      <el-form-item label="菜单名字：" required prop="name">
        <el-input v-model.trim="form.name" maxlength="60" placeholder="菜单名字" />
        <span class="fz12 red ml10">用于国际化处理</span>
      </el-form-item>
      <el-form-item label="菜单路径：" required prop="path">
        <el-input v-model.trim="form.path" maxlength="100" placeholder="菜单路径" />
        <span class="fz12 red ml10">提示：路径必须要对应路由页面，
          请在项目工程目录下新建/views/{{ viewPath }}.vue 路径才可以访问页面</span>
      </el-form-item>
      <el-form-item label="菜单排序：" required prop="sort">
        <el-input-number v-model="form.sort" placeholder="菜单排序" />
      </el-form-item>
      <el-form-item label="菜单图标：" required prop="icon">
        <!-- 因为外层的 KeepAlive 有缓存的效果，所以用下面的 :key="form.id" 来解决，使得每一个都是唯一的 -->
        <add-icon v-model:icon="form.icon" :key="form.id" style="width:100%" />
      </el-form-item>
      <el-form-item label="是否隐藏：" required prop="hidden">
        <el-radio-group v-model="form.hidden">
          <el-radio :label="1">显示</el-radio>
          <el-radio :label="0">隐藏</el-radio>
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
import { SaveData, GetById, UpdateData, LoadRootData } from '@/api/sysmenus.js'
import KVA from '@/utils/kva.js'
import AddIcon from '@/views/sys/components/AddIcon.vue'
import debounce from '@/utils/debounce'
import { reactive } from 'vue'

const emits = defineEmits(["load"])
// 控制弹窗得显示
const dialogFormVisible = ref(false)
// 这个数据模型
let form = ref({
  parentId: 0,
  path: "",
  name: "",
  hidden: 1,
  sort: 1,
  icon: "Aim",
  isDeleted: 0,
  title: ""
})

const viewPath = ref("")
// 弹出标题
const title = ref("保存菜单")
const formLoading = ref(false)
const suLoading = ref(false)
const parentMenus = ref([{ id: 0, title: "根目录", parentId: 0 }])
// 表单ref, 用于提交最终得验证处理
const formRef = ref({})
// 表单验证
const rules = reactive({
  "title": [{ required: true, message: '请输入菜单标题', trigger: 'blur' }],
  "name": [{ required: true, message: '请输入菜单名称', trigger: 'blur' }],
  "path": [{ required: true, message: '请输入路径', trigger: 'blur' }],
  "icon": [{ required: true, message: '请选择一个图标', trigger: 'blur' }]
})

// 监听form，是在编写path，我要告诉用户你在view要定义一个什么样子的页面
watch(form, (newValue) => {
  var path = newValue.path.split("/").filter(c => c != '')
  if (path && path.length > 0) {
    viewPath.value = path.length == 1 ? (path[0] + '/Index') : (path[0] + "/" + (path[1].substring(0, 1).toUpperCase()) + (path[1].substring(1)))
  } else {
    viewPath.value = ""
  }
}, { immediate: true, deep: true })


// 打开修改密码的弹窗
const handleOpen = async (data, flag, sort) => {
  // 每次打开头重置
  dialogFormVisible.value = true;
  // 查询主菜单
  await handleLoadRoot()
  if (data && flag === "edit") { // data 是一个 id
    title.value = "编辑菜单"
    formLoading.value = true;
    const resp = await GetById(data)
    formLoading.value = false;
    // 这里地方要注意，一定要用reactive进行包裹处理
    form.value = resp.data.data
    title.value = "编辑菜单【" + form.value.title + "】"
  } else if (data && flag === 'child') { // data 是一个 item 数据
    form.value.sort = (sort || 0) + 1
    title.value = "添加子菜单"
    form.value.parentId = data.id
    form.value.path = data.path + "/"
  } else {
    form.value.sort = sort + 1
  }
}

// 保存用户
const handleSubmit = () => {
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
        // 提示valid失败
        KVA.notifyError('valid失败')
        suLoading.value = false
      }
    })
  }, 1000)
}

const handleClose = () => {
  delete form.value.id
  form.value = {
    parentId: 0,
    path: "",
    name: "",
    hidden: 1,
    sort: 1,
    icon: "Aim",
    isDeleted: 0,
    title: ""
  }
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