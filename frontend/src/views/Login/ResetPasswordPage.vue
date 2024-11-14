<template>
  <!-- 固定定位的 header，確保圖片維持在最上方 -->
  <div class="flex flex-col h-screen w-full ">
    <div class="flex-shrink-0">
      <img src="../../images/forgetpassword.jpg"/>
    </div>  



    <form class="flex flex-col space-y-10 px-10 mt-3">

      <div class="flex items-center w-full mt-7 relative">
          <router-link to="/forgot-password" class="flex-shrink-0"><PhArrowUUpLeft :size="32" weight="bold" color="#4d4d4d"  /></router-link>
          <label class="absolute left-1/2 transform -translate-x-1/2 text-H1 font-bold">密碼重設</label>  
      </div>      
      <!-- 密碼輸入框 -->
      <div class="flex border-2 rounded-md shadow-md p-1 justify-between px-3">
        <div class="flex">
          <PhLockKey :size="32" color="#4d4d4d" class="flex-shrink-0" />
          <input class="border-none resize outline-none p-1 w-full" :type="showPassword1 ? 'text' : 'password'" placeholder="密碼" v-model="password" />          
        </div>
        <PhEyeClosed v-if="!showPassword1" @click="togglePassword1" :size="32" color="#4d4d4d" class="flex-shrink-0" />
        <PhEye v-else @click="togglePassword1" :size="32" class="flex-shrink-0" />
      </div>

      <!-- 密碼輸入框 -->
      <div class="">
        <div :class="[
    'flex border-2 rounded-md shadow-md p-1 justify-between px-3',
    password && confirmPassword
      ? (password === confirmPassword ? 'border-[#92c700]' : 'border-[#d06262]')
      : 'border-2'
  ]">
          <div class="flex">
            <PhLockKey :size="32" color="#4d4d4d" class="flex-shrink-0" />
            <input class="border-none resize outline-none p-1 w-full" :type="showPassword2 ? 'text' : 'password'" placeholder="再次輸入密碼"v-model="confirmPassword" />            
          </div>

          <PhEyeClosed v-if="!showPassword2" @click="togglePassword2" :size="32" color="#4d4d4d"class="flex-shrink-0" />
          <PhEye v-else @click="togglePassword2" :size="32" class="flex-shrink-0" />
        </div>
        <!--
        <div v-if="password && confirmPassword">
          <PhCheck v-if="password === confirmPassword" :size="32" color="#92c700" />
          <PhExclamationMark v-else :size="32" color="#d06262" />
        </div>
        -->
      </div>
        <button @click="resetPasswordHandler(password, confirmPassword)" class="bg-[#90c700] rounded shadow-md items-center text-white font-bold text-H3 p-2" type="button">確定</button>
    </form>
  </div>

</template>
<script setup lang="ts">
import {
  PhArrowUUpLeft,
  PhCheck,
  PhExclamationMark,
  PhEye,
  PhEyeClosed,
  PhLockKey,
} from "@phosphor-icons/vue";
import { validate } from "validate.js";
import { ref } from "vue";

// 密碼儲存
const password = ref("");
const confirmPassword = ref("");

// 顯示密碼的狀態
const showPassword1 = ref(false);
const showPassword2 = ref(false);

// 切換密碼顯示狀態
const togglePassword1 = () => {
  showPassword1.value = !showPassword1.value;
};

const togglePassword2 = () => {
  showPassword2.value = !showPassword2.value;
};

const resetPasswordHandler = async (password: string, confirmPassword: string) => {
  const vaildResult = await validate({ password: password, confirmPassword: confirmPassword }, {
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
    }
  })
  if(vaildResult) {
    alert(vaildResult[0])
    return
  }else if(password !== confirmPassword){
    alert("兩次輸入的密碼不相同")
    return
  }

  console.log("GO");
  
}
</script>
