<template>
  <div>
    <!-- 選擇地區部分 -->
    <div class="justify-center">
      <button class="bg-[#92c700] w-full text-H2 p-1 rounded-lg shadow-md"
      @click="() => navigateToSelection('/regionselection')"
      >
        點我選「地區」
      </button>
    </div>
    <div>
      <p class="text-H3 mb-2 mt-2">已選擇地區</p>
      <div class="flex flex-wrap gap-2">
        <div
          v-for="region in selectedRegions"
          :key="region"
          class="flex items-center bg-[#6ca000] text-white text-H3 px-3 py-1 rounded-lg shadow-md"
        >
          <span class="mr-2">{{ region }}</span>
          <button @click="removeRegion(region)" class="text-white">
            ✕
          </button>
        </div>
      </div>
    </div>

    <!-- 選擇服務部分 -->
    <div class="justify-center">
      <button class="bg-[#92c700] w-full text-H2 p-1 rounded-lg shadow-md mt-4"
      @click="() => navigateToSelection('/serviceselection')"
      >
        點我選「服務」
      </button>
    </div>
    <div>
      <p class="text-H3 mb-2 mt-2">已選擇服務</p>
      <div class="flex flex-wrap gap-2">
        <div
          v-for="service in selectedServices"
          :key="service"
          class="flex items-center bg-[#6ca000] text-white text-H3 px-3 py-1 rounded-lg shadow-md"
        >
          <span class="mr-2">{{ service }}</span>
          <button @click="removeService(service)" class="text-white">
            ✕
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { getTextByLocation, getTextByService } from '@/utils/getTextByNumber';

const route = useRoute();
const router = useRouter();

const selectedServices = ref<string[]>([]);
const selectedRegions = ref<string[]>([]);

const navigateToSelection = (path: string) => {
  router.push({
    path,
    query: route.query,
  });
};

const initializeServices = () => {
  const services = route.query.selectedServices;
  if (services) {
    const serviceNames = Array.isArray(services) ? services : services.split(',');
    selectedServices.value = serviceNames;
  }
};

const removeService = (service: string) => {
  selectedServices.value = selectedServices.value.filter((s) => s !== service);
  updateRouteParams('selectedServices', selectedServices.value);
};

const initializeRegions = () => {
  const regions = route.query.selectedRegions;
  if (regions) {
    const regionIds = Array.isArray(regions) ? regions : regions.split(',');
    selectedRegions.value = regionIds;
  }
};

const removeRegion = (regionId: string) => {
  selectedRegions.value = selectedRegions.value.filter((r) => r !== regionId);
  updateRouteParams('selectedRegions', selectedRegions.value);
};

const updateRouteParams = (key: string, values: string[]) => {
  const newQuery = { ...route.query, [key]: values.join(',') };
  router.replace({ name: route.name || '', query: newQuery });
};

initializeServices();
initializeRegions();
</script>

<style scoped>
</style>