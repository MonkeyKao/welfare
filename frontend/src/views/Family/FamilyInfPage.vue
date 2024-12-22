<template>
  <HeaderBar>
    {{family?.familyName}}
    <template v-slot:siderBar>
      <div @click="router.push('/user/family/setting/' + route.params.id)">
        <PhDotsThree :size="28" color="#000"/>
      </div>
    </template>
  </HeaderBar>
</template>

<script setup lang="ts">
import { useRoute } from 'vue-router';
import { inject, onMounted, ref } from 'vue';
import HeaderBar from '@/components/HeaderBar.vue';
import request from '@/axios';
import router from '@/router';
import { useFamilyStore } from '@/store/family';
import { PhDotsThree } from "@phosphor-icons/vue";
import { showMsgFunction } from '@/type/ShowMsg';
import Family from '@/model/family';

const showMsg: showMsgFunction = inject("showMsg")!
const route = useRoute()
const familyStore = useFamilyStore();
const family = ref<Family>();
onMounted(() => {
  const familyId = route.params.id as string
  family.value = familyStore.families.find((item) => String(item.familyId)==familyId)
})

</script>

<style scoped></style>