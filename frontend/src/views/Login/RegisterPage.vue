<template>
  <div class="h-full flex flex-col overflow-hidden">
    <header class="basis-3/12 overflow-auto">
      <!-- 固定高度的 header -->
      <img
        v-if="!doingPw"
        src="../images/login.jpg"
        class="w-full object-cover"
      />
      <img v-else src="../images/password.jpg" class="w-full object-cover" />
    </header>

    <div
      class="flex flex-col w-full md:max-w-md px-10 mt-5 sm:mt-1 flex-grow h-auto"
    >
      <!-- Flexbox 居中表单 -->
      <form class="flex flex-col gap-5">
        <!-- 限制最大宽度 -->
        <label class="flex text-4xl sm:text-3xl font-bold justify-center"
          >Register</label
        >
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
            v-model="password"
            @focus="doingPw = true"
            @blur="doingPw = false"
            class="mx-2 min-w-0 text-base max-w-xs md:max-w-md lg:max-w-lg sm:text-sm lg:text-lg"
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
        <div class="flex w-full border-2 rounded-md shadow-md p-3 items-center">
          <PhEnvelopeSimple :size="32" color="#4d4d4d" class="flex-shrink-0" />
          <input
            v-model="user.email"
            class="mx-2 min-w-0 text-base max-w-xs md:max-w-md lg:max-w-lg sm:text-sm lg:text-lg"
            type="email"
            placeholder="信箱"
          />
        </div>

        <div class="flex items-center">
          <PhSquare v-if="!check" @click="checklogin" :size="20" />
          <PhCheckSquare v-else @click="checklogin" :size="20" />
          <span class="ml-2 text-lg sm:text-base">保持登入</span>
        </div>

        <div>
          <router-link to="/verify">
            <button
              @click="registerHandler(user)"
              class="w-full bg-[#90c700] text-white font-bold text-2xl py-3 sm:py-2 rounded shadow-md"
              type="button"
            >
              創建
            </button></router-link
          >
        </div>
        <div class="pb-5">
          <div>
            <div class="flex items-center mt-44">
              <div class="flex-1 border-t"></div>
              <span class="mx-3">or</span>
              <div class="flex-1 border-t"></div>
            </div>
            <div>
              <router-link
                to="/login"
                class="flex text-lg sm:text-base text-[#92c700] justify-center"
              >
                登入
              </router-link>
            </div>
          </div>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import request from "@/axios";
import router from "@/router";
import {
  PhCheckSquare,
  PhEnvelopeSimple,
  PhEye,
  PhEyeClosed,
  PhLockKey,
  PhSquare,
  PhUser,
} from "@phosphor-icons/vue";
import { ref } from "vue";

class form {
  account: string = "";
  password: string = "";
  email: string = "";
}
const user = ref<form>(new form());

const registerHandler = async (user: form) => {
  try {
    const json = JSON.stringify(user);
    const result = await request.post("/users/register", json)
    console.log(result.data.msg);
    
    localStorage.setItem("token", result.data.msg)
    router.push("verify")
  } catch (error:any) {
    alert(error.response.data)
  }


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
