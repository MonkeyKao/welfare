<template>
  <div class="container mx-auto p-4">
    <div class="overflow-x-auto flex space-x-4">
      <!-- 子項目，顯示圖片和文字 -->
      <div v-for="item in items" :key="item.id"
        class="w-48 h-64 bg-blue-500 text-white flex-shrink-0 flex flex-col items-center justify-center rounded">

        <!-- 圖片顯示：設定3比4比例 -->
        <img :src="item.image" alt="Item Image" class="w-full h-3/4 object-cover rounded-t-md" />
        <!-- 文字顯示 -->
        <div class="mt-2 text-center">
          {{ item.name }}
        </div>
      </div>
    </div>
  </div>
</template>

<script>

export default {
  data() {
  return {
    items: [
      { id: 1, name: "Item 1", image: "../images/testphoto.jpg" },
      { id: 2, name: "Item 2", image: "../images/forgetpassword.jpg" },
      { id: 3, name: "Item 3", image: "../images/forgetpassword.jpg" },
      { id: 4, name: "Item 4", image: "../images/forgetpassword.jpg" },
      { id: 5, name: "Item 5", image: "../images/forgetpassword.jpg" },
      { id: 6, name: "Item 6", image: "../images/forgetpassword.jpg" },
      // 可以繼續添加更多項目
    ],
  };
},  
};

</script>
