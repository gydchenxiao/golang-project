<template>
  <div>
    <span class="gva-icon" style="position: absolute; z-index: 9999; padding: 3px 10px 0; ">
      <el-icon>
        <component :is="metaData" />
      </el-icon>
    </span>

    <!-- 菜单图标: 图标和图标名过于靠近的 bug -->
    <span style="margin-left: 10px;"></span>

    <el-select v-model="metaData" @change="handleChange" style="width:100%" clearable filterable class="gva-select"
      placeholder="请选择">
      <el-option v-for="item in options" :key="item.key" class="select__option_item" :label="item.key" :value="item.key">
        <span class="gva-icon" style=" padding: 3px 0 0; " :class="item.label">
          <el-icon>
            <component :is="item.label" />
          </el-icon>
        </span>
        <span style="text-align: left">{{ item.key }}</span>
      </el-option>
    </el-select>
  </div>
</template>
  
<script setup>
// 1: 导入所有的图标
import * as ElementPlusIconsVue from '@element-plus/icons-vue'

// 2：下拉列表图标的容器
const options = reactive([])
// 3：具名响应数据 v-model:icon="form.icon"
const props = defineProps({
  icon: {
    default: "Aim",
    type: String
  },
})

// 把传递进来属性转变响应式数据
const metaData = ref(props.icon)
// 开始处理
const handleLoadIcon = () => {
  // 把所有的图标放入options图标容器中
  for (let key in ElementPlusIconsVue) {
    options.push({ label: key, key: key })
  }
  // 如果没有传递，默认把第一个选中
  if (!metaData.value) {
    metaData.value = options[0].label
  }
}


// 定义一个自定义事件
const emits = defineEmits(["update:icon"])
// 选择一个图标，就把图标用自定义事件的传送出去
const handleChange = () => {
  emits("update:icon", metaData.value)
}

onMounted(() => {
  handleLoadIcon()
})

</script>
  
<script>
export default {
  name: 'Icon',
}
</script>
  
<style lang="scss">
.gva-icon {
  color: rgb(132, 146, 166);
  font-size: 14px;
  margin-right: 10px;
}

.gva-select .el-input__inner {
  padding: 0 30px !important
}

.select__option_item {
  display: flex;
  align-items: center;
  justify-content: flex-start;
}
</style>
  