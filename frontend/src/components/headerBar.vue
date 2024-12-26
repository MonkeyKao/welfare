<template>
  <div class="flex h-16 bg-[#90c700] items-center justify-between p-4 fixed-header">
    <div class="flex gap-2">
      <div @click="router.go(-1)" class="forgot-password-link flex items-center">
        <PhArrowUUpLeft :size="28" color="#000" />
      </div>
      <label class="text-H2 font-bold">
        <slot></slot>
      </label>
    </div>

    <div>
      <slot name="siderBar"></slot>
    </div>
  </div>
</template>


<script setup lang="ts">
import router from "@/router";
import { PhArrowUUpLeft } from "@phosphor-icons/vue";
</script>

<style scoped>
  .fixed-header {
    position: fixed; /* 固定元素位置 */
    top: 0; /* 固定在頁面頂端 */
    left: 0; /* 左邊對齊 */
    width: 100%; /* 寬度為 100% */
    z-index: 1000; /* 確保它位於最上層 */
  }

  .forgot-password-link {
    cursor: pointer; /* 點擊手勢 */
    /* 123 */
  }
</style>