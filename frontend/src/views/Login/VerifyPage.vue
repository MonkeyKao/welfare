<template>
  <!-- 固定定位的 header，確保圖片維持在最上方 -->
  <header>
    <img src="/forgetpassword.jpg" class="w-full" />
  </header>

  <body>
    <router-link to="/" class="forgot-password-link"
      ><PhArrowUUpLeft :size="32" color="#4d4d4d" class="ml-5"
    /></router-link>
    <form class="flex flex-col w-full h-full px-10 gap-5">
      <label class="text-3xl font-bold">驗證碼已寄至</label>
      <label class="text-3xl font-bold">email</label>

      <div class="flex w-full border-2 rounded-md shadow-md p-3 items-center">
        <PhShieldCheck :size="32" class="flex-shrink-0" />
        <input
          class="mx-2 min-w-0"
          type="password"
          v-model="verificationCode"
          placeholder="驗證碼"
        />
        <button
          class="flex text-lg text-[#92c700] border-l-2 border-[#92c700] whitespace-nowrap px-2"
        >
          重新發送
        </button>
      </div>

      <button
        class="w-full bg-[#90c700] text-white font-bold text-2xl py-3 rounded shadow-md"
        type="button"
        @click="sendVerificationCode(verificationCode)"
      >
        繼續
      </button>
    </form>
  </body>
</template>
<script setup lang="ts">
import request from "@/axios";
import router from "@/router";
import { PhArrowUUpLeft, PhShieldCheck } from "@phosphor-icons/vue";
import { ref } from "vue";
const verificationCode = ref<string>("");

const sendVerificationCode = async (verificationCode: string) => {
  try {
    const result = await request.post("/users/verify",{"email":localStorage.getItem("email"),"code": verificationCode});
    localStorage.setItem("token",result.data.msg)
    router.push("account/create-profile")
  } catch (err:any) {
    console.log(err);
    
    alert(err.response.data.error)
  }

};

</script>
