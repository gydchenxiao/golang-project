<template>
  <div class="header-cont">
    <div class="left">
      <h1>
        <router-link to="/">{{ t('uniLiveMangeSystem') }}</router-link>
      </h1>
    </div>

    <div class="right flex-center">
      <!--全屏处理-->
      <div class="fullbox">
        <!-- <el-icon @click="handleFullChange(true)" v-if="!screenfullFlag" color="#fff">
          <FullScreen />
        </el-icon>
        <el-icon @click="handleFullChange(false)" v-else color="#fff">
          <Aim />
        </el-icon> -->

        <!-- 使用 /fonts/iconfont.css 中的图标 -->
        <span @click="handleFullChange(true)" v-if="!screenfullFlag" class="icon-quanping fz20"></span>
        <span @click="handleFullChange(false)" v-else class="icon-quanping1 fz20"></span>
      </div>
      <!--国际化-->
      <div class="lang gap">
        <span class="item" :class="{ active: locale === 'zh-cn' }" @click="changeLanguage('zh-cn')">简体中文</span>
        /
        <span class="item" :class="{ active: locale === 'en' }" @click="changeLanguage('en')">EN</span>
      </div>

      <template v-if="isLogin">
        <div class="gap">
          <router-link to="/personal/message">
            <el-badge :is-dot="!!unReadCount">
              <el-icon>
                <message />
              </el-icon>
            </el-badge>
          </router-link>
        </div>
        <el-dropdown trigger="click" @command="handleCommand">
          <div class="flex-center cursor">
            <!-- userStore.user 对象中有当前用户的所有信息的哦 -->
            <el-avatar size="small" :src="userStore.user.avatar" />
            <span class="uname"> {{ userName }}</span>
            <el-icon>
              <caret-bottom />
            </el-icon>
          </div>
          <template #dropdown>
            <!-- <el-dropdown-menu>
              <el-dropdown-item style="color: rgb(9, 75, 241);">当前角色：{{ currentRole.name }}</el-dropdown-item>
              <el-dropdown-item v-for="(item, index) in otherRoles" :key="index">切换角色：{{ item.name }}</el-dropdown-item>
              <el-dropdown-item divided command="toPersonal"><el-icon>
                  <User />
                </el-icon>{{ t('personalCenter') }}</el-dropdown-item>
              <el-dropdown-item divided command="toLogout"><el-icon>
                  <Pointer />
                </el-icon>{{ t('logout') }}</el-dropdown-item>
            </el-dropdown-menu> -->

            <el-dropdown-menu>
              <el-dropdown-item style="color: blue;">当前角色：{{ currRoleName }}</el-dropdown-item>
              <el-dropdown-item @click="handleChangeRole(item.id, item.roleName, item.roleCode)"
                v-for="(item, index) in otherRoleList" :key="index">切换角色：{{ item.roleName }}</el-dropdown-item>
              <el-dropdown-item divided command="toPersonal"><el-icon>
                  <User />
                </el-icon>{{ t('personalCenter') }}</el-dropdown-item>
              <el-dropdown-item divided command="toLogout"><el-icon>
                  <Pointer />
                </el-icon>{{ t('logout') }}</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </template>
      <template v-else-if="$route.name !== 'Login'">
        <router-link to="/login">{{ t('login') }}</router-link>
      </template>
    </div>
  </div>
</template>

<script setup>
// 状态管理获取登录信息
import KVA from '@/utils/kva.js'
import { useUserStore } from '@/stores/user.js'

const userStore = useUserStore()
// 路由跳转
const router = useRouter();
// 国际化处理
const { locale, t } = useI18n();

// 获取登录的信息
const isLogin = computed(() => userStore.token);
const userName = computed(() => userStore.userName)
// 消息未读取的数量
const unReadCount = computed(() => 100);
// 获取第一个以后角色方便进行切换
// const currentRole = computed(() => userStore.roles && userStore.roles.length && userStore.roles[0]) // 默认给出第一个 role 作为当前用户的 role
const currRoleId = computed(() => userStore.currRoleId)
const currRoleName = computed(() => userStore.currRoleName)
// 获取角色列表
// const otherRoles = computed(() => userStore.roles && userStore.roles.length > 1 && userStore.roles.filter((c, index) => index > 0)) // 把当前角色从总的 roles 中去除
const otherRoleList = computed(() => userStore.roles.filter(c => c.id != userStore.currRoleId))

// 全屏处理
import screenfull from 'screenfull'
// 状态管理全屏按钮切换
const screenfullFlag = ref(false)
// 全屏事件处理
const handleFullChange = (flag) => {
  screenfull.toggle()
  screenfullFlag.value = flag
}

// 下拉事件处理
const commands = ({
  //个人中心跳转
  toPersonal: () => {
    router.push('/personal')
  },
  // 退出方法
  toLogout: () => {
    KVA.confirm("退出提示", "您确定要离开吗?", { icon: "error" }).then(res => {
      userStore.LoginOut();
    })
  }
});

function handleCommand(command) {
  commands[command] && commands[command]();
}

// 语言切换
function changeLanguage(lang) {
  // 把选择的语言进行切换
  locale.value = lang
  // 切换以后记得把本地缓存进行修改，否则只会生效当前，刷新就还原。
  localStorage.setItem('ksd-kva-language', lang)
}

// 用于读取本地缓存存储的语言是什么？
function initReadLocale() {
  locale.value = localStorage.getItem("ksd-kva-language") || 'zh-cn'
}

// 角色切换
const handleChangeRole = (roleId, roleName, roleCode) => {
  userStore.handlePianaRole(roleId, roleName, roleCode);
}

onMounted(() => {
  initReadLocale();
})

</script>
<style lang="scss">
.header-cont {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 100%;
  padding: 0 20px;

  a {
    color: inherit;
    text-decoration: none;
  }

  h1 {
    margin: 0;
    font-size: 20px;
  }

  .gap {
    margin-right: 20px;
  }

  .right {
    .uname {
      margin-left: 10px;
    }

    .fullbox {
      margin-right: 20px;
      cursor: pointer;
    }

    .lang {
      font-size: 14px;

      .item {
        cursor: pointer;

        &.active {
          font-size: 16px;
          font-weight: bold;
        }
      }
    }
  }

  .el-dropdown {
    color: inherit;
  }
}
</style>