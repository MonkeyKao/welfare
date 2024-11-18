<template>
<div class="h-screen bg-[#92c700] flex flex-col space-y-2">
  <!-- 快速福利查詢區塊 -->
  <div class="flex-1 basis-1/4 flex flex-col justify-center space-y-2">
    <h2 class="text-center text-H1 font-bold">快速福利查詢</h2>
    <div class="border-t-2 my-2 border-black"></div>
    <div class="grid grid-cols-2 px-5 gap-4 ">
      <button
        v-for="(item, index) in welfareItems"
        :key="index"
        class="text-H2 bg-white font-bold rounded-lg shadow-md mt-2"
        @click="showMessage(item)"
      >
        {{ item }}
      </button>
    </div>
  </div>

  <!-- 功能選項區塊 -->
  <div class="flex-3 basis-3/4 flex flex-col">
    <h2 class="text-H1 font-bold text-center">功能選項</h2>
    <div class="border-t-2 my-2 border-black"></div>
    <div class="grid grid-cols-2 gap-4 p-4 flex-grow">
      <button
        v-for="(item, index) in menuItems"
        :key="index"
        class="bg-white rounded-lg shadow-md flex flex-col items-center justify-center"
        @click="router.push(item.path)"
      >
        <component :is="item.icon" size="40" color="#6E6144" />
        <label class="text-H2 font-bold">{{ item.label }}</label>
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
