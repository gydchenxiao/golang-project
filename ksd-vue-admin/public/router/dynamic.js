import router from '@/router/index.js'
import PageMain from '@/layout/components/PageMain.vue'

// 这里是获取工程目录下的views下的所以的.vue结尾的SPA页面
const modules = import.meta.glob('../views/**/*.vue')
function addDynamicRoutes(menuTreeData, parent) {
  // 开始循环遍历菜单信息
  menuTreeData.forEach((item, index) => {
    // 准备路由数据格式
    const route = {
      path: item.path,
      name: item.name,
      // 增加访问路径的元数据信息
      meta: { name: item.name, icon: item.icon },
      children: []
    }
    // 如果存在parent,就说明有children
    if (parent) {
      if (item.parentId !== 0) {
        // 这里就开始给子菜单匹配views下面的页面spa
        const compParr = item.path.replace('/', '').split('/')
        const l = compParr.length - 1
        const compPath = compParr
          .map((v, i) => {
            return i === l ? v.replace(/\w/, (L) => L.toUpperCase()) + '.vue' : v
          })
          .join('/')
        route.path = compParr[l]
        // 设置动态组件
        route.component = modules[`../views/${compPath}`]
        console.log('route', route)
        parent.children.push(route)
      }
    } else {
      // 判断你是否有children
      if (item.children && item.children.length > 0) {
        // 这里的含义是：把匹配到菜单数据第一项作为首页的入口页面
        // /order-----redirect-----/order/list
        route.redirect = item.children[0].path
        route.component = PageMain
        // 递归
        addDynamicRoutes(item.children, route)
      } else {
        //route.component = modules[`../views/${item.name}.vue`]
        route.component = modules[`../views/${item.name.toLowerCase()}/Index.vue`]
      }
      router.addRoute('Home', route)
    }
  })
}

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

export default addDynamicRoutes
