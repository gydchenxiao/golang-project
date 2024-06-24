import request from '@/request'

// 获取验证码(验证码也是不需要 token 的哦)
export const getCapatcha = () => {
  return request.get('/code/init', { noToken: true })
}

// 测试jwt续期用的：创造条件使得访问后端，需要 jwt 验证
export const gggg = () => {
  return request.get('/video/find', { noToken: false })
}
