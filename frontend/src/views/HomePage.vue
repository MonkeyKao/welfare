<template>
  <div class="h-full flex flex-col w-full divide-y">
    <div class=" basis-1">
      <TopNav @selectRegion="(regions) => selectedRegion = regions"
        @selectService="(services) => selectedService = services" />
    </div>

    <div class="overflow-auto" @scroll="onScroll" ref="scrollContainer">
      <HomeInsideText v-for="(item, index) in displayedData" :key="index" :data="item"
        @clickFavorited="(data: Welfare) => clickFavoriteHandler(data)" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, inject, watch, onMounted } from 'vue';
import TopNav from '@/components/TopNav.vue';
import HomeInsideText from '@/components/HomeInsideText.vue';
import { useWelfareStore } from '@/store/welfareStroe';
import type Welfare from '@/model/welfare';
import { useFavoriteStore } from '@/store/favorite';
const showMsg: Function = inject("showMsg")!

const favoriteStore = useFavoriteStore()
const favoriteData = computed(() => favoriteStore.favorites)

const clickFavoriteHandler = async (welfare: Welfare) => {
  if (Array.isArray(favoriteData.value) && favoriteData.value.some((item) => item.id === welfare.id)) {
    // 如果已存在，则调用删除
    favoriteStore.deleteFavoriteHandler(welfare.id)
  } else {
    try {
      await favoriteStore.createFavoriteHanlder(welfare)
    } catch (error) {
      showMsg(error)
    }
    // 如果未收藏，则添加

  }
}

const welfareStore = useWelfareStore()
const welfareData = computed(() => welfareStore.getWelfare(selectedRegion.value, selectedService.value))
let selectedRegion = ref([])
let selectedService = ref([])

const scrollContainer = ref<HTMLElement | null>(null);
const pageSize = 50; // 每次載入的資料數量
const currentPage = ref(1);
const displayedData = ref([] as Welfare[]);

// 更新顯示資料
const updateDisplayedData = () => {
  const start = 0;
  const end = currentPage.value * pageSize;
  displayedData.value = welfareData.value.slice(start, end);
};

// 監聽 welfareData 資料變化
watch(welfareData, () => {
  currentPage.value = 1; // 重置分頁
  updateDisplayedData();
});

// 滾動事件處理
const onScroll = () => {
  if (!scrollContainer.value) return;

  const { scrollTop, scrollHeight, clientHeight } = scrollContainer.value;
  if (scrollTop + clientHeight >= scrollHeight - 50) {
    loadMoreData();
  }
};

// 載入更多資料
const loadMoreData = () => {
  if (currentPage.value * pageSize < welfareData.value.length) {
    currentPage.value += 1;
    updateDisplayedData();
  }
};

onMounted(() => {
  updateDisplayedData();
});


</script>

<style scoped></style>
