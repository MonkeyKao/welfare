<template>

  <div class="flex flex-row justify-between px-4 py-3">

    <div class="flex gap-3 flex-row">

      <div class="flex flex-col">
        <div @click="toggleRegionDropdown" class="flex">
          <p class="text-l">{{ selectedRegion.length===0?"地區":"地區("+selectedRegion.length+")" }}</p>
          <PhCaretDown :size="20" class="menu ml-1 transition-transform"
            :class="{ 'rotate-180': showRegionDropdown }" />
        </div>
        
        <transition>
          <div v-if="showRegionDropdown"
            class="fixed flex flex-col w-1/2 overflow-auto max-h-96 z-10 bg-white rounded-md border border-gray-300 p-3">
            <span v-for="region in 19" class=" select-none text-xl font-bold p-2" :class="{active:getRegionActive(region)}" :key="region" @click="selectRegion(region)">
              {{ getTextByLocation(region) }}
            </span>
          </div>
        </transition>
      </div>



      <div class="flex items-center" @click="toggleServiceDropdown">
        <p class="text-l">{{ selectedService.length===0?"服務":"服務("+selectedService.length+")" }}</p>
        <PhCaretDown :size="20" class="menu ml-1 transition-transform" :class="{ 'rotate-180': showServiceDropdown }" />
      </div>


      <Transition>
        <div v-if="showServiceDropdown"
          class="fixed flex flex-col gap-3 z-10 overflow-auto max-h-96 w-1/2 bor bg-white rounded-md border border-gray-300 p-3">
          <span v-for="service in 10" class=" select-none text-xl font-bold p-2" :key="service"  :class="{active:getServiceActive(service)}"  @click="selectService(service)">
            {{ getTextByService(service) }}
          </span>
        </div>
      </Transition>


    </div>

    <div class="flex gap-1 items-center" @click="toggleSidebar">
      <PhHourglass :size="20" />
      <span>篩選</span>
    </div>
  </div>

  <!-- 側邊欄 -->
  <div :class="{'translate-x-0': showSidebar, 'translate-x-full': !showSidebar}" 
         class="fixed top-0 right-0 space-y-4 w-64 h-full bg-white border-l border-gray-300 shadow-md p-4 transition-transform ease-in-out duration-300 z-10">
      <h3 class="font-bold text-xl text-center">篩選條件</h3>

      <div class="mt-4 space-y-5">
        <p class="text-lg font-bold">年齡</p>
        <div class="grid grid-cols-2 space-y-2" v-for="(item, index) in year" :key="index">
          <button @click="toggleSelect(item)"  class="flex justify-between items-center">
            <span :class="{ 'text-black': !item.selected, 'text-[#92c700]': item.selected }">
              {{ item.name }}
            </span>     
            <PhCheck v-if="item.selected" :size="20" color="#92c700" />
          </button>
        </div>
      </div>

      <div class="mt-4 space-y-5">
      <p class="text-lg font-bold">性別</p>
      <div class="grid grid-cols-2 gap-4">
        <button
          v-for="(item, index) in sex"
          :key="index"
          @click="toggleSelectSingle(item)"
          class="flex justify-between items-center"
        >
          <span :class="{ 'text-black': !item.selected, 'text-[#92c700]': item.selected }">
            {{ item.name }}
          </span>
          <PhCheck v-if="item.selected" :size="20" color="#92c700" />
        </button>
        </div>
      </div>

      <div class="mt-4 space-y-5">
        <p class="text-lg font-bold">收入</p>
        <div class="grid grid-cols-2 space-y-2" v-for="(item, index) in money" :key="index">
          <button @click="toggleSelectMoney(item)"  class="flex justify-between items-center">
            <span :class="{ 'text-black': !item.selected, 'text-[#92c700]': item.selected }">
              {{ item.name }}
            </span>     
            <PhCheck v-if="item.selected" :size="20" color="#92c700" />
          </button>
        </div>
      </div>

      <div class="mt-4 space-y-5">
        <p class="text-lg font-bold">身分別</p>
        <div class="grid grid-cols-2 space-y-2" v-for="(item, index) in people" :key="index">
          <button @click="toggleSelect(item)"  class="flex justify-between items-center">
            <span :class="{ 'text-black': !item.selected, 'text-[#92c700]': item.selected }">
              {{ item.name }}
            </span>     
            <PhCheck v-if="item.selected" :size="20" color="#92c700" />
          </button>
        </div>
      </div>

      <div class=" flex mt-4 justify-center">
        <button @click="applyFilters" class="bg-[#92c700] text-white px-4 py-2 rounded-md ">確定</button>
      </div>
    </div>

    <!-- 遮罩層 -->
    <div v-if="showSidebar" @click="toggleSidebar" 
         class="fixed top-0 left-0 w-full h-full bg-black opacity-50 z-5"></div>


  <!-- 下拉選單遮罩用 -->
  <div v-if="showRegionDropdown||showServiceDropdown" class=" h-screen w-screen bg-slate-300 fixed top-0 left-0 z-5 opacity-0 " @click="toggleDropdown"></div>

