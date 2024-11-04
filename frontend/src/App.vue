<script setup>
import BottomNav from "@/components/BottomNav.vue"; // 引入底部導航欄
import { onMounted } from "vue";
import { RouterView, useRoute } from "vue-router"; // 引入 RouterView 用來動態渲染路由對應的頁面
import { useUserStore } from "./store/userStroe";
import { useWelfareStore } from "./store/welfareStroe";

const route = useRoute(); // 獲取當前路由
const userStore = useUserStore();
const welfareStore = useWelfareStore();
userStore.fetchUser()
welfareStore.fetchWelfare();
</script>

<template>
  <!-- 渲染當前頁面 -->
  <div class=" overflow-y-auto h-full flex flex-col justify-between">

    <div class="overflow-auto">
      <RouterView />
    </div>

    <div class="flex flex-col-reverse h-20" v-if="
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
