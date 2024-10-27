<template>
  <div class="w-full divide-y container" >

    <header class="flex flex-col items-center my-3">
      <Avatar />
      <label class="text-3xl font-bold">{{ user?.name ? user.name : '未登錄' }}</label>
    </header>

    <div class=" ">
      <RouterLink to="/personal-data" class="w-full">
        <div class="flex items-center mx-10">
          <PhUserCircle :size="32" color="#4d4d4d" />
          <label class="text-lg font-bold flex ml-5">個人資料</label>
        </div>
      </RouterLink>

      <RouterLink to="/link-account" class="w-full">
        <div class="flex items-center mx-10">
          <PhLink :size="32" color="#4d4d4d" />
          <label class="text-lg font-bold flex ml-5">連結帳戶</label>
        </div>
      </RouterLink>

      <RouterLink to="/family" class="w-full">
        <div class="flex items-center mx-10">
          <PhUsersThree :size="32" color="#4d4d4d" />
          <label class="text-lg font-bold flex ml-5">家庭</label>
        </div>
      </RouterLink>

      <RouterLink to="/settings" class="w-full">
        <div class="flex items-center mx-10">
          <PhGear :size="32" color="#4d4d4d" />
          <label class="text-lg font-bold flex ml-5">設定</label>
        </div>
      </RouterLink>

      <RouterLink to="/qa" class="w-full">
        <div class="flex items-center mx-10">
          <PhQuestion :size="32" color="#4d4d4d" />
          <label class="text-lg font-bold flex ml-5">常見問題</label>
        </div>
      </RouterLink>

      <div class="mx-10">
        <router-link v-if="user" to="/login" class="forgot-password-link"><button
            class="w-full bg-[#90c700] text-white font-bold text-2xl py-3 rounded shadow-md" type="button"
            @click="logoutHandler">
            登出
          </button></router-link>
        <router-link v-else to="/login" class="forgot-password-link"><button
            class="w-full bg-[#90c700] text-white font-bold text-2xl py-3 rounded shadow-md" type="button">
            登入
          </button></router-link>

      </div>


    </div>

  </div>



</template>

<script setup lang="ts">
import request from "@/axios";
import Avatar from "@/components/Avatar.vue";
import User from "@/model/user";
import {
  PhGear,
  PhLink,
  PhQuestion,
  PhUserCircle,
  PhUsersThree,
} from "@phosphor-icons/vue";
import { onMounted, ref } from "vue";

const logoutHandler = () => {
  localStorage.removeItem("token")
}

const user = ref<User>()

onMounted(async () => {
  const result = await request.get("users");
  user.value = result.data

})
</script>

<style scoped>
/* 可選樣式 */
</style>
