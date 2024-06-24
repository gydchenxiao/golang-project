<!-- 面包屑的处理和思考 -->

<template>
    <div>
        <el-breadcrumb separator="/">
            <el-breadcrumb-item :to="{ path: '/' }">
                首页
            </el-breadcrumb-item>
            <el-breadcrumb-item v-if="parentName">
                <a href="javascript:void(0);">{{ t('menu.' + parentName) }}</a>
            </el-breadcrumb-item>
            <el-breadcrumb-item v-if="isChildren">
                <a href="javascript:void(0);">{{ t('menu.' + route.meta.name) }}</a>
            </el-breadcrumb-item>
        </el-breadcrumb>
        <div style="padding:15px 0">
            <slot></slot>
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue';
import { useRoute } from 'vue-router'
import { useI18n } from 'vue-i18n'
// 获取菜单数据
import { menuTreeData } from '@/mock/data.js'

// 获取到当前路由
const route = useRoute()
// 获取国际化
const { t } = useI18n();
// 判断是不是有子元素, 因为在菜单中存在一种没有子的情况，这个时候就没有第二级。
const isChildren = ref(true)

console.log('route', route)

// 开始截取当前的访问路径，比如：/sys/user
let parentPath = route.path.substring(0, route.path.indexOf('/', 2))//得到的是：/sys
if (!parentPath) {
    parentPath = route.path
    // 代表你没有子元素
    isChildren.value = false;
}
// 如果有子元素，可以把去查找菜单信息
const parentName = menuTreeData.find(obj => obj.path == parentPath).name

</script>


<style></style>