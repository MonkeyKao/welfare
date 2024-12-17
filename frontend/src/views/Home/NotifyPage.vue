<template>
  <HeaderBar>通知</HeaderBar>

  <div class="mt-[70px] space-y-5 mb-2"><!-- mb-2 在修最後一項下斜線不見的問題-->
    <div
      v-for="(message, index) in messages"
      :key="index"
      class="flex items-end space-x-2 px-4"
    >
      <div > <!--這個div在解決icon大小無法正確顯示的問題-->
        <PhUserSquare :size="40" />        
      </div>

      <div class="flex p-4 bg-white shadow-md rounded-tl-lg rounded-tr-lg rounded-br-lg">
        <p class="text-gray-800 ">{{ message.text }}</p>
      </div>
      <div class="">
        <PhHeartStraight
          :size="28"
          :weight="message.isFavorited ? 'fill' : 'regular'"
          class="cursor-pointer"
          @click="toggleFavorite(index)"
        />
      </div>    
  </div>

  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { PhHeartStraight,PhUserSquare } from '@phosphor-icons/vue';
import HeaderBar from '@/components/HeaderBar.vue';

const messages = ref([
  { text: '【文化部消息】震驚!!!114年文化幣每人可領取2000元', isFavorited: false },
  { text: '【文化部消息】花蓮振興文化幣再加碼 16-22歲青年快衝一波', isFavorited: false },
  { text: '【文化部消息】花蓮振興文化幣再加碼 16-22歲青年快衝一波', isFavorited: false },
  { text: '【文化部消息】花蓮振興文化幣再加碼 16-22歲青年快衝一波', isFavorited: false },
  { text: '【文化部消息】花蓮振興文化幣再加碼 16-22歲青年快衝一波', isFavorited: false },
  { text: '【文化部消息】花蓮振興文化幣再加碼 16-22歲青年快衝一波', isFavorited: false },
]);

const toggleFavorite = (index: number) => {
  messages.value[index].isFavorited = !messages.value[index].isFavorited;
  messages.value.splice(index, 1);
};
</script>

<style scoped>
</style>