<template>
  <div ref="colContainer" class="container flex gap-4 flex-col mx-auto p-4">
    <button @click="insertItems(ewlfareitems,0)" class="mt-4 bg-lime-500 text-white p-2 rounded">阿哞篩選</button>
    <!-- insert col -->
  </div>
  
</template>

<script setup lang="ts">
import { ref } from 'vue';

const colContainer = ref<HTMLDivElement>()

let temp: Array<Number> = [0, 0, 0];

//將元素插入到colContainer中，存在的話進createCol函數把內容附加到容器
const insertItems = (items: Array<{ "id": number, "name": string, "image": string }>,num:number) => { 
  if (colContainer.value) {
    
    colContainer.value.appendChild(createCol(items,num));
    
  }
};

//創建消息插入colContainer中
const insertMsg = (message: string) => { //用戶方
  const colDiv = document.createElement("div");
  colDiv.className = 'flex flex-row-reverse gap-4';

  //圖片區塊
 /* const imgdiv = document.createElement("div");
  imgdiv.className = "relative";
  imgdiv.appendChild(createLogo("1.jpg"));*/

  //文字區塊
  const itemDiv = document.createElement('div');
  itemDiv.className = 'bg-lime-300 p-4 rounded-md';
  itemDiv.innerHTML = message

  //圖片及文字區塊組合放進colDiv
  //colDiv.appendChild(imgdiv)
  colDiv.appendChild(itemDiv)  

  if (colContainer.value) {
    colContainer.value.appendChild(colDiv);
  }

  if (colContainer.value) {
    colContainer.value.appendChild(colDiv);
    setTimeout(() => {
      colDiv.scrollIntoView({ behavior: 'smooth', block: 'start' });
    }, 100); // 延遲確保插入完成後再滾動
  }

  return colDiv;
}

const createCol = (items: Array<{ "id": number, "name": string, "image": string }>,num:number) => { //阿哞
  const colDiv = document.createElement("div"); 
  colDiv.className = 'flex items-end'; 

  const imgdiv = document.createElement("div");
  imgdiv.className = "relative";
  imgdiv.appendChild(createLogo("avatar.jpg"));
  colDiv.appendChild(imgdiv);
  console.log('num=',num)
  if(num==1){
    const itemDiv = document.createElement('div');
    itemDiv.className = 'px-3 space-y-3';
 
    items.forEach(item => //for迴圈陣列
    itemDiv.appendChild(cardplace(item.name)),
    colDiv.appendChild(itemDiv)
  );
  }else{
    const itemDiv = document.createElement('div');
    itemDiv.className = 'overflow-x-auto flex space-x-4 ml-4';
 
    items.forEach(item => //for迴圈陣列
    itemDiv.appendChild(createCard(item.image,item.name)),
    colDiv.appendChild(itemDiv)
    );
  } 
   // 滾動到新插入的 colDiv
  if (colContainer.value) {
    colContainer.value.appendChild(colDiv);
    setTimeout(() => {
      colDiv.scrollIntoView({ behavior: 'smooth', block: 'start' });
    }, 300); // 延遲確保插入完成後再滾動
  }

  return colDiv;
}

const checkIndex = (name: string): Array<string | number> => {
  let index = 0;  // 用來記錄當前項目的索引
  let selectIndex = 0;  // 用來記錄找到的項目的索引
  let category = 0;  // 用來記錄找到的項目的類別
  let foundName = "";  // 用來記錄找到的名稱

  // 遍歷 `ewlfareitems` 資料
  ewlfareitems.value.forEach((item) => {
    index++;  
    category = 1; 
    if (item.name === name) { 
      selectIndex = index;  
      foundName = item.name; 
    }
  });

  if (selectIndex != 0) {  
    console.log(selectIndex, category, foundName); 
    return [selectIndex, category, foundName]; 
  }

  // 遍歷 `taiwanitems` 資料
  index = 0;  // 重置索引
  taiwanitems.value.forEach((item) => {
    index++; 
    category = 2; 
    if (item.name === name) { 
      selectIndex = index; 
      foundName = item.name;  
    }
  });

  if (selectIndex != 0) {  // 如果找到了匹配的項目
    console.log(selectIndex, category, foundName);  
    return [selectIndex, category, foundName];  
  }

  // 遍歷 `northitems` 資料
  index = 0;  // 重置索引
  northitems.value.forEach((item) => {
    index++; 
    category = 3; 
    if (item.name === name) {  
      selectIndex = index; 
      foundName = item.name;  
    }
  });

  if (selectIndex != 0) {  
    console.log(selectIndex, category, foundName);  
    return [selectIndex, category, foundName];  
  }

  // 遍歷 `miditems` 資料
  index = 0;  // 重置索引
  miditems.value.forEach((item) => {
    index++;  
    category = 3;  
    if (item.name === name) {  
      selectIndex = index;  
      foundName = item.name;  
    }
  });

  if (selectIndex != 0) { 
    console.log(selectIndex, category, foundName)
    return [selectIndex, category, foundName];  
  }

  // 遍歷 `southitems` 資料
  index = 0;  // 重置索引
  southitems.value.forEach((item) => {
    index++;  
    category = 3;  
    if (item.name === name) { 
      selectIndex = index; 
      foundName = item.name; 
    }
  });

  if (selectIndex != 0) { 
    console.log(selectIndex, category, foundName); 
    return [selectIndex, category, foundName];  
  }

  // 遍歷 `eastitems` 資料
  index = 0;  // 重置索引
  eastitems.value.forEach((item) => {
    index++;  // 每次遍歷時，索引加 1
    category = 3;  // 設定類別為 3，表示這是 `eastitems` 的資料
    if (item.name === name) {  // 如果當前項目的名稱與傳入的 `name` 匹配
      selectIndex = index;  // 記錄當前項目的索引
      foundName = item.name;  // 記錄找到的名稱
    }
  });

  if (selectIndex != 0) {  // 如果找到了匹配的項目
    return [selectIndex, category, foundName];  // 返回找到的索引、類別和名稱
  }

  // 如果沒有找到匹配的項目，返回預設值
  return [0, 0, ""];
};

