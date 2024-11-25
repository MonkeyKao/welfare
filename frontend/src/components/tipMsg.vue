<template>
    <div class="fixed z-50 bottom-5 flex justify-center w-screen">
        <Transition name="slide">
            <div v-if="tipState.visible" @click="tipState.onClick?.()"
                class="border shadow-xl text-lime-600 rounded-xl bg-white w-auto px-10 py-1 text-lg">
                {{ tipState.text }}
            </div>
        </Transition>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const tipState = ref<{
    text: string;
    visible: boolean;
    onClick?: () => void;
}>({ text: "", visible: false });

const showMsg = (msg: string, clickFunction?: () => void) => {
    tipState.value = { text: msg, visible: true, onClick: clickFunction };
    autoClose();
};

let timer: ReturnType<typeof setTimeout>;

const autoClose = () => {
    clearTimeout(timer);
    timer = setTimeout(() => {
        tipState.value.visible = false;
        tipState.value.onClick = undefined;
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