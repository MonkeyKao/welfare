<template>
  <!-- 固定定位的 header，確保圖片維持在最上方 -->
  <header>
    <img src="/forgetpassword.jpg" class="w-full" />
  </header>
  <body>
    <div class="backbutton">
      <router-link to="/" class="forgot-password-link">返回</router-link>
    </div>
    <br />
    <form class="form-container">
      <label class="title">忘記密碼</label>

      <div>
        <div class="inputbox">
          <input
            class="input"
            v-model="contact"
            type="text"
            placeholder="信箱"
          />
        </div>

        <div class="inputbox">
          <input
            class="input"
            type="password"
            v-model="verificationCode"
            placeholder="驗證碼"
          />
          <button @click="sendVerificationCode">發送驗證碼</button>
        </div>
      </div>
      <div class="reset-password">
        <router-link to="/reset-password" class="forgot-password-link"
          ><button class="button" type="button" @click="verifyCode">
            下一步
          </button></router-link
        >
      </div>
      <div></div>
    </form>
  </body>
</template>
<script setup lang="ts">
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
