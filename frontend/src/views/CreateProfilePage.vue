<template>
  <!-- 固定定位的 header，確保圖片維持在最上方 -->
  <header class="flex flex-col items-center mt-5 text-center">
    <label class="text-3xl font-bold">個人資料</label>
    <div class="w-24 h-24 flex justify-center items-center">
      <img src="/logo.png" class="h-full object-cover" />
    </div>
  </header>
  <main>
    <form class="flex flex-col justify-center items-center mt-5">
      <div>
        <div class="mt-5">
          <label class="block text-left text-xl mb-2">姓名</label>
          <input
            class="flex-1 focus:outline-none border-2 rounded-md px-3 shadow-md"
            type="text"
          />
        </div>
        <div class="mt-5">
          <label class="block text-left text-xl mb-2">生日</label>
          <input
            class="flex w-full focus:outline-none border-2 rounded-md px-3 shadow-md"
            type="date"
          />
        </div>
        <div class="mt-5">
          <label class="block text-left text-xl mb-2">性別</label>
          <div class="flex items-center">
            <input type="radio" id="male" name="gender" class="mr-2" />
            <label for="male">男性</label>
            <input type="radio" id="female" name="gender" class="mr-2 ml-4" />
            <label for="neutral">中性</label>
            <input type="radio" id="female" name="gender" class="mr-2 ml-4" />
            <label for="female">女性</label>
          </div>
        </div>

        <div class="mt-5">
          <label class="block text-left text-xl mb-2">地區</label>
          <select
            class="flex-1 focus:outline-none border-2 rounded-md px-3 shadow-md"
          >
            <option>台北市</option>
            <option>新北市</option>
            <option>基隆市</option>
            <option>新竹市</option>
            <option>新竹縣</option>
            <option>苗栗縣</option>
            <option>台中市</option>
            <option>彰化縣</option>
            <option>雲林縣</option>
            <option>嘉義市</option>
            <option>嘉義縣</option>
            <option>台南市</option>
            <option>高雄市</option>
            <option>屏東縣</option>
            <option>南投縣</option>
            <option>宜蘭縣</option>
            <option>花蓮縣</option>
            <option>台東縣</option>
          </select>
        </div>

        <div class="mt-5">
          <router-link to="/home"
            ><button
              class="w-full bg-[#90c700] text-white font-bold text-2xl py-3 rounded shadow-md"
              type="button"
            >
              繼續
            </button></router-link
          >
        </div>
      </div>

      <div
        class="flex mt-5 absolute bottom-5 right-5 justify-center items-center"
      >
        <router-link to="/home">先略過</router-link>
        <PhArrowCircleRight :size="32" />
      </div>
    </form>
  </main>
</template>

<script setup lang="ts">
import { PhArrowCircleRight } from "@phosphor-icons/vue";
function loginWithFacebook() {
  // 呼叫 Facebook 登入 API 的程式邏輯
}
function loginWithGoogle() {
  // 呼叫 Google 登入 API 的程式邏輯
}
</script>
