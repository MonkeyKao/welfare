<template>
  <!-- 固定定位的 header，確保圖片維持在最上方 -->
  <div>
    <img src="/forgetpassword.jpg" class="w-full" />
  </div>

  <body>
    <router-link to="/account/register" class="forgot-password-link">
      <PhArrowUUpLeft :size="32" color="#4d4d4d" class="ml-5" />
    </router-link>
    <div class="flex flex-col w-full h-full px-10 gap-5">
      <label class="text-xl font-bold">驗證碼已寄至</label>
      <label class="text-xl font-bold">{{ maskEmail(email) }}</label>

      <div class="flex w-full border-2 rounded-md shadow-md p-3 items-center">
        <PhShieldCheck :size="32" class="flex-shrink-0" />
        <input class="mx-2 min-w-0" type="password" v-model="verificationCode" placeholder="驗證碼" />
        <button :disabled="isDisabled" @click="startCountdown"
          class="flex text-lg text-[#92c700] border-l-2 border-[#92c700] whitespace-nowrap px-2">
          {{ buttonText }}
        </button>
      </div>

      <button class="w-full bg-[#90c700] text-white font-bold text-2xl py-3 rounded shadow-md"
        @click="sendVerificationCode(verificationCode)">
        繼續
      </button>
    </div>
  </body>
</template>
<script setup lang="ts">
import request from "@/axios";
import router from "@/router";
import { PhArrowUUpLeft, PhShieldCheck } from "@phosphor-icons/vue";
import validate from "validate.js";
import { inject, ref } from "vue";

const showMsg: Function = inject("showMsg")!
const verificationCode = ref<string>("");
const email: string = localStorage.getItem("email")!;

const countdown = ref<number>(60);
const isDisabled = ref<boolean>(false);
const buttonText = ref<string>("重新發送");
let countdownInterval: number;

const startCountdown = async () => {
  // 禁用按鈕並開始倒計時
  try {
    const result = await request.get("/users/getVerifyEmail", { params: { email: email } })
    showMsg("已發送驗證碼")
  } catch (err: any) {
  }

  isDisabled.value = true;
  buttonText.value = `請稍等 ${countdown.value} 秒...`;
  countdownInterval = setInterval(() => {
    countdown.value--;
    buttonText.value = `請稍等 ${countdown.value} 秒...`;

    if (countdown.value <= 0) {
      clearInterval(countdownInterval!);
      resetButton();
    }
  }, 1000);
}

const resetButton = () => {
  // 重置按鈕狀態
  isDisabled.value = false;
  countdown.value = 60;
  buttonText.value = "重新發送";
}

const maskEmail = (email: string): string => {
  const [account, domain] = email.split("@");
  if (account.length <= 2) {
    return `${account}@${domain}`; // 如果賬號少於或等於2位，直接返回
  }
  const maskedAccount = account.slice(0, 2) + "*".repeat(account.length - 2);
  return `${maskedAccount}@${domain}`;
}

const sendVerificationCode = async (verificationCode: string) => {

  const vaildResult = await validate.single(verificationCode,{format: {pattern:"^\\d{6}$",message:"驗證碼需為六位數字"}})
  if (vaildResult) {
    alert(vaildResult[0])
    return
  }

  try {
    const result = await request.post("/users/verify", {
      email: email,
      code: verificationCode,
    });
    localStorage.setItem("token", result.data.msg);
    router.push("account/create-profile");
  } catch (err: any) {
    console.log(err);

    alert(err.response.data.error);
  }
};
</script>
