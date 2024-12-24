<template>
  <div class="flex flex-col h-screen w-full justify-between">
    <div class="flex-shrink-0">
      <img src="/forgetpassword.jpg" />
    </div>

    <div class="flex-grow">

      <form @submit.prevent class="flex flex-col space-y-10 px-6 mt-3">

        <div class="flex items-center w-full mt-7 relative">
          <router-link to="/account/login" class="flex-shrink-0">
            <PhArrowUUpLeft :size="36" weight="bold" color="#4d4d4d" />
          </router-link>
          <label class="absolute left-1/2 transform -translate-x-1/2 text-H1 font-bold">忘記密碼</label>
        </div>



        <div class="flex border-2 rounded-md shadow-md y-5 p-1 items-center px-3 ">
          <PhEnvelopeSimple :size="28" color="#4d4d4d" class="flex-shrink-0" />
          <input v-model="email" type="email" placeholder="信箱" class="border-none resize outline-none p-1 w-full"
            :disabled="emailInput" />
        </div>

        <div class="flex border-2  rounded-md shadow-md p-1 px-3 justify-between">
          <div class="flex">
            <PhShieldCheck :size="28" color="#4d4d4d" class="flex-shrink-0" />
            <input class="border-none resize outline-none p-1 mx-1 w-full" type="text" v-model="verificationCode"
              placeholder="驗證碼" />
          </div>
          <div class="flex-shrink-0 flex items-center justify-center border-l-2 border-[#92c700] px-2">
            <button @click="sendVerificationCode" class="text-[#92c700] ">重設驗證碼</button>
          </div>
        </div>

        <button class=" bg-[#90c700] rounded shadow-md items-center text-white font-bold text-H3 p-2"
          @click="verifyCode">下一步</button>
      </form>
    </div>
  </div>
</template>
<script setup lang="ts">
import request from "@/axios";
import router from "@/router";
import { AlertColor, showMsgFunction } from "@/type/ShowMsg";
import { PhArrowUUpLeft, PhEnvelopeSimple, PhShieldCheck } from "@phosphor-icons/vue";
import { isAxiosError } from "axios";
import validate from "validate.js";
import { inject, ref } from "vue";
const showMsg: showMsgFunction = inject("showMsg")!;
const email = ref("");
const verificationCode = ref("");
const emailInput = ref(false)

const sendVerificationCode = async () => {
  const vaildResult = await validate({ email: email.value }, {
    email: {
      email: {
        message: "郵箱輸入有誤"
      }
    }
  })

  if (vaildResult) {
    showMsg(vaildResult[0], AlertColor.waring)
    return
  }

  try {
    await request.get("/users/getVerifyEmail", { params: { "email": email.value } })
    emailInput.value = true
  } catch (err: any) {
    if (isAxiosError(err)) {
      showMsg(err.response?.data.error, AlertColor.error)
    }
  }
};

const verifyCode = async () => {
  const form = {
    email: email.value,
    code: verificationCode.value,
  }

  const vaildResult = await validate(form, {
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
    showMsg(vaildResult[0], AlertColor.waring)
    return
  }

  try {
    const result = await request.post("/users/verify", form)
    router.push("/account/reset-password")
    localStorage.setItem('token', result.data.msg)

  } catch (err: any) {
    if(isAxiosError(err)) {
      showMsg(err.response?.data.error, AlertColor.error)
    }
  }
};
</script>
