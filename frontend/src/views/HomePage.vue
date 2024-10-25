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

<script>
import TopNav from '@/components/TopNav.vue';
import HomeInsideText from '@/components/HomeInsideText.vue';

export default {
  name: 'HomePage',
  components : {
    TopNav,
    HomeInsideText,
  },
};
</script>

<style scoped>
/* 可選樣式 */
</style>