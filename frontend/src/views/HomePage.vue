<template>
  <div class="h-full flex flex-col w-full divide-y">
    <div class=" basis-1">
      <TopNav @selectRegion="(regions) => selectedRegion = regions"
        @selectService="(services) => selectedService = services" />
    </div>

    <div class="overflow-auto">
      <!-- 根據 API 回應資料動態生成 HomeInsideText 元件 -->
      <HomeInsideText @clickFavorited="(data) => { }" v-for="(item, index) in welfareData" :key="index" :data="item" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import TopNav from '@/components/TopNav.vue';
import HomeInsideText from '@/components/HomeInsideText.vue';
import { useWelfareStore } from '@/store/welfareStroe';

const welfareStore = useWelfareStore()
const welfareData = computed(() => welfareStore.getWelfare(selectedRegion.value, selectedService.value))
let selectedRegion = ref([])
let selectedService = ref([])
</script>

<style scoped></style>
