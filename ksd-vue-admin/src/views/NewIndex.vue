<!-- 测试页面 -->
<template>
    <div>
        <animated-number from="0" :to="toNum1" :key="toNum"></animated-number>
        <animated-number from="0" :to="toNum2" :key="toNum"></animated-number>
        <animated-number from="0" :to="toNum3" :key="toNum"></animated-number>
        <hr>

        <div>{{ t('ksdAdminManage') }}</div>
        <div>{{ t('loginButton') }}</div>
        <hr>
        <button @click="handleChangeLang('zh-cn')">切换成中文</button>
        <button @click="handleChangeLang('en')">切换成英文</button>
    </div>
</template>

<script setup>
import { ref } from 'vue'
import { useI18n } from 'vue-i18n'
import Cookies from 'js-cookie'

const toNum1 = ref(4654)
const toNum2 = ref(1152)
const toNum3 = ref(5486)

const { locale, t } = useI18n()
// 语言切换
const handleChangeLang = (lang) => {
    locale.value = lang
    Cookies.set("ksd-kva-language", lang)
}
</script>