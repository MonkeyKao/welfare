<template>

    <button v-if="props.showOpenBtn" type="button" class="" @click="openModal">
        <slot name="open-slot">{{ props.openBtnTitle
            }}</slot>
    </button>


    <dialog :id="ModalId" class=" z-50 bg-transparent ">

        <div class="rounded-lg bg-white p-10">
            <slot></slot>
            <div class="modal-action">
                <form class="flex gap-3" method="dialog">
                    <button v-if="displayCloseBtn" class="btn">關閉</button>
                    <button @click="emit('onClickConfirm')" class="btn btn-primary">{{ closeBtnTitle ? closeBtnTitle :
                        '關閉' }}</button>
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

defineExpose({
    openModal,
})

</script>

<style scoped></style>