//卡片格式by地區選擇
const cardplace = (name:string)=>{
  const itemDiv = document.createElement('div');
  itemDiv.className = 'w-36 h-8 border-b shadow-sm flex items-center justify-center'; 

  const nameDiv = document.createElement('div');
  nameDiv.className = "text-H3";
  nameDiv.innerText = name;

  itemDiv.appendChild(nameDiv);  // 將 nameDiv 添加到 itemDiv

   // 添加点击事件
   itemDiv.addEventListener('click', () => {

  let index = checkIndex(name)
  //進全區
  if (index[1] == 1) {
  insertMsg(index[2].toString())
  insertItems(taiwanitems.value,1)
  }   
  //看說要進北中南東哪區
  if (index[1] == 2) {
    insertMsg(index[2].toString())
    if (index[0] == 1) {
      insertItems(northitems.value,1)
    } else if (index[0] == 2) {
      insertItems(miditems.value,1)
    } else if (index[0] == 3) {
      insertItems(southitems.value,1)
    } else if (index[0] == 4) {
      insertItems(eastitems.value,1)
    }
  }
  if (index[1] == 3) {
    console.log("enter",index[2].toString())
    insertMsg(index[2].toString())
  }
  });
  // 将 imgElement 和 nameDiv 添加到 itemDiv 中

  itemDiv.appendChild(nameDiv);
  return itemDiv
  // 返回整個 itemDiv 以便在 createCol 中使用
  
}

//卡片格式by服務選擇
const createCard = (image: string, name: string) => {
  const itemDiv = document.createElement('div');
  itemDiv.className = 'flex-none w-40 h-56 text-H3 bold flex flex-col items-center justify-center rounded';

  // 创建 img 元素并设置属性
  const imgElement = document.createElement('img');
  imgElement.src = image; // 设置图像源
  imgElement.alt = "Item Image"; // 设置替代文本
  imgElement.className = "w-full h-3/4 object-cover rounded-t-md"; // 设置样式类

  const nameDiv = document.createElement('div');
  nameDiv.className = "mt-4 text-center ";
  nameDiv.innerText = name;

  // 添加点击事件
  itemDiv.addEventListener('click', () => {

    let index = checkIndex(name)
    //進全區
    if (index[1] == 1) {
      insertMsg(index[2].toString())
      insertItems(taiwanitems.value,1)
    }   
   
  });
  // 将 imgElement 和 nameDiv 添加到 itemDiv 中
  itemDiv.appendChild(imgElement);
  itemDiv.appendChild(nameDiv);
  return itemDiv
}

//logo
const createLogo = (image: string) => {
  const itemPhoto = document.createElement('div');
  itemPhoto.className = 'flex-none w-10 h-10 bg-blue-500 text-white flex flex-col items-center justify-center rounded';
  itemPhoto.innerHTML = `
        <img src="${image}" class="w-10 h-10 object-cover rounded-t-md" />
      `;
  return itemPhoto;
}

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
  { id: 1, name: '台中市', image: 'login.jpg' },
  { id: 1, name: '苗栗縣', image: 'password.jpg' },
  { id: 1, name: '彰化縣', image: 'login.jpg' },
  { id: 1, name: '南投縣', image: 'password.jpg' },
  { id: 1, name: '雲林縣', image: 'login.jpg' },
])

const southitems = ref([
  { id: 1, name: '高雄市', image: 'login.jpg' },
  { id: 1, name: '台南市', image: 'password.jpg' },
  { id: 1, name: '嘉義市', image: 'login.jpg' },
  { id: 1, name: '嘉義縣', image: 'password.jpg' },
  { id: 1, name: '屏東縣', image: 'login.jpg' },
])

const eastitems = ref([
  { id: 1, name: '花蓮縣', image: 'login.jpg' },
  { id: 1, name: '台東縣', image: 'password.jpg' },
])


</script>
