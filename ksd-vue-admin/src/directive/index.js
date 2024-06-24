import { useUserStore } from '@/stores/user.js'

// 插件初始化---指令的定义
export default {
  install(app) {
    // app 为根组件实例 v-permission="100"
    // app.directive('permission',{
    //   // 在绑定元素的 attribute 前
    // 	// 或事件监听器应用前调用
    // 	created(el, binding, vnode, prevVnode) {
    // 		// 下面会介绍各个参数的细节
    // 	},
    // 	// 在元素被插入到 DOM 前调用
    // 	beforeMount(el, binding, vnode, prevVnode) {},
    // 	// 在绑定元素的父组件
    // 	// 及他自己的所有子节点都挂载完成后调用
    // 	mounted(el, binding, vnode, prevVnode) {},
    // 	// 绑定元素的父组件更新前调用
    // 	beforeUpdate(el, binding, vnode, prevVnode) {},
    // 	// 在绑定元素的父组件
    // 	// 及他自己的所有子节点都更新后调用
    // 	updated(el, binding, vnode, prevVnode) {},
    // 	// 绑定元素的父组件卸载前调用
    // 	beforeUnmount(el, binding, vnode, prevVnode) {},
    // 	// 绑定元素的父组件卸载后调用
    // 	unmounted(el, binding, vnode, prevVnode) {}
    // })

    // 权限指令
    app.directive('permission', (el, binding) => {
      var bindvalue = binding.value
      var value = []
      if (typeof bindvalue === 'string') {
        value.push(bindvalue)
      }

      if (Array.isArray(bindvalue)) {
        value = [...bindvalue]
      }

      //  当前角色拥有的所有的权限信息
      const userStore = useUserStore()
      if (userStore.currRoleCode === 'root') return //

      const permissionCode = userStore?.permissionCode
      // 这个是当前的 includes 找到就返回true
      let hasAuth = value.findIndex((v) => permissionCode?.includes(v))
      if (hasAuth == -1) {
        // el.parentNode 在js中删除元素只能通过父元素对象删除子元素。自己不能删除自己
        el.parentNode.removeChild(el)
      }
    })
  }
}
