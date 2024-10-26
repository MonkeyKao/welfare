<template>
  <header class="top-0 left-0 right-0 flex items-center p-5 pt-14 bg-[#92c700]">
    <router-link to="/user" class="forgot-password-link">
      <PhArrowUUpLeft :size="32" color="#000" class="ml-5" />
    </router-link>
    <label class="ml-5 text-3xl font-bold">個人資訊</label>
  </header>
  <body>
    <div class="flex-col items-center justify-center">
      <Avatar />
      <form class="flex flex-col justify-center w-full max-w-sm">
        <div class="w-full">
          <div class="w-full mt-5">
            <label class="text-left text-xl mb-2">姓名:</label>
            <label
              class="w-full flex focus:outline-none border-2 rounded-md px-3 shadow-md"
            />
          </div>
          <div class="mt-5">
            <label class="block text-left text-xl mb-2">生日:</label>
            <label
              class="flex w-full focus:outline-none border-2 rounded-md px-3 shadow-md"
            />
          </div>
          <div class="mt-5">
            <label class="block text-left text-xl mb-2">性別:</label>
            <label
              class="flex w-full focus:outline-none border-2 rounded-md px-3 shadow-md"
            />
          </div>

          <div class="mt-5">
            <label class="block text-left text-xl mb-2">地區:</label>
            <label
              class="flex w-full focus:outline-none border-2 rounded-md px-3 shadow-md"
            />
          </div>

          <div class="mt-5">
            <router-link to="/edit-personal-data"
              ><button
                class="w-full bg-[#90c700] text-white font-bold text-2xl py-3 rounded shadow-md"
                type="button"
              >
                編輯
              </button></router-link
            >
          </div>
        </div>
      </form>
    </div>
  </body>
</template>
<script setup lang="ts">
import Avatar from "@/components/Avatar.vue";
import { PhArrowUUpLeft } from "@phosphor-icons/vue";
</script>
