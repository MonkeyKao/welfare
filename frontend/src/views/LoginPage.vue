<template>
  <div class="relative h-screen flex flex-col overflow-hidden">
    <header>
      <img v-if="!doingPw" src="../images/login.jpg" class="w-full h-auto" />
      <img v-else src="../images/password.jpg" class="w-full h-auto" />
    </header>

    <!-- Flexbox 居中表單 -->
    <div
      class="flex flex-col w-full md:max-w-md px-10 mt-5 sm:mt-1 flex-grow h-auto"
    >
      <form class="flex flex-col gap-5">
        <!-- 標題 -->
        <label class="flex text-4xl sm:text-3xl font-bold justify-center"
          >Login</label
        >

        <!-- 帳號輸入框 -->
        <div class="flex w-full border-2 rounded-md shadow-md p-3 items-center">
          <PhUser :size="32" color="#4d4d4d" class="flex-shrink-0" />
          <input
            v-model="user.account"
            class="mx-2 min-w-0 text-base max-w-xs md:max-w-md lg:max-w-lg sm:text-sm lg:text-lg"
            type="text"
            placeholder="帳號"
          />
        </div>

        <!-- 密碼輸入框 -->
        <div class="flex w-full border-2 rounded-md shadow-md p-3 items-center">
          <PhLockKey :size="32" color="#4d4d4d" class="flex-shrink-0" />
          <input
            v-model="user.password"
            @focus="doingPw = true"
            @blur="doingPw = false"
            class="mx-2 min-w-0 text-base"
            :type="showPassword ? 'text' : 'password'"
            placeholder="密碼"
          />
          <PhEyeClosed
            v-if="!showPassword"
            @click="togglePassword"
            :size="32"
            color="#4d4d4d"
            class="flex-shrink-0"
          />
          <PhEye
            v-else
            @click="togglePassword"
            :size="32"
            class="flex-shrink-0"
          />
        </div>

        <!-- 按鈕 -->
        <button
          @click="loginHandler(user)"
          class="w-full bg-[#90c700] text-white font-bold text-2xl py-3 sm:py-2 rounded shadow-md"
          type="button"
        >
          登入
        </button>

        <router-link
          to="/forgot-password"
          class="flex text-lg sm:text-base text-[#92c700] justify-center"
        >
          忘記密碼？
        </router-link>
      </form>

      <!-- 底部連結區域 -->
      <div class="mt-auto pb-5">
        <div class="flex flex-col justify-center">
          <div>
            <label class="flex text-sm sm:text-base justify-center"
              >其他登入方式</label
            >
            <div class="flex justify-center space-x-4 mt-4">
              <button
                class="w-12 h-12 sm:w-10 sm:h-10 rounded-full flex justify-center items-center shadow bg-[#92c700]"
              >
                <PhGoogleLogo :size="24" color="#fff" />
              </button>
              <button
                class="w-12 h-12 sm:w-10 sm:h-10 rounded-full flex justify-center items-center shadow bg-[#92c700]"
              >
                <PhFacebookLogo :size="24" color="#fff" />
              </button>
            </div>
          </div>
          <div class="flex items-center">
            <div class="flex-1 border-t"></div>
            <span class="mx-3">or</span>
            <div class="flex-1 border-t"></div>
          </div>
          <router-link
            to="/register"
            class="flex text-lg sm:text-base text-[#92c700] justify-center"
          >
            創建帳號
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import request from "@/axios";
import router from "@/router";
import {
  PhEye,
  PhEyeClosed,
  PhFacebookLogo,
  PhGoogleLogo,
  PhLockKey,
  PhUser,
} from "@phosphor-icons/vue";
import { ref } from "vue";
class form {
  account: string = "";
  password: string = "";
}
const user = ref<form>(new form());

const loginHandler = async (user: form) => {
  try {
    const result = await request.post("users/doLogin", JSON.stringify(user));
    localStorage.setItem(result.data.tokenName, result.data.tokenValue);
  } catch (err) {
    console.log("123");
    return;
  }

  router.push("/home");
};

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
};

const check = ref(false);
const checklogin = () => {
  check.value = !check.value;
};
</script>
