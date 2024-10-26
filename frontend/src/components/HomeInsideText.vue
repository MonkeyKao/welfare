<template>
  <div class="p-2">
    <p class="ml-5 mt-5 text-[#7F8689] text-sm">{{ data.city }} / {{ data.category }}</p>
    <div class="flex justify-between items-center">
      <p class="ml-5 mt-3 text-base font-semibold w-3/4">{{ data.title }}</p>
      <PhHeartStraight
        :size="28"
        :weight="isFavorited ? 'fill' : 'regular'"  
        class="mr-4 icon justify-center" 
        @click="toggleFavorite"
      />
    </div>
  </div>
</template>

<script>
import { PhHeartStraight } from '@phosphor-icons/vue';

export default {
  components: {
    PhHeartStraight,
  },
  props: {
    data: {
      type: Object,
      required: true,
    },
  },
  data() {
    return {
      isFavorited: false,
    };
  },
  methods: {
    toggleFavorite() {
      this.isFavorited = !this.isFavorited;
    },
  },
};
</script>

<style scoped>
.icon {
  transition: color 0.3s;
}
</style>