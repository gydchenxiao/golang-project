import defaultSettings from '@/settings'
import i18n from '@/i18n'
const title = defaultSettings.title || 'KVA后台管理系统'

export default function getPageTitle(key) {
  const hasKey = i18n.t(`${key}`)
  if (hasKey) {
    const pageName = i18n.t(`${key}`)
    return `${pageName} - ${title}`
  }
  return `${title}`
}
