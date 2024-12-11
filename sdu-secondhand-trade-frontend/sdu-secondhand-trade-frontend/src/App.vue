<script setup>
import { RouterLink, RouterView } from 'vue-router'
import { useStaticStore } from './stores/static';
import { onMounted } from 'vue';

const staticStore = useStaticStore()

onMounted(()=>{
  staticStore.getCampusList()
  staticStore.getGenderList()
})
</script>

<template>
  <RouterView />
</template>

<style scoped>

</style>
