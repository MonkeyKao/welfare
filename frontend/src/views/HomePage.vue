<template>
  <div>
    <TopNav />
    <div class="mt-40">
    <HomeInsideText />
    <HomeInsideText />
    <HomeInsideText />
    <HomeInsideText />
    <HomeInsideText />
    <HomeInsideText />
    <HomeInsideText />
    <HomeInsideText />
    <HomeInsideText />
    <HomeInsideText />
    <HomeInsideText />
    <HomeInsideText />
    </div>
  </div>
</template>

<script setup>
import TopNav from '@/components/TopNav.vue';
import HomeInsideText from '@/components/HomeInsideText.vue';
import request from '@/axios';

const getWelfare = async () => {
  const result = await request.get("/welfare")
  console.log(result);
}

getWelfare()
</script>
``  
<style scoped>
/* 可選樣式 */
</style>