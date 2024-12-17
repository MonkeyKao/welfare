<template>
  <div class="flex justify-start items-center gap-3 p-4 bg-[#92c700]">
    <div @click="router.go(-1)">
      <PhArrowUUpLeft :size="32" color="" />
    </div>
    <label class="text-H2 font-bold">福利查詢</label>
  </div>

  <div class="overflow-auto">
    <!-- 根據 API 回應資料動態生成 HomeInsideText 元件 -->
    <HomeInsideCard @clickFavorited="() => { }" v-for="(item, index) in welfareData" :key="index" :data="item" />
  </div>
</template>

<script setup lang="ts">
import { PhArrowUUpLeft } from '@phosphor-icons/vue';
import { ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import HomeInsideCard from '@/components/Home/HomeInsideCard.vue';
import { useWelfareStore } from '@/store/welfareStroe';
import router from '@/router';
import type Welfare from '@/model/welfare';

// 获取路由信息
const route = useRoute();
const welfareStore = useWelfareStore();
const welfareData = ref<Array<Welfare>>([])

// 从路由的查询参数中提取选中的地区和服务
onMounted(() => {
  // const regions = route.query.selectedRegions!.toString().split(',').map((item) => Number(item))
  const services = route.query.selectedServices?route.query.selectedServices.toString().split(',').map((item) => Number(item)):[];
  const regions = route.query.selectedRegions?route.query.selectedRegions.toString().split(',').map((item) => Number(item)) : []


  welfareData.value = welfareStore.getWelfare(regions, services)
});

</script>

<style scoped>
/* 自定义样式 */
</style>