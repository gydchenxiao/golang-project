import { defineStore } from 'pinia'
import router from '@/router'
import request from '@/request'
// import { menuTreeData } from '@/mock/data.js'
import { handleLogout } from '@/api/logout.js'
import { ChangeRoleIdMenus } from '@/api/sysroles.js'
import { useSkeletonStore } from '@/stores/skeleton.js' // 骨架频的状态管理
import { useMenuTabStore } from '@/stores/menuTab.js' // 切换用户角色需要清空导航栏
import addDynamicRoutes from '@/router/dynamic.js'

//https://blog.csdn.net/weixin_62897746/article/details/129124364
//https://prazdevs.github.io/pinia-plugin-persistedstate/guide/
export const useUserStore = defineStore('user', {
  // 定义状态
  state: () => ({
    routerLoaded: false,
    // 登录用户
    user: {},
    username: '',
    userId: '',
    // 登录token
    token: '',
    // 挤下线使用
    uuid: '',
    // 当前角色
    currRoleName: '',
    currRoleCode: '',
    currRoleId: 0,
    // 获取用户对应的角色列表
    roles: [],
    // 获取角色对应的权限
    permissions: [],
    // 获取角色对应的菜单
    menuTree: []
  }),

  // 就是一种计算属性的机制，定义的是函数，使用的是属性就相当于 computed(只读属性)
  getters: {
    roleName(state) {
      return state.roles.map((r) => r.userName).join(',')
    },

    // 判断当前是否为已登陆状态
    isLogin(state) {
      return state.token !== '' ? true : false
    },

    permissionCode(state) {
      return state.permissions && state.permissions.map((r) => r.code).join(',')
    },

    permissionPath(state) {
      return state.permissions && state.permissions.map((r) => r.path).join(',')
    }
  },

  // 定义动作
  actions: {
    // 设置 token
    setToken(newtoken) {
      this.token = newtoken
    },

    // 获取 token
    getToken() {
      return this.token
    },

    // 登陆
    async toLogin(loginUser) {
      // 登陆页面是不需要 token 验证的
      const resp = await request.post('/login/toLogin', loginUser, { noToken: true })

      if (resp.data.code === 20000) {
        // 这个会回退，回退登录页
        // var { permissions, roleMenus, roles, token, user } = resp.data.data // 使用解构也是可以的哦
        console.log('toLogin ==> resp', resp)

        // 把数据放入到状态管理中
        this.user = resp.data.data.user
        this.userId = this.user.id
        this.username = this.user.username
        this.token = resp.data.data.token
        this.uuid = resp.data.data.uuid
        this.roles = resp.data.data.roles
        this.permissions = resp.data.data.permissions
        // 登录成功以后获取到菜单信息。 ---> 可以统一在登陆接口中把对应角色的菜单从数据库中查询出来的哦
        this.menuTree = resp.data.data.roleMenus //  || menuTreeData
        // 把roles列表中的角色的第一个作为，当前角色
        this.currRoleId = this.roles && this.roles.length > 0 ? this.roles[0].id : 0
        this.currRoleName = this.roles && this.roles.length > 0 ? this.roles[0].roleName : ''
        this.currRoleCode = this.roles && this.roles.length > 0 ? this.roles[0].roleCode : ''
        console.log('toLogin ===> this.menuTree', this.menuTree)

        return Promise.resolve(resp)
      } else {
        return Promise.reject(resp)
      }
    },

    // 登出
    async LoginOut() {
      // 执行服务端退出
      await handleLogout() // 后端销毁 token 到黑名单
      // 清除状态信息
      this.token = ''
      this.uuid = ''
      this.user = {}
      this.userName = ''
      this.userId = ''
      this.role = []
      this.permissions = []
      this.menuTree = []
      // 清除自身的本地存储
      localStorage.removeItem('ksd-kva-language')
      localStorage.removeItem('kva-pinia-userstore')
      // removeItem 并把骨架屏的状态恢复到true的状态
      sessionStorage.removeItem('kva-pinia-skeleton')
      useSkeletonStore().setLoading(true)
      // console.log(useSkeletonStore().skLoading)

      localStorage.removeItem('isWhitelist')
      // 然后跳转到登录
      router.push({ name: 'Login', replace: true })
    },

    // 改变用户角色的时候把对应菜单和权限查询出来，进行覆盖
    async handlePianaRole(roleId, roleName, roleCode) {
      if (roleId > 0 && roleId != this.currRoleId) {
        this.currRoleId = roleId
        this.currRoleName = roleName
        this.currRoleCode = roleCode

        // 获取到导航菜单，切换以后直接全部清空掉
        const menuTabStore = useMenuTabStore()
        menuTabStore.clear()
      }

      const resp = await ChangeRoleIdMenus({ roleId: this.currRoleId })
      // console.log('handlePianaRole ===> resp', resp)
      // 对应的权限和菜单进行覆盖
      this.permissions = resp.data.data.permissions
      this.menuTree = resp.data.data.roleMenus.sort((a, b) => a.sort - b.sort)
      // this.user = resp.data.data.user // 切换了当前用户的角色，前端对应的user是不用切换的
      // 动态路由进行重新注册
      addDynamicRoutes(this.menuTree)
      // console.log('ChangeRoleIdMenus', this.currRoleId, this.currRoleName)
      // 路由改变是会进入 router.beforeEach 的，所以上一行 addDynamicRoutes(this.menuTree) 是可以不写的
      router.push(this.menuTree[0].path)
    }
  },

  // 解决问题：刷新会丢失状态数据或者打开新的浏览器会丢失状态数据(pinia 的本地存储)
  persist: {
    enabled: true, //开启数据持久化
    strategies: [
      {
        key: 'useUserStore', //给一个要保存的名称
        storage: localStorage //sessionStorage / localStorage 存储方式
      }
    ]
  }
})
