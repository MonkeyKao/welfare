<template>
  <!-- 固定定位的 header，確保圖片維持在最上方 -->
  <header>
    <img src="/forgetpassword.jpg" class="w-full" />
  </header>
  <body>
    <router-link to="/forgot-password" class="forgot-password-link"
      ><PhArrowUUpLeft :size="32" color="#4d4d4d" class="ml-5"
    /></router-link>
    <form class="flex flex-col justify-center text-center">
      <label class="text-3xl font-bold">密碼重設</label>
      <!-- 密碼輸入框 -->
      <div
        class="flex items-center border-2 rounded-md px-3 ml-10 mr-10 shadow-md mt-4"
      >
        <PhLockKey :size="32" color="#4d4d4d" />
        <input
          class="flex-1 p-3 focus:outline-none"
          :type="showPassword ? 'text' : 'password'"
          placeholder="密碼"
        />
        <PhEyeClosed
          v-if="!showPassword"
          @click="togglePassword"
          :size="32"
          color="#4d4d4d"
        />
        <PhEye v-else @click="togglePassword" :size="32" />
      </div>

      <!-- 密碼輸入框 -->
      <div
        class="flex items-center border-2 rounded-md px-3 ml-10 mr-10 shadow-md mt-4"
      >
        <PhLockKey :size="32" color="#4d4d4d" />
        <input
          class="flex-1 p-3 focus:outline-none"
          :type="showPassword ? 'text' : 'password'"
          placeholder="再次輸入密碼"
        />
        <PhEyeClosed
          v-if="!showPassword"
          @click="togglePassword"
          :size="32"
          color="#4d4d4d"
        />
        <PhEye v-else @click="togglePassword" :size="32" />
      </div>

      <div class="mt-5 mx-10">
        <router-link to="/reset-password" class="forgot-password-link"
          ><button
            class="w-full bg-[#90c700] text-white font-bold text-2xl py-3 rounded shadow-md"
            type="button"
          >
            確定
          </button></router-link
        >
      </div>
    </form>
  </body>
</template>
<script setup lang="ts">
import { PhEye, PhEyeClosed, PhLockKey } from "@phosphor-icons/vue";
import { ref } from "vue";

//圖片更動
const password = ref(""); // 用于存储密码
const doingPw = ref(false); // 用于指示是否处于密码输入状态

//點眼睛密碼明文顯示
const showPassword = ref(false);
const togglePassword = () => {
  showPassword.value = !showPassword.value; // 切换 showPassword 的值
};
</script>
<style scoped>
/* 固定 header 區域 */
header.img {
  width: 100%; /* 讓圖片填滿 header */
  display: block; /* 移除任何潛在的空白間隙 */
}

/* Label 樣式 */
.title {
  font-size: 30px;
}

.form-container {
  display: flex;
  flex-direction: column;
  justify-content: center; /* 水平置中 */
  margin-top: 20px; /* 與圖片間距 */
  text-align: center;
}
.inputbox {
  display: flex;
  justify-content: center; /* 水平置中 */
  margin-top: 20px; /* 與標題之間的垂直間距 */
}

/* Input 樣式 */
.input {
  margin-top: 20px; /* 與標題之間的垂直間距 */
  border-radius: 0.25rem; /* 圓角 */
  padding: 0.5rem 1rem; /* 內距，確保文字不靠近邊緣 */
  color: #4a5568; /* 文字顏色 */
  line-height: 1.25; /* 行高，讓內容更緊湊 */
  box-shadow: 0px 4px 6px rgba(0, 0, 0, 0.1); /* 陰影效果 */
}

.button {
  /* 背景色和文字色 */
  background-color: #90c700; /* 對應 bg-purple-500 */
  color: #ffffff; /* 對應 text-white */

  /* 字體與按鈕外觀 */
  font-weight: bold; /* 對應 font-bold */
  padding: 0.5rem 1rem; /* 對應 py-2 px-4 */
  border-radius: 0.25rem; /* 對應 rounded */

  /* 陰影 */
  box-shadow: 0px 1px 3px rgba(0, 0, 0, 0.1), 0px 1px 2px rgba(0, 0, 0, 0.06); /* 對應 shadow */
  outline: none; /* 對應 focus:outline-none */
}
.backbutton {
  margin-top: 20px; /* 與標題之間的垂直間距 */
}
</style>
