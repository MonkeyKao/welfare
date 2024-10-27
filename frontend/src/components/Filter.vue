<template>

  <div class="flex flex-row justify-between px-4 py-3">

    <div class="flex gap-3 flex-row">

      <div class="flex items-center" @click="toggleRegionDropdown">
        <p>{{ selectedRegion ? selectedRegion : "地區" }}</p>
        <PhCaretDown :size="20" class="menu ml-1 transition-transform" :class="{ 'rotate-180': showRegionDropdown }" />
      </div>

      <div v-if="showRegionDropdown" class=" bg-white border border-gray-300 z-10 p-2">
        <p v-for="region in regions" :key="region" @click="selectRegion(region)"
          class="cursor-pointer hover:bg-gray-200">
          {{ region }}</p>
      </div>

      <div class="flex items-center" @click="toggleServiceDropdown">
        <p>{{ selectedService ? selectedService : "服務" }}</p>
        <PhCaretDown :size="20" class="menu ml-1 transition-transform" :class="{ 'rotate-180': showServiceDropdown }" />
      </div>

      <div v-if="showServiceDropdown" class="bg-white border border-gray-300 z-10 p-2">
        <p v-for="service in services" :key="service" @click="selectService(service)"
          class="cursor-pointer hover:bg-gray-200">{{ service }}</p>
      </div>

    </div>

    <div class="flex gap-1 items-center">
      <PhHourglass :size="20" />
      <span>篩選</span>
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
      services: [
        '育兒津貼',
        '托育補助',
        '孕婦補助',
        '單親家庭津貼',
        '學費減免',
        '獎助學金',
        '特殊教育支援',
        '低收入戶子女教育補助',
        '醫療補助',
        '長照補助',
        '精神健康補助',
        '老人年金'], // 根據需要替換
    };
  },
  methods: {
    toggleRegionDropdown() {
      this.showRegionDropdown = !this.showRegionDropdown; // 切換狀態
      if (this.showRegionDropdown) {
        this.showServiceDropdown = false; // 關閉服務下拉選單
      }
    },
    toggleServiceDropdown() {
      this.showServiceDropdown = !this.showServiceDropdown; // 切換狀態
      if (this.showServiceDropdown) {
        this.showRegionDropdown = false; // 關閉地區下拉選單
      }
    },
    selectRegion(region) {
      this.selectedRegion = region;
      console.log(region);

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
  transition: transform 0.2s;
  /* 動畫過渡效果 */
}
</style>