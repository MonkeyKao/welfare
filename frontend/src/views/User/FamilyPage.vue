<template>
  <HeaderBar>家庭</HeaderBar>

  <div class="p-4 flex flex-col gap-4">
    <div v-for="family in familyStroe.families" class="w-full p-3 flex justify-between shadow-md">
      <div>
        <span class="text-H2 font-bold flex-2/3 ">{{ family.familyName }}</span>
      </div>

      <button class="custom-button flex-1/3" @click="router.push('/user/family/' + family.familyId)">
        進入家庭頁面
      </button>

    </div>

    <div>
      <form @submit.prevent="joinFmaily(familyCode)">
        <input class="border-2" type="text" v-model="familyCode">
        <button class="custom-button" type="submit">手動加入</button>
      </form>
    </div>
  </div>

  <div>
    <button class="custom-button" @click="openCamera">掃描二維碼</button>
  </div>

  <div class="fixed bottom-5 right-10">
    <Modal>
      <template v-slot:open-slot>
        <PhPlus :size="32" />
      </template>
      <form @submit.prevent="familyStroe.createFamily(familyName)">
        <input class="border-2" type="text" v-model="familyName">
        <button class="custom-button" type="submit">創建家庭</button>
      </form>
    </Modal>
  </div>

</template>
<script setup lang="ts">
import { PhPlus } from "@phosphor-icons/vue";
import { computed, inject, onMounted, ref } from "vue";
import request from "@/axios";
import Modal from "@/components/modal.vue";
import HeaderBar from "@/components/headerBar.vue";
import router from "@/router";
import { useFamilyStore } from "@/store/family";
const familyName = ref<string>("");
const familyCode = ref<string>("");
const showMsg: Function = inject("showMsg")!
const familyStroe = useFamilyStore()
onMounted(() => {
  (window as any).receiveQRCodeResult = async (result: string) => {
    await joinFmaily(result)
  };
})

const openCamera = () => {
  if ((window as any).FlutterChannel) {
    (window as any).FlutterChannel.postMessage("openCamera");
  } else {
    console.error("FlutterChannel 不可用！");
  }
}

const joinFmaily = async (code: string) => {
  try{
    await familyStroe.joinFamily(code);
    showMsg("添加成功")
  }catch(err:any){
    showMsg(err)
  }

}

const createFamilyHandler = async (name: string) => {
  await request.post("/family/" + name)
  showMsg("創建成功")
}

</script>
