<template>

  <div class="flex flex-row justify-between px-4 py-3">

    <div class="flex gap-3 flex-row">

      <div class="flex flex-col">
        <div @click="toggleRegionDropdown" class="flex">
          <p class="text-l">{{ selectedRegion ? selectedRegion : "地區" }}</p>
          <PhCaretDown :size="20" class="menu ml-1 transition-transform"
            :class="{ 'rotate-180': showRegionDropdown }" />
        </div>

        <transition>
          <div v-if="showRegionDropdown"
            class="fixed flex flex-col gap-3 w-1/3 bor bg-white rounded-md border border-gray-300 p-3">
            <span v-for="region in regions" class=" text-xl font-bold" :key="region" @click="selectRegion(region)">
              {{ region }}</span>
          </div>
        </transition>

      </div>



      <div class="flex items-center" @click="toggleServiceDropdown">
        <p class="text-l">{{ selectedService ? selectedService : "服務" }}</p>
        <PhCaretDown :size="20" class="menu ml-1 transition-transform" :class="{ 'rotate-180': showServiceDropdown }" />
      </div>


      <Transition>
        <div v-if="showServiceDropdown"
          class="fixed flex flex-col gap-3 overflow-auto w-1/3 bor bg-white rounded-md border border-gray-300 p-3">
          <span v-for="service in services" class="text-xl font-bold" :key="service" @click="selectService(service)">{{
            service }}</span>
        </div>
      </Transition>


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
        '全選',
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

.v-enter-active,
.v-leave-active {
  transition: all 0.3s ease;
}

.v-enter-from,
.v-leave-to {
  transform: translateY(-20px);
  opacity: 0;
}
</style>