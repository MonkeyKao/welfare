<template>
  <header class="top-0 left-0 right-0 flex items-center p-5 pt-14 bg-[#92c700]">
    <router-link to="/user" class="forgot-password-link">
      <PhArrowUUpLeft :size="32" color="#000" class="ml-5" />
    </router-link>
    <label class="ml-5 text-3xl font-bold">家庭</label>
  </header>

  <div class="p-4 flex flex-col gap-4">
    <div v-for="family in familes" class="w-full p-3 flex justify-between shadow-md">
      <div>
        <span class="text-H2 font-bold">{{ family.family_name }}</span>
        <div v-for="people in family.users">
          <span>{{ people.user_name }}-</span>
          <span>{{ people.role }}</span>
        </div>
      </div>

      <div class="flex flex-col gap-2 flex-nowrap justify-center">
        <Modal>
          <template v-slot:open-slot>
            <button class="rounded border-2 px-4 border-black">
              產生QRCODE
            </button>
          </template>
          <QRCODE :id="family.family_id" ref="qrCodeRef"></QRCODE>
        </Modal>
        <button  @click="deleteFamily(family.family_id)" class="rounded border-2 px-4 border-black">
          刪除家庭
        </button>
      </div>

    </div>

    <div>
      <form @submit.prevent="joinFmaily(familyCode)">
        <input class="border-2" type="text" v-model="familyCode">
        <button class=" custom-button" type="submit">手動加入</button>
      </form>
    </div>


  </div>

  <div class="fixed bottom-16 right-10">
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
import { PhArrowUUpLeft, PhPlus } from "@phosphor-icons/vue";
import { inject, onMounted, ref } from "vue";
import QRCODE from "@/views/User/QRCODE.vue";
import request from "@/axios";
import Modal from "./modal.vue";
const familyName = ref<string>("");
const familyCode = ref<string>("");

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

const deleteFamily = async (id: number) => {
  try {
    await request.delete("/family/" + id)
    showMsg("刪除成功")
  } catch (err: any) {
    showMsg(err.response.data)
  }
}

const joinFmaily = async (code: string) => {
  try {
    await request.post("/family/join/" + code)
    showMsg("加入成功")
  } catch (err: any) {

    showMsg(err.response.data)
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
