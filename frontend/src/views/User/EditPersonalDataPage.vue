<template>
  <div class=" flex flex-col justify-between">

    <HeaderBar>編輯個人資訊</HeaderBar>

    <!-- 头像部分 -->
    <div class="flex flex-col bg-white justify-center items-center">
      <!-- 用头像触发文件选择 -->
      <div @click="openImagePicker" class="cursor-pointer">
        <Avatar :src="previewUrl" />
      </div>
    </div>

    <div class="flex flex-col items-center px-3 overflow-auto m-2">
      <form @submit.prevent="updataDataHandler(user)" class="flex flex-col w-full h-full p-6 gap-3">
        <div class=" space-y-2">
          <label class="flex text-H2 font-bold">帳號</label>
          <input class="border-2 w-full p-2 shadow-md outline-none pl-2 text-H3 rounded-md" type="text"
            v-model="user.account" disabled />
        </div>

        <div class=" space-y-2">
          <label class="flex text-H2 font-bold">姓名</label>
          <input class="border-2 w-full p-2 shadow-md outline-none pl-2 text-H3 rounded-md" type="text"
            v-model="user.name" />
        </div>


        <div class=" space-y-2">
          <label class="flex text-H2 font-bold">生日</label>
          <VDatePicker v-model="user.birthday">
            <template #default="{ inputValue, inputEvents }">
              <input class=" border-2 w-full p-2 shadow-md outline-none pl-2 text-H3 rounded-md" :value="inputValue"
                v-on="inputEvents" :format="formatDate" />
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
          <select v-model="user.location"
            class="border-2 w-full p-2 shadow-md outline-none pl-2  text-H3 rounded-md bg-white">
            <option :value="item" v-for="item in 19">{{ getTextByLocation(item) }}</option>
          </select>
        </div>


        <button class="custom-button w-full mt-2" type="submit">
          確認
        </button>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import Avatar from "@/components/Avatar.vue";
import { PhArrowUUpLeft, PhSealCheck, PhSeal } from "@phosphor-icons/vue";
import { computed, onMounted, ref } from "vue";

import "@vuepic/vue-datepicker/dist/main.css";
import { getTextByLocation } from "@/utils/getTextByNumber";
import { useUserStore } from "@/store/userStroe";

import { DatePicker as VDatePicker } from 'v-calendar';
import HeaderBar from "@/components/headerBar.vue";
import Modal from "@/components/modal.vue";
import type User from "@/model/user";
import router from "@/router";
import dayjs from "dayjs";
import request from "@/axios";

const formatDate = "yyyy/MM/dd";

const userStore = useUserStore()
const user = computed(() => userStore.user)

const selectedFile = ref<File | null>(null);
const previewUrl = ref("");
const fileInput = ref();

onMounted(() => {
  (window as any).receiveImage = async (result: string) => {
    // const res = await joinFmaily(result)
    const blob = await (await fetch(`data:image/jpeg;base64,${result}`)).blob();
    const file = new File([blob], "avatar.jpg", { type: blob.type });

    selectedFile.value = file;
    previewUrl.value = URL.createObjectURL(file); // 使用 Blob 生成预览 URL
    
  };
})

const openImagePicker = () => {
  if ((window as any).FlutterChannel) {
    (window as any).FlutterChannel.postMessage("openImagePicker");
  } else {
    console.error("FlutterChannel 不可用！");
  }
}

const uploadAvatar = async () => {
  if (!selectedFile.value) return;
  const formData = new FormData();
  formData.append("avatar", selectedFile.value);

  try {
    const response = await request.post("/users/upload-avatar", formData, {
      headers: { "Content-Type": "multipart/form-data" },
    });
    console.log(response.data);
  } catch (error) {
  }
};

// 点击选择性别时更新 `user.female`
const selectGender = (value: number) => {
  user.value.female = value; // 设置为选中的性别值
};

const updataDataHandler = async (user: User) => {
  user.birthday = dayjs(user.birthday).add(7, 'h').format('YYYY-MM-DD')
  try {
    // 先上傳頭像
    

    await userStore.saveUserHanlder(JSON.stringify(user))
    await uploadAvatar();
    router.push('/user/personal-data')
  } catch (err: any) {

  }

}
</script>