</template>

<script setup lang="ts">
import { getTextByLocation, getTextByService } from '@/utils/getTextByNumber';
import { PhCaretDown, PhHourglass,PhCheck } from '@phosphor-icons/vue';
import { ref } from 'vue';
let showRegionDropdown = ref<boolean>(false);
let showServiceDropdown = ref<boolean>(false);

const selectedRegion = ref<Array<number>>([])
const selectedService = ref<Array<number>>([])

const emit = defineEmits(['selectRegion','selectService'])

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
  }else{
    selectedRegion.value.splice(index,1);
  }

  emit('selectRegion',selectedRegion.value)
}

const getRegionActive = (region: number):boolean => {
  return selectedRegion.value.indexOf(region) !== -1
}

const selectService = (service: number) => {
  const index: number = selectedService.value.indexOf(service)
  if (index === -1) {
    selectedService.value.push(service)
  }else{
    selectedService.value.splice(index,1);
  }
  emit('selectService',selectedService.value)
}

const getServiceActive = (region: number):boolean => {
  return selectedService.value.indexOf(region) !== -1
}

const toggleDropdown = () => {
  showRegionDropdown.value = false
  showServiceDropdown.value = false
}

const showSidebar = ref(false); // 控制側邊欄的顯示

const toggleSidebar = () => {
  showSidebar.value = !showSidebar.value; // 切換側邊欄顯示狀態
};

const applyFilters = () => {
  // 在這裡處理篩選應用邏輯
  console.log('應用篩選');
  toggleSidebar(); // 篩選後關閉側邊欄
};

const year = ref([
  { id: 1, name: '20歲以下',selected: false },
  { id: 2, name: '20歲~65歲',selected: false },
  { id: 3, name: '65歲以上',selected: false },
])

const sex = ref([
  { id:1, name:"男性",selected: false},
  { id:2, name:"女性",selected: false},
])

const money = ref([
  { id:1, name:"中低收入戶",selected: false},
  { id:2, name:"低收入戶",selected: false},
])

const people = ref([
  { id:1, name:"榮民",selected: false},
  { id:2, name:"身心障礙者",selected: false},
])

const toggleSelect = (selectedItem: { id: number, name: string, selected: boolean }) => {
  // 如果當前選項已選擇，則取消選擇
  selectedItem.selected = !selectedItem.selected;
};
const toggleSelectSingle = (selectedItem: { id: number, name: string, selected: boolean }) => {
  if (!selectedItem.selected) {
    // 取消所有選項的選擇，然後選中當前項目
    sex.value.forEach(item => {
      item.selected = false;
    });
    selectedItem.selected = true; // 設置當前選項為選中狀態
  } else {
    // 如果當前項目已選擇，則取消選擇
    selectedItem.selected = false;
  }
};
const toggleSelectMoney = (selectedItem: { id: number, name: string, selected: boolean }) => {
  if (!selectedItem.selected) {
    // 取消所有選項的選擇，然後選中當前項目
    money.value.forEach(item => {
      item.selected = false;
    });
    selectedItem.selected = true; // 設置當前選項為選中狀態
  } else {
    // 如果當前項目已選擇，則取消選擇
    selectedItem.selected = false;
  }
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

.active:after {
  content: "✔";
  position: absolute;
  right: 10px;
}

.active {
  color: brown;
}
</style>