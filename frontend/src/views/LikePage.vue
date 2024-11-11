<template>
  <div class="flex justify-start items-center gap-3 p-4 bg-[#92c700]">
    <router-link to="/home" class="forgot-password-link">
      <PhArrowUUpLeft :size="32" color="#000" />
    </router-link>
    <label class="text-3xl font-bold">收藏</label>
  </div>

  <div
    v-for="(message, index) in messages"
    :key="index"
    class="flex items-end p-4"
  >
    <img src="@/images/user-square.png" alt="類別圖片" class="w-10 h-10" />
    <div class="flex p-4 bg-white shadow-md rounded-tl-lg rounded-tr-lg rounded-br-lg">
      <p class="text-gray-800 ">{{ message.text }}</p>
    </div>
    <div class="ml-4">
      <PhHeartStraight
        :size="28"
        :weight="message.isFavorited ? 'fill' : 'regular'"
        class="cursor-pointer"
        @click="toggleFavorite(index)"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { PhArrowUUpLeft, PhHeartStraight } from '@phosphor-icons/vue';

const messages = ref([
  { text: '【文化部消息】震驚!!!114年文化幣每人可領取2000元', isFavorited: true },
  { text: '【文化部消息】花蓮振興文化幣再加碼 16-22歲青年快衝一波', isFavorited: true },
]);

const toggleFavorite = (index: number) => {
  messages.value[index].isFavorited = !messages.value[index].isFavorited;
  messages.value.splice(index, 1);
};
</script>

<style scoped>
</style>