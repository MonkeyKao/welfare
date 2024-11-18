<template>
  <div class="flex flex-col bg-[#92c700] h-screen p-4 gap-5">
    <!-- 快速福利查詢區塊 -->
    <div class="text-center mt-2">
      <h2 class="text-H1 font-bold">快速福利查詢</h2>
      <div class="border-t-2 my-2 border-black"></div>
      <div class="flex flex-wrap justify-between gap-2">
        <button
          v-for="(item, index) in welfareItems"
          :key="index"
          class="text-H2 bg-white font-bold py-2 rounded-lg shadow-md w-[48%]"
          @click="showMessage(item)"
        >
          {{ item }}
        </button>
      </div>
    </div>

    <!-- 功能選項區塊 -->

    <div class="text-center">
      <h2 class="text-H1 font-bold">功能選項</h2>
      <div class="border-t-2 my-2 border-black"></div>
      <div class="flex flex-wrap justify-between gap-3">
        <button
          v-for="(item, index) in menuItems"
          :key="index"
          class="flex flex-col items-center bg-white p-2 rounded-md shadow-xl w-[48%]"
          @click="router.push(item.path)"
        >
          <component :is="item.icon" size="60" color="#6E6144" />
          <label class="text-H1 font-bold">{{ item.label }}</label>
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import router from "@/router";
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
const welfareItems = ["醫療補助", "老人年金", "長照補助", "老人生活津貼"];
const menuItems = [
  { path: "/favorites", label: "收藏", icon: PhHeart },
  { path: "/notifications", label: "通知", icon: PhBell },
  { path: "/elderlysearch", label: "福利查詢", icon: PhMagnifyingGlass },
  { path: "/ai", label: "機器查詢", icon: PhRobot },
  { path: "/user/personal-data", label: "個人資料", icon: PhUser },
  { path: "/user/family", label: "家庭", icon: PhUsersThree },
  { path: "/user/qa", label: "常見問題", icon: PhQuestion },
  { path: "/user/settings", label: "設定", icon: PhGear },
];
const showMessage = (item: string) => {
  console.log(`已點擊: ${item}`);
};
</script>
