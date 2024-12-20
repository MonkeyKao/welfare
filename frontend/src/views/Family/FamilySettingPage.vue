<template>
    <div class="">
        <div class=" grid grid-cols-3 m-3">
            <div @click="router.go(-1)" class="forgot-password-link flex items-center">
                <PhArrowUUpLeft :size="28" color="#000" />
            </div>
            <span class=" text-center text-H2 font-bold">家庭設定</span>

        </div>

        <div class="flex  space-x-3 overflow-x-auto mt-4 mx-3">
            <div class="flex-shrink-0 items-center justify-center">
                <img src="../../../public/logo.png" alt="家庭圖片" class="w-16">
                <span class="text-center block text-H4">毛</span>
            </div>
            <div class="flex-shrink-0 items-center justify-center">
                <img src="../../../public/logo.png" alt="家庭圖片" class="w-16">
                <span class="text-center block  text-H4">毛</span>
            </div>
            <div class="flex-shrink-0 items-center justify-center">
                <img src="../../../public/logo.png" alt="家庭圖片" class="w-16">
                <span class="text-center block  text-H4">毛</span>
            </div>
        </div>
        <div class="border-2 mt-5 border-gray-100 "></div>

        <div class=" flex flex-col items-start mx-5">
            <div class="flex border-b w-full text-H3 py-3 justify-between">
                <span>群組名稱</span>
                <div>
                    <PhCaretRight :size="24" color="#8a8d8f" />
                </div>
            </div>
            <div class="flex border-b w-full text-H3 py-3 justify-between">
                <span>開啟通知</span>
                <input type="checkbox" class="toggle toggle-success " />


            </div>
            <span class=" w-full text-H3 py-3" @click="showQRCODE = !showQRCODE;">產生QRCODE</span>

        </div>

        <div class="border-2 mt-5 border-gray-100 "></div>

        <!-- 弹出式QRCODE卡片 -->
        <div v-if="showQRCODE" class="fixed inset-0 bg-gray-500 bg-opacity-50 flex justify-center items-center z-50">
            <div
                class=" flex bg-white items-start p-6 rounded-lg shadow-lg transform transition-all duration-300 scale-100 hover:scale-105">
                <!-- QRCode 组件 -->
                <QRCODE :id="route.params.id" />

                <!-- 關閉按鈕 -->
                <button @click="showQRCODE = false"class=" items-start justify-start ml-5">
                    <PhX :size="32" color="#1e1f1f" />
                </button>
            </div>
        </div>
        <div class=" flex flex-col items-start mx-5">
            <span @click="deleteFamily(Number(route.params.id.toString()))"
                class="  w-full text-H3 py-3 text-center text-[#D06262]">
                離開家庭
            </span>
            <span @click="deleteFamily(Number(route.params.id.toString()))"
                class="  w-full text-H3 py-3 text-center text-[#D06262]">
                刪除家庭
            </span>
        </div>

    </div>
</template>

<script setup lang="ts">
import router from '@/router';
import { useFamilyStore } from '@/store/family';
import { inject, ref } from 'vue';
import { useRoute } from 'vue-router';
import QRCODE from '@/components/QRCODE.vue';
import { PhArrowUUpLeft, PhCaretRight,PhX } from "@phosphor-icons/vue";
//import HeaderBar from '@/components/HeaderBar.vue';

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

<style scoped>
/* 覆盖 toggle 成功状态的背景颜色 */
.toggle:checked {
    background-color: #73AA00 !important;
    /* 设置选中时的背景颜色 */
}

.toggle:checked+.toggle-bg {
    background-color: #73AA00 !important;
    /* 同时设置选中时的背景 */
}
</style>