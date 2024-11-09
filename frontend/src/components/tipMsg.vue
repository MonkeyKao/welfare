<template>
    <div class="fixed z-50 bottom-5 flex justify-center w-screen">
        <Transition name="slide" @after-enter="autoClose">
            <div v-if="showTip" class="border shadow-xl text-lime-600 rounded-xl bg-white w-auto px-10 py-1 text-lg">
                {{showText}}
            </div>
        </Transition>
    </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue';

const showText=ref<string>("");


const showMsg = (msg:string,clickFunction?:Function) => {
    showText.value = msg;
    showTip.value = true
} 

const showTip = ref(false);

const autoClose = () => {
    setTimeout(() => {
        showTip.value = false;
    }, 1500);
}

defineExpose({
    showMsg,
})
</script>

<style scoped>
.slide-enter-active, .slide-leave-active {
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