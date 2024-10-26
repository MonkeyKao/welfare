<template>
  <header class="top-0 left-0 right-0 flex items-center p-5 pt-14 bg-[#92c700]">
    <router-link to="/user" class="forgot-password-link">
      <PhArrowUUpLeft :size="32" color="#000" class="ml-5" />
    </router-link>
    <label class="ml-5 text-3xl font-bold">個人資訊</label>
  </header>
  <body>
    <div>
      <Avatar />
      <form class="flex flex-col justify-center items-center mt-5">
        <div>
          <div class="mt-5">
            <label class="block text-left text-xl mb-2">姓名</label>
            <input
              class="flex-1 focus:outline-none border-2 rounded-md px-3 shadow-md"
              type="text"
            />
          </div>
          <div class="mt-5">
            <label class="block text-left text-xl mb-2">生日</label>
            <input
              class="flex w-full focus:outline-none border-2 rounded-md px-3 shadow-md"
              type="date"
            />
          </div>
          <div class="mt-5">
            <label class="block text-left text-xl mb-2">性別</label>
            <div class="flex items-center">
              <input type="radio" id="male" name="gender" class="mr-2" />
              <label for="male">男性</label>
              <input type="radio" id="female" name="gender" class="mr-2 ml-4" />
              <label for="neutral">中性</label>
              <input type="radio" id="female" name="gender" class="mr-2 ml-4" />
              <label for="female">女性</label>
            </div>
          </div>

          <div class="mt-5">
            <label class="block text-left text-xl mb-2">地區</label>
            <select
              class="flex-1 focus:outline-none border-2 rounded-md px-3 shadow-md"
            >
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

          <div class="mt-5">
            <router-link to="/home"
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
