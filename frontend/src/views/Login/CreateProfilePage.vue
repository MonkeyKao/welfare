<template>
  <!-- 固定定位的 header，確保圖片維持在最上方 -->

    <div class="text-center">
      <label class="text-H1 font-bold">建立個人檔案</label>
      <Avatar />            
    </div>


    <div class=" flex flex-col mt-3 space-y-8 px-10">
      
      <div class=" space-y-2">
        <label class="flex text-H2 font-bold">姓名</label>
        <input v-model="user.name" class=" border-2 w-full p-2 shadow-md outline-none pl-2 text-H3" type="text" />
      </div>

      <div class=" space-y-2">
        <label class=" flex text-H2 font-bold">生日</label>
        <Datepicker 
          v-model="user.birthday" 
          :format="formatDate" 
          :enable-time-picker="false"
          class="datepicker-custom w-ful" 
        />
      </div>

      <div  class=" space-y-2">
        <label class="flex text-2xl font-bold">性別</label>
        <div class="flex justify-around">
          <div>
            <input v-model="user.female" type="radio" id="female" name="gender" :value=1 />
            <label for="male" class="text-H3 ml-2">男性</label>
          </div>
          <div>
            <input v-model="user.female" type="radio" id="female" name="gender" :value=2 />
            <label for="neutral" class="text-H3 ml-2">中性</label>
          </div>
          <div>
            <input v-model="user.female" type="radio" id="female" name="gender" :value=3 />
            <label for="female" class="text-H3 ml-2">女性</label>
          </div>
        </div>
      </div>

      <div>
        <label class="flex text-2xl font-bold">地區</label>
        <select v-model="user.location" class="flex w-full rounded-md shadow-md p-3 text-2xl">
          <option :value="item" v-for="item in 19">{{ getTextByLocation(item) }}</option>
        </select>
      </div>

      <div>
        <button @click="saveUserHanlder(user)"
          class="w-full bg-[#90c700] text-white font-bold text-2xl py-3 rounded shadow-md" type="button">
          繼續
        </button>
      </div>

      <div class="flex justify-end">
        <RouterLink to="/home" class="flex items-center gap-2">
          <label class="text-xl font-bold">先略過</label>
          <PhArrowCircleRight :size="32" />
        </RouterLink>
      </div>
</div>

</template>

<script setup lang="ts">
import Avatar from "@/components/Avatar.vue";
import User from "@/model/user";
import router from "@/router";
import { useUserStore } from "@/store/userStroe";
import { getTextByLocation } from "@/utils/getTextByNumber";
import { PhArrowCircleRight } from "@phosphor-icons/vue";
import Datepicker from "@vuepic/vue-datepicker";
import "@vuepic/vue-datepicker/dist/main.css";
import { ref } from "vue";
import { RouterLink } from "vue-router";

const formatDate = "yyyy-MM-dd";
const user = ref<User>({
  ID: 0,
  account: "",
  name: "",
  password: "",
  birthday: new Date(),
  female: 1,
  location: 1,
  email: "",
});
const userStore = useUserStore(); 
const saveUserHanlder = async (user: User) => {
  try {
    userStore.saveUserHanlder(JSON.stringify(user))
    router.push("/home");
  } catch (err) {
    console.log(err);
  }
};
</script>

<style>
.datepicker-custom {
  border: none !important;
  outline: none !important;
  box-shadow: none !important;
}
</style>