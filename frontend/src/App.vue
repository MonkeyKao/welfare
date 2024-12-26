<script setup lang="ts">
import BottomNav from "./components/BottomNav.vue";
import { ComponentPublicInstance, computed, onMounted, provide, ref, watch } from "vue";
import { RouterView, useRoute } from "vue-router"; // 引入 RouterView 用來動態渲染路由對應的頁面
import TipMsg from "./components/TipMsg.vue";

const route = useRoute(); // 獲取當前路由

type TipMsgInstance = ComponentPublicInstance<{},{showMsg : () => void}>

// 计算属性，用于判断是否显示底部导航
const showBottomBar = computed(() => route.meta.showBottomBar ?? true);

const tipmsg = ref<TipMsgInstance|null>(null);
onMounted(() => {
  if (tipmsg.value && tipmsg.value.showMsg) {
    provide("showMsg", tipmsg.value.showMsg);
  }
});
</script>

<template>
  <!-- 渲染當前頁面 -->
  <div class="h-screen flex flex-col">
    <!-- 动态渲染路由页面 -->
    <div v-if="['AiPage'].includes(String(route.name))" class="overflow-auto pb-10 flex-grow">
      <RouterView />
    </div>
    <div v-else class="overflow-auto flex-grow">
      <RouterView />
    </div>

    <!-- 底部导航 -->
    <div class="flex-shrink-0">
      <div v-if="showBottomBar">
        <BottomNav />
      </div>
    </div>
  </div>

  <TipMsg ref="tipmsg"></TipMsg>
</template>
