import { defineStore } from 'pinia'

// PageMenuTab.vue/PageMenu.vue 页面的骨架屏设置
// 为什么 PageMenuTab.vue/PageMenu.vue 不和 PageSidebar.vue 页面设置同一个 pinia 状态管理(skLoading)？
// 因为 PageSidebar.vue 页面是管理角色菜单的，只有在重新登陆的时候才看到效果，在刷新的时候不用在加载的哦。
// 而 PageMenuTab.vue/PageMenu.vue 页面在刷新的时候也得看到效果的哦
export const useSkeletonStore = defineStore('skeleton', {
  // 定义状态
  state: () => ({
    skLoading: true
  }),

  // 定义动作
  actions: {
    /* 设置loading */
    setLoading(loading) {
      this.skLoading = loading
    }
  },
  persist: {
    // enabled: true, // 开启数据持久化（不能开启的哦，不然刷新的时候是没有效果的。因为sessionStorage存储在浏览器中了，那么只有得重新登陆才会有效果的）
    key: 'kva-pinia-skeleton', // 和其他的得不同的哦（记得登出的时候得删除的哦）
    // 注意下面为什么使用的是 sessionStorage 而不是 localStorage？
    //    和 localStorage 是不一样的，因为重新登陆的时候,刷新的时候得有效果的。
    //    如果是 localStorage 那么已经是本地存储了，登出再重新登陆是不会影响 this.skLoading 的值的就看不到效果了
    storage: sessionStorage
  }
})
