<template>
  <div class="w-full divide-y-2 container">

    <header class="flex flex-col items-center gap-2 mt-10 mb-5">
      <Avatar />
      <label class="text-3xl font-bold">{{ user?.name ? user.name : '未登錄' }}</label>
    </header>

    <div class="p-10 flex flex-col gap-4 ">
      <div v-for="(item, index) in menuItems" :key="index" class="flex items-center" @click="router.push(item.path)">
        <component :is="item.icon" :size="40" color="#4d4d4d" />
        <label class="text-lg font-bold flex ml-3">{{item.label}}</label>
      </div>


      <div class="hidden">
        <router-link v-if="user" to="/login" class="forgot-password-link"><button
            class="w-full bg-[#90c700] text-white font-bold text-2xl py-3 rounded shadow-md" type="button"
            @click="logoutHandler">
            登出
          </button></router-link>
        <router-link v-else to="/login" class="forgot-password-link"><button
            class="w-full bg-[#90c700] text-white font-bold text-2xl py-3 rounded shadow-md" type="button">
            登入
          </button></router-link>

      </div>


    </div>

  </div>



</template>

<script setup lang="ts">
import request from "@/axios";
import Avatar from "@/components/Avatar.vue";
import User from "@/model/user";
import router from "@/router";
import {
  PhGear,
  PhLink,
  PhQuestion,
  PhUserCircle,
  PhUsersThree,
} from "@phosphor-icons/vue";
import { onMounted, ref } from "vue";

const menuItems = [
  { path: "/personal-data", label: "個人資料", icon: PhUserCircle },
  { path: "/link-account", label: "連結帳戶", icon: PhLink },
  { path: "/family", label: "家庭", icon: PhUsersThree },
  { path: "/settings", label: "設定", icon: PhGear },
  { path: "/qa", label: "常見問題", icon: PhQuestion },
]

const logoutHandler = () => {
  localStorage.removeItem("token")
  router.push("/login")
}

const user = ref<User>()

onMounted(async () => {
  const result = await request.get("users");
  user.value = result.data

})
</script>

<style scoped>
/* 可選樣式 */
</style>
