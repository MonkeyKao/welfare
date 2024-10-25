<script setup>
import { RouterView } from 'vue-router'; // 引入 RouterView 用來動態渲染路由對應的頁面
import BottomNav from '@/components/BottomNav.vue'; // 引入底部導航欄
import { useRoute } from 'vue-router'; // 引入 useRoute

const route = useRoute(); // 獲取當前路由
</script>

<template>
  <!-- 渲染當前頁面 -->
  <RouterView />

  <BottomNav />
</template>

<style scoped>
/* 你可以在這裡根據需要定義樣式 */
</style>