<template>
  <div class="relative overflow-hidden p-3 flex flex-col gap-2" @touchstart="handleTouchStart"
    @touchmove="handleTouchMove" @touchend="handleTouchEnd">

    <!-- 左滑的主要內容 -->
    <div :style="{ transform: `translateX(${translateX}px)` }" class="transition-transform duration-300">
      <div class="flex justify-between">
        <div class="w-full" @click="router.push('/welfare/' + props.data.id)">
          <p class="text-[#7F8689] text-sm">
            {{ getTextByLocation(props.data.city) }} /
            {{ props.data.category
              ? getTextByService(props.data.category[0])
              : "其他服務" }}
          </p>
          <p class="text-base font-semibold basis-3/4">{{ props.data.title }}</p>
        </div>
        <!-- PhCircle 按鈕 -->
        <div class="flex items-center">
          <!-- 根據 isSwiped 控制 margin-right -->
          <PhCircle :size="20" :color="getCanGetColor(props.data.canGet)" weight="fill" 
            :style="{ marginRight: isSwiped ? '18px' : '6px' }" />
        </div>
      </div>

    </div>

    <!-- 隱藏的按鈕，加入動畫效果 -->
    <div class="absolute right-0 top-0 h-full flex items-center px-4 gap-3 bg-customGreen"
      :style="{
        transform: isSwiped ? 'translateX(0)' : 'translateX(100%)',
        transition: 'transform 0.3s ease-in-out'
      }">
      <!-- 第一顆按鈕 -->
      <div class="flex items-center justify-center w-10 " @click="toggleFavorite">
        <PhHeartStraight :size="28" color="#FFFFFF" :weight="isFavorite(props.data.id) ? 'fill' : 'regular'" />
      </div>

      <!-- 第二顆按鈕 -->
      <div class="flex items-center justify-center w-10" @click="handleSecondButtonClick">
        <PhShareFat :size="28" color="#FFFFFF" :weight="isClicked ? 'fill' : 'regular'" />
      </div>
    </div>

  </div>
</template>

<script setup lang="ts">
import { getTextByLocation, getTextByService } from "@/utils/getTextByNumber";
import { PhHeartStraight, PhCircle, PhShareFat } from "@phosphor-icons/vue";
import { computed, ref } from "vue";
import { useFavoriteStore } from "@/store/favorite";
import router from "@/router";

const props = defineProps(["data"]); 
const emit = defineEmits(["clickFavorited"]);  

const favoriteStore = useFavoriteStore();  
const favoriteData = computed(() => favoriteStore.favorites);  

const isSwiped = ref(false); // 控制隱藏按鈕顯示
const startX = ref(0); // 初始觸控點
const translateX = ref(0); // 內容的偏移量

const isFavorite = (welfareId: number): boolean => { // 判斷是否已加入收藏
  return Array.isArray(favoriteData.value) && favoriteData.value.some((item) => item.id === welfareId);
};

const toggleFavorite = () => {
  emit("clickFavorited", props.data);
};

const handleTouchStart = (event: TouchEvent) => {
  startX.value = event.touches[0].clientX; // 記錄觸控起始位置
};

const handleTouchMove = (event: TouchEvent) => {
  const deltaX = event.touches[0].clientX - startX.value; // 計算滑動距離
  translateX.value = Math.min(0, deltaX); // 限制偏移只能向左滑
  if (deltaX < -50) {
    isSwiped.value = true; // 左滑超過 50px 顯示按鈕
  } else {
    isSwiped.value = false; // 未達到滑動門檻
  }
};

const handleTouchEnd = () => {
  if (isSwiped.value) {
    translateX.value = -100; // 左滑後固定內容偏移
  } else {
    translateX.value = 0; // 滑動回彈
  }
};

const getCanGetColor = (canGet: number): string => {
  switch (canGet) {
    case 1:
      return "#92C700"; // 綠色=快去領
    case 2:
      return "#FFEA53"; // 黃色=不確定能不能領
    case 3:
      return "#D06262"; // 紅色=沒得領啦
    default:
      return "#cccccc"; // 預設灰色
  }
};

const isClicked = ref(false); 
const handleSecondButtonClick = () => { 
  isClicked.value = !isClicked.value;  
  console.log("分享被點擊了！", isClicked.value);  
};
</script>

<style scoped>
.relative {
  position: relative; /*相對定位*/
}

.transition-transform {
  transition: transform 0.3s;  /*動畫過渡效果*/
}

.overflow-hidden {
  overflow: hidden; /*超出部分隱藏*/
}

.absolute {
  position: absolute; /*絕對定位*/
}

.bg-customGreen {
  background-color: #92C700; /* 假設的綠色 */
}

.text-white {
  color: white;
}
</style>
