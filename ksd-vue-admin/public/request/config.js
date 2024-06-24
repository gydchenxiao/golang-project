export default {
  method: 'get',
  // 基础url前缀
  baseURL: 'http://localhost:9899' + '/api',
  // baseURL: import.meta.env.VITE_BASE_PATH + import.meta.env.VITE_BASE_API, // 这个环境变量的设置好像并没有起到作用的哦，就用上面的方式了
  // 请求头信息
  headers: {
    'Content-Type': 'application/json;charset=UTF-8'
  },
  // 设置超时时间
  timeout: 30000,
  // 返回数据类型
  responseType: 'json'
}
