<template>
    <div class="flex h-20 bg-[#90c700] items-center space-x-2">
    <div @click="router.go(-1)" class="forgot-password-link">
      <PhArrowUUpLeft :size="28" color="#000" class="ml-5" />
    </div>
    <label class="text-H2 font-bold"><slot></slot></label>
  </div>
</template>

<script setup>
import router from "@/router";
import { PhArrowUUpLeft } from "@phosphor-icons/vue";
</script>

<style lang="scss" scoped>

</style>