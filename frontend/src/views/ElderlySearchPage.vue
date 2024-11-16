<template>
  <div class="flex justify-start items-center gap-3 p-4 bg-[#92c700]">
    <router-link to="/elderlysearch" class="forgot-password-link">
      <PhArrowUUpLeft :size="32" color="#000" />
    </router-link>
    <label class="text-H2 font-bold">福利查詢</label>
  </div>

  <div class="flex flex-col justify-center p-4">

    <ElderlyFilter  />

    <button 
      class="bg-[#92c700] text-H2 w-full p-1 rounded-lg shadow-md mt-4"
      @click="handleSearch">
      搜索
    </button>
  </div>
</template>

<script setup lang="ts">
import { PhArrowUUpLeft } from '@phosphor-icons/vue';
import { ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import ElderlyFilter from '@/components/ElderlyFilter.vue';

const router = useRouter();

const selectedRegions = router.currentRoute.value.query.selectedRegions || '';
const selectedServices = router.currentRoute.value.query.selectedServices || '';

const emit = defineEmits(['selectedRegions','selectedServices'])

const handleSearch = () => {
  const selectedRegions = router.currentRoute.value.query.selectedRegions || '';
  const selectedServices = router.currentRoute.value.query.selectedServices || '';

  const searchQuery = {
    selectedRegions,
    selectedServices
  };

  router.push({
    path: '/searchresult',
    query: searchQuery
  });
};
</script>

<style scoped>
button {
  outline: none;
}
</style>