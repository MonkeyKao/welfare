<template>
  <!-- 固定定位的 header，確保圖片維持在最上方 -->
  <header>
    <img src="/login.jpg" class="w-full" />
  </header>
  <body class="h-screen overflow-y-auto relative">
    <form class="flex flex-col justify-center mt-5 text-center">
      <label class="text-3xl font-bold">Login</label>

      <div
        class="flex items-center border-2 rounded-md px-3 ml-10 mr-10 shadow-md mt-4"
      >
        <PhUser :size="32" color="#4d4d4d" />
        <input
          class="flex-1 p-2 focus:outline-none"
          type="text"
          placeholder="帳號"
        />
      </div>
      <div
        class="flex items-center border-2 rounded-md px-3 ml-10 mr-10 shadow-md mt-4"
      >
        <PhLockKey :size="32" color="#4d4d4d" />
        <input
          class="flex-1 p-2 focus:outline-none"
          type="password"
          placeholder="密碼"
        />
        <PhEyeClosed :size="32" color="#4d4d4d" />
      </div>

      <div class="mt-5">
        <label class="flex items-center">
          <input type="checkbox" class="mr-2 ml-10 font-H3-regular" /> 保持登入
        </label>
      </div>

      <div class="mt-5 mx-10">
        <button
          class="w-full bg-[#90c700] text-white font-bold text-2xl py-3 rounded shadow-md"
          type="button"
        >
          登入
        </button>
      </div>

      <div class="mt-5">
        <router-link
          to="/forgot-password"
          class="text-[#90c700] underline hover:text-[#7ab600]"
        >
          忘記密碼？
        </router-link>
      </div>

      <div class="mt-20">
        <div class="flex flex-col items-center">
          <label>其他登入方式</label>
          <div class="flex justify-center space-x-4 mt-4">
            <button
              @click="loginWithGoogle"
              class="w-12 h-12 rounded-full flex justify-center items-center shadow bg-[#92c700]"
            >
              <PhGoogleLogo :size="32" color="#fff" />
            </button>
            <button
              @click="loginWithFacebook"
              class="w-12 h-12 rounded-full flex justify-center items-center shadow bg-[#92c700]"
            >
              <PhFacebookLogo :size="32" color="#fff" />
            </button>
          </div>
        </div>

        <div class="flex items-center mt-5">
          <div class="flex-1 border-t"></div>
          <span class="mx-3">or</span>
          <div class="flex-1 border-t"></div>
        </div>

        <div class="mt-5">
          <router-link
            to="/register"
            class="text-[#90c700] underline hover:text-[#7ab600]"
          >
            創建帳號
          </router-link>
        </div>
      </div>
    </form>
  </body>
</template>

<script setup lang="ts">
import {
  PhEyeClosed,
  PhFacebookLogo,
  PhGoogleLogo,
  PhLockKey,
  PhUser,
} from "@phosphor-icons/vue";

function loginWithFacebook() {
  // 呼叫 Facebook 登入 API 的程式邏輯
}
function loginWithGoogle() {
  // 呼叫 Google 登入 API 的程式邏輯
}
</script>
