import { createRouter, createWebHistory, createWebHashHistory } from 'vue-router'
import NProgress from 'nprogress'
import { useUserStore } from '@/stores/user.js'
import { useMenuTabStore } from '@/stores/menuTab.js'
import { menuTreeData } from '@/mock/data.js'

import Layout from '@/layout/Index.vue'
import PageMain from '@/layout/components/PageMain.vue'
import settings from '@/settings.js'
import addDynamicRoutes from '@/router/dynamic.js'

// 显示右上角螺旋加载提示
NProgress.configure({ showSpinner: true })

const router = createRouter({
  // history: createWebHistory(import.meta.env.BASE_URL),
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes: [
    // 首页
    {
      path: '/',
      name: 'Home',
      component: Layout,
      redirect: '/dashboard'
      // children: []
    },
    // 登陆页面
    {
      path: '/login',
      name: 'Login',
      meta: { title: 'login' },
      component: () => import('@/views/Login.vue')
    }
  ]
})

// 类似与中间件的效果，不用在login.vue等等以后的每一个页面都要写一次
router.beforeEach(async (to) => {
  //开启进度条
  NProgress.start()
  const userStore = useUserStore()
  // 获取菜单tab的状态信息
  const menuTabStore = useMenuTabStore()

  // 如果当前是登录状态，用户访问又是登录，属于无用操作，应该跳转到首页去
  if (to.path === '/login') {
    if (userStore.isLogin) {
      return { name: 'Home' }
    }
    return true
  }

  // 判断是否登录
  if (!userStore.isLogin && to.name !== 'Login') {
    // 这里的query就是为了记录用户最后一次访问的路径，这个路径是通过to的参数获取
    // 后续在登录成功以后，就可以根据这个path的参数，然后调整到你最后一次访问的路径
    return { name: 'Login', query: { path: to.path } }
  }

  // 动态加载路由---这里需要耗时---db--ajax-
  await addDynamic()

  // 如果刷新出现空白的问题，那么就使用下面这行代码
  if (!to.name && hasRoute(to)) {
    return { ...to }
  }

  // 如果访问的是首页，就跳转到/dashboard页面
  if (to.path === '/') {
    // 读取默认菜单的默认页面,需要从数据库的菜单表中去读取
    return settings.defaultPage
  }

  menuTabStore.addMenuTab(to)
  // 查询是否注册
  return true
})

// 404 页面的路由
const router404 = {
  path: '/:pathMatch(.*)*',
  name: 'NotFound',
  component: () => import('@/views/error/NotFound.vue')
}

// 动态路由
function addDynamic() {
  const userStore = useUserStore()
  // 404可以这样处理
  router.addRoute(router404)
  // 必须服务器返回的菜单和views去碰撞形成一个完整的route信息，然后注册到home下
  if (userStore.menuTree && userStore.menuTree.length > 0) {
    // console.log('===============>1111111111111', userStore.menuTree)
    addDynamicRoutes(userStore.menuTree)
  }
}

// // 动态路由
// function addDynamic() {
//   // 1. 同步数据到上面的 router.routes 中。是得点击对应的菜单，访问对应的路径生效、不写的话虽然url改变了，但是不会生效的
//   // 404可以这样处理
//   router.addRoute(router404) // 第一个参数没有填写一个路由的 name ---> 默认是添加到根路由下（404页面就是这样的）
//   // 必须服务器返回的菜单和views去碰撞形成一个完整的route信息，然后注册到home下
//   addDynamicRoutes(menuTreeData)

//   // 2. 是得后面的 PageSidebar.vue 页面可以拿到 pinia 状态管理 userStore.menuTree 的数据，生成菜单栏，点击菜单可跳转到对应的url
//   // 同时同步到状态管理中
//   const userStore = useUserStore()
//   userStore.menuTree = menuTreeData
// }

// 放到外面去了：import addDynamicRoutes from '@/router/dynamic.js'
// // 这里是获取工程目录下的views下的所以的.vue结尾的SPA页面
// const modules = import.meta.glob('../views/**/*.vue')
// function addDynamicRoutes(menuTreeData, parent) {
//   // 开始循环遍历菜单信息
//   menuTreeData.forEach((item, index) => {
//     // 准备路由数据格式
//     const route = {
//       path: item.path,
//       name: item.name,
//       meta: { name: item.name, icon: item.icon },
//       children: []
//     }
//     // 如果存在parent,就说明有children
//     if (parent) {
//       if (item.parentId !== 0) {
//         // 这里就开始给子菜单匹配views下面的页面spa
//         const compParr = item.path.replace('/', '').split('/')
//         const l = compParr.length - 1
//         const compPath = compParr
//           .map((v, i) => {
//             return i === l ? v.replace(/\w/, (L) => L.toUpperCase()) + '.vue' : v
//           })
//           .join('/')
//         route.path = compParr[l]
//         // 设置动态组件
//         route.component = modules[`../views/${compPath}`]
//         parent.children.push(route)
//       }
//     } else {
//       // 判断你是否有children
//       if (item.children && item.children.length > 0) {
//         // 这里的含义是：把匹配到菜单数据第一项作为首页的入口页面
//         // /order-----redirect-----/order/list
//         route.redirect = item.children[0].path
//         route.component = PageMain
//         // 递归
//         addDynamicRoutes(item.children, route)
//       } else {
//         //route.component = modules[`../views/${item.name}.vue`]

//         // 1. 点击没有子菜单的菜单，那么跳转到对应目录的 Index.vue 页面中。 比如：src\views\dashboard\Index.vue 就是一个典型的例子
//         // 2. 但是只有一个子菜单的菜单，也是可以没有 Index.vue 页面的，因为 children 中有设置内容（上面的 dashboard 菜单中的 children 就没有设置，就跳转到 Index.vue 页面了）
//         route.component = modules[`../views/${item.name.toLowerCase()}/Index.vue`]
//       }
//       router.addRoute('Home', route)
//     }
//   })
// }

// 判断当前路由是否存在动态添加的路由数据中
function hasRoute(to) {
  const item = router.getRoutes().find((item) => item.path === to.path)
  return !!item
}

router.afterEach(() => {
  //完成进度条
  NProgress.done()
})

/*
// 类似与中间件的效果，不用在login.vue等等以后的每一个页面都要写一次
router.beforeEach((to, from, next) => {
  //开启进度条
  NProgress.start()
  const useStore = useUserStore()

  // 判断是否登录
  if (!useStore.isLogin && to.name !== 'Login') {
    // (UseStore.isLogin === false)不是登陆状态 && (to.name !== '/login')当前不是登陆页面 ---> 拦截直接去登陆

    // 注意：下面是使用 next 的哦，而不是使用 return 什么的

    // 这里的query就是为了记录用户最后一次访问的路径，这个路径是通过to的参数获取。后续在登录成功以后，就可以根据这个path的参数，然后调整到你最后一次访问的路径
    next({ name: 'Login', query: { path: to.path } }) // 修改当前的 path 参数去登陆
  } else {
    next()
  }
})
*/

export default router
