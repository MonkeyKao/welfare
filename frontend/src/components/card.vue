<template>
  <div class="w-full h-full flex-shrink-0 flex flex-col items-center justify-center rounded">
    <!-- 圖片顯示：設定3比4比例 -->
    <img :src="props.imageUrl" alt="Item Image" class="w-full h-3/4 object-cover rounded-t-md" />
    <!-- 文字顯示 -->
    <div class="mt-4 text-center text-2xl">
      <slot></slot>
    </div>
  </div>
</template>

<script setup lang="ts">
import { defineProps } from 'vue';

const props = defineProps(["imageUrl"]);
</script>

<style scoped></style>