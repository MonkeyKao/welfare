<template>

  
    <div class="text-center flex flex-col p-3 fixed z-10 w-full px-6 bg-white">
      <label class="text-H1 font-bold">建立個人檔案</label>
      <Avatar />            
    </div>

    <div class=" flex flex-col space-y-8 px-10 mt-[200px] mb-[20px]">
      
      <div class=" space-y-2">
        <label class="flex text-H2 font-bold">姓名</label>
        <input v-model="user.name" class=" border-2 w-full p-2 shadow-md outline-none pl-2 text-H3 rounded-md" type="text" />
      </div>

      <div class=" space-y-2">
        <label class=" flex text-H2 font-bold">生日</label>      
        <VDatePicker v-model="user.birthday">
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
      <div  class=" space-y-3">
        <label class="flex text-2xl font-bold">性別</label>
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

      <div  class=" space-y-2">
        <label class="flex text-2xl font-bold">地區</label>
        <select v-model="user.location" class=" border-2 w-full p-2 shadow-md outline-none pl-2 text-H3 rounded-md">
          <option :value="item" v-for="item in 19">{{ getTextByLocation(item) }}</option>
        </select>
      </div>

      <div class="flex flex-col pt-5">
        <button @click="saveUserHanlder(user)"
          class="w-full bg-[#92C700] text-white text-H2 py-1 rounded shadow-md" type="button">
          繼續
        </button>
        <router-link to="/home" class="flex text-H3 text-[#92c700] justify-center pt-2">先略過</router-link>          
      </div>
</div>

</template>

<script setup lang="ts">
import Avatar from "@/components/Avatar.vue";
import User from "@/model/user";
import router from "@/router";
import { useUserStore } from "@/store/userStroe";
import { getTextByLocation } from "@/utils/getTextByNumber";
import { PhSealCheck,PhSeal } from "@phosphor-icons/vue";
import { ref, watch  } from "vue";
import { RouterLink } from "vue-router";

import { DatePicker as VDatePicker } from 'v-calendar';
import dayjs from "dayjs";

const formatDate = "yyyy-MM-dd";

const user = ref<User>({
  ID: 0,
  account: "",
  name: "",
  password: "",
  birthday: dayjs().format('YYYY/MM/DD'),
  female: 1,
  location: 1,
  email: "",
  avatar: ""
});

const userStore = useUserStore(); 
const saveUserHanlder = async (user: User) => {
  try {
    user.birthday = dayjs(user.birthday).format("YYYY/MM/DD")
    userStore.saveUserHanlder(JSON.stringify(user))
    router.push("/home");
  } catch (err) {
    console.log(err);
  }
};

// 監聽 `user.birthday` 的變化，並在控制台輸出選擇的日期
watch(() => user.value.birthday, (newDate) => {
  console.log('選擇的生日日期:', newDate);
});

watch(() => user.value.female, (newfemale) => {
  console.log('選擇的生日日期:', newfemale);
});

const selectGender = (value: number) => {
  user.value.female = value; // 设置为选中的性别值
};
</script>

<style>
.datepicker-custom {
  border: none !important;
  outline: none !important;
  box-shadow: none !important;
}
</style>