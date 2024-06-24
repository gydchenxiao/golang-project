import { createI18n } from 'vue-i18n'
import zhLocale from './lang/zh'
import enLocale from './lang/en'
import zhCn from 'element-plus/es/locale/lang/zh-cn' // element-plus 组件下载的时候，随便是把国际化的各种语言也下载下来了的
import en from 'element-plus/es/locale/lang/en' //

const i18n = createI18n({
  legacy: false,
  fallbackLocale: 'zh',
  locale: localStorage.getItem('ksd-kva-language') || 'zh-cn', // 设置地区
  messages: {
    en: {
      ...enLocale,
      ...zhCn
    },
    'zh-cn': {
      ...zhLocale,
      ...en
    }
  }
})

export default i18n

export const elementLocales = {
  'zh-cn': zhCn,
  en
}
