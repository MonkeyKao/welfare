<template>
  <div class="h-full flex flex-col w-full divide-y">
    <div class=" basis-1">
      <TopNav @selectRegion="(regions) => selectedRegion = regions"
        @selectService="(services) => selectedService = services" />
    </div>

    <div class="overflow-auto">
      <!-- 根據 API 回應資料動態生成 HomeInsideText 元件 -->
      <HomeInsideText @clickFavorited="(data: Welfare) => clickFavoriteHandler(data)"
        v-for="(item, index) in welfareData" :key="index" :data="item" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, inject } from 'vue';
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


</script>

<style scoped></style>
