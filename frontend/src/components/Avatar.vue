<template>
  <div class="flex flex-col items-center mt-5 text-center">
    <div class="w-24 h-24 flex justify-center items-center">
      <!-- 使用 Pinia 存储的头像或默认头像 -->
      <img :src="computedSrc" class="w-full h-full object-cover rounded-full" alt="頭像" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { useUserStore } from "@/store/userStroe";
import { computed } from "vue";

const userStore = useUserStore();
const props = defineProps(['src']);

// 默认头像
const defaultAvatar = '/logo.png';

// 计算头像的实际来源
const computedSrc = computed(() => {
  return props.src || (userStore.avatarBase64 ? `data:image/png;base64,${userStore.avatarBase64}` : defaultAvatar);
});
</script>