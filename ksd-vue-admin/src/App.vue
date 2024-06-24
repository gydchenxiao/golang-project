<template>
  <el-config-provider :locale="elementLocales[locale]">
    <RouterView />
  </el-config-provider>
</template>

<script setup>
import { elementLocales } from '@/i18n'
import { useSkeletonStore } from '@/stores/skeleton.js'

const skeletonStore = useSkeletonStore()
const { locale } = useI18n();
locale.value = localStorage.getItem('locale') || 'zh-cn'; // localStorage 是浏览器的本地存储

// onMounted(() => {
//   setTimeout(() => {
//     skeletonStore.skLoading = false;
//   }, 1000)
// })
</script>
