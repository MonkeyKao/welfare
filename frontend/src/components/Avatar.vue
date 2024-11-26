<template>
  <div class="flex flex-col items-center mt-5 text-center">
    <div class="w-24 h-24 flex justify-center items-center">
      <!-- 使用 Pinia 存储的头像 -->
      <img :src="src?src:'data:image/png;base64,'+userStore.avatarBase64" class="w-full h-full object-cover rounded-full" alt="頭像" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { useUserStore } from "@/store/userStroe";

const userStore = useUserStore();

const props = defineProps(['src'])
</script>