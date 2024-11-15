<template>
  <div class="flex justify-start items-center gap-3 p-4 bg-[#92c700]">
    <router-link to="/home" class="forgot-password-link">
      <PhArrowUUpLeft :size="32" color="#000" />
    </router-link>
    <label class="text-3xl font-bold">收藏</label>
  </div>

  <div
    v-for="(welfare, index) in favoriteData"
    :key="index"
    class="flex items-end p-4"
  >
    <img src="@/images/user-square.png" alt="類別圖片" class="w-10 h-10" />
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
</template>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import { PhArrowUUpLeft, PhHeartStraight } from '@phosphor-icons/vue';
import request from '@/axios';
import type Welfare from '@/model/welfare';
import router from '@/router';
import { useFavoriteStore } from '@/store/favorite';
const favoriteStore = useFavoriteStore()
const favoriteData = computed(() => favoriteStore.favorites)


</script>

<style scoped>
</style>