<template>
  <div class="login-box">
    <div class="imgbox">
      <div class="bgblue"></div>
      <!-- 设置动态效果 -->
      <ul class="circles">
        <li></li>
        <li></li>
        <li></li>
        <li></li>
        <li></li>
        <li></li>
        <li></li>
        <li></li>
        <li></li>
        <li></li>
      </ul>
      <div class="w-full max-w-md">
        <div class="fz48 cof fw">欢迎光临</div>
        <div class="fz14 cof" style="margin-top:10px;line-height:24px;">欢迎来到好玩俱乐部，在这里和志同道合的朋友一起分享有趣的故事，一起组织有趣的活动...
        </div>
      </div>
    </div>
    <div class="loginbox">
      <div class="login-wrap">
        <h1 class="header fz32">{{ ctitle }}</h1>
        <form action="#">
          <div class="ksd-el-items"><input type="text" v-model="loginUser.account" class="ksd-login-input"
              placeholder="请输入账号"></div>
          <div class="ksd-el-items"><input type="password" v-model="loginUser.password" class="ksd-login-input"
              placeholder="请输入密码" @keydown.enter="handleSubmit"></div>
          <div class="ksd-el-items pr">
            <input type="text" class="ksd-login-input" maxlength="6" v-model="loginUser.code" placeholder="请输入验证码"
              @keydown.enter="handleSubmit">
            <img v-if="codeURL" class="codeurl" :src="codeURL" @click="handleGetCapatcha">
          </div>
          <div class="ksd-el-items" v-loading="btnLoading"><input type="button" @click.prevent="handleSubmit"
              class="ksd-login-btn" value="登录">
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { onMounted, reactive, ref } from 'vue';
import { useRouter, useRoute } from 'vue-router'
import { useUserStore } from '@/stores/user.js'
import { getCapatcha } from '@/api/code.js'
import KVA from "@/utils/kva.js";
import NProgress from 'nprogress'


// 按钮动画(使得点击登陆的时候，如果长时间没有反应需要等待的话，按钮是不能被点击的 ---> 间接防抖)
const btnLoading = ref(false)

// 显示右上角螺旋加载提示
NProgress.configure({ showSpinner: true })


// 定义一个路由对象 // 用做js跳转路由页面
const router = useRouter()
// 获取当前路由信息，用于获取当前路径参数，路径，HASH等
const route = useRoute();
const userStore = useUserStore();

// .env 配置文件中设置的 VITE_APP_TITLE
const ctitle = ref(import.meta.env.VITE_APP_TITLE)
// 准备接受图像验证码
const codeURL = ref("");
// 获取用户输入账号和验证码信息
const loginUser = reactive({
  code: "",
  account: "admin",
  password: "123456",
  codeId: ""
})

// 业务问题：看上面重复点击登陆按钮和点击验证码图片的时候是会刷新验证码的哦(防抖动可以解决这个问题)，以及提交的account和password可能会不安全的问题(加密之后再post请求服务端)


// 根据axios官方文档开始调用生成验证码的接口
const handleGetCapatcha = async () => {
  const resp = await getCapatcha() // 封装的获取验证码的请求
  // console.log('code resp', resp)
  const { b64sImg, codeId } = resp.data.data // 解构
  codeURL.value = b64sImg
  loginUser.codeId = codeId

  // 不结构的话可以使用下面的方式的哦
  // codeURL.value = resp.data.data.b64sImg
  // loginUser.codeId = resp.data.data.codeId
}

// 提交表单
const handleSubmit = async () => {
  // axios.post ---application/json---gin-request.body
  if (!loginUser.code) {
    btnLoading.value = false
    KVA.notifyError("请输入验证码")
    return;
  }
  if (!loginUser.account) {
    btnLoading.value = false
    KVA.notifyError("请输入账号")
    return;
  }
  if (!loginUser.password) {
    btnLoading.value = false
    KVA.notifyError("请输入密码")
    return;
  }

  // 把数据放入到状态管理中
  try {
    btnLoading.value = true
    await userStore.toLogin(loginUser)
    btnLoading.value = false
    // 这个会回退，回退登录页
    var path = route.query.path || "/"
    router.push(path)
    // 重定向，浏览器回退按钮不会到登录
    //router.push({ path: '/', replace: true })

    //完成进度条
    NProgress.done()
  } catch (e) {
    //完成进度条
    NProgress.done()
    btnLoading.value = false
    if (e.code === 60002) {
      loginUser.code = ""
      handleGetCapatcha()
    }
  }
}
// 用生命周期去加载生成验证码
onMounted(() => {
  handleGetCapatcha()
})

