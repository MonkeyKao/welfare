<template>
  <div class="h-full flex flex-col w-full divide-y">
    <div class="basis-1">
      <TopNav />
    </div>

    <div class="overflow-auto">
      <!-- 根據 API 回應資料動態生成 HomeInsideText 元件 -->
      <HomeInsideText
        @clickFavorited="(data) => {}"
        v-for="(item, index) in welfareData"
        :key="index"
        :data="item"
      />
    </div>
  </div>
</template>

<script setup>
import request from "@/axios";
import HomeInsideText from "@/components/HomeInsideText.vue";
import TopNav from "@/components/TopNav.vue";
import { onMounted, ref } from "vue";

const welfareData = ref([]); // 定義響應式數據

const getWelfare = async () => {
  try {
    const result = await request.get("/welfare"); // 發送請求
    welfareData.value = result.data; // 存儲數據
    console.log(result.data);
  } catch (error) {
    console.error("error", error); // 處理錯誤
  }
};

onMounted(() => {
  getWelfare(); // 在組件掛載後調用函數
});
</script>

<style scoped>
/* 可選樣式 */
</style>
