<script setup>
import { RouterLink, RouterView } from 'vue-router'
import { useStaticStore } from './stores/static';
import { onMounted } from 'vue';
import { useCategoryStore } from './stores/category';

const staticStore = useStaticStore()
const categoryStore = useCategoryStore()

onMounted(() => {
  staticStore.getCampusList()
  staticStore.getGenderList()
  categoryStore.getCategory()
})


</script>

<template>
  <RouterView />
</template>

<style scoped></style>
