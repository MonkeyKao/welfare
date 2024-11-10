<template>
  <div ref="colContainer" class="container flex gap-4 flex-col mx-auto p-4">
    <button @click="insertItems(ewlfareitems)" class="mt-4 bg-blue-500 text-white p-2 rounded">阿哞篩選</button>
    <!-- insert col -->

  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';


const colContainer = ref<HTMLDivElement>()

let temp: Array<Number> = [0, 0, 0];

const insertItems = (items: Array<{ "id": number, "name": string, "image": string }>) => {
  if (colContainer.value) {
    colContainer.value.appendChild(createCol(items));
  }
};

const insertMsg = (message: string) => {
  const colDiv = document.createElement("div");
  colDiv.className = 'flex flex-row-reverse gap-4';

  const imgdiv = document.createElement("div");
  imgdiv.className = "relative";
  imgdiv.appendChild(createLogo("avatar.jpg"));


  const itemDiv = document.createElement('div');
  itemDiv.className = 'bg-gray-300 p-4 rounded-md';
  itemDiv.innerHTML = message
colDiv.appendChild(imgdiv)
  colDiv.appendChild(itemDiv)

  if (colContainer.value) {
    colContainer.value.appendChild(colDiv);
  }
  return colDiv;
}

const createCol = (items: Array<{ "id": number, "name": string, "image": string }>) => {
  const colDiv = document.createElement("div");
  colDiv.className = 'flex items-end';

  const imgdiv = document.createElement("div");
  imgdiv.className = "relative";
  imgdiv.appendChild(createLogo("avatar.jpg"));
  colDiv.appendChild(imgdiv);

  const itemDiv = document.createElement('div');
  itemDiv.className = 'overflow-x-auto flex space-x-4 ml-4';
  items.forEach(item =>
    itemDiv.appendChild(createCard(item.image, item.name))
  );

  colDiv.appendChild(itemDiv)
  return colDiv;
}

const checkIndex = (name: string): Array<number> => {
  let index = 0;

  let selectIndex = 0;
  let category = 0;



  ewlfareitems.value.forEach((item) => {
    index++;
    category = 1
    if (item.name === name) {
      selectIndex = index
    }
  })

  if (selectIndex != 0) {
    return [selectIndex, category]
  }

  index = 0
  taiwanitems.value.forEach((item) => {
    index++;
    category = 2;
    if (item.name === name) {
      selectIndex = index
    }
  })

  if (selectIndex != 0) {
    return [selectIndex, category]
  }

  index = 0
  northitems.value.forEach((item) => {
    index++;
    category = 3;
    if (item.name === name) {
      selectIndex = index
    }
  })

  if (selectIndex != 0) {
    return [selectIndex, category]
  }

  index = 0
  miditems.value.forEach((item) => {
    index++;
    category = 3;
    if (item.name === name) {
      selectIndex = index
    }
  })

  if (selectIndex != 0) {
    return [selectIndex, category]
  }

  index = 0
  southitems.value.forEach((item) => {
    index++;
    category = 3;
    if (item.name === name) {
      selectIndex = index

    }
  })

  if (selectIndex != 0) {
    return [selectIndex, category]
  }

  index = 0
  eastitems.value.forEach((item) => {
    index++;
    category = 3;
    if (item.name === name) {
      selectIndex = index
    }
  })

  if (selectIndex != 0) {
    return [selectIndex, category]
  }
  return [0, 0]
}

//卡片格式
const createCard = (image: string, name: string) => {
  const itemDiv = document.createElement('div');
  itemDiv.className = 'flex-none w-48 h-64 text-H3 bold flex flex-col items-center justify-center rounded';

  // 创建 img 元素并设置属性
  const imgElement = document.createElement('img');
  imgElement.src = image; // 设置图像源
  imgElement.alt = "Item Image"; // 设置替代文本
  imgElement.className = "w-full h-3/4 object-cover rounded-t-md"; // 设置样式类

  const nameDiv = document.createElement('div');
  nameDiv.className = "mt-4 text-center";
  nameDiv.innerText = name;

  // 添加点击事件
  itemDiv.addEventListener('click', () => {


    let index = checkIndex(name)

    

    if (temp[0] == 0) {
      insertMsg(index.toString())
      insertItems(taiwanitems.value)
    }if (index[1] == 3 && temp[2]==0) {
      insertMsg(index.toString())
    }

    if (index[1] == 2 && temp[1] == 0) {
      insertMsg(index.toString())
      if (index[0] == 1) {
        insertItems(northitems.value)
      } else if (index[0] == 2) {
        insertItems(miditems.value)
      } else if (index[0] == 3) {
        insertItems(southitems.value)
      } else if (index[0] == 4) {
        insertItems(eastitems.value)
      }
    }

    if (index[1] == 3 && temp[2]==0) {
      insertMsg(index.toString())
    }

    temp[index[1] - 1] = index[0];

    


  });

  // 将 imgElement 和 nameDiv 添加到 itemDiv 中
  itemDiv.appendChild(imgElement);
  itemDiv.appendChild(nameDiv);
  return itemDiv
}

//logo
const createLogo = (image: string) => {
  const itemPhoto = document.createElement('div');
  itemPhoto.className = 'flex-none w-12 h-12 bg-blue-500 text-white flex flex-col items-center justify-center rounded';
  itemPhoto.innerHTML = `
        <img src="${image}" class="w-12 h-12 object-cover rounded-t-md" />
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
