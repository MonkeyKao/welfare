<template>
  <HeaderBar>設定</HeaderBar>

  <div class="mt-[80px] px-4 flex flex-col space-y-8">
    <div v-for="(item, index) in menuItems" :key="index" class="flex items-center" @click="router.push(item.path)">
      <component :is="item.icon" :size="28" color="#4d4d4d" />
      <label class="text-H2 flex ml-3">{{ item.label }}</label>
    </div>
    <div class="flex items-center justify-between">
      <div class="flex items-center">
        <PhAddressBook :size="28" color="#4d4d4d" />
        <label class="text-H2 flex ml-3">自動篩選用戶資料</label>
      </div>

      <input type="checkbox" class="toggle toggle-success " />
    </div>
  </div>
</template>
<script setup lang="ts">
import HeaderBar from "@/components/HeaderBar.vue";
import router from "@/router";
import { PhLink, PhDetective, PhAddressBook } from "@phosphor-icons/vue";
const menuItems = [
  { path: "/user/link-account", label: "連結帳戶", icon: PhLink },
  { path: "/accessibility/home", label: "老人模式", icon: PhDetective },
];
</script>
