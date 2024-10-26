<template>
  <!-- 固定定位的 header，確保圖片維持在最上方 -->
  <header>
    <img src="/login.jpg" class="w-full" />
  </header>

  <body>
    <form class="flex flex-col justify-center mt-5 text-center">
      <label class="text-3xl font-bold">Registar</label>
      <div class="flex items-center border-2 rounded-md px-3 mx-10 shadow-md mt-4">
        <PhUser :size="32" color="#4d4d4d" />
        <input v-model="user.account" class="flex-1 p-2 focus:outline-none" type="text" placeholder="帳號" />
      </div>
      <div class="flex items-center border-2 rounded-md px-3 mx-10 shadow-md mt-4">
        <PhLockKey :size="32" color="#4d4d4d" />
        <input v-model="user.password" class="flex-1 p-2 focus:outline-none" type="password" placeholder="密碼" />
        <PhEyeClosed :size="32" color="#4d4d4d" />
      </div>
      <div class="flex items-center border-2 rounded-md px-3 mx-10 shadow-md mt-4">
        <PhEnvelopeSimple :size="32" color="#4d4d4d" />
        <input v-model="user.email" class="flex-1 p-2 focus:outline-none" type="email" placeholder="信箱" />
      </div>

      <div class="mt-5">
        <label class="flex items-center">
          <input type="checkbox" class="mr-2 ml-10" /> 保持登入
        </label>
      </div>

      <div class="mt-5 mx-10">
        <router-link to="/verify">
          <button @click="registerHandler(user)"
            class="w-full bg-[#90c700] text-white font-bold text-2xl py-3 rounded shadow-md" type="button">
            創建
          </button></router-link>
      </div>

      <div class="loginbottom inset-x-0 fixed bottom-5">
        <div class="flex flex-col items-center">
          <div class="flex items-center mt-5">
            <div class="flex-1 border-t"></div>
            <span class="mx-3">or</span>
            <div class="flex-1 border-t"></div>
          </div>
          <div class="mt-5">
            <router-link to="/login" class="text-[#90c700] underline hover:text-[#7ab600]">
              登入
            </router-link>
          </div>
        </div>
      </div>
    </form>
  </body>
</template>

<script setup lang="ts">
import request from "@/axios";
import {
  PhEnvelopeSimple,
  PhEyeClosed,
  PhLockKey,
  PhUser,
} from "@phosphor-icons/vue";
import { ref } from "vue";

class form {
  account: string = "";
  password: string = "";
  email: string = ""
}
const user = ref<form>(new form());

const registerHandler = async (user: form) => {
  const json = JSON.stringify(user);
  const result = await request.post("/users/register", json)
  localStorage.setItem("s",result.data)

}

function loginWithFacebook() {
  // 呼叫 Facebook 登入 API 的程式邏輯
}
function loginWithGoogle() {
  // 呼叫 Google 登入 API 的程式邏輯
}
</script>
