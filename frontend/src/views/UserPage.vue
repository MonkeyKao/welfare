<template>
  <div class="flex flex-col w-full divide-y-2 container">
    <div class="flex flex-col items-center gap-2 mt-10 mb-5 flex-shrink-0">
      <Avatar />
      <label class="text-3xl font-bold">{{
        user?.name ? user.name : "未登錄"
      }}</label>
    </div>

    <div class="p-10 flex flex-col gap-5">
      <div
        v-for="(item, index) in menuItems"
        :key="index"
        class="flex items-center"
        @click="router.push(item.path)"
      >
        <component :is="item.icon" :size="40" color="#4d4d4d" />
        <label class="text-lg font-bold flex ml-3">{{ item.label }}</label>
      </div>

      <router-link v-if="user" to="/account/login" class="forgot-password-link"
        ><button
          class="w-full bg-[#90c700] text-white font-bold text-2xl py-3 rounded shadow-md"
          type="button"
          @click="logoutHandler"
        >
          登出
        </button></router-link
      >
      <router-link v-else to="/account/login" class="forgot-password-link"
        ><button
          class="w-full bg-[#90c700] text-white font-bold text-2xl py-3 rounded shadow-md"
          type="button"
        >
          登入
        </button></router-link
      >
    </div>
  </div>
</template>

<script setup lang="ts">
import Avatar from "@/components/Avatar.vue";
import router from "@/router";
import { useUserStore } from "@/store/userStroe";
import {
  PhGear,
  PhQuestion,
  PhUserCircle,
  PhUsersThree,
} from "@phosphor-icons/vue";
import { computed } from "vue";

const userStroe = useUserStore();
const user = computed(() => userStroe.user);
const menuItems = [
  { path: "/user/personal-data", label: "個人資料", icon: PhUserCircle },
  // { path: "/user/link-account", label: "連結帳戶", icon: PhLink },
  { path: "/user/family", label: "家庭", icon: PhUsersThree },
  { path: "/user/settings", label: "設定", icon: PhGear },
  { path: "/user/qa", label: "常見問題", icon: PhQuestion },
];

const logoutHandler = () => {
  localStorage.removeItem("token");
  router.push("/login");
};
</script>

<style scoped>
/* 可選樣式 */
</style>
