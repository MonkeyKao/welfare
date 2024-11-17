<script setup>
import BottomNav from "@/components/BottomNav.vue"; // 引入底部導航欄
import { onMounted, ref } from "vue";
import { RouterView, useRoute } from "vue-router"; // 引入 RouterView 用來動態渲染路由對應的頁面
import { useUserStore } from "./store/userStroe";
import { useWelfareStore } from "./store/welfareStroe";
import { provide } from 'vue'
import TipMsg from "./components/tipMsg.vue";

import aiinput from "@/components/aiinput.vue";
import { useFavoriteStore } from "./store/favorite";

const route = useRoute(); // 獲取當前路由
const userStore = useUserStore();
const welfareStore = useWelfareStore();
const favoriteStore = useFavoriteStore();
userStore.fetchUser()
welfareStore.fetchWelfare();
favoriteStore.fetchFavorite();


const tipmsg = ref(null)

onMounted(() => {
  provide('showMsg',tipmsg.value.showMsg)
})
</script>

<template>
  <!-- 渲染當前頁面 -->
  <div class="overflow-y-auto h-full flex flex-col justify-between">
    <div v-if="['AiPage'].includes(route.name)" class="overflow-auto pb-10">
      <RouterView />
    </div>
    <div v-else class="overflow-auto">
      <RouterView />
    </div>
    <div class="flex flex-col-reverse">
      <div      
        v-if="
          ![
            'LoginPage',
            'RegisterPage',
            'ForgetPasswordPage',
            'ResetPasswordPage',
            'CreateProfile',
            'VerifyPage',
            'LikePage',
            'NotifyPage',
            'CreateProfilePage',
            'QAPage',
            'Settings',
            'LinkAccountPage',
            'ElderlySearchPage',
            'RegionSelectionPage',
            'ServiceSelectionPage',
            'SearchResultPage',
            'EditPersonalDataPage',
            'PersonalDataPage',
          ].includes(route.name)
        "
      >
        <BottomNav />
      </div>
      
    </div>

    
  </div>

  <TipMsg ref="tipmsg"></TipMsg>
</template>

<style scoped>
/* 你可以在這裡根據需要定義樣式 */
</style>
<style scoped></style>
