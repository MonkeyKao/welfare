<template>
  <HeaderBar>設定</HeaderBar>

  <div class="px-8 flex flex-col  mt-4 space-y-8">
    <div v-for="(item, index) in menuItems" :key="index" class="flex items-center" @click="router.push(item.path)">
      <component :is="item.icon" :size="28" color="#4d4d4d" />
      <label class="text-H2 flex ml-3">{{ item.label }}</label>
    </div>
  </div>
</template>
<script setup lang="ts">
import HeaderBar from "@/components/headerBar.vue";
import router from "@/router";
import { PhArrowUUpLeft, PhLink, PhDetective } from "@phosphor-icons/vue";
const menuItems = [
  { path: "/user/link-account", label: "連結帳戶", icon: PhLink },
  { path: "/older-home", label: "老人模式", icon: PhDetective },
];
</script>
