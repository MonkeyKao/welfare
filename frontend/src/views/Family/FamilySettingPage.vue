<template>
    <div v-if="family" class="">
        <div class="grid grid-cols-3 m-3">
            <div @click="router.go(-1)" class="forgot-password-link flex items-center">
                <PhArrowUUpLeft :size="28" color="#000" />
            </div>
            <span class=" text-center text-H2 font-bold">家庭設定</span>

        </div>

        <div class="flex space-x-3 overflow-x-auto mt-4 mx-3">
            <div v-for="(item, index) in people" :key="index" class="avatar flex flex-col items-center">
                <div class="w-12 h-12 rounded-lg overflow-hidden">
                    <img :src="item.image" alt="Avatar" />

                </div>
                <span>{{ item.name }}</span>
            </div>
        </div>
        <div class="border-2 mt-5 border-gray-100 "></div>

        <div class=" flex flex-col items-start mx-5">
            <div class="flex border-b w-full text-H3 py-3 justify-between">
                <span>群組名稱:{{ family.familyName }}</span>
                <div>
                    <PhCaretRight :size="24" color="#8a8d8f" />
                </div>
            </div>
            <div class="flex border-b w-full text-H3 py-3 justify-between">
                <span>開啟通知</span>
                <input type="checkbox" class="toggle toggle-success " />


            </div>

            <div class="w-full text-H3 py-3" @click="showQRCODE = true">
                <Modal close-btn-title="關閉" @on-click-confirm="showQRCODE = false" open-btn-title="產生QRCODE">
                    <span class="text-H2 font-bold text-center">掃描QRCode加入家庭</span>
                    <QRCODE v-if="showQRCODE"></QRCODE>
                </Modal>
            </div>

        </div>

        <div class="border-2 mt-5 border-gray-100 "></div>

        <div class=" flex flex-col-reverse items-start mx-5">
            <span v-if="family.users.find((item) => item.role == 3)?.user_name === userStore.user.name"
                @click="deleteFamily(Number(route.params.id))" class="  w-full text-H3 py-3 text-center text-[#D06262]">
                刪除家庭
            </span>
            <span v-else @click="leaveFamily(Number(route.params.id))"
                class="  w-full text-H3 py-3 text-center text-[#D06262]">
                離開家庭
            </span>
        </div>
    </div>
    <div v-else>
        <HeaderBar>家庭</HeaderBar>
        讀取資料失敗
    </div>
</template>

<script setup lang="ts">
import router from '@/router';
import { useFamilyStore } from '@/store/familyStore';
import { inject, onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import QRCODE from '@/components/QRCODE.vue';
import { PhArrowUUpLeft, PhCaretRight, PhX } from "@phosphor-icons/vue";
import { AlertColor, showMsgFunction } from '@/type/ShowMsg';
import Family from '@/model/family';
import { useUserStore } from '@/store/userStore';
import HeaderBar from '@/components/HeaderBar.vue';
import Modal from '@/components/modal.vue';
const showMsg: showMsgFunction = inject("showMsg")!
const showQRCODE = ref<boolean>(false)
const route = useRoute()
const familyStore = useFamilyStore();
const family = ref<Family>()
const userStore = useUserStore()

const people = [
    { image: "../../../public/mao.jpg", name: "劉采沅" },
    { image: "../../../public/liao.jpg", name: "廖柏安" },
    { image: "../../../public/yun.jpg", name: "林子芸" },
];

const toggleQrcode = () => {
    showQRCODE.value = !showQRCODE.value
}

const deleteFamily = async (id: number) => {
    try {
        familyStore.deleteFamily(id);
        router.push('/user/family')
        showMsg("刪除成功")
    } catch (err: any) {
        showMsg(err.response.data, AlertColor.error)
    }
}

const leaveFamily = async (id: number) => {
    try {
        familyStore.deleteFamily(id);
        router.push('/user/family')
        showMsg("離開成功", AlertColor.waring)
    } catch (err: any) {
        showMsg(err.response.data, AlertColor.error)
    }
}

onMounted(() => {
    const familyId = route.params.id as string
    family.value = familyStore.families.find((item) => String(item.familyId) == familyId)
})
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