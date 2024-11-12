<template>
  <div class="flex items-center p-5 pt-14 bg-[#92c700]">
    <router-link to="/user" class="forgot-password-link">
      <PhArrowUUpLeft :size="32" color="#000" class="ml-5" />
    </router-link>
    <label class="ml-5 text-3xl font-bold">設定</label>
  </div>

  <div class="p-10 flex flex-col gap-2 divide-y">
    <div
      v-for="(item, index) in menuItems"
      :key="index"
      class="flex items-center"
      @click="router.push(item.path)"
    >
      <component :is="item.icon" :size="40" color="#4d4d4d" />
      <label class="text-lg font-bold flex ml-3">{{ item.label }}</label>
    </div>
  </div>
</template>
<script setup lang="ts">
import router from "@/router";
import { PhArrowUUpLeft, PhLink } from "@phosphor-icons/vue";
const menuItems = [
  { path: "/user/link-account", label: "連結帳戶", icon: PhLink },
];
</script>
