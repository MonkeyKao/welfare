<template>
  <!-- 固定定位的 header，確保圖片維持在最上方 -->
  <header>
    <img src="/forgetpassword.jpg" class="w-full" />
  </header>

  <body>
    <router-link to="/login">
      <PhArrowUUpLeft :size="32" color="#4d4d4d" class="ml-5" />
    </router-link>
    <form @submit.prevent class="flex flex-col w-full h-full px-10 gap-5">
      <label class="text-3xl font-bold">忘記密碼</label>
      <div class="flex w-full border-2 rounded-md shadow-md p-3 items-center">
        <PhEnvelopeSimple :size="32" color="#4d4d4d" class="flex-shrink-0" />
        <input v-model="email" type="email" placeholder="信箱" class="ml-2" :disabled="emailInput" />
      </div>

      <div class="flex w-full border-2 rounded-md shadow-md p-3 items-center">
        <PhShieldCheck :size="32" color="#4d4d4d" class="flex-shrink-0" />
        <input class="mx-2 min-w-0" type="password" v-model="verificationCode" placeholder="驗證碼" />
        <button @click="sendVerificationCode"
          class="flex text-lg text-[#92c700] border-l-2 border-[#92c700] whitespace-nowrap px-2">
          發送驗證碼
        </button>
      </div>

      <div>
        <button class="w-full bg-[#90c700] text-white font-bold text-2xl py-3 rounded shadow-md" @click="verifyCode">
          下一步
        </button>
      </div>
    </form>

  </body>
</template>
<script setup lang="ts">
import request from "@/axios";
import router from "@/router";
import {
  PhArrowUUpLeft,
  PhEnvelopeSimple,
  PhShieldCheck,
} from "@phosphor-icons/vue";
import validate from "validate.js";
import { ref } from "vue";
const email = ref("");
const verificationCode = ref("");
const emailInput = ref(false)

const sendVerificationCode = async () => {
  const vaildResult =  validate.async({email:email.value}, {
    email: {
      email: {
        message: "郵箱輸入有誤"
      }
    }
  }).catch((err) => {
    console.log(err);
    
    return
  })
  

  try {
    const result = await request.get("/users/getVerifyEmail", { params: { "email": email.value } })
    emailInput.value = true
  } catch (error) {
    console.log("發送驗證碼時發生錯誤:", error);
  }
};

const verifyCode = async () => {
  const form = {
    email: email.value,
    code: verificationCode.value,
  }

  const vaildResult = await validate.async(form, {
    email: {
      email: {
        message: "郵箱輸入有誤"
      }
    },
    code: {
      format: { pattern: "^\\d{6}$", message: "驗證碼需為六位數字" }
    }
  })

  if (vaildResult) {
    alert(vaildResult[0])
    return
  }

  try {
    const result = await request.post("/users/verify", form)
    router.push("/account/reset-password")
    localStorage.setItem('token', result.data.msg)
    console.log("成功");

  } catch (error: any) {
    alert(error.response.data.error)
  }
};
</script>
