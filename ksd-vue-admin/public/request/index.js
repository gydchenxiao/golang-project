// 1: 导入axios异步请求组件
import axios from 'axios'
// 2: 把axios请求的配置剥离成一个config/index.js的文件
import config from './config'
// 3: 获取路由对象--原因是：在加载的过程中先加载的pinia所以不能useRouter机制。
import router from '@/router'
// 4: elementplus消息框
import { ElMessage } from 'element-plus'
import KVA from '@/utils/kva.js'
// 5: 获取登录的token信息
import { useUserStore } from '@/stores/user.js'

import { useRouter, useRoute } from 'vue-router'

/*
const instance = axios.create({
  // 基础路径，如果你在使用过程URL过程是相对路径。就自动把baseURL+URL
  // 基础路径，如果你在使用过程URL过程是绝对路径。直接使用URL
  baseURL: 'http://localhost:9899',
  timeout: 3000, // 如果一个接口超过10秒。那么就说明你接口要优化了。
  headers: {
    // Authorization: userStore.token,
    'Content-Type': 'application/json;charset=UTF-8'
  },
  // 返回数据类型
  responseType: 'json'
})
*/
// 6: 然后创建一个axios的实例
const instance = axios.create({ ...config })

// 添加请求拦截器
instance.interceptors.request.use(
  function (config) {
    // console.log('config', config)
    // 在发送请求之前做些什么，判断是否需要使用 token 过期验证
    if (!config.noToken) {
      const userStore = useUserStore()
      //   alert(userStore.token)
      const isLogin = userStore.isLogin
      if (!isLogin) {
        router.push('/login')
        return
      } else {
        // 这里给请求头增加参数.request--header，在服务端可以通过request的header可以获取到对应参数
        // 比如go: c.GetHeader("Authorization")
        // 比如java: request.getHeader("Authorization")
        config.headers.Authorization = userStore.getToken()
        config.headers.KsdUUID = userStore.uuid // 踢下线的检测
      }
    }

    return config
  },
  function (error) {
    // 判断请求超时
    if (error.code === 'ECONNABORTED' && error.message.indexOf('timeout') !== -1) {
      ElMessage({ message: '请求超时', type: 'error', showClose: true })
      // 这里为啥不写return
    }
    return Promise.reject(error)
  }
)

// 添加响应拦截器
instance.interceptors.response.use(
  (response) => {
    // 2xx 范围内的状态码都会触发该函数。
    // 对响应数据做点什么
    // 在这里可以获取到server传递过来的头部信息（可以开始续期等操作了）
    // console.log('response', response)

    // 是否需要更新 token
    if (response.headers['new-authorization']) {
      const UseStore = useUserStore()
      UseStore.setToken(response.headers['new-authorization'])
    }

    // 挤下线问题
    if (response.data.code === 4001) {
      const userStore = useUserStore()
      userStore.LoginOut()
      KVA.notifyError(response.data.message) // 没有提示框的效果
      location.reload() // 直接刷新页面，不然应该下线的用户还是可以多点击几次页面的(好像是名优实现很快的同步，所以使用刷新一下是很快会被同步的)
      return
    }

    // 响应是否成功
    if (response.data.code === 20000) {
      console.log('response', response)

      return response
    } else if (response.data.code === 80001) {
      KVA.notifyError(response.data.message) // 权限不足
      // 看情况是否需要跳转页面
      // router.push('nopermission')
      return
    } else {
      console.log('response', response)

      if (response.data.message) {
        // ElMessage({
        //   message: response.data.message,
        //   type: 'error',
        //   showClose: true
        // })

        // 使用封装了的信息提示
        KVA.notifyError(response.data.message)
      }
      return Promise.reject(response) // 返回接口返回的错误信息
    }
  },
  (err) => {
    // 这里还要返回含义是：告诉你可以继续在接口的调用为止来做判断和处理
    // 如果是.then.catch 这里error就会进入到catch
    // 如果是async/await的写法，请使用try/catch去捕捉error

    return Promise.reject(err) // 返回接口返回的错误信息
  }
)

export default instance
