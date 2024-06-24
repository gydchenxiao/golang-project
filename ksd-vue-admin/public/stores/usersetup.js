import { defineStore } from 'pinia'
import axios from 'axios'
import { ref, computed } from 'vue'

//https://blog.csdn.net/weixin_62897746/article/details/129124364
//https://prazdevs.github.io/pinia-plugin-persistedstate/guide/
export const useUserStore = defineStore('user', () => {
  // 相当于  state
  const user = ref({})
  const username = ref('')
  const userId = ref('')
  const token = ref('')
  const roles = ref([])
  const age = ref(10)
  const permissions = ref([])

  // 使用计算属性处理状态数据
  const roleName = computed(() => {
    return roles.value.map((r) => r.name).join(',')
  })

  const permissionCode = computed(() => {
    return permissions.value.map((r) => r.code).join(',')
  })

  // 定义函数---actions
  const toLogin = async (loginUser) => {
    const resp = await axios.post('http://localhost:9899/login/toLogin', loginUser)
    if (resp.data.code === 20000) {
      // 这个会回退，回退登录页
      var result = resp.data.data
      console.log('result', result)
      // 把数据放入到状态管理中
      user.value = result.user
      userId.value = result.user.id
      username.value = result.user.name
      token.value = result.token
      roles.value = result.roles
      permissions.value = result.permissions
      return Promise.resolve(resp)
    } else {
      return Promise.reject(resp)
    }
  }

  return {
    age,
    user,
    userId,
    username,
    token,
    toLogin,
    roles,
    permissions,
    roleName,
    permissionCode
  }
})
