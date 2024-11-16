<template>
  <header class="top-0 left-0 right-0 flex items-center p-5 pt-14 bg-[#92c700]">
    <router-link to="/user" class="forgot-password-link">
      <PhArrowUUpLeft :size="32" color="#000" class="ml-5" />
    </router-link>
    <label class="ml-5 text-3xl font-bold">家庭</label>
  </header>

  <div class="p-4">
    <div v-for="people in family">
      <span>{{ people.name }}</span>
    </div>

    <div class="w-full p-5 flex flex-col border-2 shadow-lg" v-if="showQRCode">
      <QRCODE></QRCODE>
    </div>

    <div class="flex flex-col w-full p-3 gap-3">
      <button @click="genrateQRCODE" class="custom-button">產生QRCode及代碼</button>
      <button class="custom-button">掃描QRCode</button>
      <button class="custom-button">手動輸入代碼</button>
    </div>
  </div>
</template>
<script setup lang="ts">
import type User from "@/model/user";
import { useUserStore } from "@/store/userStroe";
import { PhArrowUUpLeft } from "@phosphor-icons/vue";
import { computed, onMounted, ref } from "vue";
import QRCODE from "@/views/User/QRCODE.vue";
const userStroe = useUserStore()

const family = ref<Array<User>>([])
const showQRCode = ref<boolean>(false)


const genrateQRCODE = () => {
  showQRCode.value = true
  
}
onMounted(() => {
  family.value.push(userStroe.user)
})
</script>
