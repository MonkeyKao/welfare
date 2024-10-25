<template>
  <div class="flex justify-start pb-1 mb-2 ml-6">
    <div class="dropdown relative">
      <div class="flex items-center" @click="toggleRegionDropdown">
        <p class="text-lg">地區</p>
        <PhCaretDown
          :size="20" 
          class="menu ml-1 transition-transform" 
          :class="{ 'rotate-180': showRegionDropdown }" 
        />
      </div>
      <div v-if="showRegionDropdown" class="absolute top-full left-0 bg-white border border-gray-300 z-10 p-2 w-40">
        <p v-for="region in regions" :key="region" @click="selectRegion(region)" class="cursor-pointer hover:bg-gray-200">{{ region }}</p>
      </div>
    </div>

    <div class="relative dropdown ml-2">
      <div class="flex items-center" @click="toggleServiceDropdown">
        <p class="text-lg">服務</p>
        <PhCaretDown
          :size="20" 
          class="menu ml-1 transition-transform" 
          :class="{ 'rotate-180': showServiceDropdown }" 
        />
      </div>
      <div v-if="showServiceDropdown" class="absolute top-full left-0 bg-white border border-gray-300 z-10 p-2 w-40">
        <p v-for="service in services" :key="service" @click="selectService(service)" class="cursor-pointer hover:bg-gray-200">{{ service }}</p>
      </div>
    </div>

    <div class="flex items-center ml-40">
      <PhHourglass :size="20" />
      <span class="text-lg ml-2">篩選</span>
    </div>
  </div>
</template>

<script>
import { PhCaretDown, PhHourglass } from '@phosphor-icons/vue';

export default {
  components: {
    PhCaretDown,
    PhHourglass,
  },
  data() {
    return {
      showRegionDropdown: false,
      showServiceDropdown: false,
      selectedRegion: null,
      selectedService: null,
      regions: ['全選', '台北市', '新北市', '桃園市', '新竹市', '新竹縣', '苗栗縣', '台中市', '彰化縣', '南投縣', '雲林縣', '嘉義縣', '嘉義市', '台南市', '高雄市', '屏東縣', '基隆市'], // 根據需要替換
      services: ['服務1', '服務2', '服務3'], // 根據需要替換
    };
  },
  methods: {
    toggleRegionDropdown() {
      this.showRegionDropdown = !this.showRegionDropdown;
    },
    toggleServiceDropdown() {
      this.showServiceDropdown = !this.showServiceDropdown;
    },
    selectRegion(region) {
      this.selectedRegion = region;
      this.showRegionDropdown = false; // 選擇後關閉下拉選單
    },
    selectService(service) {
      this.selectedService = service;
      this.showServiceDropdown = false; // 選擇後關閉下拉選單
    },
  },
};
</script>

<style scoped>
.rotate-180 {
  transform: rotate(180deg);
  transition: transform 0.2s; /* 動畫過渡效果 */
}
</style>