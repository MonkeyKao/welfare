<template>
  <div class="p-3 flex flex-col gap-2">
    <div class="flex justify-between">
      <div class="w-full" @click="router.push('/welfare/' + props.data.id)" >
        <p class=" text-[#7F8689] text-sm">{{ data.city }} / {{ data.category ? getTextByService(data.category[0]) : "其他服務" }}
        </p>
        <p  class="text-base font-semibold basis-3/4">{{ data.title }}</p>
      </div>

      <div class="flex flex-row-reverse mr-3">
        <PhHeartStraight :size="30" :weight="isFavorite(data.id) ? 'fill' : 'regular'"
          @click="emit('clickFavorited', props.data);" />
      </div>

    </div>
  </div>
</template>

<script setup lang="ts">
import { getTextByService } from '@/utils/getTextByNumber';
import { PhHeartStraight } from '@phosphor-icons/vue';
import { computed, ref } from 'vue';
import { useFavoriteStore } from '@/store/favorite';
import router from '@/router';

const favoriteStore = useFavoriteStore()
const favoriteData = computed(() => favoriteStore.favorites)

const isFavorited = ref<boolean>(false);

const props = defineProps(["data"])

const emit = defineEmits(["clickFavorited"]);

let isFavorite = (welfareId: number): boolean => {
  if (Array.isArray(favoriteData.value)) {
    return favoriteData.value.some((item) => item.id === welfareId)
  }
  return false
}

const toggleFavorite = () => {
  isFavorited.value = !isFavorited.value;
  emit("clickFavorited", props.data);
}


</script>

<style scoped>
.icon {
  transition: color 0.3s;
}
</style>