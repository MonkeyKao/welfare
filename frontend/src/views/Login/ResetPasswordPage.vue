<template>
  <!-- 固定定位的 header，確保圖片維持在最上方 -->
  <header>
    <img src="/forgetpassword.jpg" class="w-full" />
  </header>
  <body>
    <router-link to="/forgot-password" class="forgot-password-link"
      ><PhArrowUUpLeft :size="32" color="#4d4d4d" class="ml-5"
    /></router-link>

    <form class="flex flex-col w-full h-full px-10 gap-5">
      <label class="text-3xl font-bold">密碼重設</label>
      <!-- 密碼輸入框 -->
      <div class="flex w-full border-2 rounded-md shadow-md p-3 items-center">
        <PhLockKey :size="32" color="#4d4d4d" class="flex-shrink-0" />
        <input
          class="mx-2 min-w-0"
          :type="showPassword1 ? 'text' : 'password'"
          placeholder="密碼"
          v-model="password1"
        />
        <PhEyeClosed
          v-if="!showPassword1"
          @click="togglePassword1"
          :size="32"
          color="#4d4d4d"
          class="flex-shrink-0"
        />
        <PhEye
          v-else
          @click="togglePassword1"
          :size="32"
          class="flex-shrink-0"
        />
      </div>

      <!-- 密碼輸入框 -->
      <div class="flex">
        <div
          class="flex w-full border-2 rounded-md shadow-md p-3 items-center flex-shrink-0"
        >
          <PhLockKey :size="32" color="#4d4d4d" class="flex-shrink-0" />
          <input
            class="mx-2 min-w-0"
            :type="showPassword2 ? 'text' : 'password'"
            placeholder="再次輸入密碼"
            v-model="password2"
          />
          <PhEyeClosed
            v-if="!showPassword2"
            @click="togglePassword2"
            :size="32"
            color="#4d4d4d"
            class="flex-shrink-0"
          />
          <PhEye
            v-else
            @click="togglePassword2"
            :size="32"
            class="flex-shrink-0"
          />
        </div>
        <div v-if="password1 && password2">
          <PhCheck v-if="password1 === password2" :size="32" color="#92c700" />
          <PhExclamationMark v-else :size="32" color="#d06262" />
        </div>
      </div>
      <div>
        <router-link to="/reset-password" class="forgot-password-link"
          ><button
            class="w-full bg-[#90c700] text-white font-bold text-2xl py-3 rounded shadow-md"
            type="button"
          >
            確定
          </button></router-link
        >
      </div>
    </form>
  </body>
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
import { ref } from "vue";

// 密碼儲存
const password1 = ref("");
const password2 = ref("");

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
</script>
