<template>
  <!-- 固定定位的 header，確保圖片維持在最上方 -->
  <header class="flex flex-col gap-3 p-4 basis-1">
    <div class="text-center">
      <label class="text-4xl font-bold">建立個人檔案</label>
    </div>
  </header>

  <body>
    <Avatar />
    <form class="flex flex-col w-full h-full p-10 gap-5">
      <div>
        <label class="flex text-2xl font-bold">姓名</label>
        <input
          v-model="user.name"
          class="flex w-full border-2 rounded-md shadow-md p-3"
          type="text"
        />
      </div>
      <div>
        <label class="flex text-2xl font-bold">生日</label>
        <Datepicker
          v-model="birthday"
          :format="formatDate"
          :enable-time-picker="false"
          class="w-full rounded-md shadow-md"
        />
      </div>
      <div>
        <label class="flex text-2xl font-bold">性別</label>
        <div class="flex justify-around">
          <div>
            <input v-model="user.female" type="radio" id="male" name="gender" />
            <label for="male" class="text-2xl ml-2">男性</label>
          </div>
          <div>
            <input
              v-model="user.female"
              type="radio"
              id="female"
              name="gender"
            />
            <label for="neutral" class="text-2xl ml-2">中性</label>
          </div>
          <div>
            <input
              v-model="user.female"
              type="radio"
              id="female"
              name="gender"
            />
            <label for="female" class="text-2xl ml-2">女性</label>
          </div>
        </div>
      </div>
      <div>
        <label class="flex text-2xl font-bold">地區</label>
        <select
          v-model="user.location"
          class="flex w-full border-2 rounded-md shadow-md p-3 text-2xl"
        >
          <option>台北市</option>
          <option>新北市</option>
          <option>基隆市</option>
          <option>新竹市</option>
          <option>新竹縣</option>
          <option>苗栗縣</option>
          <option>台中市</option>
          <option>彰化縣</option>
          <option>雲林縣</option>
          <option>嘉義市</option>
          <option>嘉義縣</option>
          <option>台南市</option>
          <option>高雄市</option>
          <option>屏東縣</option>
          <option>南投縣</option>
          <option>宜蘭縣</option>
          <option>花蓮縣</option>
          <option>台東縣</option>
        </select>
      </div>

      <div>
        <button
          @click="saveUserHanlder(user)"
          class="w-full bg-[#90c700] text-white font-bold text-2xl py-3 rounded shadow-md"
          type="button"
        >
          繼續
        </button>
      </div>

      <div class="flex justify-end">
        <RouterLink to="/home" class="flex items-center gap-2">
          <label class="text-xl font-bold">先略過</label>
          <PhArrowCircleRight :size="32" />
        </RouterLink>
      </div>
    </form>
  </body>
</template>

<script setup lang="ts">
import request from "@/axios";
import Avatar from "@/components/Avatar.vue";
import User from "@/model/user";
import router from "@/router";
import { PhArrowCircleRight } from "@phosphor-icons/vue";
import Datepicker from "@vuepic/vue-datepicker";
import "@vuepic/vue-datepicker/dist/main.css";
import { ref } from "vue";
import { RouterLink } from "vue-router";

const birthday = ref<Date | undefined>(undefined);
const formatDate = "yyyy-MM-dd";
const user = ref<User>({
  _id: 0,
  account: "",
  name: "",
  password: "",
  birthday: new Date(),
  female: 0,
  location: 0,
  email: "",
});

const saveUserHanlder = async (user: User) => {
  try {
    const result = await request.post(
      "/users/updateuser",
      JSON.stringify(user)
    );
    router.push("/home");
  } catch (err) {
    console.log(123);
  }
};
</script>
