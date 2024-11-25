<template>
    <div>
        {{ route.params.code }}
        123
    </div>
</template>

<script setup lang="ts">
import router from '@/router';
import { useRoute } from 'vue-router';

const route = useRoute()
console.log(route.params.code.toString());

</script>

<style scoped>

</style>