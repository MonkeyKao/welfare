<template>
  <header class="top-0 left-0 right-0 flex items-center p-5 pt-14 bg-[#92c700]">
    <router-link to="/user" class="forgot-password-link">
      <PhArrowUUpLeft :size="32" color="#000" class="ml-5" />
    </router-link>
    <label class="ml-5 text-3xl font-bold">家庭</label>
  </header>

  <div class="p-4">
    <div v-for="family in familes">
      <span>{{ family.family_name }}</span>
      <div v-for="people in family.users" >
        <div class="flex flex-col shadow-lg">
          <span>職位：{{ people.role }}</span>
          <span>姓名：{{ people.user_name }}</span>
        </div>

      </div>
      <button @click="genrateQRCODE" class="custom-button">產生QRCode及代碼</button>
    </div>

    <div class="w-full p-5 flex flex-col border-2 shadow-lg" v-if="showQRCode">
      <QRCODE></QRCODE>
    </div>

    <div class="flex flex-col w-full p-3 gap-3">
      <button class="custom-button">掃描QRCode</button>
      <button class="custom-button">手動輸入代碼</button>
    </div>
  </div>

  <form @submit.prevent="createFamilyHandler(familyName)">
    <input class="border-2" type="text" v-model="familyName">
    <button type="submit">創建家庭</button>
  </form>

</template>
<script setup lang="ts">
import type User from "@/model/user";
import { useUserStore } from "@/store/userStroe";
import { PhArrowUUpLeft } from "@phosphor-icons/vue";
import { computed, onMounted, ref } from "vue";
import QRCODE from "@/views/User/QRCODE.vue";
import request from "@/axios";
import { result } from "validate.js";
const userStroe = useUserStore()
const familyName = ref<string>("");

class user {
  role: string;
  user_name: string;

  constructor(role: string, user_name: string) {
    this.role = role;
    this.user_name = user_name;
  }
}

class family {
  family_name: string;
  users: Array<user>;

  constructor(family_name: string, users: Array<user>) {
    this.family_name = family_name;
    this.users = users;
  }
}

const familes = ref<Array<family>>([])
const showQRCode = ref<boolean>(false)

const createFamilyHandler = async (name: string) => {
  const result = await request.post("/family/" + name)
}

const genrateQRCODE = () => {
  showQRCode.value = true

}
onMounted(async () => {
  const result = await request.get("/family")
  familes.value = JSON.parse(result.data);
})
</script>
