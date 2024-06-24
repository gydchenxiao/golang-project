import request from '@/request/index.js'

/**
 * 退出登录
 */
export const handleLogout = () => {
  request.post('/login/logout')
}
