<template>

  <div class="flex bg-[#90c700] p-3 items-center fixed z-10 w-full px-6">
    <div class=" " @click="returnSelection">
      <PhArrowUUpLeft :size="32" color="#000" />
    </div>
    <label class="ml-5 text-H2 font-bold">選擇地區</label>
  </div>

  <div class="flow flow-col p-4 space-y-5 mt-[50px] mb-[50px]">
    <div v-for="(area, index) in areas" :key="index" class=" p-3">
      <p class="text-H2 font-bold pb-2">{{ area.name }}</p>
      <div v-for="(row, rowIndex) in area.regions" :key="rowIndex" class="grid grid-cols-3 gap-x-5 gap-y-8  mb-3">
        <button
          v-for="region in row"
          :key="region"
          :class="[ 
            'flex p-2 rounded-lg shadow-md text-H2 items-center justify-center gap-2 transition-colors',
          isRegionSelected(region) ? 'bg-[#73AA00] text-white' : 'bg-white text-black'
          ]"
          @click="toggleSelection(region)"
        >
          {{ getTextByLocation(region) }}
        </button>
      </div>
    </div>

  </div>
<!-- 確定按鈕容器，加入 z-20 確保按鈕位於其他元素之上 -->
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
import { getTextByLocation } from '@/utils/getTextByNumber'; // 確保這個函數是正確的

const router = useRouter();

// 地區數據 (代號對應文字的邏輯在 getTextByLocation 函數中處理)
const areas = ref([
  {
    name: "北區",
    regions: [
      [1, 2, 3],  // 台北市, 新北市, 基隆市
      [4, 5, 6],  // 桃園市, 新竹市, 新竹縣
      [17]        // 宜蘭縣
    ]
  },
  {
    name: "中區",
    regions: [
      [8, 9, 16], // 台中市, 彰化縣, 南投縣
      [10, 7]     // 雲林縣, 苗栗縣
    ]
  },
  {
    name: "南區",
    regions: [
      [14, 13, 15], // 高雄市, 台南市, 屏東縣
      [11, 12]      // 嘉義市, 嘉義縣
    ]
  },
  {
    name: "東區",
    regions: [
      [18, 19] // 花蓮縣, 台東縣
    ]
  }
]);

const selectedRegions = ref<string[]>([]);

const toggleSelection = (region: string) => {
  if (selectedRegions.value.includes(region)) {
    selectedRegions.value = selectedRegions.value.filter((r) => r !== region);
  } else {
    selectedRegions.value.push(region);
  }
};

// 判斷地區是否被選中
const isRegionSelected = (region: string) => selectedRegions.value.includes(region);

const confirmSelection = () => {
  // 從當前路由的 query 中提取 selectedServices
  const selectedServices = router.currentRoute.value.query.selectedServices || '';

 const selectedRegionNames = selectedRegions.value.map((regionNumber) =>
    getTextByLocation(regionNumber)
  );

  // 將選中的地區代碼轉換為字符串，並加入路由參數
  router.push({
    name: 'ElderlySearchPage',
    query: {
      selectedRegions: selectedRegionNames.join(','), // 更新選中的地區
      selectedServices: selectedServices // 保留之前選中的服務
    },
  });
};

const returnSelection = () => {
  // 如果有新的選擇，直接返回上一頁
  if (selectedRegions.value.length > 0) {
    router.back();
    return;
  }

  // 如果沒有新選擇，保留原本的選擇並返回
  const currentQuery = router.currentRoute.value.query;
  const selectedServices = currentQuery.selectedServices || '';
  const previousRegions = (currentQuery.selectedRegions as string || '').split(',');

  router.push({
    name: 'ElderlySearchPage',
    query: {
      selectedRegions: previousRegions.join(','), // 保留之前的選擇
      selectedServices: selectedServices // 保留之前選中的服務
    },
  });
};
</script>