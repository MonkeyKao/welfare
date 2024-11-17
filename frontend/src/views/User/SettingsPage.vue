<template>

  <div class="flex bg-[#90c700] p-3 items-center fixed z-10 w-full px-6">
    <router-link to="/user" class=""><PhArrowUUpLeft :size="32"  weight="bold"  color="#000"/></router-link>
    <label class="ml-5 text-H2 font-bold">設定</label>
  </div>

  <div class="px-8 flex flex-col  mt-[80px] space-y-8">
    <div
      v-for="(item, index) in menuItems"
      :key="index"
      class="flex items-center"
      @click="router.push(item.path)"
    >
      <component :is="item.icon" :size="28" color="#4d4d4d" />
      <label class="text-H2 flex ml-3">{{ item.label }}</label>
    </div>
  </div>
</template>
<script setup lang="ts">
import router from "@/router";
import { PhArrowUUpLeft, PhLink, PhDetective } from "@phosphor-icons/vue";
const menuItems = [
  { path: "/user/link-account", label: "連結帳戶", icon: PhLink },
  { path: "/elderlysearch", label: "老人模式", icon: PhDetective },
  { path: "/user/link-account", label: "連結帳戶", icon: PhLink },
  { path: "/user/link-account", label: "連結帳戶", icon: PhLink },
  { path: "/user/link-account", label: "連結帳戶", icon: PhLink },
  { path: "/user/link-account", label: "連結帳戶", icon: PhLink },
];
</script>
