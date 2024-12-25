<template>

    <div v-if="props.showOpenBtn" type="button" class="w-full" @click="openModal">
        <slot name="open-slot">{{ props.openBtnTitle
            }}</slot>
    </div>


    <dialog :id="ModalId" class=" z-50 bg-transparent ">

        <div class="rounded-lg bg-white p-6">
            <slot></slot>
            <div class="modal-action justify-center">
                <form class="flex" method="dialog">
                    <button @click="closeModal();emit('onClickConfirm')" class=" text-[#D06262]">{{ closeBtnTitle ? closeBtnTitle :
                        '取消' }}</button>
                </form>
            </div>
        </div>
    </dialog>
</template>

<script setup lang="ts">
const emit = defineEmits(['onClickConfirm'])

const props = defineProps({
    openBtnTitle: {
        type: String,
        default: () => "開啓"
    },
    closeBtnTitle: {
        type: String,
        default: () => "關閉"
    },
    showOpenBtn: {
        type: Boolean,
        default: () => true
    },
    displayCloseBtn: {
        type: Boolean,
        default: () => false
    }
});

const ModalId = "ID" + Math.round(Math.random() * 100000)

const openModal = () => {
    const modalDom = document.getElementById(ModalId) as HTMLDialogElement;
    if (modalDom) {
        modalDom.showModal();
    } else {
        console.error(`Modal with id ${ModalId} not found`);
    }
};

const closeModal = () => {
    const modalDom = document.getElementById(ModalId) as HTMLDialogElement;
    if (modalDom) {
        modalDom.close();
    } else {
        console.error(`Modal with id ${ModalId} not found`);
    }
}

defineExpose({
    openModal,
    closeModal
})

</script>

<style scoped></style>