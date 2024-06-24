<template>
  <el-skeleton animated :loading="skLoading">
    <template #template>
      <div class="page-sidebar">
        <div class="el-menu sidemenu" style="background: #fff;padding: 10px;">
          <el-skeleton-item variant="div" v-for="(item, i) in menuTree" :key="i" style="height: 45px" />
        </div>
      </div>
    </template>
    <template #default>
      <div class="page-sidebar" v-show="isHidden">
        <div class="collape-bar">
          <el-icon class="cursor" @click="isCollapse = !isCollapse">
            <expand v-if="isCollapse" />
            <fold v-else />
          </el-icon>
        </div>
        <el-menu active-text-color="#333" background-color="#ffffff" text-color="#333" router
          :default-active="defaultActive" class="sidemenu" :collapse="isCollapse">
          <template v-for="(item, i) in menuTree" :key="i">
            <template v-if="item.children && item.children?.length && item.hidden == 1">
              <el-sub-menu :index="item.path">
                <template #title>
                  <el-icon v-if="item.icon">
                    <component :is="item.icon"></component>
                  </el-icon>
                  <span>{{ t(`menu.${item.name}`) }}</span>
                </template>
                <template v-for="(child, ci) in item.children?.sort((a, b) => a.sort - b.sort)" :key="ci">
                  <el-menu-item :index="child.path" v-if="child.hidden == 1">
                    <el-icon>
                      <component :is="child.icon"></component>
                    </el-icon>
                    {{ t(`menu.${child.name}`) }}
                  </el-menu-item>
                </template>
              </el-sub-menu>
            </template>
            <template v-else>
              <el-menu-item :index="item.path" v-if="item.hidden == 1">
                <el-icon v-if="item.icon">
                  <component :is="item.icon"></component>
                </el-icon>
                <span>{{ t(`menu.${item.name}`) }}</span>
              </el-menu-item>
            </template>
          </template>
        </el-menu>
      </div>
    </template>
  </el-skeleton>
</template>
  
<script  setup>
import { useUserStore } from '@/stores/user.js'
import settings from '@/settings.js'
import { useSkeletonStore } from '@/stores/skeleton.js'

const skeletonStore = useSkeletonStore()
const skLoading = computed(() => skeletonStore.skLoading)
const swidth = ref(settings.slideWidth + "px")
// 这个是用来获取当前访问的路由信息,
const route = useRoute();
const { t } = useI18n();
// 默认情况下不折叠
const isCollapse = ref(false)
const isHidden = ref(true)
// 根据当前路由来激活菜单
const defaultActive = computed(() => (route.path))
// 获取状态管理的菜单信息
const userStore = useUserStore();
// 如何获取菜单数据呢？
const menuTree = computed(() => userStore.menuTree)

// 获取屏幕宽度
const screenWidth = ref(window.innerWidth)
onMounted(() => {
  // 然后监听浏览器的窗口resize事件，只要浏览器发生大小的变化就会触发。
  window.addEventListener("resize", () => {
    screenWidth.value = window.innerWidth
  })
})

//watch监听屏幕宽度的变化，进行侧边栏的收缩和展开
watch(screenWidth, (newValue, oldValue) => {
  // 如果浏览器的宽度小于992px时候就会菜单的处理折叠状态
  isCollapse.value = newValue < settings.collapseWidth
  isHidden.value = !(newValue < settings.hiddenWidth)
}, { immediate: true })

</script>
<style lang="scss">
.page-sidebar {
  height: calc(100vh - 90px);
  overflow: hidden auto;

  .sidemenu.el-menu,
  .sidemenu .el-sub-menu>.el-menu {
    --el-menu-text-color: #ccc;
    --el-menu-hover-bg-color: #060251;
    --el-menu-border-color: transparent;
    --el-menu-bg-color: #001529;

    .el-icon:first-child {
      top: -2px
    }

    .el-menu-item {
      .el-icon:first-child {
        top: -2px
      }

      &.is-active {
        background-color: #e6f7ff;
        color: #1890ff
      }
    }
  }

  /* elmenu菜单的折叠效果是通过属性：
      :collapse="isCollapse"  原理就在控制在不停切换elmenu="el-menu--collapse"样式信息
      1：ture 就折叠，就会使用图标宽度+padding作为菜单宽度
      2: false 就不折叠，那么就使用默认宽度：200px

      下面这行css是什么意思：
      如果菜单上存在el-menu--collapse样式就说明是折叠状态，就使用图标宽度+padding作为菜单宽度
      否则：就用我的width:200作为菜单宽度
   */
  .sidemenu.el-menu:not(.el-menu--collapse) {
    width: v-bind(swidth);
  }

  .collape-bar {
    color: #333;
    font-size: 16px;
    line-height: 36px;
    position: fixed;
    z-index: 2;
    left: 20px;
    bottom: 0;

    .c-icon {
      cursor: pointer;
    }
  }
}
</style>