<!--HTML-->
<template>
  <HeaderBar>常見問題</HeaderBar>

  <div class="mt-[70px] flex flex-col mb-5">
    <div
      v-for="item in QAs"
      :key="item.Id"
      class="mx-8 gap-5 p-4 px-4 mt-3 cursor-pointer border border-gray-300 rounded-lg shadow-md "
      
      @click="toggle(item.Id)"
    >
      <div class="flex justify-between space-x-1">
        <div class=" flex ">
          <span class=" text-H3 font-bold ">Q </span>
          <span class="mx-2">:</span>
          <span class=" font-bold">{{ item.Question }}</span>          
        </div>
        <span>
          <PhCaretRight
            v-if="!getIsActive(item.Id)"
            :size="25"
            color="#92c700"
          />
          <PhCaretDown v-if="getIsActive(item.Id)" :size="25" color="#92c700" />
        </span>
      </div>
      <div class="">
        <transition name="fade" class=" ">
          <div
            v-if="getIsActive(item.Id)"
            class="animate__animated animate__fadeIn  flex items-start mx-5 pt-5 "
          >
            <span class="font-bold">A</span>
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
import { PhCaretDown, PhCaretRight } from "@phosphor-icons/vue";
import { ref } from "vue";
import HeaderBar from "@/components/HeaderBar.vue";
import 'animate.css';


const QAs = ref<QA[]>([]); // 明確指定 QAs 為 QA 型別的陣列
const openItems = ref<number[]>([]);

const getQa = async () => {
  const result = await request.get("QA");
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
