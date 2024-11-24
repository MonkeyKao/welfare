<template>
  <div class="flex h-20 bg-[#90c700] items-center space-x-2">
    <div @click="router.push('/older-home')" class="forgot-password-link">
      <PhArrowUUpLeft :size="28" color="#000" class="ml-5" />
    </div>
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
import { useRoute, useRouter } from 'vue-router';
import HeaderBar from '@/components/headerBar.vue';
import ElderlyFilter from '@/components/ElderlyFilter.vue';
import { PhArrowUUpLeft } from "@phosphor-icons/vue";

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