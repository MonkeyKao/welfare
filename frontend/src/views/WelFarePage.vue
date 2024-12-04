<template>
    <div class="flex bg-[#90c700] p-3 items-center w-full px-6">
        <router-link to="/home" class="">
            <PhArrowUUpLeft :size="32" weight="bold" color="#000" />
        </router-link>
        <label class="ml-5 text-H2 font-bold">福利詳細資料</label>
    </div>

    <div class="flex flex-col">
        <span class=" text-H1 font-bold">{{ welfare.title }}</span>
        <span>{{ welfare.city }}</span>
        <span class="text-H3">發佈日期:{{ welfare.date ? welfare.date : "無法獲取日期" }}</span>
        <div class="flex flex-col">
            <span>申請條件:</span>
            <span>1.</span>
            <span>2.</span>
            <span>3.</span>
        </div>

        <div class="flex flex-col">
            <span>可獲得之福利:</span>
            <span>1.</span>
            <span>2.</span>
            <span>3.</span>
        </div>


        <span><a class=" italic font" :href="welfare.url">原文鏈接</a></span>
        <span class="text-H3">原文:{{ welfare.detail ? welfare.detail : "無法獲取原文" }}</span>

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
    welfare.value = result.data



})





</script>

<style scoped></style>