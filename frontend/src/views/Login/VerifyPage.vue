<template>

  <div class="flex-shrink-0">
      <img src="/forgetpassword.jpg" />
  </div>

  
  <form @submit.prevent="sendVerificationCode(verificationCode)" class="flex flex-col space-y-10 px-10 mt-3">

    <div class="flex items-center w-full mt-7 relative">
      <router-link to="/account/register" class=""><PhArrowUUpLeft :size="36" color="#4d4d4d" weight="bold" /></router-link>
      <label class="absolute left-1/2 transform -translate-x-1/2 text-H1 font-bold">寄驗證碼</label>
    </div>


      <label class="text-xl font-bold">驗證碼已寄至 {{ maskEmail(email) }}</label>

      <div class="flex border-2  rounded-md shadow-md p-1 px-3 justify-between">
        
        <PhShieldCheck :size="28" class="flex-shrink-0" />
        <input class="border-none resize outline-none p-1 mx-1 w-full" type="password" v-model="verificationCode" placeholder="驗證碼" />
        <div class="flex-shrink-0 flex items-center justify-center border-l-2 border-[#92c700] px-2">
          <button type="button" :disabled="isDisabled" @click="startCountdown" class="text-[#92c700]">
            {{ buttonText }}
          </button>          
        </div>

      </div>

      <button type="submit" class=" bg-[#90c700] rounded shadow-md items-center text-white font-bold text-H3 p-2">繼續</button>
  </form>

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
const buttonText = ref<string>("重設驗證碼");
let countdownInterval: number;

const startCountdown = async () => {
  isDisabled.value = true;
  // 禁用按鈕並開始倒計時
  try {
    const result = await request.get("/verify", { params: { email: email } })
    showMsg("已發送驗證碼")
  } catch (err: any) {
  }

  
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
    showMsg(vaildResult[0])
    return
  }

  try {
    const result = await request.post("/verify", {
      email: email,
      code: verificationCode,
    });
    localStorage.setItem("token", result.data.msg);
    localStorage.removeItem("email")
    router.push("/account/create-profile");
  } catch (err: any) {
    showMsg(err.response.data.error)
  }
};
</script>
