import 'normalize.css/normalize.css' // 组件原生样式的重置
import './assets/base.css' // * 通配符的形式，对于很多标签的时候是很不友好的哦
// import '../public/css/main.css'
import 'nprogress/nprogress.css' // 实现头部nprogress动画

import { createApp } from 'vue'
import { createPinia } from 'pinia' // pinia 路由/页面间的数据共享
import piniaPersist from 'pinia-plugin-persist'

import App from './App.vue'
import router from './router'
import ElementPlus from 'element-plus' // element-plus 组件
import 'element-plus/dist/index.css' // element-plus 图形库的 .css 文件
import * as ElementPlusIconsVue from '@element-plus/icons-vue' // element-plus 图形库得单独再安装的哦
import AnimatedNumber from 'animated-number-vue3' // 实现数字动画
import i18n from '@/i18n' // 实现国际化处理
import KVAComponents from '@/components' // 全局组件
import KVADirective from '@/directive/index.js' // 菜单权限设置

const app = createApp(App)

const pinia = createPinia()
pinia.use(piniaPersist)
app.use(pinia)
app.use(ElementPlus)
app.use(AnimatedNumber)
app.use(i18n)
app.use(router)
app.use(KVAComponents)
app.use(KVADirective)

for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}

app.mount('#app')
