<template>
  <nav class="fixed bottom-0 left-0 right-0 bg-white shadow-md">
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
        <router-link to="/user" class="flex flex-col items-center text-center">
          <PhUserGear
            :size="32"
            :weight="isActive('/user') ? 'fill' : 'regular'"
            :class="{ 'text-green-500': isActive('/user') }"
            :style="{ color: isActive('/user') ? '#92C700' : 'black' }"
          />
          <span v-if="isActive('/user')" class="text-gray-800">用戶</span>
        </router-link>
      </li>
    </ul>
  </nav>
</template>

<script>
import { PhHouse, PhRobot, PhUserGear } from "@phosphor-icons/vue";

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
