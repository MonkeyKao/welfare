<template>
  <!-- 固定定位的 header，確保圖片維持在最上方 -->
  <header>
    <img src="/forgetpassword.jpg" class="w-full" />
  </header>
  <body>
    <router-link to="/" class="forgot-password-link"
      ><PhArrowUUpLeft :size="32" color="#4d4d4d" class="ml-5"
    /></router-link>
    <form class="flex flex-col justify-center text-center">
      <label class="text-3xl font-bold">忘記密碼</label>
      <div
        class="flex items-center border-2 rounded-md px-3 mx-10 shadow-md mt-4"
      >
        <PhEnvelopeSimple :size="32" color="#4d4d4d" />
        <input
          class="flex-1 p-2 focus:outline-none"
          type="email"
          placeholder="信箱"
        />
      </div>

      <div
        class="flex items-center border-2 rounded-md px-3 mx-10 shadow-md mt-4"
      >
        <PhShieldCheck :size="32" color="#4d4d4d" />
        <input
          class="flex-1 p-2 focus:outline-none"
          type="password"
          v-model="verificationCode"
          placeholder="驗證碼"
        />
        <button @click="sendVerificationCode" class="text-[#92c700]">
          發送驗證碼
        </button>
      </div>

      <div class="mt-5 mx-10">
        <router-link to="/reset-password" class="forgot-password-link"
          ><button
            class="w-full bg-[#90c700] text-white font-bold text-2xl py-3 rounded shadow-md"
            type="button"
            @click="verifyCode"
          >
            下一步
          </button></router-link
        >
      </div>
    </form>
  </body>
</template>
<script setup lang="ts">
import {
  PhArrowUUpLeft,
  PhEnvelopeSimple,
  PhShieldCheck,
} from "@phosphor-icons/vue";
import { ref } from "vue";
const contact = ref("");
const verificationCode = ref("");

const sendVerificationCode = async () => {
  try {
    const response = await fetch("/api/send-verification-code", {
      method: "POST",
      body: JSON.stringify({ contact: contact.value }),
      headers: {
        "Content-Type": "application/json",
      },
    });
    const result = await response.json();
    console.log(result.message);
  } catch (error) {
    console.error("發送驗證碼時發生錯誤:", error);
  }
};

const verifyCode = async () => {
  try {
    const response = await fetch("/api/verify-code", {
      method: "POST",
      body: JSON.stringify({
        contact: contact.value,
        code: verificationCode.value,
      }),
      headers: {
        "Content-Type": "application/json",
      },
    });
    const result = await response.json();
    console.log(result.message);
  } catch (error) {
    console.error("驗證碼驗證時發生錯誤:", error);
  }
};
</script>
<style scoped>
/* 固定 header 區域 */
header.img {
  width: 100%; /* 讓圖片填滿 header */
  display: block; /* 移除任何潛在的空白間隙 */
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
</style>
