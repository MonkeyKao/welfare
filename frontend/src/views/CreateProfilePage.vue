<template>
  <!-- 固定定位的 header，確保圖片維持在最上方 -->
  <header class="flex flex-col items-center mt-5 text-center">
    <label class="title">個人資料</label>
    <div class="w-24 h-24 flex justify-center items-center">
      <img src="/logo.png" />
    </div>
  </header>
  <body>
    <form class="form-container">
      <div class="column mt-5">
        <label class="text-left text-[20px] mb-2">姓名</label>
        <div class="inputbox">
          <input class="input" type="text" />
        </div>
      </div>
      <div class="column mt-5">
        <label class="text-left text-[20px] mb-2">生日</label>
        <div class="inputbox">
          <input class="input" type="date" />
        </div>
      </div>
      <div class="column mt-5">
        <label class="text-left text-[20px] mb-2">性別</label>
        <div class="inputbox">
          <input class="input" type="radio" />
        </div>
      </div>
      <div class="column mt-5">
        <label class="text-left text-[20px] mb-2">電話號碼</label>
        <div class="inputbox">
          <input class="input" type="number" />
        </div>
      </div>
      <div class="column mt-5">
        <label class="text-left text-[20px] mb-2">地區</label>
        <div class="inputbox">
          <ul class="input">
            <li>台北市</li>
            <li>新北市</li>
          </ul>
        </div>
      </div>
      <div class="inputbox">
        <input class="input" type="email" placeholder="信箱" />
      </div>

      <div class="mt-5">
        <RouterLink to="verify">
          <button class="button" type="button">繼續</button></RouterLink
        >
      </div>

      <div class="inset-x-0 fixed bottom-5">
        <router-link to="/home" class="forgot-password-link">略過</router-link>
      </div>
    </form>
  </body>
</template>

<script setup lang="ts">
function loginWithFacebook() {
  // 呼叫 Facebook 登入 API 的程式邏輯
}
function loginWithGoogle() {
  // 呼叫 Google 登入 API 的程式邏輯
}
</script>

<style scoped>
/* 固定 header 區域 */
header img {
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
}

/* Input 樣式 */
.input {
  margin-top: 10px; /* 與標題之間的垂直間距 */
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

/* 連結樣式 */
.forgot-password {
  margin-top: 10px;
}

.forgot-password-link {
  color: #90c700;
  text-decoration: underline;
  cursor: pointer;
}
.forgot-password-link:hover {
  color: #7ab600;
}
</style>
