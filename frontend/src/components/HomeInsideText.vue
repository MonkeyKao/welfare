<template>
<div class="p-2">
<p class="ml-5 mt-5 text-[#7F8689] text-sm">台北市 / 嬰兒照顧服務</p>
<div class="flex justify-between items-center">
    <p class="ml-5 mt-3 text-base font-semibold">嬰幼兒托育地圖(找托育資源)</p>
    <PhHeartStraight
      :size="28"
      :weight="isFavorited ? 'fill' : 'regular'"  
      class="mr-4 icon justify-center" 
      @click="toggleFavorite"
    />
  </div>
</div>
</template>

<script>
import { PhHeartStraight } from '@phosphor-icons/vue';

export default {
  components: {
    PhHeartStraight,
  },
  data() {
    return {
      isFavorited: false,  // 初始為未被喜愛
    };
  },
  methods: {
    toggleFavorite() {
      this.isFavorited = !this.isFavorited;  // 切換喜愛狀態
    },
  },
};
</script>

<style scoped>
.icon {
  transition: color 0.3s;  /* 添加平滑過渡效果 */
}
</style>