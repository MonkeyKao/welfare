<template>
<div class="flex bg-[#90c700] p-3 items-center fixed z-10 w-full px-6 space-x-2">
    <router-link to="/home" class="forgot-password-link">
      <PhArrowUUpLeft :size="32" color="" />
    </router-link>
    <label class="text-3xl font-bold">收藏</label>
  </div>

  <div class="mt-[70px] space-y-5 mb-2">
    <div
      v-for="(welfare, index) in favoriteData"
      :key="index"
      class="flex items-end space-x-2 px-4"
    >
    <div > <!--這個div在解決icon大小無法正確顯示的問題-->
        <PhUserSquare :size="40" />        
      </div>
      <div     @click="router.push(welfare.url)" class="flex p-4 bg-white shadow-md rounded-tl-lg rounded-tr-lg rounded-br-lg">
        <p class="text-gray-800 ">{{ welfare.title }}</p>
      </div>
      <div class="ml-4" @click="favoriteStore.deleteFavoriteHandler(welfare.id)">
        <PhHeartStraight
          :size="28"
          weight= 'fill'
          class="cursor-pointer"
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import { PhArrowUUpLeft, PhHeartStraight,PhUserSquare } from '@phosphor-icons/vue';
import request from '@/axios';
import type Welfare from '@/model/welfare';
import router from '@/router';
import { useFavoriteStore } from '@/store/favorite';
const favoriteStore = useFavoriteStore()
const favoriteData = computed(() => favoriteStore.favorites)


</script>

<style scoped>
</style>