</script>

<style scoped lang="scss">
.pr {
  position: relative;
}

.codeurl {
  position: absolute;
  top: 5px;
  right: 5px;
  width: 140px;
}

.ksd-el-items {
  margin: 15px 0;
}

.ksd-login-input {
  border: 1px solid #eee;
  padding: 16px 8px;
  width: 100%;
  box-sizing: border-box;
  outline: none;
  border-radius: 4px;
}

.ksd-login-btn {
  border: 1px solid #eee;
  padding: 16px 8px;
  width: 100%;
  box-sizing: border-box;
  background: #2196F3;
  color: #fff;
  border-radius: 6px;
  cursor: pointer;
}

.ksd-login-btn:hover {
  background: #1789e7;
}

.login-box {
  display: flex;
  flex-wrap: wrap;
  background: #fff;

  .imgbox {
    width: 65%;
    height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    position: relative;
    background: url("../assets/imgs/login_background.jpg");
    background-size: cover;
    background-repeat: no-repeat;

    .bgblue {
      background-image: linear-gradient(to bottom, #4f46e5, #3b82f6);
      position: absolute;
      top: 0;
      left: 0;
      bottom: 0;
      right: 0;
      opacity: 0.75;
    }
  }

  .loginbox {
    width: 35%;
    height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;

    .header {
      margin-bottom: 30px;
    }

    .login-wrap {
      width: 560px;
      height: 444px;
      padding: 20px 100px;
      box-sizing: border-box;
      border-radius: 8px;
      box-shadow: 0 0 10px #fafafa;
      background: rgba(255, 255, 255, 0.6);
      text-align: center;
      display: flex;
      flex-direction: column;
      justify-content: center;
    }
  }
}

.circles {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  overflow: hidden;
}

.circles li {
  position: absolute;
  display: block;
  list-style: none;
  width: 20px;
  height: 20px;
  background: rgba(255, 255, 255, 0.2);
  animation: animate 25s linear infinite;
  bottom: -150px;
}

.circles li:nth-child(1) {
  left: 25%;
  width: 80px;
  height: 80px;
  animation-delay: 0s;
}

.circles li:nth-child(2) {
  left: 10%;
  width: 20px;
  height: 20px;
  animation-delay: 2s;
  animation-duration: 12s;
}

.circles li:nth-child(3) {
  left: 70%;
  width: 20px;
  height: 20px;
  animation-delay: 4s;
}

.circles li:nth-child(4) {
  left: 40%;
  width: 60px;
  height: 60px;
  animation-delay: 0s;
  animation-duration: 18s;
}

.circles li:nth-child(5) {
  left: 65%;
  width: 20px;
  height: 20px;
  animation-delay: 0s;
}

.circles li:nth-child(6) {
  left: 75%;
  width: 110px;
  height: 110px;
  animation-delay: 3s;
}

.circles li:nth-child(7) {
  left: 35%;
  width: 150px;
  height: 150px;
  animation-delay: 7s;
}

.circles li:nth-child(8) {
  left: 50%;
  width: 25px;
  height: 25px;
  animation-delay: 15s;
  animation-duration: 45s;
}

.circles li:nth-child(9) {
  left: 20%;
  width: 15px;
  height: 15px;
  animation-delay: 2s;
  animation-duration: 35s;
}

.circles li:nth-child(10) {
  left: 85%;
  width: 150px;
  height: 150px;
  animation-delay: 0s;
  animation-duration: 11s;
}

@keyframes animate {
  0% {
    transform: translateY(0) rotate(0deg);
    opacity: 1;
    border-radius: 0;
  }

  100% {
    transform: translateY(-1000px) rotate(720deg);
    opacity: 0;
    border-radius: 50%;
  }
}

.max-w-md {
  max-width: 28rem;
  position: relative;
  z-index: 10;
}
</style>