<template>
  <div>
    <div class="flex flex-col">
      <span class="text-H3">剩餘有效時間:{{ formattedTime }}</span>
      <span class="text-H3">代碼:{{ code }}</span>
      <img :src="qrCodeUrl" alt="QR Code" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue';
import request from '@/axios';

const qrCodeUrl = ref('');
const code = ref('');
const time = ref(300);  // 5 分鐘
const showQRCode = ref(false);

// 格式化倒計時
const formattedTime = computed(() => {
  const minutes = Math.floor(time.value / 60);
  const seconds = time.value % 60;
  return `${String(minutes).padStart(2, '0')}:${String(seconds).padStart(2, '0')}`;
});

onMounted(() => {
  fetchQRCode();
})
// 取得 QR code
const fetchQRCode = async () => {
  try {
    const response = await request.get('/users/bind');
    const data = response.data;

    // 更新 QR code 和 code
    qrCodeUrl.value = data.image; // Base64 編碼的圖片
    code.value = data.code; // 生成的代碼

    // 開始倒計時
    showQRCode.value = true;
    const interval = setInterval(() => {
      if (time.value > 0) {
        time.value -= 1;
      } else {
        clearInterval(interval); // 停止倒計時
      }
    }, 1000); // 每秒更新一次
  } catch (error) {
    console.error("Error fetching QR code:", error);
  }
};
</script>