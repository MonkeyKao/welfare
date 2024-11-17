<template>
  <div class=" flex flex-col justify-between">
    
    <div class="flex bg-[#92c700] p-3 items-center fixed z-10 w-full px-6">
      <router-link to="/user" class=""><PhArrowUUpLeft :size="32"  weight="bold"  color="#000"/></router-link>
      <label class="ml-5 text-H2 font-bold">編輯個人資訊</label>
    </div>

    <!-- 头像部分 -->
    <div class="flex bg-white justify-center items-center fixed top-[60px] left-1/2 transform -translate-x-1/2 z-20 w-full px-3">
      <Avatar />
    </div>

    <div class="flex flex-col items-center  mt-[170px] px-3 overflow-auto">

      <form class="flex flex-col w-full h-full p-6 space-y-8">

        <div class=" space-y-2">
          <label class="flex text-H2 font-bold">帳號</label>
          <input class="border-2 w-full p-2 shadow-md outline-none pl-2 text-H3 rounded-md" type="text" v-model="user.account" disabled/>
        </div>

        <div class=" space-y-2">
          <label class="flex text-H2 font-bold">姓名</label>
          <input class="border-2 w-full p-2 shadow-md outline-none pl-2 text-H3 rounded-md" type="text" v-model="user.name" />
        </div>


        <div class=" space-y-2">
          <label class="flex text-H2 font-bold">生日</label>
          <VDatePicker v-model="user.birthday" >
            <template #default="{ inputValue, inputEvents }">
              <input 
                class=" border-2 w-full p-2 shadow-md outline-none pl-2 text-H3 rounded-md"
                :value="inputValue"
                v-on="inputEvents"
                :format="formatDate"  
              />
            </template>
          </VDatePicker>
        </div>

        <div class=" space-y-3">
          <label class="flex text-H2 font-bold">性別</label>
          <div class="flex justify-around">
            <div @click="selectGender(1)" class="flex items-center cursor-pointer">        
              <PhSealCheck v-if="user.female === 1" :size="20" color="#92c700" weight="fill" />
              <PhSeal v-else :size="20" color="#0d0d0c" weight="thin" />
              <label for="male" class="text-H3 ml-2">男性</label>
            </div>

            <div @click="selectGender(2)" class="flex items-center cursor-pointer">
              <PhSealCheck v-if="user.female === 2" :size="20" color="#92c700" weight="fill" />
              <PhSeal v-else :size="20" color="#0d0d0c" weight="thin" />
              <label for="neutral" class="text-H3 ml-2">中性</label>
            </div>

            <div @click="selectGender(3)" class="flex items-center cursor-pointer">
              <PhSealCheck v-if="user.female === 3" :size="20" color="#92c700" weight="fill" />
              <PhSeal v-else :size="20" color="#0d0d0c" weight="thin" />
              <label for="female" class="text-H3 ml-2">女性</label>
            </div>
          </div>
        </div>

        <div class=" space-y-2">
          <label class="flex text-H2 font-bold">地區</label>
          <select v-model="user.location" class="border-2 w-full p-2 shadow-md outline-none pl-2  text-H3 rounded-md bg-white">
            <option :value="item" v-for="item in 19">{{ getTextByLocation(item) }}</option>
          </select>
        </div>
      </form>
    <div class=" w-full px-6 ">
      <router-link to="/user/personal-data">
        <button @click="userStore.saveUserHanlder(JSON.stringify(user))" class="w-full bg-[#92C700] text-white text-H2 py-1 mt-10 rounded shadow-md" type="button">
          確認
        </button>
      </router-link>
    </div>      
    </div>


  </div>
</template>

<script setup lang="ts">
import Avatar from "@/components/Avatar.vue";
import { PhArrowUUpLeft,PhSealCheck,PhSeal } from "@phosphor-icons/vue";
import Datepicker from "@vuepic/vue-datepicker";
import { computed, ref } from "vue";

import "@vuepic/vue-datepicker/dist/main.css";
import { getTextByLocation } from "@/utils/getTextByNumber";
import { useUserStore } from "@/store/userStroe";

import { DatePicker as VDatePicker } from 'v-calendar';

const formatDate = "yyyy-MM-dd";

const userStore = useUserStore()
const user = computed(() => userStore.user)

// 点击选择性别时更新 `user.female`
const selectGender = (value: number) => {
  user.value.female = value; // 设置为选中的性别值
  console.log(user.value.female);
};
</script>

