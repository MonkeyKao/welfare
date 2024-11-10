<template>
  <div class="flex flex-col h-screen w-full">
    <div class="flex-shrink-0">
      <!-- 固定高度的 header -->
      <img v-if="!doingPw" src="../../images/login.jpg" />
      <img v-else src="../../images/password.jpg" />
    </div>

    <div class="flex flex-col flex-grow">
      <!-- Flexbox 居中表单 -->
      <form class="flex flex-col gap-5 px-10 mt-5">
        <!-- 限制最大宽度 -->
        <label class="flex text-4xl sm:text-3xl font-bold justify-center">Register</label>
        <div class="flex border-2 rounded-md shadow-md p-3 items-center">
          <PhUser :size="32" color="#4d4d4d" class="flex-shrink-0" />
          <input v-model="user.account" class="mx-2 min-w-0 text-base" type="text" placeholder="帳號" />
        </div>
        <!-- 密碼輸入框 -->
        <div class="flex border-2 rounded-md shadow-md p-3 items-center">
          <PhLockKey :size="32" color="#4d4d4d" class="flex-shrink-0" />
          <input v-model="user.password" @focus="doingPw = true" @blur="doingPw = false" class="mx-2 min-w-0 text-base"
            :type="showPassword ? 'text' : 'password'" placeholder="密碼" />
          <PhEyeClosed v-if="!showPassword" @click="togglePassword" :size="32" color="#4d4d4d" class="flex-shrink-0" />
          <PhEye v-else @click="togglePassword" :size="32" class="flex-shrink-0" />
        </div>
        <div class="flex w-full border-2 rounded-md shadow-md p-3 items-center">
          <PhEnvelopeSimple :size="32" color="#4d4d4d" class="flex-shrink-0" />
          <input v-model="user.email" class="mx-2 min-w-0 text-base" type="email" placeholder="信箱" />
        </div>

        <div class="flex items-center">
          <PhSquare v-if="!check" @click="checklogin" :size="20" />
          <PhCheckSquare v-else @click="checklogin" :size="20" />
          <span class="ml-2 text-lg">保持登入</span>
        </div>

        <div>

            <button @click="registerHandler(user)"
              class="w-full bg-[#90c700] text-white font-bold text-2xl py-3 sm:py-2 rounded shadow-md" type="button">
              創建
            </button>
        </div>
      </form>
    </div>
    <div class="flex-shrink-0">
      <div class="flex items-center mt-44">
        <div class="flex-1 border-t"></div>
        <span class="mx-3">or</span>
        <div class="flex-1 border-t"></div>
      </div>
      <div>
        <router-link to="/account/login" class="flex text-lg text-[#92c700] justify-center">
          登入
        </router-link>
      </div>
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
import validate from "validate.js";
import { ref } from "vue";

class form {
  account: string = "";
  password: string = "";
  email: string = "";
}
const user = ref<form>(new form());

var constraints = {
  account: {
    length: {
      minimum: 3,
      maximum: 10,
      message: "賬號長度需在3-10個字元間"
    }
  },
  password: {
    format: {
      pattern: /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d).+$/,
      message: "密碼必须包含至少一个大写字母、一个小写字母和一个数字"
    },
    length: {
      minimum: 3,
      maximum: 10,
      message: "密碼長度需在3-10個字元間"
    }
  },
  email: {
    email: {
      message: "郵箱輸入有誤"
    }
    
  }
};

const registerHandler = async (user: form) => {
  const vaildResult = await validate(user, constraints)
  if (vaildResult) {
    alert(vaildResult[0])
    return
  }

  try {
    const json = JSON.stringify(user);
    localStorage.setItem("email", user.email);
    const result = await request.post("/users/register", json);
    router.push("/account/verify");
  } catch (error: any) {
    alert(error.response.data.error);
  }
};

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
