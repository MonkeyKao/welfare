<template>
  <div class="flex flex-col h-screen w-full justify-between">
    <div class="flex-shrink-0">
      <!-- 固定高度的 header -->
      <img v-show="!doingPw" src="../../images/login.jpg" />
      <img v-show="doingPw" src="../../images/password.jpg" />
    </div>

    <div class="flex-grow">
      <!-- Flexbox 居中表单 -->
      <form class="flex flex-col px-10 mt-3 space-y-5">
        <!-- 限制最大宽度 -->
        <label class="flex text-H1 font-bold justify-center">Register</label>

        <div class="flex border-2 rounded-md shadow-md y-2 p-1 px-3 items-center">
          <PhUser :size="32" color="#4d4d4d" class="flex-shrink-0" />
          <input v-model="user.account" class=" border-none outline-none mx-2 min-w-0 text-base" type="text" placeholder="帳號" />
        </div>
        <!-- 密碼輸入框 -->
        <div class="flex border-2 rounded-md shadow-md p-1 px-3 justify-between ">
            <div class="flex">
              <PhLockKey :size="32" color="#4d4d4d" class="flex-shrink-0" />
              <input
                v-model="user.password"
                @focus="doingPw = true"
                @blur="doingPw = false"
                :type="showPassword ? 'text' : 'password'"
                placeholder="密碼"
                class="border-none resize outline-none p-1 w-full px-2"
              >
            </div>
            <PhEyeClosed v-if="!showPassword" @click="togglePassword" :size="32" color="#4d4d4d" class="flex-shrink-0" />
            <PhEye v-else @click="togglePassword" :size="32" class="flex-shrink-0" />
          </div>     

        <div class="flex border-2 rounded-md shadow-md y-2 p-1 px-3 items-center">
          <PhEnvelopeSimple :size="32" color="#4d4d4d" class="flex-shrink-0" />
          <input v-model="user.email" class=" border-none outline-none mx-2  text-base" type="email" placeholder="信箱" />
        </div>

        <div class="flex items-center px-2">
          <PhSquare v-if="!check" @click="checklogin" :size="20" />
          <PhCheckSquare v-else @click="checklogin" :size="20" />
          <span class="ml-2 text-lg">保持登入</span>
        </div>

        <div  class="flex flex-col">
          <button @click="registerHandler(user)" class="bg-[#90c700] text-white text-H3 py-2 rounded shadow-md" type="button">創建</button>
        </div>
      </form>
    </div>

    <div class="flex-shrink-0">
      <div class="flex flex-col justify-center">
        <div class="flex items-center">
          <div class="flex-1 border-t"></div>
          <span class="mx-3">or</span>
          <div class="flex-1 border-t"></div>
        </div>
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
      maximum: 12,
      message: "帳號長度需在3-12個字元間"
    }
  },
  password: {
    format: {
      pattern: /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d).+$/,
      message: "密碼必須包含至少一個大寫字母、一個小寫字母和一個數字"
    },
    length: {
      minimum: 3,
      maximum: 15,
      message: "密碼長度需在3-15個字元間"
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
