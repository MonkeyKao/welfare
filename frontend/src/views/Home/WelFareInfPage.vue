<template>
    <div class="flex bg-[#90c700] p-3 items-center w-full px-6">
        <router-link to="/home" class="">
            <PhArrowUUpLeft :size="32" weight="bold" color="#000" />
        </router-link>
        <label class="ml-5 text-H2 font-bold ">福利詳細資料</label>
    </div>

    <div class="flex flex-col p-4">
        <span class=" text-H2 font-bold">{{ welfare.title }}</span>
        <!--  
        <span class=" text-H3 font-bold">{{ getTextByLocation(welfare.city) }}</span>         
        -->
        <span class="text-H3 mt-3">發佈日期 : {{ welfare.date ? welfare.date : "無法獲取日期" }}</span>
        <div class="flex flex-col my-2 space-y-2">
            <span class=" font-bold">申請條件</span>
            <div v-for="(item, index) in welfare.detailCondition" :key="index"
                class="grid grid-cols-6 items-center gap-4 py-2 mx-3 border-b border-gray-300 last:border-0">
                <div class="grid text-center col-span-1">{{ index + 1 }}</div>
                <div class="col-span-5">{{ item }}</div>
            </div>
        </div>

        <div class="flex flex-col space-y-2 my-3 ">
            <span class=" font-bold ">可獲得之福利</span>
            <div v-for="(item, index) in welfare.forward" :key="index"
                class="grid grid-cols-6 items-center gap-4 py-2 mx-3 border-b border-gray-300 last:border-0">
                <div class="grid text-center col-span-1">{{ index + 1 }}</div>
                <div class=" col-span-5">{{ item }}</div>
            </div>
        </div>
        <div class="flex items-center justify-between my-3">
            <div class=" flex justify-center items-center">
                <span class="font-bold ">您是否符合申請條件</span>
                <PhInfo :size="20" class=" cursor-pointer mx-2" @click="showTooltip = !showTooltip" />
                <PhCircle :size="20" :color="getCanGetStatus(welfare.canGet).color" weight="fill" class="mx-2" />
                <span>{{ getCanGetStatus(welfare.canGet).text }}</span>
            </div>
            <div class="relative">
                <div v-if="showTooltip"
                    class="absolute right-32 bottom-8 w-44 bg-slate-50 space-y-2  text-sm p-2 rounded shadow-md">
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
        <div class="flex flex-col space-y-2 my-3 ">
            <span class=" font-bold ">可獲得福利之家人</span>
            <div v-for="(item, index) in people" :key="index"
                class="grid grid-cols-6 items-center gap-4 py-2 mx-3 border-b border-gray-300 last:border-0">
                <div class="grid text-center col-span-1 avatar">
                    <div class="w-8 rounded-full">
                        <img :src="item.image" alt="Avatar" />                        
                    </div>
                </div>
                <div class=" col-span-5">{{ item.name }}</div>
            </div>
        </div>
        <span class=" font-bold">點此前往-> <a class=" italic font my-2 decoration-double underline text-blue-500"
                :href="welfare.url">原文鏈接</a></span>
        <span class="text-H3 font-bold mt-2 leading-loose">簡要原文 :</span>
        <span class="text-H3 leading-loose"> {{ welfare.detail ? welfare.detail : "無法獲取原文" }}</span>
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

const getCanGetStatus = (canGet: number): { color: string, text: string } => {
    switch (canGet) {
        case 1:
            return { color: "#92C700", text: "符合領取資格 !" }; // 綠色=快去領
        case 2:
            return { color: "#FFEA53", text: "請再次確認能否領取" }; // 黃色=不確定能不能領
        case 3:
            return { color: "#D06262", text: "無法領取" }; // 紅色=沒得領啦
        default:
            return { color: "#cccccc", text: "未知狀態" }; // 預設灰色
    }
};

const people = [
  { image:"../../../public/logo.png",name:" 大哥"},
  { image:"../../../public/password.jpg",name:" 猴子"},
  { image:"../../../public/forgetpassword.jpg",name:" 子芸"},
  { image:"../../../public/login.jpg",name:" 茹茵"},
];
const showTooltip = ref(false); 
</script>

<style scoped></style>