<template>
  <div class="flex flex-col h-screen w-full justify-between">
    <div class="flex-shrink-0">
      <img v-show="!doingPw" src="../../images/login.jpg" />
      <img v-show="doingPw" src="../../images/password.jpg" />
    </div>

    <!-- Flexbox 居中表單 -->
    <div class="flex-grow">
      <form class="flex flex-col px-10 mt-3 space-y-3">
        <!-- 標題 -->
        <label class="flex text-H1 font-bold justify-center">Login</label>

        <!-- 帳號輸入框 -->
        <div class="flex border-2 rounded-md shadow-md y-2 p-1 items-center">
          <PhUser :size="32" color="#4d4d4d" class="flex-shrink-0" />
          <input v-model="user.account" class="border-none resize outline-none p-1 w-full" type="text" placeholder="帳號" />
        </div>

        <div class="">
          <router-link to="/account/forgot-password" class="flex text-H3 text-[#92c700] justify-end">忘記密碼？</router-link>
          <!-- 密碼輸入框 -->
          <div class="flex border-2 rounded-md shadow-md p-1 justify-between">
            <div class="flex">
              <PhLockKey :size="32" color="#4d4d4d" class="flex-shrink-0" />
              <input
                v-model="user.password"
                @focus="doingPw = true"
                @blur="doingPw = false"
                :type="showPassword ? 'text' : 'password'"
                placeholder="密碼"
                class="border-none resize outline-none p-1 w-full"
              >
            </div>
            <PhEyeClosed v-if="!showPassword" @click="togglePassword" :size="32" color="#4d4d4d" class="flex-shrink-0" />
            <PhEye v-else @click="togglePassword" :size="32" class="flex-shrink-0" />
          </div>          
        </div>

        <div class=" flex flex-col pt-5">
          <button @click="loginHandler(user)" class=" bg-[#90c700] text-white text-H3 py-2 rounded shadow-md">登入</button>
          <router-link to="/home" class="flex text-H3 text-[#92c700] justify-center pt-2">訪客登入</router-link>          
        </div>

        <!-- 按鈕 -->
      </form>
      
    </div>

    <!-- 底部連結區域 -->
    <div class="flex-shrink-0">
      <div class="flex flex-col justify-center">
        <div>
          <label class="flex text-sm justify-center">其他登入方式</label>
          <div class="flex justify-center space-x-4 mt-4">
            <button class="w-12 h-12 rounded-full flex justify-center items-center shadow bg-[#92c700]">
              <PhGoogleLogo :size="24" color="#fff" />
            </button>
            <button class="w-12 h-12 rounded-full flex justify-center items-center shadow bg-[#92c700]">
              <PhFacebookLogo :size="24" color="#fff" />
            </button>
          </div>
        </div>
        <div class="flex items-center">
          <div class="flex-1 border-t"></div>
          <span class="mx-3">or</span>
          <div class="flex-1 border-t"></div>
        </div>
        <router-link to="/account/register" class="flex text-lg text-[#92c700] justify-center pb-2">
          創建帳號
        </router-link>
        
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import request from "@/axios";
import router from "@/router";
import { useUserStore } from "@/store/userStroe";
import {
  PhEye,
  PhEyeClosed,
  PhFacebookLogo,
  PhGoogleLogo,
  PhLockKey,
  PhUser,
} from "@phosphor-icons/vue";
import validate from "validate.js";
import { ref } from "vue";

const userStroe = useUserStore();
class form {
  account: string = "";
  password: string = "";
}
const user = ref<form>(new form());

var constraints = {
  account: {
    length: {
      minimum: 3,
      maximum: 12,
      message: "賬號長度需在3-12個字元間"
    }
  },
  password: {
    format: {
      pattern: /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d).+$/,
      message: "密碼必须包含至少一个大写字母、一个小写字母和一个数字"
    },
    length: {
      minimum: 3,
      maximum: 15,
      message: "密碼長度需在3-15個字元間"
    }
  }
};



const loginHandler = async (user: form) => {
  const vaildResult = await validate(user, constraints)
  if (vaildResult) {
    alert(vaildResult[0])
    return
  }

  try {
    const result = await request.post("users/doLogin", JSON.stringify(user));
    localStorage.setItem(result.data.tokenName, result.data.tokenValue);
    await userStroe.fetchUser();
    router.push("/home");
  } catch (err: any) {
    alert(err.response.data.error);
    return;
  }
};

function loginWithFacebook() {
  // 呼叫 Facebook 登入 API 的程式邏輯
}
function loginWithGoogle() {
  // 呼叫 Google 登入 API 的程式邏輯
}

//圖片更動
const doingPw = ref(false); // 用于指示是否处于密码输入状态

//點眼睛密碼明文顯示
const showPassword = ref(false);
const togglePassword = () => {
  showPassword.value = !showPassword.value; // 切换 showPassword 的值
};


</script>
