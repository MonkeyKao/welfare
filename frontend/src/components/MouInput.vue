<template>
  <div class="flex justify-between p-3 bg-white">
    <input v-model="msg" type="text" placeholder="請輸入您的問題" class="outline-none border-none focus:ring-0 bg-transparent px-4 py-2 text-gray-700 w-full">
    <PhPaperPlaneRight  @click="emit('clickSendMsg',msg);msg=''" :size="28" class="m-3" color="#92c700" weight="fill" />
  </div>
</template>

<script setup lang="ts">
import {PhPaperPlaneRight} from '@phosphor-icons/vue'
import { ref } from 'vue';
const msg = ref<string>("");
const emit = defineEmits(["clickSendMsg"])
</script>