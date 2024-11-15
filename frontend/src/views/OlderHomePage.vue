<template>
  <div class="flex flex-col bg-[#92c700] h-screen p-4 gap-10">
    <!-- 快速福利查詢區塊 -->
    <div class="text-center">
      <h2 class="text-H1 font-bold">快速福利查詢</h2>
      <div class="border-t-2 border-[#3e3e3e] my-2"></div>
      <div class="flex flex-wrap justify-between gap-2">
        <button
          class="text-H2 bg-white font-bold py-2 rounded shadow-md w-[48%]"
        >
          醫療補助
        </button>
        <button
          class="text-H2 bg-white font-bold py-2 rounded shadow-md w-[48%]"
        >
          老人年金
        </button>
        <button
          class="text-H2 bg-white font-bold py-2 rounded shadow-md w-[48%]"
        >
          長照補助
        </button>
        <button
          class="text-H2 bg-white font-bold py-2 rounded shadow-md w-[48%]"
        >
          老人生活津貼
        </button>
      </div>
    </div>

    <!-- 功能選項區塊 -->

    <div class="text-center">
      <h2 class="text-H1 font-bold">功能選項</h2>
      <div class="border-t-2 border-[#3e3e3e] my-2"></div>
      <div class="flex flex-wrap justify-between gap-2">
        <button
          class="flex flex-col items-center bg-white p-4 rounded shadow-md w-[48%]"
        >
          <PhHeart size="50" color="#3e3e3e" />
          <span class="text-H1 font-bold">收藏</span>
        </button>
        <button
          class="flex flex-col items-center bg-white p-4 rounded shadow-md w-[48%]"
        >
          <PhBell size="50" color="#3e3e3e" />
          <span class="text-H1 font-bold">通知</span>
        </button>
        <button
          class="flex flex-col items-center bg-white p-4 rounded shadow-md w-[48%]"
        >
          <PhMagnifyingGlass size="50" color="#3e3e3e" />
          <span class="text-H1 font-bold">福利查詢</span>
        </button>
        <button
          class="flex flex-col items-center bg-white p-4 rounded shadow-md w-[48%]"
        >
          <PhRobot size="50" color="#3e3e3e" />
          <span class="text-H1 font-bold">語音查詢</span>
        </button>
        <button
          class="flex flex-col items-center bg-white p-4 rounded shadow-md w-[48%]"
        >
          <PhUser size="50" color="#3e3e3e" />
          <span class="text-H1 font-bold">個人資料</span>
        </button>
        <button
          class="flex flex-col items-center bg-white p-4 rounded shadow-md w-[48%]"
        >
          <PhUsersThree size="55" color="#3e3e3e" />
          <span class="text-H1 font-bold">家庭</span>
        </button>
        <button
          class="flex flex-col items-center bg-white p-4 rounded shadow-md w-[48%]"
        >
          <PhQuestion size="50" color="#3e3e3e" />
          <span class="text-H1 font-bold">常見問題</span>
        </button>
        <button
          class="flex flex-col items-center bg-white p-4 rounded shadow-md w-[48%]"
        >
          <PhGear size="50" color="#3e3e3e" />
          <span class="text-H1 font-bold">設定</span>
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {
  PhBell,
  PhGear,
  PhHeart,
  PhMagnifyingGlass,
  PhQuestion,
  PhRobot,
  PhUser,
  PhUsersThree,
} from "@phosphor-icons/vue";
</script>
