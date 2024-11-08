<template>
  <header class="flex items-center gap-3 p-4 bg-[#92c700] basis-1">
    <router-link to="/user" class="forgot-password-link">
      <PhArrowUUpLeft :size="32" color="#000" class="ml-5" />
    </router-link>
    <label class="text-3xl font-bold">編輯個人資訊</label>
  </header>
  <div class="flex flex-col items-center gap-2 mt-5">
    <Avatar />
    <form class="flex flex-col w-full h-full p-10 gap-5">
      <div>
        <label class="flex text-2xl font-bold">姓名</label>
        <input
          class="flex w-full border-2 rounded-md shadow-md p-3"
          type="text"
        />
      </div>
      <div>
        <label class="flex text-2xl font-bold">帳號</label>
        <input
          class="flex w-full border-2 rounded-md shadow-md p-3"
          type="text"
        />
      </div>
      <div>
        <label class="flex text-2xl font-bold">生日</label>
        <Datepicker
          v-model="birthday"
          :format="formatDate"
          :enable-time-picker="false"
          class="w-full rounded-md shadow-md"
        />
      </div>
      <div>
        <label class="flex text-2xl font-bold">性別</label>
        <div class="flex justify-around">
          <div>
            <input type="radio" id="male" name="gender" class="mr-2" />
            <label for="male" class="text-2xl">男性</label>
          </div>
          <div>
            <input type="radio" id="female" name="gender" class="mr-2 ml-4" />
            <label for="neutral" class="text-2xl">中性</label>
          </div>
          <div>
            <input type="radio" id="female" name="gender" class="mr-2 ml-4" />
            <label for="female" class="text-2xl">女性</label>
          </div>
        </div>
      </div>
      <div>
        <label class="flex text-2xl font-bold">地區</label>
        <select class="flex w-full border-2 rounded-md shadow-md p-3 text-2xl">
          <option>台北市</option>
          <option>新北市</option>
          <option>基隆市</option>
          <option>新竹市</option>
          <option>新竹縣</option>
          <option>苗栗縣</option>
          <option>台中市</option>
          <option>彰化縣</option>
          <option>雲林縣</option>
          <option>嘉義市</option>
          <option>嘉義縣</option>
          <option>台南市</option>
          <option>高雄市</option>
          <option>屏東縣</option>
          <option>南投縣</option>
          <option>宜蘭縣</option>
          <option>花蓮縣</option>
          <option>台東縣</option>
        </select>
      </div>
      <div>
        <router-link to="/personal-data">
          <button
            class="w-full bg-[#90c700] text-white font-bold text-2xl py-3 rounded shadow-md"
            type="button"
          >
            確認
          </button>
        </router-link>
      </div>
    </form>
  </div>
</template>

<script setup lang="ts">
import Avatar from "@/components/Avatar.vue";
import { PhArrowUUpLeft } from "@phosphor-icons/vue";
import Datepicker from "@vuepic/vue-datepicker";
import { ref } from "vue";

import "@vuepic/vue-datepicker/dist/main.css";

const birthday = ref<Date | undefined>(undefined);
const formatDate = "yyyy-MM-dd";
</script>
