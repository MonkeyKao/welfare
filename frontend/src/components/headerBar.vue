<template>
  <div class="flex h-20 bg-[#90c700] items-center justify-between p-4">
    <div class="flex gap-2">
      <div @click="router.go(-1)" class="forgot-password-link flex items-center">
        <PhArrowUUpLeft :size="28" color="#000" />
      </div>
      <label class="text-H2 font-bold">
        <slot></slot>
      </label>
    </div>

    <div>
      <slot name="siderBar"></slot>
    </div>

  </div>

</template>

<script setup lang="ts">
import router from "@/router";
import { PhArrowUUpLeft } from "@phosphor-icons/vue";
</script>

<style lang="scss" scoped></style>