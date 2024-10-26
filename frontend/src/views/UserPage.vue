<template>
  <header class="flex flex-col items-center mt-5">
    <Avatar />
    <label class="text-3xl font-bold">用戶名</label>
    <hr
      class="border-0 h-px bg-gradient-to-r from-gray-300 via-gray-800 to-gray-300 w-full mt-4"
    />
  </header>

  <div class="flex flex-col items-center justify-center mt-5">
    <RouterLink to="/personal-data" class="w-full">
      <div class="flex items-center mt-5 mx-10">
        <PhUserCircle :size="32" color="#4d4d4d" />
        <label class="text-lg font-bold flex ml-5">個人資料</label>
      </div>
    </RouterLink>

    <RouterLink to="/link-account" class="w-full">
      <div class="flex items-center mt-5 mx-10">
        <PhLink :size="32" color="#4d4d4d" />
        <label class="text-lg font-bold flex ml-5">連結帳戶</label>
      </div>
    </RouterLink>

    <RouterLink to="/family" class="w-full">
      <div class="flex items-center mt-5 mx-10">
        <PhUsersThree :size="32" color="#4d4d4d" />
        <label class="text-lg font-bold flex ml-5">家庭</label>
      </div>
    </RouterLink>

    <RouterLink to="/settings" class="w-full">
      <div class="flex items-center mt-5 mx-10">
        <PhGear :size="32" color="#4d4d4d" />
        <label class="text-lg font-bold flex ml-5">設定</label>
      </div>
    </RouterLink>

    <RouterLink to="/qa" class="w-full">
      <div class="flex items-center mt-5 mx-10">
        <PhQuestion :size="32" color="#4d4d4d" />
        <label class="text-lg font-bold flex ml-5">常見問題</label>
      </div>
    </RouterLink>
  </div>
  <div class="mt-5 mx-10">
    <router-link to="/login" class="forgot-password-link"
      ><button
        class="w-full bg-[#90c700] text-white font-bold text-2xl py-3 rounded shadow-md"
        type="button"
      >
        登出
      </button></router-link
    >
  </div>
</template>

<script setup lang="ts">
import Avatar from "@/components/Avatar.vue";
import {
  PhGear,
  PhLink,
  PhQuestion,
  PhUserCircle,
  PhUsersThree,
} from "@phosphor-icons/vue";
</script>

<style scoped>
/* 可選樣式 */
</style>
