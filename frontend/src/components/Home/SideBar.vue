<template>
    <!-- 側邊欄 -->
    <div :class="{ 'translate-x-0': showSidebar, 'translate-x-full': !showSidebar }"
        class="fixed top-0 right-0 space-y-4 w-64 h-full bg-white border-l border-gray-300 shadow-md p-4 transition-transform ease-in-out duration-300 z-10">
        <h3 class="font-bold text-xl text-center">篩選條件</h3>
        <div class="mt-4 space-y-5">
            <p class="text-lg font-bold">年齡</p>
            <div class="grid grid-cols-2 space-y-2" v-for="(item, index) in year" :key="index">
                <button @click="toggleSelect(item)" class="flex justify-between items-center">
                    <span :class="{ 'text-black': !item.selected, 'text-[#92c700]': item.selected }">
                        {{ item.name }}
                    </span>
                    <PhCheck v-if="item.selected" :size="20" color="#92c700" />
                </button>
            </div>
        </div>

        <div class="mt-4 space-y-5">
            <p class="text-lg font-bold">性別</p>
            <div class="grid grid-cols-2 gap-4">
                <button v-for="(item, index) in sex" :key="index" @click="toggleSelectSingle(item)"
                    class="flex justify-between items-center">
                    <span :class="{ 'text-black': !item.selected, 'text-[#92c700]': item.selected }">
                        {{ item.name }}
                    </span>
                    <PhCheck v-if="item.selected" :size="20" color="#92c700" />
                </button>
            </div>
        </div>

        <div class="mt-4 space-y-5">
            <p class="text-lg font-bold">收入</p>
            <div class="grid grid-cols-2 space-y-2" v-for="(item, index) in money" :key="index">
                <button @click="toggleSelectMoney(item)" class="flex justify-between items-center">
                    <span :class="{ 'text-black': !item.selected, 'text-[#92c700]': item.selected }">
                        {{ item.name }}
                    </span>
                    <PhCheck v-if="item.selected" :size="20" color="#92c700" />
                </button>
            </div>
        </div>

        <div class="mt-4 space-y-5">
            <p class="text-lg font-bold">身分別</p>
            <div class="grid grid-cols-2 space-y-2" v-for="(item, index) in people" :key="index">
                <button @click="toggleSelect(item)" class="flex justify-between items-center">
                    <span :class="{ 'text-black': !item.selected, 'text-[#92c700]': item.selected }">
                        {{ item.name }}
                    </span>
                    <PhCheck v-if="item.selected" :size="20" color="#92c700" />
                </button>
            </div>
        </div>


        <div class=" flex mt-4 justify-center">
            <button @click="applyFilters" class="bg-[#92c700] text-white px-4 py-2 rounded-md ">確定</button>
        </div>
    </div>

    <!-- 遮罩層 -->
    <div v-if="showSidebar" @click="toggleSidebar" class="fixed top-0 left-0 w-full h-full bg-black opacity-50 z-5">
    </div>



</template>

<script setup lang="ts">
import { useUserStore } from '@/store/userStroe';
import { PhCheck } from '@phosphor-icons/vue';
import { onMounted, ref } from 'vue';

const showSidebar = ref(false); // 控制側邊欄的顯示
const emit = defineEmits(["selectStatus"])
const toggleSidebar = () => {
    showSidebar.value = !showSidebar.value; // 切換側邊欄顯示狀態
};

const applyFilters = () => {
    // 在這裡處理篩選應用邏輯
    let selectedItem = [
        year.value.filter((item) => item.selected)[0]?.id || null,
        sex.value.filter((item) => item.selected)[0]?.id || null,
        money.value.filter((item) => item.selected)[0]?.id || null,
        people.value.filter((item) => item.selected)[0]?.id || null,
    ]
    emit("selectStatus", selectedItem)
    toggleSidebar(); // 篩選後關閉側邊欄
};

const year = ref([
    { id: 7, name: '20歲以下', selected: false },
    { id: 8, name: '20歲~65歲', selected: false },
    { id: 9, name: '65歲以上', selected: false },
])

const sex = ref([
    { id: 10, name: "男性", selected: false },
    { id: 11, name: "女性", selected: false },
])

const money = ref([
    { id: 1, name: "中低收入戶", selected: false },
    { id: 2, name: "低收入戶", selected: false },
])

const people = ref([
    { id: 3, name: "榮民", selected: false },
    { id: 4, name: "身心障礙者", selected: false },
    { id: 5, name: "原住民", selected: false },
    { id: 6, name: "外籍配偶家庭", selected: false }
])

const toggleSelect = (selectedItem: { id: number, name: string, selected: boolean }) => {
    // 如果當前選項已選擇，則取消選擇
    selectedItem.selected = !selectedItem.selected;
};

const toggleSelectSingle = (selectedItem: { id: number, name: string, selected: boolean }) => {
    if (!selectedItem.selected) {
        // 取消所有選項的選擇，然後選中當前項目
        sex.value.forEach(item => {
            item.selected = false;
        });
        selectedItem.selected = true; // 設置當前選項為選中狀態
    } else {
        // 如果當前項目已選擇，則取消選擇
        selectedItem.selected = false;
    }
};

const toggleSelectMoney = (selectedItem: { id: number, name: string, selected: boolean }) => {
    if (!selectedItem.selected) {
        // 取消所有選項的選擇，然後選中當前項目
        money.value.forEach(item => {
            item.selected = false;
        });
        selectedItem.selected = true; // 設置當前選項為選中狀態
    } else {
        // 如果當前項目已選擇，則取消選擇
        selectedItem.selected = false;
    }
}

defineExpose({ toggleSidebar })
</script>

<style scoped></style>