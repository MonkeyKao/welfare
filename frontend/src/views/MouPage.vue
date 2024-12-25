<template>
  <div ref="colContainer" class="container mb-20  flex gap-4 flex-col mx-auto p-4 ">
    <!-- 这里是内容区域 -->
  </div>
  
  <div class="fixed bottom-12 w-full z-20">
    <MouInput @click-send-msg="(msg) => reciveAccountMsg(msg)" />
  </div>
</template>

<script setup lang="ts">
import { mouRequest } from '@/axios';
import MouInput from '@/components/MouInput.vue';
import { AlertColor } from '@/type/ShowMsg';
import { inject, onMounted, ref, watch, nextTick } from 'vue';
const showMsg: Function = inject("showMsg")!

const loading = ref<boolean>(false)

const reciveAccountMsg = async (msg: string) => {
  insertAccountMsg(msg)
  // insertResultTextCard("加載中...")
  insertHtmlContent("<span>加載中<span class='loading loading-spinner text-success'></span></span>")
  try {
    const result = await sendMessageToModel(msg)
    colContainer.value.removeChild(colContainer.value.lastChild as Node)
    insertResultTextCard(result)
  } catch (error) {
    insertResultTextCard("發生錯誤，請重新輸入")
    
  }
}

let chatID: string = "";
const getChatId = async () => {
  const result = await mouRequest.get("application/6236a802-a99f-11ef-86e8-0242ac110002/chat/open")
  chatID = result.data.data
}

const sendMessageToModel = async (message: string): Promise<string> => {
  const result = await mouRequest.post("/application/chat_message/" + chatID, { "message": message, "re_chat": false, "stream": false })
  return result.data.data.content
}

const colContainer = ref<HTMLDivElement>(document.createElement('div'))

//創建用戶消息
const insertAccountMsg = (message: string) => { //用戶
  const colDiv = document.createElement("div");
  colDiv.className = 'flex flex-row-reverse gap-4';

  // 創建頭像div
  const imgdiv = document.createElement("div");
  imgdiv.className = "relative";
  imgdiv.appendChild(createAvatar("avatar.jpg"));

  //文字區塊
  const itemDiv = document.createElement('div');
  itemDiv.className = 'bg-lime-300 p-4 rounded-md';
  itemDiv.innerHTML = message

  //圖片及文字區塊組合放進colDiv
  colDiv.appendChild(itemDiv)
  colContainer.value.appendChild(colDiv);

  colContainer.value.appendChild(colDiv);
  setTimeout(() => {
    colDiv.scrollIntoView({ behavior: 'smooth', block: 'start' });
  }, 100); // 延遲確保插入完成後再滾動

  return colDiv;
}

// 創建橫排服務選擇卡片
const insertServiceCard = (items: Array<{ "id": number, "name": string, "image": string }>) => {
  const colDiv = document.createElement("div");
  colDiv.className = 'flex items-end';

  // 創建頭像div
  const imgdiv = document.createElement("div");
  imgdiv.className = "relative";
  imgdiv.appendChild(createAvatar("avatar.jpg"));
  colDiv.appendChild(imgdiv);

  const itemDiv = document.createElement('div');
  itemDiv.className = 'overflow-x-auto snap-x flex space-x-4 ml-4';

  items.forEach(item => //for迴圈陣列
    itemDiv.appendChild(createServiceCard(item.image, item.name)),
    colDiv.appendChild(itemDiv)
  );

  colContainer.value.appendChild(colDiv);
  setTimeout(() => {
    colDiv.scrollIntoView({ behavior: 'smooth', block: 'start' });
  }, 300); // 延遲確保插入完成後再滾動
}

// 創建豎排地區選擇卡片 
const insertPlaceCard = (items: Array<{ "id": number, "name": string, "image": string }>) => { //阿哞
  const colDiv = document.createElement("div");
  colDiv.className = 'flex items-end';

  const imgdiv = document.createElement("div");
  imgdiv.className = "relative";
  imgdiv.appendChild(createAvatar("avatar.jpg"));
  colDiv.appendChild(imgdiv);

  const itemDiv = document.createElement('div');
  itemDiv.className = 'px-3 space-y-3';

  items.forEach(item => //for迴圈陣列
    itemDiv.appendChild(createPlaceCard(item.name)),
    colDiv.appendChild(itemDiv)
  );

  // 滾動到新插入的 colDiv
  if (colContainer.value) {
    colContainer.value.appendChild(colDiv);
    setTimeout(() => {
      colDiv.scrollIntoView({ behavior: 'smooth', block: 'start' });
    }, 300); // 延遲確保插入完成後再滾動
  }

  return colDiv;
}

/**
 * 根據傳入參數，找到點擊對應之id
 * @param name 傳入參數
 * 
 * @returns [返回卡片之id,返回點擊之層級,返回點擊之name，返回選取之福利id]
 */
