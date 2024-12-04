<template>
    <HeaderBar>
        家庭詳情頁面
        <template v-slot:siderBar>
            123
        </template>
    </HeaderBar>
    <div>
        <button class="rounded border-2 px-4 border-black" @click="showQRCODE = true;">產生QRCODE</button>
        <button @click="deleteFamily(Number(route.params.id.toString()))" class="rounded border-2 px-4 border-black">
            刪除家庭
        </button>
    </div>
    <div>

        <QRCODE v-if="showQRCODE" :id="route.params.id" ref="qrCodeRef"></QRCODE>
    </div>

</template>

<script setup lang="ts">
import { useRoute } from 'vue-router';
import QRCODE from '@/components/QRCODE.vue';
import { inject, ref } from 'vue';
import HeaderBar from '@/components/headerBar.vue';
import request from '@/axios';
import router from '@/router';
import { useFamilyStore } from '@/store/family';
const showQRCODE = ref<boolean>(false)
const showMsg: Function = inject("showMsg")!
const route = useRoute()
const familyStore = useFamilyStore();

const deleteFamily = async (id: number) => {
  try {
    familyStore.deleteFamily(id);
    router.push('/user/family')
    showMsg("刪除成功")
  } catch (err: any) {
    showMsg(err.response.data)
  }
}
</script>

<style scoped></style>