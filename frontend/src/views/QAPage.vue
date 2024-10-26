<!--HTML-->
<template>
  <header class="top-0 left-0 right-0 flex items-center p-5 pt-14 bg-[#92c700]">
    <router-link to="/home" class="forgot-password-link">
      <PhArrowUUpLeft :size="32" color="#fff" class="ml-5" />
    </router-link>
    <label class="ml-5 text-3xl font-bold text-white">常見問題 Q&A</label>
  </header>

  <div class="flex flex-col mt-2">
    <div
      v-for="item in QAs"
      :key="item.Id"
      class="mx-10 gap-5 p-4 px-6 mt-3 cursor-pointer border border-gray-300 rounded-lg shadow-md"
      @click="toggle(item.Id)"
    >
      <div class="flex justify-between items-center">
        <span class="text-lg font-semibold">Q : {{ item.Question }}</span>
        <span>
          <PhCaretRight
            v-if="!getIsActive(item.Id)"
            :size="25"
            color="#92c700"
          />
          <PhCaretDown v-if="getIsActive(item.Id)" :size="25" color="#92c700" />
        </span>
      </div>
      <div class="bg-[#92c700]">
        <transition name="fade">
          <div
            v-if="getIsActive(item.Id)"
            class="flex items-start space-x-2 m-7 bg-[#92c700]"
          >
            <span class="font-semibold">A</span>
            <span class="mx-2">:</span>
            <span>{{ item.Answer }}</span>
          </div>
        </transition>
      </div>
    </div>
  </div>
</template>

<!--JS/TS-->
<script setup lang="ts">
import request from "@/axios";
import QA from "@/model/qa";
import { PhArrowUUpLeft, PhCaretDown, PhCaretRight } from "@phosphor-icons/vue";
import { ref } from "vue";

const QAs = ref<QA[]>([]); // 明確指定 QAs 為 QA 型別的陣列
const openItems = ref<number[]>([]);

const getQa = async () => {
  const result = await request.get("welfare/QA");
  QAs.value = result.data; // 將 API 回傳的資料儲存到 QAs 中
};

getQa();

const toggle = (id: number) => {
  if (openItems.value.includes(id)) {
    openItems.value = openItems.value.filter((itemId) => itemId !== id);
  } else {
    openItems.value.push(id);
  }
};

const getIsActive = (id: number): boolean => openItems.value.includes(id);
</script>

<!--樣式-->
<style scoped></style>
