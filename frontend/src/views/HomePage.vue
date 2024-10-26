<template>
  <div>
    <TopNav />
    <div class="mt-40">
      <!-- 根據 API 回應資料動態生成 HomeInsideText 元件 -->
      <HomeInsideText
        v-for="(item, index) in welfareData"
        :key="index"
        :data="item"
      />
    </div>
  </div>
</template>

<<<<<<< HEAD
<script>
=======
<script setup>
import { ref, onMounted } from 'vue';
>>>>>>> origin/loveandnotify
import TopNav from '@/components/TopNav.vue';
import HomeInsideText from '@/components/HomeInsideText.vue';

<<<<<<< HEAD
export default {
  name: 'HomePage',
  components : {
    TopNav,
    HomeInsideText,
  },
};
=======
const welfareData = ref([]); // 定義響應式數據

const getWelfare = async () => {
  try {
    const result = await request.get("/welfare"); // 發送請求
    welfareData.value = result.data; // 存儲數據
  } catch (error) {
    console.error("error", error); // 處理錯誤
  }
};

onMounted(() => {
  getWelfare(); // 在組件掛載後調用函數
});
>>>>>>> origin/loveandnotify
</script>

<style scoped>
/* 可選樣式 */
</style>