<template>
  <!-- 固定定位的 header，確保圖片維持在最上方 -->
  <header>
    <img src="/forgetpassword.jpg" class="w-full" />
  </header>
  <body>
    <router-link to="/login"
      ><PhArrowUUpLeft :size="32" color="#4d4d4d" class="ml-5"
    /></router-link>
    <form class="flex flex-col w-full h-full px-10 gap-5">
      <label class="text-3xl font-bold">忘記密碼</label>
      <div class="flex w-full border-2 rounded-md shadow-md p-3 items-center">
        <PhEnvelopeSimple :size="32" color="#4d4d4d" class="flex-shrink-0" />
        <input type="email" placeholder="信箱" class="ml-2" />
      </div>

      <div class="flex w-full border-2 rounded-md shadow-md p-3 items-center">
        <PhShieldCheck :size="32" color="#4d4d4d" class="flex-shrink-0" />
        <input
          class="mx-2 min-w-0"
          type="password"
          v-model="verificationCode"
          placeholder="驗證碼"
        />
        <button
          @click="sendVerificationCode"
          class="flex text-lg text-[#92c700] border-l-2 border-[#92c700] whitespace-nowrap px-2"
        >
          發送驗證碼
        </button>
      </div>

      <div>
        <router-link to="/reset-password" class="forgot-password-link"
          ><button
            class="w-full bg-[#90c700] text-white font-bold text-2xl py-3 rounded shadow-md"
            type="button"
            @click="verifyCode"
          >
            下一步
          </button></router-link
        >
      </div>
    </form>
  </body>
</template>
<script setup lang="ts">
import {
  PhArrowUUpLeft,
  PhEnvelopeSimple,
  PhShieldCheck,
} from "@phosphor-icons/vue";
import { ref } from "vue";
const contact = ref("");
const verificationCode = ref("");

const sendVerificationCode = async () => {
  try {
    const response = await fetch("/api/send-verification-code", {
      method: "POST",
      body: JSON.stringify({ contact: contact.value }),
      headers: {
        "Content-Type": "application/json",
      },
    });
    const result = await response.json();
    console.log(result.message);
  } catch (error) {
    console.error("發送驗證碼時發生錯誤:", error);
  }
};

const verifyCode = async () => {
  try {
    const response = await fetch("/api/verify-code", {
      method: "POST",
      body: JSON.stringify({
        contact: contact.value,
        code: verificationCode.value,
      }),
      headers: {
        "Content-Type": "application/json",
      },
    });
    const result = await response.json();
    console.log(result.message);
  } catch (error) {
    console.error("驗證碼驗證時發生錯誤:", error);
  }
};
</script>
