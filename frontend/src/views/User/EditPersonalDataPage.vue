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
        <input class="flex w-full border-2 rounded-md shadow-md p-3" type="text" v-model="user.name" />
      </div>
      <div>
        <label class="flex text-2xl font-bold">帳號</label>
        <input class="flex w-full border-2 rounded-md shadow-md p-3" type="text" v-model="user.account" />
      </div>
      <div>
        <label class="flex text-2xl font-bold">生日</label>
        <Datepicker v-model="user.birthday" :format="formatDate" :enable-time-picker="false"
          class="w-full rounded-md shadow-md" />
      </div>
      <div>
        <label class="flex text-2xl font-bold">性別</label>
        <div class="flex justify-around">
          <div>
            <input v-model="user.female" type="radio" id="female" name="gender" :value=1 />
            <label for="male" class="text-2xl ml-2">男性</label>
          </div>
          <div>
            <input v-model="user.female" type="radio" id="female" name="gender" :value=2 />
            <label for="neutral" class="text-2xl ml-2">中性</label>
          </div>
          <div>
            <input v-model="user.female" type="radio" id="female" name="gender" :value=3 />
            <label for="female" class="text-2xl ml-2">女性</label>
          </div>
        </div>
      </div>
      <div>
        <label class="flex text-2xl font-bold">地區</label>
        <select v-model="user.location" class="flex w-full border-2 rounded-md shadow-md p-3 text-2xl">
          <option :value="item" v-for="item in 19">{{ getTextByLocation(item) }}</option>
        </select>
      </div>
      <div>
        <router-link to="/user/personal-data">
          <button @click="userStore.saveUserHanlder(JSON.stringify(user))" class="w-full bg-[#90c700] text-white font-bold text-2xl py-3 rounded shadow-md" type="button">
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
import { computed, ref } from "vue";

import "@vuepic/vue-datepicker/dist/main.css";
import { getTextByLocation } from "@/utils/getTextByNumber";
import { useUserStore } from "@/store/userStroe";

const formatDate = "yyyy-MM-dd";

const userStore = useUserStore()
const user = computed(() => userStore.user)


</script>
