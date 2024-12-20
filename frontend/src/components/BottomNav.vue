<template>
  <nav class=" bg-gray-50">
    <ul class="flex justify-around items-center">
      <li>
        <router-link to="/mou" class="flex flex-col items-center text-center">
          <img v-if="isActive('/mou')" src="/amuoclick.png" class="h-7" />
          <img v-else src="/amuo.png" class="h-7" />
          <span v-if="isActive('/mou')" class="text-gray-800 text-H4">阿哞</span>
        </router-link>
      </li>
      <li>
        <router-link to="/home" class="flex flex-col items-center text-center">
          <PhHouse :size="32" :weight="isActive('/home') ? 'fill' : 'regular'"
            :class="{ 'text-green-500': isActive('/home') }"
            :style="{ color: isActive('/home') ? '#92C700' : 'black' }" />
          <span v-if="isActive('/home')" class="text-gray-800 text-H4">首頁</span>
        </router-link>
      </li>
      <li>
        <router-link to="/user" class="flex flex-col items-center text-center">
          <PhUserGear :size="32" :weight="isActive('/user') ? 'fill' : 'regular'"
            :class="{ 'text-green-500': isActive('/user') }"
            :style="{ color: isActive('/user') ? '#92C700' : 'black' }" />
          <span v-if="isActive('/user')" class="text-gray-800 text-H4">用戶</span>
          <span v-if="isActive('/settings')" class="text-gray-800 text-H4">用戶</span>
        </router-link>
      </li>
    </ul>
  </nav>
</template>

<script setup lang="ts">
import { PhHouse, PhUserGear } from "@phosphor-icons/vue";
import { useRoute } from "vue-router";

const route = useRoute();

const isActive = (path:string) => {
  return route.path === path;
}

</script>
