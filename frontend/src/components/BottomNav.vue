<template>
  <nav class=" bg-white shadow-md">
    <ul class="flex justify-around items-center p-4">
      <li>
        <router-link to="/ai" class="flex flex-col items-center text-center">
          <PhRobot
            :size="32"
            :weight="isActive('/ai') ? 'fill' : 'regular'" 
            :class="{ 'text-green-500': isActive('/ai') }"
            :style="{ color: isActive('/ai') ? '#92C700' : 'black' }" 
          />
          <span v-if="isActive('/ai')" class="text-gray-800">阿哞</span>
        </router-link>
      </li>
      <li>
        <router-link to="/home" class="flex flex-col items-center text-center">
          <PhHouse
            :size="32"
            :weight="isActive('/home') ? 'fill' : 'regular'" 
            :class="{ 'text-green-500': isActive('/home') }"
            :style="{ color: isActive('/home') ? '#92C700' : 'black' }" 
          />
          <span v-if="isActive('/home')" class="text-gray-800">首頁</span>
        </router-link>
      </li>
      <li>
        <router-link to="/settings" class="flex flex-col items-center text-center">
          <PhUserGear
            :size="32"
            :weight="isActive('/settings') ? 'fill' : 'regular'" 
            :class="{ 'text-green-500': isActive('/settings') }"
            :style="{ color: isActive('/settings') ? '#92C700' : 'black' }" 
          />
          <span v-if="isActive('/settings')" class="text-gray-800">用戶</span>
        </router-link>
      </li>
    </ul>
  </nav>
</template>

<script>
import { PhRobot, PhUserGear, PhHouse } from '@phosphor-icons/vue';

export default {
  components: {
    PhRobot,
    PhUserGear,
    PhHouse,
  },
  methods: {
    isActive(route) {
      return this.$route.path === route;
    },
  },
};
</script>