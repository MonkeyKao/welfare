<template>
    <div class="fixed z-50 bottom-5 flex justify-center w-screen">
        <Transition name="slide">
            <div v-if="tipState.visible" @click="tipState.onClick?.()"
                :class="['border', 'shadow-xl', 'rounded-xl', 'bg-white', 'w-auto', 'px-10', 'py-1', 'text-lg', tipState.color]">
                {{ tipState.text }}
            </div>
        </Transition>
    </div>
</template>

<script setup lang="ts">
import { AlertColor } from '@/type/ShowMsg';
import { ref } from 'vue';
const tipState = ref<{
    text: string;
    visible: boolean;
    color: string;
    onClick?: () => void;
}>({
    text: "",
    visible: false,
    color: 'text-lime-600'
});

const showMsg = (msg: string, color?: AlertColor, clickFunction?: () => void) => {
    tipState.value.text = msg
    tipState.value.onClick = clickFunction;
    tipState.value.visible = true;
    switch (color) {
        case AlertColor.success:
            tipState.value.color = "text-lime-600"
            break;
        case AlertColor.error:
            tipState.value.color = "text-red-400"
            break;
        case AlertColor.info:
            tipState.value.color = "text-black"
            break;
        case AlertColor.waring:
            tipState.value.color = "text-amber-400"
    }
    autoClose();
};

let timer: ReturnType<typeof setTimeout>;

const autoClose = () => {
    clearTimeout(timer);
    timer = setTimeout(() => {
        tipState.value.visible = false;
        tipState.value.onClick = undefined;
        tipState.value.color = "text-lime-600"
    }, 1500);
};

defineExpose({
    showMsg,
});
</script>

<style scoped>
.slide-enter-active,
.slide-leave-active {
    transition: transform 0.5s ease, opacity 0.5s ease;
}

.slide-enter {
    transform: translateY(100%);
    opacity: 0;
}

.slide-leave-to {
    transform: translateY(100%);
    opacity: 0;
}
</style>