<template>
  <HeaderBar>家庭</HeaderBar>

  <div class="p-4 flex flex-col gap-4">
    <div v-for="family in familyStroe.families" class="w-full p-3 flex justify-between shadow-md rounded-md">
      <div>
        <span class="text-H2 flex-2/3 ">{{ family.familyName }}</span>
      </div>

      <button class="flex-1/3" @click="router.push('/user/family/' + family.familyId)">
        <PhCaretRight :size="28" color="#92c700" />
      </button>

    </div>


  </div>

  <div 
  v-if="addF" 
  class="fixed inset-0 flex items-center justify-center bg-gray-900 bg-opacity-50 z-50"
>
  <!-- 卡片容器 -->
  <div class="bg-white rounded-lg shadow-lg p-5 w-96 mx-6 relative">
    <!-- 關閉按鈕 -->

    <!-- 表單 -->
    <form @submit.prevent="joinFmaily(familyCode)" class="space-y-4 ">
      <div class="grid  grid-cols-5">   
        <span class="text-H2 font-bold text-center col-start-2 col-span-3">加入家庭</span>
        <button class=" grid justify-end col-start-5 " @click="addF = false">        
          <PhX :size="32" color="#10110d " />
        </button>            
      </div>

      <input 
        class="border-2 w-full p-2 rounded-md outline-none" 
        type="text " 
        v-model="familyCode" 
        placeholder="輸入家庭代碼"
      />
      <button 
        class="w-full bg-[#92C700] text-white py-2 rounded-md"
        type="submit"
      >
        手動加入
      </button>
    </form>
    <!-- 分隔線 -->
    <div class="flex items-center my-2">
      <div class="flex-1 border-t"></div>
      <span class="mx-3">or</span>
      <div class="flex-1 border-t"></div>
    </div>
    <!-- 二維碼按鈕 -->
    <div class="flex justify-center">
      <button 
        class="bg-[#73AA00] text-white py-2 px-4 rounded-md hover:bg-blue-600"
        @click="openCamera"
      >
        掃描QR code
      </button>
    </div>
  </div>
</div>



  <div class="fixed bottom-5 right-5 flex flex-col items-end space-y-3">
    <!-- 子按鈕容器 -->
    <div v-if="menuOpen" class="flex flex-col items-center space-y-3 -mt-20">
      <Modal>
        <template v-slot:open-slot>
          <div class="flex items-center flex-row ">
            <h3 class=" bg-white p-1 rounded mx-2">創建家庭</h3>
            <div class="w-12 h-12 bg-white rounded-full flex items-center justify-center shadow-lg hover:bg-[#5A8700] hover:text-white"><PhFolderPlus :size="28"  weight="thin"/>                 </div>    
          </div>
        </template>
        <form @submit.prevent="familyStroe.createFamily(familyName)" class=" flex flex-col space-y-5 ">
          <span class=" text-H2 text-center">創建家庭</span>
          <div class=" flex flex-col space-y-2">
            <span class=" text-H3">名稱</span>
            <input class="border-2 px-2 outline-none" type="text" v-model="familyName">              
          </div>
          <button class="custom-button text-H3 " type="submit">創建</button>
        </form>
      </Modal>     

      
      <div class="flex items-center flex-row">
        <h3 class=" bg-white p-1 rounded mx-2">加入家庭</h3> 
        <!-- 加入家庭按鈕 -->
        <button 
          @click="joinFamily" 
          class="w-12 h-12 bg-white rounded-full flex items-center justify-center shadow-lg hover:bg-[#5A8700] hover:text-white"
        >
          <PhFolderSimplePlus :size="28" weight="thin"/>
          
        </button>
      </div>

    </div>    
    <!-- 主懸浮按鈕 -->
    <button 
      @click="toggleMenu"
      class="w-14 h-14 bg-white text-white rounded-full flex items-center justify-center shadow-lg transform transition-all duration-300 hover:scale-110"
    >
      <PhPlus :size="32" color="#90c700"/>
    </button>
  </div>
</template>
<script setup lang="ts">
import { PhPlus,PhFolderSimplePlus,PhFolderPlus,PhX,PhCaretRight} from "@phosphor-icons/vue";
import { computed, inject, onMounted, ref } from "vue";
import request from "@/axios";
import Modal from "@/components/modal.vue";
import HeaderBar from "@/components/headerBar.vue";
import router from "@/router";
import { useFamilyStore } from "@/store/family";
const familyName = ref<string>("");
const familyCode = ref<string>("");
const showMsg: Function = inject("showMsg")!
const familyStroe = useFamilyStore()
const addF = ref(false);
onMounted(() => {
  (window as any).receiveQRCodeResult = async (result: string) => {
    await joinFmaily(result)
  };
})

const openCamera = () => {
  if ((window as any).FlutterChannel) {
    (window as any).FlutterChannel.postMessage("openCamera");
  } else {
    console.error("FlutterChannel 不可用！");
  }
}

const joinFmaily = async (code: string) => {
  try{
    await familyStroe.joinFamily(code);
    showMsg("添加成功")
  }catch(err:any){
    showMsg(err)
  }

}

const createFamilyHandler = async (name: string) => {
  await request.post("/family/" + name)
  showMsg("創建成功")
}

// 控制菜單開關
const menuOpen = ref(false);

const toggleMenu = () => {
  menuOpen.value = !menuOpen.value;
};

// 點擊事件處理
const createFamily = () => {
  console.log('創建家庭被點擊！');
};

const joinFamily = () => {
  addF.value=true
  console.log('加入家庭被點擊！');
};

</script>
