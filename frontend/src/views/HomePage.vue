<template>
  <div class="h-full flex flex-col w-full divide-y">
    <!-- 顶部导航，用于筛选条件 -->
    <div class="basis-1">
      <TopNav @selectRegion="(regions: number[]) => searchCondition.city = regions"
        @selectService="(services: number[]) => searchCondition.category = services"
        @selectTitle="(title: string) => searchCondition.title = title" />
    </div>

    <!-- 数据展示区 -->
    <div class="overflow-auto" @scroll="onScroll" ref="scrollContainer">
      <HomeInsideCard v-for="(item, index) in displayedData" :key="index" :data="item"
        @clickFavorited="(data: Welfare) => clickFavoriteHandler(data)" />
      <div v-if="loading" class="loading-indicator">加载中...</div>
      <div v-else-if="!hasMore" class="no-more-indicator">没有更多数据了</div>
    </div>
  </div>
</template>


<script setup lang="ts">
import Welfare from '@/model/welfare';
import { ref, computed, inject, watch, onMounted } from 'vue';
import TopNav from '@/components/Home/TopNav.vue';
import HomeInsideCard from '@/components/Home/HomeInsideCard.vue';
import { useFavoriteStore } from '@/store/favorite';
import { AlertColor, showMsgFunction } from '@/type/ShowMsg';
import { isAxiosError } from 'axios';
import request from '@/axios';

const showMsg: showMsgFunction = inject("showMsg")!

const favoriteStore = useFavoriteStore();
const favoriteData = computed(() => favoriteStore.favorites)

type Condition = {
  city?: Array<number>
  category?: Array<number>
  title?: string
}

const searchCondition = ref<Condition>({})

const clickFavoriteHandler = async (welfare: Welfare) => {
  if (Array.isArray(favoriteData.value) && favoriteData.value.some((item) => item.id === welfare.id)) {
    // 如果已存在，则调用删除
    favoriteStore.deleteFavoriteHandler(welfare.id)
  } else {
    try {
      await favoriteStore.createFavoriteHanlder(welfare)
    } catch (err: unknown) {
      if (isAxiosError(err))
        showMsg(err.response?.data.error, AlertColor.error)
    }
  }
}

// 分页状态
const page = ref(1);
const pageSize = ref(10);
const hasMore = ref(true); // 是否还有更多数据
const loading = ref(false); // 是否正在加载

// 当前展示数据
const displayedData = ref<Welfare[]>([]);

// 滚动容器引用
const scrollContainer = ref<HTMLElement | null>(null);

// 搜索功能
const getWelfare = async () => {
  if (loading.value || !hasMore.value) return;

  loading.value = true;
  try {
    const response = await request.get("welfare", {
      params: {
        ...searchCondition.value,
        page: page.value,
        pageSize: pageSize.value
      }
    })
    const data = response.data
    // 判断是否还有更多数据
    if (data.length < pageSize.value) {
      hasMore.value = false;
    }

    // 添加数据到展示列表
    displayedData.value.push(...data);
    page.value += 1; // 增加页码
  } catch (error) {
    showMsg('加载数据失败，请重试！', AlertColor.error);
  } finally {
    loading.value = false;
  }
};

// 滚动事件处理
const onScroll = () => {
  if (!scrollContainer.value || loading.value || !hasMore.value) return;

  const { scrollTop, scrollHeight, clientHeight } = scrollContainer.value;
  if (scrollTop + clientHeight >= scrollHeight - 10) {
    getWelfare(); // 滚动到底部时加载新数据
  }
};

// 监听筛选条件变化
watch(searchCondition, async () => {
  // 重置分页状态
  page.value = 1;
  hasMore.value = true;
  displayedData.value = [];

  // 加载新的数据
  await getWelfare();
}, { deep: true });

// 初始加载数据
onMounted(() => {
  getWelfare();
});
</script>

<style scoped></style>
