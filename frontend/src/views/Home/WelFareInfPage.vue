<template>
    <div class="flex bg-[#90c700] p-3 items-center w-full px-6">
        <router-link to="/home" class="">
            <PhArrowUUpLeft :size="32" weight="bold" color="#000" />
        </router-link>
        <label class="ml-5 text-H2 font-bold ">福利詳細資料</label>
    </div>

    <div class="flex flex-col p-4">
        <span class=" text-H2 text-center font-bold">{{ welfare.title }}</span>
        <span class=" text-H3 font-bold">{{ getTextByLocation(welfare.city) }}</span>
        <span class="text-H3">發佈日期 : {{ welfare.date ? welfare.date : "無法獲取日期" }}</span>
        <div class="flex flex-col my-2 space-y-2">
            <span class=" font-bold">申請條件</span>
            <div v-for="(item, index) in welfare.detailCondition" :key="index"
                class="grid grid-cols-6 items-center gap-4 py-2 mx-3 border-b border-gray-300 last:border-0">
                <div class="grid text-center col-span-1">{{ index + 1 }}</div>
                <div class="col-span-5">{{ item }}</div>
            </div>
        </div>
        <div class="flex items-center justify-between my-2">
            <div class=" flex justify-center items-center">
                <span class="font-bold ">是否符合申請條件</span>
                <PhInfo :size="20" class=" cursor-pointer mx-2" @click="showTooltip = !showTooltip" />
                <PhCircle :size="20" :color="getCanGetColor(welfare.canGet)" weight="fill" class="mx-5" />
            </div>
            <div class="relative">
                <div v-if="showTooltip"
                    class="absolute right-0 bottom-8 w-44 bg-slate-50 space-y-2  text-sm p-2 rounded shadow-md">
                    <div class="flex items-center space-x-2">
                        <PhCircle :size="16" color="#92C700" weight="fill" />
                        <span>符合領取資格</span>
                    </div>
                    <div class="flex items-center space-x-2">
                        <PhCircle :size="16" color="#FFEA53" weight="fill" />
                        <span>不一定符合領取資格</span>
                    </div>
                    <div class="flex items-center space-x-2">
                        <PhCircle :size="16" color="#D06262" weight="fill" />
                        <span>不符合領取資格</span>
                    </div>
                </div>
            </div>
        </div>
        <div class="flex flex-col space-y-2 my-2 ">
            <span class=" font-bold ">可獲得之福利</span>
            <div v-for="(item, index) in welfare.forward" :key="index"
                class="grid grid-cols-6 items-center gap-4 py-2 mx-3 border-b border-gray-300 last:border-0">
                <div class="grid text-center col-span-1">{{ index + 1 }}</div>
                <div class=" col-span-5">{{ item }}</div>
            </div>
        </div>


        <span class=" font-bold">點此前往-> <a class=" italic font my-2 decoration-double underline text-blue-500"
                :href="welfare.url">原文鏈接</a></span>
        <span class="text-H3 my-2 leading-loose">簡要原文 : {{ welfare.detail ? welfare.detail : "無法獲取原文" }}</span>

    </div>
</template>

<script setup lang="ts">
import request from '@/axios';
import Welfare from '@/model/welfare';
import { onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import { PhArrowUUpLeft, PhCircle, PhInfo } from "@phosphor-icons/vue";
import { getTextByLocation } from '@/utils/getTextByNumber';

const route = useRoute()
const welfare = ref<Welfare>(new Welfare())

onMounted(async () => {
    const result = await request.get("/welfare/" + route.params.id)
    console.log(result, '詳細福利')
    welfare.value = result.data

})

const getCanGetColor = (canGet: number): string => {
    switch (canGet) {
        case 1:
            return "#92C700"; // 綠色=快去領
        case 2:
            return "#FFEA53"; // 黃色=不確定能不能領
        case 3:
            return "#D06262"; // 紅色=沒得領啦
        default:
            return "#cccccc"; // 預設灰色
    }
};

const showTooltip = ref(false); 
</script>

<style scoped></style>