<template>
    <HeaderBar>
        家庭設定頁面
    </HeaderBar>

    <div>
        <button @click="deleteFamily(Number(route.params.id.toString()))" class="rounded border-2 px-4 border-black">
            刪除家庭
        </button>
        <button class="rounded border-2 px-4 border-black" @click="showQRCODE = !showQRCODE;">產生QRCODE</button>
    </div>

    <QRCODE v-if="showQRCODE" :id="route.params.id"></QRCODE>


</template>

<script setup lang="ts">
import router from '@/router';
import { useFamilyStore } from '@/store/family';
import { inject, ref } from 'vue';
import { useRoute } from 'vue-router';
import QRCODE from '@/components/QrCode.vue';
import HeaderBar from '@/components/HeaderBar.vue';

const showMsg: Function = inject("showMsg")!
const showQRCODE = ref<boolean>(false)
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