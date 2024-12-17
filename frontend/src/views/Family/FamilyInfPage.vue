<template>
  <HeaderBar>
    家庭詳情頁面
    <template v-slot:siderBar>
      <div @click="router.push('/user/family/setting/' + route.params.id)">
        <PhDotsThree :size="28" color="#000"/>
      </div>
    </template>
  </HeaderBar>

  我應該要可以聊天

</template>

<script setup lang="ts">
import { useRoute } from 'vue-router';
import { inject, ref } from 'vue';
import HeaderBar from '@/components/HeaderBar.vue';
import request from '@/axios';
import router from '@/router';
import { useFamilyStore } from '@/store/family';
import { PhDotsThree } from "@phosphor-icons/vue";

const showMsg: Function = inject("showMsg")!
const route = useRoute()
const familyStore = useFamilyStore();


</script>

<style scoped></style>