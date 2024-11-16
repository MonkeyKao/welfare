<template>
  <div class="flex justify-start items-center gap-3 p-4 bg-[#92c700]">
    <router-link to="/elderlysearch" class="forgot-password-link">
      <PhArrowUUpLeft :size="32" color="#000" />
    </router-link>
    <label class="text-H2 font-bold">福利查詢</label>
  </div>

  <div class="overflow-auto">
      <!-- 根據 API 回應資料動態生成 HomeInsideText 元件 -->
      <HomeInsideText @clickFavorited="(data) => { }" v-for="(item, index) in welfareData" :key="index" :data="item" />
    </div>
</template>

<script setup lang="ts">
import { PhArrowUUpLeft } from '@phosphor-icons/vue';
import { ref, computed, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import HomeInsideText from '@/components/HomeInsideText.vue';
import { useWelfareStore } from '@/store/welfareStroe';

// 获取路由信息
const route = useRoute();
const welfareStore = useWelfareStore();

// 初始化地区和服务
let selectedRegion = ref<string[]>([]);
let selectedService = ref<string[]>([]);

// 从路由的查询参数中提取选中的地区和服务
onMounted(() => {
  const regions = route.query.selectedRegions;
  const services = route.query.selectedServices;

  if (regions) {
    selectedRegion.value = Array.isArray(regions) ? regions : regions.split(',');
  }

  if (services) {
    selectedService.value = Array.isArray(services) ? services : services.split(',');
  }
});

const welfareData = computed(() => {
  return welfareStore.getWelfare(selectedRegion.value, selectedService.value);
});
</script>

<style scoped>
/* 自定义样式 */
</style>