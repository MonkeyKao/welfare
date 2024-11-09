<template>
  <header class="flex items-center gap-3 p-4 bg-[#92c700] basis-1">
    <router-link to="/user" class="forgot-password-link">
      <PhArrowUUpLeft :size="32" color="#000" class="ml-5" />
    </router-link>
    <label class="text-3xl font-bold">個人資訊</label>
  </header>

  <body>
    <div class="flex flex-col items-center gap-2 mt-5">
      <Avatar />
      <label class="text-sm text-[#7F8689]">{{
        user?.name ? user.name : "未登錄"
      }}</label>
      <form class="flex flex-col w-full h-full p-10 gap-5">
        <div class="flex rounded-md">
          <label class="flex text-2xl font-bold">姓名:</label>
          <label class="text-2xl font-bold">{{ user?.name }}</label>
        </div>
        <div class="flex rounded-md">
          <label class="flex text-2xl font-bold">帳號:</label>
          <label class="text-2xl font-bold">{{ user?.account }}</label>
        </div>
        <div class="flex rounded-md">
          <label class="flex text-2xl font-bold">生日:</label>
          <label class="text-2xl font-bold">{{ user?.birthday }}</label>
        </div>
        <div class="flex rounded-md">
          <label class="flex text-2xl font-bold">性別:</label>
          <label class="text-2xl font-bold">{{ getTextByGender(user.female) }}</label>
        </div>
        <div class="flex rounded-md gap-2">
          <label class="flex text-2xl font-bold">地區:</label>
          <label class="text-2xl font-bold">{{ getTextByLocation(user.location) }}</label>
        </div>
        <div>
          <router-link to="/user/edit-personal-data"><button
              class="w-full bg-[#90c700] text-white font-bold text-2xl py-3 rounded shadow-md" type="button">
              編輯
            </button></router-link>
        </div>
      </form>
    </div>
    {{user}}
  </body>
</template>
<script setup lang="ts">
import request from "@/axios";
import Avatar from "@/components/Avatar.vue";
import User from "@/model/user";
import { useUserStore } from "@/store/userStroe";
import { getTextByGender, getTextByLocation } from "@/utils/getTextByNumber";
import { PhArrowUUpLeft } from "@phosphor-icons/vue";
import { computed, onMounted, ref } from "vue";
const userStore = useUserStore()
const user = computed(() => userStore.user)

</script>
