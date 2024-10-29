<template>
  <div class="h-full flex flex-col overflow-hidden">
    <header class=" basis-3/12 overflow-auto"> <!-- 固定高度的 header -->
      <img v-if="!doingPw" src="../images/login.jpg" class="w-full object-cover" />
      <img v-else src="../images/password.jpg" class="w-full object-cover" />
    </header>

    <main class="flex-grow flex-col flex items-center"> <!-- Flexbox 居中表单 -->
      <form class="flex flex-col justify-center text-center w-full max-w-sm"> <!-- 限制最大宽度 -->
        <label class="text-4xl font-bold my-3">Register</label>
        <div class="flex items-center border-2 rounded-md px-3 my-3 ml-10 mr-10 shadow-md">
          <PhUser :size="32" color="#4d4d4d" />
          <input v-model="user.account" class="flex-1 p-3 focus:outline-none" type="text" placeholder="帳號" />
        </div>
        <!-- 密碼輸入框 -->
        <div class="flex items-center border-2 rounded-md px-3 ml-10 my-3 mr-10 shadow-md">
          <PhLockKey :size="32" color="#4d4d4d" />
          <input v-model="password" @focus="doingPw = true" @blur="doingPw = false"
            class="flex-1 p-3 focus:outline-none" :type="showPassword ? 'text' : 'password'" placeholder="密碼" />
          <PhEyeClosed v-if="!showPassword" @click="togglePassword" :size="32" color="#4d4d4d" />
          <PhEye v-else @click="togglePassword" :size="32" />
        </div>
        <div class="flex items-center border-2 rounded-md px-3 ml-10 my-3 mr-10 shadow-md">
          <PhEnvelopeSimple :size="32" color="#4d4d4d" />
          <input v-model="user.email" class="flex-1 p-3 focus:outline-none" type="email" placeholder="信箱" />
        </div>

        <div class="mt-5 mx-10 flex items-center">
          <PhSquare v-if="!check" @click="checklogin" :size="20" />
          <PhCheckSquare v-else @click="checklogin" :size="20" />
          <span class="ml-2">保持登入</span>
        </div>

        <div class="mt-5 mx-10">

          <button @click="registerHandler(user)"
            class="w-full bg-[#90c700] text-white font-bold text-2xl py-3 rounded shadow-md" type="button">
            創建
          </button>
        </div>
        <div class="  ">
          <div class="flex items-center mt-44">
            <div class="flex-1 border-t"></div>
            <span class="mx-3">or</span>
            <div class="flex-1 border-t"></div>
          </div>
          <div class="mb-5">
            <router-link to="/login" class="text-[#90c700] underline hover:text-[#7ab600]">
              登入
            </router-link>
          </div>
        </div>

      </form>
    </main>
  </div>

</template>

<script setup lang="ts">
import request from "@/axios";
import router from "@/router";
import {
  PhEnvelopeSimple,
  PhEyeClosed,
  PhLockKey,
  PhUser,
  PhEye,
  PhSquare,
  PhCheckSquare,
} from "@phosphor-icons/vue";
import { ref } from "vue";

class form {
  account: string = "";
  password: string = "";
  email: string = ""
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
const password = ref("");  // 用于存储密码
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
