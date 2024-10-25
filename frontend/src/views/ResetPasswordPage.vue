<template>
  <!-- 固定定位的 header，確保圖片維持在最上方 -->
  <header>
    <img src="/forgetpassword.jpg" class="w-full" />
  </header>
  <body>
    <div class="backbutton">
      <router-link to="/forgot-password" class="forgot-password-link"
        >返回</router-link
      >
    </div>
    <br />
    <form class="form-container">
      <label class="title">密碼重設</label>

      <div class="inputbox">
        <input class="input" type="text" placeholder="密碼" />
      </div>
      <div class="inputbox">
        <input class="input" type="password" placeholder="再次輸入密碼" />
      </div>
      <div>
        <button class="button" type="button">確定</button>
      </div>
    </form>
  </body>
</template>
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