let selectedService: number = 0;
const checkIndex = (name: string): Array<string | number> => {
  const items = [
    { data: ewlfareitems.value, index: 1 },
    { data: taiwanitems.value, index: 2 },
    { data: northitems.value, index: 3 },
    { data: miditems.value, index: 3 },
    { data: southitems.value, index: 3 },
    { data: eastitems.value, index: 3 },
  ];

  for (let { data, index } of items) {
    const foundItem = data.find(item => item.name === name);
    if (foundItem) {
      if (index == 1) {
        selectedService = foundItem.id
      }
      return [foundItem.id, index, foundItem.name, selectedService];
    }
  }

  // 如果未找到匹配的項目，回傳空陣列
  return [];

};

//創建地區單個卡片
const createPlaceCard = (name: string) => {
  const itemDiv = document.createElement('div');
  itemDiv.className = 'w-36 h-8 border-b shadow-sm flex items-center justify-center';
  itemDiv.addEventListener('click', (click) => clickPlaceHandler(name));

  const nameDiv = document.createElement('div');
  nameDiv.className = "text-H3";
  nameDiv.innerText = name;

  itemDiv.appendChild(nameDiv);  // 將 nameDiv 添加到 itemDiv
  return itemDiv
  // 返回整個 itemDiv 以便在 insertPlaceCard 中使用

}

// 點擊地區卡片處理
const clickPlaceHandler = (name: string) => {
  let index = checkIndex(name)
  //進全區
  if (index[1] == 1) {
    insertAccountMsg(index[2].toString())
    insertPlaceCard(taiwanitems.value)
  }
  //看說要進北中南東哪區
  if (index[1] == 2) {
    insertAccountMsg(index[2].toString())
    if (index[0] == 1) {
      insertPlaceCard(northitems.value)
    } else if (index[0] == 2) {
      insertPlaceCard(miditems.value)
    } else if (index[0] == 3) {
      insertPlaceCard(southitems.value)
    } else if (index[0] == 4) {
      insertPlaceCard(eastitems.value)
    }
  }
  if (index[1] == 3) {
    insertAccountMsg(index[2].toString())
    ResultInfHandler(index)
  }
}

// 最終資料處理方法
const ResultInfHandler = (input: Array<string | number>) => {
  // welfareStroe.getWelfare([1, 2, 3, 4, 5], [1, 2, 3, 4, 5])
  // const result = welfareStroe.getWelfare([Number(input[0])], [Number(input[3])])
  // if (result.length > 0) {
  //   insertResultInfCard(result)
  // } else {
  //   insertResultInfCard([{ title: "未找到相關福利\n點擊返回主界面", url: "home" }])
  // }
}

const insertResultTextCard = (text:string) => {
  const colDiv = document.createElement("div");
  colDiv.className = 'flex items-end';

  const imgdiv = document.createElement("div");
  imgdiv.className = "relative";
  imgdiv.appendChild(createAvatar("avatar.jpg"));
  colDiv.appendChild(imgdiv);

  const itemDiv = document.createElement('div');
  itemDiv.className = 'px-3 space-y-3';
  itemDiv.appendChild(createResultTextCard(text))
  colDiv.appendChild(itemDiv)

  colContainer.value.appendChild(colDiv);
  setTimeout(() => {
    colDiv.scrollIntoView({ behavior: 'smooth', block: 'start' });
  }, 300); // 延遲確保插入完成後再滾動
}

const createResultTextCard = (text:string) => {
  const itemDiv = document.createElement('div');
  itemDiv.className = 'border-b shadow-sm flex items-center justify-center';

  const nameDiv = document.createElement('div');
  nameDiv.className = "text-H3 ";
  nameDiv.innerText = text;

  itemDiv.appendChild(nameDiv);  // 將 nameDiv 添加到 itemDiv
  return itemDiv
}

const insertHtmlContent = (htmlContent: string) => {
  const colDiv = document.createElement("div");
  colDiv.className = 'flex items-end';

  const imgdiv = document.createElement("div");
  imgdiv.className = "relative";
  imgdiv.appendChild(createAvatar("avatar.jpg"));
  colDiv.appendChild(imgdiv);

  const itemDiv = document.createElement('div');
  itemDiv.className = 'px-3 space-y-3';
  itemDiv.innerHTML = htmlContent;
  colDiv.appendChild(itemDiv);

  colContainer.value.appendChild(colDiv);
  setTimeout(() => {
    colDiv.scrollIntoView({ behavior: 'smooth', block: 'start' });
  }, 300); // 延遲確保插入完成後再滾動
}

// 插入最終篩選資料
const insertResultInfCard = (items: Array<{ title: string, url: string }>) => {
  const colDiv = document.createElement("div");
  colDiv.className = 'flex items-end';

  const imgdiv = document.createElement("div");
  imgdiv.className = "relative";
  imgdiv.appendChild(createAvatar("avatar.jpg"));
  colDiv.appendChild(imgdiv);

  const itemDiv = document.createElement('div');
  itemDiv.className = 'px-3 space-y-3';

  items.forEach(item => //for迴圈陣列
    itemDiv.appendChild(createResultInfCard(item.title, item.url)),
    colDiv.appendChild(itemDiv)
  );


  colContainer.value.appendChild(colDiv);
  setTimeout(() => {
    colDiv.scrollIntoView({ behavior: 'smooth', block: 'start' });
  }, 300); // 延遲確保插入完成後再滾動
}

