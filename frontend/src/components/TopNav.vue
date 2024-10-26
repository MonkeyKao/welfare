<template>
  <div class="flexd overflow-auto">
    <header class="fixed top-0 left-0 right-0 flex justify-between items-center p-5 pt-14" style="background-color: #92C700;">
      <div class="search-box flex items-center bg-gray-50 rounded-lg p-1 border border-gray-300">
        <PhMagnifyingGlass size="24" color="#9ca3af" class="icon-search" />
        <input type="text" placeholder="Ex.租屋補助" class="p-1 text-base bg-transparent border-none outline-none" />
      </div>

      <div class="nav-icons flex space-x-4">
        <router-link to="/favorites" class="icon-link" title="收藏">
          <PhHeartStraight size="32" color="#6b7280" class="icon" />
        </router-link>
        <router-link to="/notifications" class="icon-link" title="通知">
          <PhBell size="32" color="#6b7280" class="icon" />
        </router-link>
      </div>
    </header>
    <div class="fixed top-28 left-0 right-0 bg-white z-10 shadow-md"> <!-- 預留空間以避免與固定的 header 重疊 -->
      <Filter class="filter mt-2 ml-10 flexd" /> 
      <hr>
    </div>
  </div>
</template>

<script>
import { PhHeartStraight, PhBell, PhMagnifyingGlass } from '@phosphor-icons/vue';
import Filter from './Filter.vue';

export default {
  components: {
    PhHeartStraight,
    PhBell,
    PhMagnifyingGlass,
    Filter, // 引入篩選器元件
  },
};
</script>