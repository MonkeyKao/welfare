<template>
  <div class="flex bg-[#90c700] p-3 items-center fixed z-10 w-full px-6">
    <div class="forgot-password-link" @click="returnSelection">
      <PhArrowUUpLeft :size="32" color="#000" />
    </div>
    <label class="ml-5 text-H2 font-bold">選擇服務</label>
  </div>

  <div class="flow flow-col p-4 space-y-5 mt-[60px] mb-[50px]">
    <div class="grid grid-cols-1 gap-5 mb-5">
      <button
        v-for="(service, serviceIndex) in services"
        :key="serviceIndex"
        :class="[ 
          'flex p-3 rounded-lg shadow-md text-H2 mx-3 px-5 gap-2 transition-colors',
          selectedServices.includes(service) ? 'bg-[#73AA00] text-white' : 'bg-white text-black'
        ]"
        @click="toggleServiceSelection(service)"
      >
        {{ getTextByService(service) }}
      </button>
    </div>
  </div>

  <div class="flex flex-col h-[10%] justify-center fixed bottom-0 w-full px-6 z-20 bg-white">
      <button class="bg-[#92c700] text-H2 p-1 rounded-lg shadow-md text-white" @click="confirmSelection">
        確定
      </button>
    </div>
</template>

<script setup lang="ts">
import { PhArrowUUpLeft, PhCheck } from '@phosphor-icons/vue';
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { getTextByService } from '@/utils/getTextByNumber'; // 引入服務映射函數

const router = useRouter();

const services = [
  1, 2, 3, 4, 5, 6, 7, 8, 9, 10
];

// 已選服務列表 (存儲服務編號)
const selectedServices = ref<string[]>([]);

// 切換服務選擇
const toggleServiceSelection = (service: string) => {
  if (selectedServices.value.includes(service)) {
    selectedServices.value = selectedServices.value.filter((s) => s !== service);
  } else {
    selectedServices.value.push(service);
  }
};

const confirmSelection = () => {
  // 從當前路由的 query 中提取 selectedRegions
  const selectedRegions = router.currentRoute.value.query.selectedRegions || '';

  // 獲取選中的服務名稱
  const selectedServiceNames = selectedServices.value.map((serviceNumber) =>
    getTextByService(serviceNumber)
  );

  // 將 selectedRegions 和 selectedServices 一起帶回 /elderlysearch
  router.push({
    name: 'ElderlySearchPage',
    query: {
      selectedRegions: selectedRegions, // 保留之前的地區參數
      selectedServices: selectedServiceNames.join(',') // 新增服務參數
    },
  });
};

const returnSelection = () => {
  // 如果有新的選擇，直接返回上一頁
  if (selectedServices.value.length > 0) {
    router.back();
    return;
  }

  // 如果沒有新選擇，保留原本的選擇並返回
  const currentQuery = router.currentRoute.value.query;
  const previousServices = (currentQuery.selectedServices as string || '').split(',');
  const previousRegions = (currentQuery.selectedRegions as string || '').split(',');

  router.push({
    name: 'ElderlySearchPage',
    query: {
      selectedRegions: previousRegions.join(','), // 保留之前的地區選擇
      selectedServices: previousServices.join(',') // 保留之前的服務選擇
    },
  });
};
</script>
