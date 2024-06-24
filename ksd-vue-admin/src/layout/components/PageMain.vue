<template>
  <!-- Transition 是展示动画效果使用的（可以使用下面的 fade、slide-fade、bounce 样式） -->
  <!-- Vue3 内置组件 KeepAlive：功能是在多个组件间动态切换时缓存被移除的组件实例 -->
  <!-- Vue3 内置组件 Suspense -->
  <RouterView v-slot="{ Component }">
    <template v-if="Component">
      <Transition name="fade">
        <KeepAlive>
          <Suspense>
            <component :is="Component"></component>
            <template #fallback>
              <H1>加载中</H1>
            </template>
          </Suspense>
        </KeepAlive>
      </Transition>
    </template>
  </RouterView>
</template>

<script>
export default {

}
</script>

<style lang="scss">
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}


.slide-fade-enter-active {
  transition: all 0.3s ease-out;
}

.slide-fade-leave-active {
  transition: all 0.8s cubic-bezier(1, 0.5, 0.8, 1);
}

.slide-fade-enter-from,
.slide-fade-leave-to {
  transform: translateX(20px);
  opacity: 0;
}



.bounce-enter-active {
  animation: bounce-in 0.5s;
}

.bounce-leave-active {
  animation: bounce-in 0.5s reverse;
}

@keyframes bounce-in {
  0% {
    transform: scale(0);
  }

  50% {
    transform: scale(1.25);
  }

  100% {
    transform: scale(1);
  }
}
</style>
