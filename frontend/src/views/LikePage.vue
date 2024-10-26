<template>
  <header class="top-0 left-0 right-0 flex items-center p-5 pt-14 bg-[#92c700]">
    <router-link to="/home" class="forgot-password-link">
      <PhArrowUUpLeft :size="32" color="#000" class="ml-5" />
    </router-link>
    <label class="ml-5 text-3xl font-bold">收藏</label>
  </header>
  
  <div class="flex items-end p-4">
    <img src="@/images/user-square.png" alt="類別圖片" class="w-10 h-10" />
    <div class="flex-1 ml-4 p-3 bg-white shadow-md rounded-tl-lg rounded-tr-lg rounded-br-lg">
      <p class="text-gray-800">
        哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈哈
      </p>
    </div>
    <div class="ml-4">
      <PhHeartStraight
        :size="28"
        :weight="isFavorited ? 'fill' : 'regular'"
        class="cursor-pointer"
        @click="toggleFavorite"
      />
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { PhArrowUUpLeft, PhHeartStraight } from '@phosphor-icons/vue';

const isFavorited = ref(true);

function toggleFavorite() {
  isFavorited.value = !isFavorited.value;
}
</script>

<style scoped>
.icon {
  transition: color 0.3s;
}
</style>