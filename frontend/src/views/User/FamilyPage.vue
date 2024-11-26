<template>
  <HeaderBar>家庭</HeaderBar>

  <div class="p-4 flex flex-col gap-4">
    <div v-for="family in familes" class="w-full p-3 flex justify-between shadow-md">
      <div>
        <span class="text-H2 font-bold flex-2/3 ">{{ family.family_name }}</span>
      </div>

      <button class="custom-button flex-1/3" @click="router.push('/user/family/'+family.family_id)">
        進入家庭頁面        
      </button>

      <!-- <div class="flex flex-col gap-2 flex-nowrap justify-center">
        <Modal>
          <template v-slot:open-slot>
            <button class="rounded border-2 px-4 border-black">
              產生QRCODE
            </button>
          </template>
          <QRCODE :id="family.family_id" ref="qrCodeRef"></QRCODE>
        </Modal>
        <button @click="deleteFamily(family.family_id)" class="rounded border-2 px-4 border-black">
          刪除家庭
        </button>
      </div> -->

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
      <form @submit.prevent="createFamilyHandler(familyName)">
        <input class="border-2" type="text" v-model="familyName">
        <button class="custom-button" type="submit">創建家庭</button>
      </form>
    </Modal>
  </div>



</template>
<script setup lang="ts">
import { PhPlus } from "@phosphor-icons/vue";
import { inject, onMounted, ref } from "vue";
import QRCODE from "@/components/QRCODE.vue";
import request from "@/axios";
import Modal from "@/components/modal.vue";
import HeaderBar from "@/components/headerBar.vue";
import router from "@/router";
const familyName = ref<string>("");
const familyCode = ref<string>("");

onMounted(() => {
  (window as any).receiveQRCodeResult = async (result: string) => {
    const res = await joinFmaily(result)
  };
})



const openCamera = () => {
  if ((window as any).FlutterChannel) {
    (window as any).FlutterChannel.postMessage("openCamera");
  } else {
    console.error("FlutterChannel 不可用！");
  }
}

class user {
  role: string;
  user_name: string;

  constructor(role: string, user_name: string) {
    this.role = role;
    this.user_name = user_name;
  }
}

const showMsg: Function = inject("showMsg")!

class family {
  family_name: string;
  family_id: number;
  users: Array<user>;

  constructor(family_name: string, family_id: number, users: Array<user>) {
    this.family_name = family_name;
    this.family_id = family_id
    this.users = users;
  }
}

const familes = ref<Array<family>>([])


const joinFmaily = async (code: string) => {
  try {
    await request.post("/family/join/" + code)
    showMsg("加入成功")
  } catch (err: any) {
    showMsg("加入失敗")
  }

}

const createFamilyHandler = async (name: string) => {
  await request.post("/family/" + name)
  showMsg("創建成功")
}

onMounted(async () => {
  const result = await request.get("/family")
  familes.value = JSON.parse(result.data);

})
</script>