const createResultInfCard = (title: string, url: string) => {
  const itemDiv = document.createElement('div');
  itemDiv.className = 'border-b shadow-sm flex items-center justify-center';
  itemDiv.addEventListener('click', () => {
    try {
      // 檢查URL是否合法
      const validUrl = new URL(url);
      window.open(validUrl.toString(), '_blank');
    } catch (error) {
      showMsg("該頁面路徑有誤，請自行搜尋")
    }

    //  window.open(url, '_blank');
  });

  const nameDiv = document.createElement('div');
  nameDiv.className = "text-H3 ";
  nameDiv.innerText = title;

  itemDiv.appendChild(nameDiv);  // 將 nameDiv 添加到 itemDiv
  return itemDiv
}

//創建服務單個卡片
const createServiceCard = (image: string, name: string) => {
  const itemDiv = document.createElement('div');
  itemDiv.className = 'flex-none snap-start w-40 h-56 text-H3 bold flex flex-col items-center justify-center rounded';

  // 创建 img 元素并设置属性
  const imgElement = document.createElement('img');
  imgElement.src = image; // 设置图像源
  imgElement.alt = "Item Image"; // 设置替代文本
  imgElement.className = "w-full h-3/4 object-cover rounded-t-md"; // 设置样式类

  const nameDiv = document.createElement('div');
  nameDiv.className = "mt-4 text-center ";
  nameDiv.innerText = name;

  // 添加点击事件
  itemDiv.addEventListener('click', () => clickServiceHandler(name));
  // 将 imgElement 和 nameDiv 添加到 itemDiv 中
  itemDiv.appendChild(imgElement);
  itemDiv.appendChild(nameDiv);
  return itemDiv
}

// 點擊服務卡片處理
const clickServiceHandler = (name: string) => {
  let index = checkIndex(name)

  //進全區
  if (index[1] == 1) {
    insertAccountMsg(index[2].toString())
    insertPlaceCard(taiwanitems.value)
  }
}

//創建頭像div
const createAvatar = (image: string) => {
  const AvatarDiv = document.createElement('div');
  AvatarDiv.className = 'flex-none w-10 h-10 bg-blue-500 text-white flex flex-col items-center justify-center rounded';
  AvatarDiv.innerHTML = `
        <img src="${image}" class="w-10 h-10 object-cover rounded-t-md" />
      `;
  return AvatarDiv;
}

onMounted(() => {
  insertServiceCard(ewlfareitems.value)
  getChatId();
})

const ewlfareitems = ref([
  { id: 1, name: '家庭與育兒福利', image: 'login.jpg' },
  { id: 2, name: '教育福利', image: 'password.jpg' },
  { id: 3, name: '健康與醫療福利', image: 'login.jpg' },
  { id: 4, name: '老人與退休福利', image: 'password.jpg' },
  { id: 5, name: '低收入戶與弱勢族群', image: 'login.jpg' },
  { id: 6, name: '殘疾與特殊需求福利', image: 'password.jpg' },
  { id: 7, name: '就業與創業福利', image: 'login.jpg' },
  { id: 8, name: '社會安全與基本生活支援', image: 'password.jpg' },
  { id: 9, name: '兒童及少年福利', image: 'login.jpg' },
  { id: 10, name: '其他特定族群福利', image: 'password.jpg' },
  // 添加更多項目
]);

const taiwanitems = ref([
  { id: 1, name: '北區', image: 'login.jpg' },
  { id: 2, name: '中區', image: 'password.jpg' },
  { id: 3, name: '南區', image: 'login.jpg' },
  { id: 4, name: '東區', image: 'password.jpg' },
])

const northitems = ref([
  { id: 1, name: '台北市', image: 'login.jpg' },
  { id: 2, name: '新北市', image: 'password.jpg' },
  { id: 3, name: '基隆市', image: 'login.jpg' },
  { id: 4, name: '桃園市', image: 'password.jpg' },
  { id: 5, name: '宜蘭縣', image: 'login.jpg' },
  { id: 6, name: '新竹縣', image: 'password.jpg' },
  { id: 7, name: '新竹市', image: 'login.jpg' },
])

const miditems = ref([
  { id: 8, name: '台中市', image: 'login.jpg' },
  { id: 9, name: '苗栗縣', image: 'password.jpg' },
  { id: 10, name: '彰化縣', image: 'login.jpg' },
  { id: 11, name: '南投縣', image: 'password.jpg' },
  { id: 12, name: '雲林縣', image: 'login.jpg' },
])

const southitems = ref([
  { id: 13, name: '高雄市', image: 'login.jpg' },
  { id: 14, name: '台南市', image: 'password.jpg' },
  { id: 15, name: '嘉義市', image: 'login.jpg' },
  { id: 16, name: '嘉義縣', image: 'password.jpg' },
  { id: 17, name: '屏東縣', image: 'login.jpg' },
])

const eastitems = ref([
  { id: 18, name: '花蓮縣', image: 'login.jpg' },
  { id: 19, name: '台東縣', image: 'password.jpg' },
])


</script>
