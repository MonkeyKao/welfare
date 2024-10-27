<script setup>
import BottomNav from "@/components/BottomNav.vue"; // 引入底部導航欄
import { RouterView, useRoute } from "vue-router"; // 引入 RouterView 用來動態渲染路由對應的頁面

const route = useRoute(); // 獲取當前路由
</script>

<template>
  <!-- 渲染當前頁面 -->
  <div class=" overflow-y-auto h-full flex flex-col">

    <div class=" overflow-auto basis-11/12">
      <RouterView />
    </div>

    <div class="basis-1/12" v-if="
      ![
        'LoginPage',
        'RegisterPage',
        'ForgetPasswordPage',
        'ResetPasswordPage',
        'CreateProfile',
        'VerifyPage',
        'LikePage',
        'NotifyPage',
      ].includes(route.name)
    ">
      <BottomNav />
    </div>
  </div>


</template>

<style scoped>
/* 你可以在這裡根據需要定義樣式 */
</style>
<style scoped></style>
