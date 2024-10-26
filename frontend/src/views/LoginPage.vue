<template>
  <!-- 固定定位的 header，確保圖片維持在最上方 -->
  <header>
    <img v-if="!doingPw" src="../images/login.jpg" />
    <img v-else src="../images/password.jpg" />
  </header>

  <main class="flex-grow flex items-center justify-center"> <!-- Flexbox 居中表单 -->
    <form class="flex flex-col justify-center text-center w-full max-w-sm"> <!-- 限制最大宽度 -->
      <label class="text-4xl font-bold my-3">Login</label>
      <div class="flex items-center border-2 rounded-md px-3 my-3 ml-10 mr-10 shadow-md">
        <PhUser :size="32" color="#4d4d4d" />
        <input v-model="user.account" class="flex-1 p-3 focus:outline-none" type="text" placeholder="帳號" />
      </div>

      <!-- 密碼輸入框 -->
      <div class="flex items-center border-2 rounded-md px-3 ml-10 my-3 mr-10 shadow-md">
        <PhLockKey :size="32" color="#4d4d4d" />
        <input v-model="password" @focus="doingPw = true" @blur="doingPw = false" class="flex-1 p-3 focus:outline-none"
          :type="showPassword ? 'text' : 'password'" placeholder="密碼" />
        <PhEyeClosed v-if="!showPassword" @click="togglePassword" :size="32" color="#4d4d4d" />
        <PhEye v-else @click="togglePassword" :size="32" />
      </div>

      <div class="mt-5 mx-10 flex items-center">
        <PhSquare v-if="!check" @click="checklogin" :size="20" />
        <PhCheckSquare v-else @click="checklogin" :size="20" />
        <span class="ml-2">保持登入</span>
      </div>

      <div class="mt-5 mx-10">
        <button @click="loginHandler(user)"
          class="w-full bg-[#90c700] text-white font-bold text-2xl py-3 rounded shadow-md" type="button">
          登入
        </button>
      </div>

      <div class="mt-5">
        <router-link to="/forgot-password" class="text-[#90c700] underline hover:text-[#7ab600]">
          忘記密碼？
        </router-link>
      </div>

      <div class="mt-20">
        <div class="flex flex-col items-center">
          <label>其他登入方式</label>
          <div class="flex justify-center space-x-4 mt-4">
            <button @click="loginWithGoogle"
              class="w-12 h-12 rounded-full flex justify-center items-center shadow bg-[#92c700]">
              <PhGoogleLogo :size="32" color="#fff" />
            </button>
            <button @click="loginWithFacebook"
              class="w-12 h-12 rounded-full flex justify-center items-center shadow bg-[#92c700]">
              <PhFacebookLogo :size="32" color="#fff" />
            </button>
          </div>
        </div>

        <div class="flex items-center mt-5">
          <div class="flex-1 border-t"></div>
          <span class="mx-3">or</span>
          <div class="flex-1 border-t"></div>
        </div>

        <div class="mb-5">
          <router-link to="/register" class="text-[#90c700] underline hover:text-[#7ab600]">
            創建帳號
          </router-link>
        </div>
      </div>
    </form>
  </main>

</template>



<script setup lang="ts">
import request from "@/axios";
import { User } from "@/model/user";
import router from "@/router";
import {
  PhEye,
  PhEyeClosed,
  PhFacebookLogo,
  PhGoogleLogo,
  PhLockKey,
  PhUser,
  PhSquare,
  PhCheckSquare,
} from "@phosphor-icons/vue";
import { ref } from "vue";
class form {
  account: string = "";
  password: string = "";
}
const user = ref<form>(new form());

const loginHandler = async (user: form) => {
  const result = await request.post("users/doLogin", JSON.stringify(user))
  localStorage.setItem(result.data.tokenName, result.data.tokenValue)
  router.push("/home")

}

function loginWithFacebook() {
  // 呼叫 Facebook 登入 API 的程式邏輯
}
function loginWithGoogle() {
  // 呼叫 Google 登入 API 的程式邏輯
}

//圖片更動
const password = ref(""); // 用于存储密码
const doingPw = ref(false); // 用于指示是否处于密码输入状态

//點眼睛密碼明文顯示
const showPassword = ref(false);
const togglePassword = () => {
  showPassword.value = !showPassword.value; // 切换 showPassword 的值
}

const check = ref(false);
const checklogin = () => {
  check.value = !check.value;
}

</script>
