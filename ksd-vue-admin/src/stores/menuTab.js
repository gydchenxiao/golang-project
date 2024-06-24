import { defineStore } from 'pinia'
import router from '@/router/index.js'
import KVA from '@/utils/kva.js'

export const useMenuTabStore = defineStore('menuTab', {
  // 定义状态
  state: () => ({
    skLoading: true,
    activePath: '/dashboard',
    menuList: [{ name: 'DashBoard', path: '/dashboard' }]
  }),

  // 定义动作
  actions: {
    // 点击改变菜单
    changeMenuTab(path) {
      this.activePath = path
      router.push(path)
    },

    // 清空导航栏
    clear() {
      this.skLoading = false
      // this.menuList = [{ name: 'DashBoard', path: '/dashboard' }] // 不一定会有 dashboard
      this.menuList = []
    },

    // 添加菜单
    addMenuTab({ name, path }) {
      if (name && path) {
        var index = this.menuList.findIndex((m) => m.path == path)
        // 激活当前路径
        this.activePath = path
        // 没有找到就插入 menuList 中，有的就不重复加入了
        if (index === -1) {
          this.menuList.push({ name, path })
        }
      }
    },
    // 移除
    removeMenuTab(path) {
      var tabs = this.menuList
      // 1. 关闭自己当前页面，需要再选一个页面激活
      var activeTabPath = this.activePath
      if (activeTabPath === path) {
        tabs.forEach((tab, index) => {
          if (tab.path == path) {
            const nextTab = tabs[index + 1] || tabs[index - 1] // activeTabPath 激活路径，在删除位置的左右选一个没有存在的路径
            if (nextTab) {
              activeTabPath = nextTab.path
            }
          }
        })
      }

      // 2. 不关闭当前页面，不需要再选一个页面激活
      this.activePath = activeTabPath
      var index = this.menuList.findIndex((m) => m.path == path)
      // 开始删除
      this.menuList.splice(index, 1)
      // 删除激活控制面板
      router.push(activeTabPath)
    }
  },
  persist: {
    // key: 'kva-pinia-menutabs',
    // storage: localStorage // sessionStorage

    // 上面两行的写法是有问题的哦，没有实现正真的 localStorage 本地存储 ---> enabled: true 是很关键的
    enabled: true, // 开启数据持久化
    strategies: [
      {
        key: 'kva-pinia-menutabs', // 给一个要保存的名称
        storage: localStorage // sessionStorage / localStorage 存储方式
      }
    ]
  }
})
