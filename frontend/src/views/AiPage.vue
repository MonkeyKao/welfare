<template>
  <div class="container mx-auto p-4">
    <button @click="insertItems" class="mt-4 bg-blue-500 text-white p-2 rounded">插入項目</button>
    <div class="flex items-end"> 
      <div class="relative">
        <!-- 用於插入不受滾動條影響的 img -->
        <div ref="logoContainer"></div>
      </div>
      <div class="overflow-x-auto flex space-x-4 ml-4" ref="itemsContainer">
        <!-- 用於插入卡片的容器 -->
      </div>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue';

export default {
  setup() {
    const itemsContainer = ref(null); // 用於獲取容器的 DOM 引用
    const logoContainer = ref(null);
    const ewlfareitems = ref([
      { id: 1, name: '家庭與育兒福利', image: 'https://via.placeholder.com/150' },
      { id: 2, name: '教育福利', image: 'https://via.placeholder.com/150' },
      { id: 3, name: '健康與醫療福利', image: 'https://via.placeholder.com/150' },
      { id: 4, name: '老人與退休福利', image: 'https://via.placeholder.com/150' },
      { id: 5, name: '低收入戶與弱勢族群', image: 'https://via.placeholder.com/150' },
      { id: 6, name: '殘疾與特殊需求福利', image: 'https://via.placeholder.com/150' },
      { id: 7, name: '就業與創業福利', image: 'https://via.placeholder.com/150' },
      { id: 8, name: '社會安全與基本生活支援', image: 'https://via.placeholder.com/150' },
      { id: 9, name: '兒童及少年福利', image: 'https://via.placeholder.com/150' },
      { id: 10, name: '其他特定族群福利', image: 'https://via.placeholder.com/150' },
      // 添加更多項目
    ]);

    const taiwanitems = ref([
      { id: 1, name: '北區', image: 'https://via.placeholder.com/150' },
      { id: 2, name: '中區', image: 'https://via.placeholder.com/150' },
      { id: 3, name: '南區', image: 'https://via.placeholder.com/150' },
      { id: 4, name: '東區', image: 'https://via.placeholder.com/150' },
    ])

    const northitems = ref([
      { id: 1, name: '台北市', image: 'https://via.placeholder.com/150' },
      { id: 2, name: '新北市', image: 'https://via.placeholder.com/150' },
      { id: 3, name: '基隆市', image: 'https://via.placeholder.com/150' },
      { id: 4, name: '桃園市', image: 'https://via.placeholder.com/150' },
      { id: 5, name: '宜蘭縣', image: 'https://via.placeholder.com/150' },
      { id: 6, name: '新竹縣', image: 'https://via.placeholder.com/150' },
      { id: 7, name: '新竹市', image: 'https://via.placeholder.com/150' },
    ])

    const miditems = ref([
      { id: 1, name: '台中市', image: 'https://via.placeholder.com/150' },
      { id: 1, name: '苗栗縣', image: 'https://via.placeholder.com/150' },
      { id: 1, name: '彰化縣', image: 'https://via.placeholder.com/150' },
      { id: 1, name: '南投縣', image: 'https://via.placeholder.com/150' },
      { id: 1, name: '雲林縣', image: 'https://via.placeholder.com/150' },
    ])

    const southitems = ref([
      { id: 1, name: '高雄市', image: 'https://via.placeholder.com/150' },
      { id: 1, name: '台南市', image: 'https://via.placeholder.com/150' },
      { id: 1, name: '嘉義市', image: 'https://via.placeholder.com/150' },
      { id: 1, name: '嘉義縣', image: 'https://via.placeholder.com/150' },
      { id: 1, name: '屏東縣', image: 'https://via.placeholder.com/150' },
    ])

    const eastitems = ref([
      { id: 1, name: '花蓮縣', image: 'https://via.placeholder.com/150' },
      { id: 1, name: '台東縣', image: 'https://via.placeholder.com/150' },
    ])


    const insertItems = () => {
      // 確保 logoContainer 只插入一次 avatar
      if (logoContainer.value && logoContainer.value.children.length === 0) {
        logoContainer.value.appendChild(createLogo());
      }

      // 遍歷 ewlfareitems 並插入卡片
      if (itemsContainer.value) {
        ewlfareitems.value.forEach(item => {
          itemsContainer.value.appendChild(createCard(item.image, item.name));
        });
      }
    };
    //卡片格式
    const createCard = (image, name) => {
      const itemDiv = document.createElement('div');
      itemDiv.className = 'flex-none w-48 h-64 bg-blue-500 text-white flex flex-col items-center justify-center rounded';
      itemDiv.innerHTML = `
        <img :src="${image}" alt="Item Image" class="w-full h-3/4 object-cover rounded-t-md" />
        <div class="mt-4 text-center ">${name}</div>
  
          `;
      return itemDiv;
    }

    const createLogo = () => {
      const itemPhoto = document.createElement('div');
      itemPhoto.className = 'flex-none w-12 h-12 bg-blue-500 text-white flex flex-col items-center justify-center rounded';
      itemPhoto.innerHTML = `
        <img src="../images/avatar.jpg" class="w-12 h-12 object-cover rounded-t-md" />
      `;
      return itemPhoto;
    }
    return {
      itemsContainer,
      logoContainer,
      insertItems
    };
  }
}
</script>
