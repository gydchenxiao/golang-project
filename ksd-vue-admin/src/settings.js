export default {
  // 配置首页访问的时候，自动跳转到你指定defaultPage
  defaultPage: { path: '/dashboard', replace: true },
  // 指定菜单导航的个数
  menuCount: 20,
  // 菜单折叠屏幕宽度的大小
  collapseWidth: 992,
  // 菜单隐藏屏幕宽度的大小
  hiddenWidth: 640,
  // 是否显示菜单导航
  showMenuTab: true,
  // 菜单的宽度
  slideWidth: 180,
  // 表格高度
  tableHeight() {
    return this.showMenuTab ? 'calc(100vh - 242px)' : 'calc(100vh - 202px)'
  }
}
