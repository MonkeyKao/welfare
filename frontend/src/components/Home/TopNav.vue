<template>
  <div class="w-full flex flex-col justify-between">
    <div class="flex justify-between items-center gap-3 p-4 " style="background-color: #92C700;">

      <div class="flex items-center flex-grow bg-gray-50 rounded-lg p-1 border border-gray-300">
        <PhMagnifyingGlass size="24" color="#9ca3af" class="icon-search" />
        <input type="text" placeholder="Ex.租屋補助" class="p-1 w-full text-base bg-transparent border-none outline-none" />
      </div>

      <div class="flex space-x-3">
        <router-link to="/favorites" class="icon-link" title="收藏">
          <PhHeartStraight size="28" color="#424E2A" class="icon" />
        </router-link>
        <router-link to="/notifications" class="icon-link" title="通知">
          <PhBell size="28" color="#424E2A" class="icon" />
        </router-link>
      </div>
    </div>



    <Filter @select-region="(region) => emit('selectRegion',region)" @select-service="(service) => emit('selectService',service)" />

  </div>
</template>

<script setup lang="ts">
import { PhHeartStraight, PhBell, PhMagnifyingGlass } from '@phosphor-icons/vue';
import Filter from './Filter.vue';
const emit = defineEmits(['selectRegion','selectService'])
</script>