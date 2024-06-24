<!-- 测试页面 -->
<template>
    我是一个首页
    {{ userStore.userId }}
    {{ userStore.userName }}
    {{ userStore.token }}
    {{ userStore.roles }}
    {{ userStore.permissions }}
    {{ userStore.roleName }}
    {{ userStore.permissionCode }}

    <hr>
    <el-icon :size="size" :color="color">
        <Edit />
    </el-icon>
    <!-- 或者独立使用它，不从父级获取属性 -->
    <!-- <Edit /> -->

    <hr>
    <div>
        <span style="float: left"><el-icon :size="20">
                <component :is="iconvalue"></component>
            </el-icon></span>
        <el-select v-model="iconvalue" placeholder="请选择图标" @change="handleChangeIcon">
            <el-option v-for="(key, value, index) in ElementPlusIconsVue" :key="index" :label="value" :value="value">
                <span style="float: left"><el-icon :size="20">
                        <component :is="value"></component>
                    </el-icon></span>
                <span style="float: left">{{ value }}</span>
            </el-option>
        </el-select>
    </div>

    <hr>
    <button @click="handleChangeLang('en')">切换语言-EN</button>
    <button @click="handleChangeLang('zh-cn')">切换语言-ZH-CN</button>
</template>


<script setup>
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import { computed, ref } from 'vue';
import { mapState, mapActions } from 'pinia'
import { useUserStore } from '@/stores/user.js'

const userStore = useUserStore()
const iconvalue = ref('')

</script>