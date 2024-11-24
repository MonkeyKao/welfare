<template>
  <div class="flex flex-col h-screen justify-between">
    <div>
      <div>
        <HeaderBar>個人資訊</HeaderBar>
        <div class="flex flex-col justify-center items-center space-y-2">
          <Avatar />          
        </div>
      </div>
      <div class="flex flex-col space-y-5 m-5 px-3">
        <div class="flex space-x-2">
          <label class="flex text-H2 font-bold">姓名</label>
          <label class="flex text-H2 font-bold">:</label>
          <label class="text-2xl ">{{ user?.name }}</label>
        </div>
        <div class="flex space-x-2">
          <label class="flex text-H2 font-bold">帳號</label>
          <label class="flex text-H2 font-bold">:</label>
          <label class="text-2xl ">{{ user?.account }}</label>
        </div>
        <div class="flex space-x-2">
          <label class="flex text-H2 font-bold">生日</label>
          <label class="flex text-H2 font-bold">:</label>
          <label class="text-2xl ">{{ user?.birthday }}</label>
        </div>
        <div class="flex space-x-2">
          <label class="flex text-H2 font-bold">性別</label>
          <label class="flex text-H2 font-bold">:</label>
          <label class="text-2xl ">{{ getTextByGender(user.female) }}</label>
        </div>
        <div class="flex space-x-2">
          <label class="flex text-H2 font-bold">地區</label>
          <label class="flex text-H2 font-bold">:</label>
          <label class="text-2xl ">{{ getTextByLocation(user.location) }}</label>
        </div>
      </div>
    </div>

    <div class="  flex justify-center space-x-5 m-10">

      <router-link to="/user/edit-personal-data"><button
          class="bg-[#90c700] text-white text-H3 p-2 px-5 rounded shadow-md w-full flex items-center space-x-5 " type="button">
          <PhPencilLine :size="24" color="#fcfcfc" />
          編輯
      </button></router-link> 

      <div>
        <router-link v-if="user.ID" to="/account/login" class="forgot-password-link">
          <button
            class="bg-[#90c700] text-white text-H3 p-2 px-5 rounded shadow-md w-full flex items-center space-x-5" type="button"
            @click="logoutHandler">
           <PhSignOut :size="24" color="#fcfcfc" />            
            登出
          </button>
        </router-link>
      </div>              
    </div>
    


  </div>
</template>
<script setup lang="ts">
import Avatar from "@/components/Avatar.vue";
import User from "@/model/user";
import { useUserStore } from "@/store/userStroe";
import { getTextByGender, getTextByLocation } from "@/utils/getTextByNumber";
import { PhArrowUUpLeft } from "@phosphor-icons/vue";
import { computed, onMounted, ref } from "vue";
import {
  PhSignOut,
  PhPencilLine,
} from "@phosphor-icons/vue";
import router from "@/router";
import { useFavoriteStore } from "@/store/favorite";
import HeaderBar from "@/components/headerBar.vue";
const userStore = useUserStore()
const favoriteStore = useFavoriteStore()
const user = computed(() => userStore.user)

const userStroe = useUserStore();

const logoutHandler = () => {
  localStorage.removeItem("token");
  userStroe.user = new User();
  favoriteStore.favorites = [] 
  router.push("/account/login");
};

</script>
