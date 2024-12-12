<template>
    <div class="flex bg-[#90c700] p-3 items-center w-full px-6">
        <router-link to="/home" class="">
            <PhArrowUUpLeft :size="32" weight="bold" color="#000" />
        </router-link>
        <label class="ml-5 text-H2 font-bold ">福利詳細資料</label>
    </div>

    <div class="flex flex-col p-4">
        <span class=" text-H1 font-bold">{{ welfare.title }}</span>
        <span class=" text-H3 font-bold">{{ welfare.city }}</span>
        <span class="text-H3">發佈日期 : {{ welfare.date ? welfare.date : "無法獲取日期" }}</span>
        <div class="flex flex-col my-2 space-y-2">
            <span class=" font-bold">申請條件:</span>
            <div v-for="(item, index) in welfare.detailCondition" :key="index" class="grid grid-cols-6 items-center gap-4 py-2 mx-3 border-b border-gray-300 last:border-0">
                <div class="grid text-center ">{{ index+1 }}</div>
                <div class=" col-span-3">{{ item }}</div>
            </div>
        </div>

        <div class="flex flex-col space-y-2 my-2 ">
            <span class=" font-bold ">可獲得之福利:</span>
            <div v-for="(item, index) in welfare.forward" :key="index" class="grid grid-cols-6 items-center gap-4 py-2 mx-3 border-b border-gray-300 last:border-0">
                <div class="grid text-center ">{{ index+1 }}</div>
                <div class=" col-span-3">{{ item }}</div>
            </div>
        </div>


        <span class=" font-bold">點此前往-> <a class=" italic font my-2 decoration-double underline text-blue-500" :href="welfare.url">原文鏈接</a></span>
        <span class="text-H3 my-2 leading-loose">簡要原文 : {{ welfare.detail ? welfare.detail : "無法獲取原文" }}</span>

    </div>
</template>

<script setup lang="ts">
import request from '@/axios';
import Welfare from '@/model/welfare';
import { onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import { PhArrowUUpLeft } from "@phosphor-icons/vue";

const route = useRoute()
const welfare = ref<Welfare>(new Welfare())

onMounted(async () => {
    const result = await request.get("/welfare/" + route.params.id)
    console.log(result,'詳細福利')
    welfare.value = result.data

})


</script>

<style scoped></style>