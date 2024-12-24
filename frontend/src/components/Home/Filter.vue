<template>

  <div class="flex flex-row justify-between px-4 py-3">

    <div class="flex gap-3 flex-row">

      <div class="flex flex-col">
        <div @click="toggleRegionDropdown" class="flex items-center">
          <p class="text-l">{{ selectedRegion.length === 0 ? "地區" : "地區(" + selectedRegion.length + ")" }}</p>
          <PhCaretDown :size="32" class=" menu ml-1 transition-transform"
            :class="{ 'rotate-180': showRegionDropdown }" />
        </div>

        <transition>
          <div v-if="showRegionDropdown"
            class="fixed flex flex-col w-1/2 overflow-auto max-h-96 z-30 bg-white rounded-md border border-gray-300 p-3">
            <span v-for="region in 19" class=" select-none text-xl font-bold p-2"
              :class="{ active: getRegionActive(region) }" :key="region" @click="selectRegion(region)">
              {{ getTextByLocation(region) }}
            </span>
          </div>
        </transition>
      </div>



      <div class="flex items-center" @click="toggleServiceDropdown">
        <p class="text-l">{{ selectedService.length === 0 ? "服務" : "服務(" + selectedService.length + ")" }}</p>
        <PhCaretDown :size="32" class="menu ml-1 transition-transform" :class="{ 'rotate-180': showServiceDropdown }" />
      </div>


      <Transition>
        <div v-if="showServiceDropdown"
          class="fixed flex flex-col gap-3 z-30 overflow-auto max-h-96 w-1/2 bor bg-white rounded-md border border-gray-300 p-3">
          <span v-for="service in 10" class=" select-none text-xl font-bold p-2" :key="service"
            :class="{ active: getServiceActive(service) }" @click="selectService(service)">
            {{ getTextByService(service) }}
          </span>
        </div>
      </Transition>


    </div>

    <div class="flex gap-1 items-center" @click="SideBarTemp.toggleSidebar">
      <PhHourglass :size="20" />
      <span>篩選</span>
    </div>

    <SideBar @select-status="(status) => emit('selectStatus',status)" ref="SideBarTemp"></SideBar>

    <!-- 下拉選單遮罩用 -->
    <div v-if="showRegionDropdown || showServiceDropdown"
      class=" h-screen w-screen bg-gray-500 fixed top-0 left-0 z-20 opacity-25 " @click="toggleDropdown"></div>
  </div>

</template>

<script setup lang="ts">
import { getTextByLocation, getTextByService } from '@/utils/getTextByNumber';
import { PhCaretDown, PhHourglass, PhCheck } from '@phosphor-icons/vue';
import SideBar from './SideBar.vue';
import { ref } from 'vue';
let showRegionDropdown = ref<boolean>(false);
let showServiceDropdown = ref<boolean>(false);

const SideBarTemp = ref();

const selectedRegion = ref<Array<number>>([])
const selectedService = ref<Array<number>>([])

const emit = defineEmits(['selectRegion', 'selectService','selectStatus'])

const toggleRegionDropdown = () => {
  showRegionDropdown.value = !showRegionDropdown.value; // 切換狀態
  if (showRegionDropdown.value) {
    showServiceDropdown.value = false; // 關閉服務下拉選單
  }
}
const toggleServiceDropdown = () => {
  showServiceDropdown.value = !showServiceDropdown.value; // 切換狀態
  if (showServiceDropdown.value) {
    showRegionDropdown.value = false; // 關閉地區下拉選單
  }
}
const selectRegion = (region: number) => {
  const index: number = selectedRegion.value.indexOf(region)
  if (index === -1) {
    selectedRegion.value.push(region)
  } else {
    selectedRegion.value.splice(index, 1);
  }

  emit('selectRegion', selectedRegion.value)
}

const getRegionActive = (region: number): boolean => {
  return selectedRegion.value.indexOf(region) !== -1
}

const selectService = (service: number) => {
  const index: number = selectedService.value.indexOf(service)
  if (index === -1) {
    selectedService.value.push(service)
  } else {
    selectedService.value.splice(index, 1);
  }
  emit('selectService', selectedService.value)
}

const getServiceActive = (region: number): boolean => {
  return selectedService.value.indexOf(region) !== -1
}

const toggleDropdown = () => {
  showRegionDropdown.value = false
  showServiceDropdown.value = false
}
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

.active:after {
  content: "✔";
  position: absolute;
  right: 10px;
}

.active {
  color: brown;
}
</style>