<script setup>
import LayoutNav from './components/LayoutNav.vue'
import LayoutHeader from './components/LayoutHeader.vue'
import LayoutFooter from './components/LayoutFooter.vue'
import LayoutFixed from './components/LayoutFixed.vue';
import { useCategoryStore } from '@/stores/category';
import { onMounted } from 'vue';

const categoryStore = useCategoryStore()
onMounted(() => {
    categoryStore.getCategory()
})
</script>

<template>
    <div class="layout-container">
        <LayoutFixed />
        <LayoutNav />
        <LayoutHeader />
        <RouterView />
        <LayoutFooter />
    </div>
</template>

<style scoped>
.layout-container {
    width: 100vw;
    max-width: 100%;
    min-height: 100vh;
    display: flex;
    flex-direction: column;
}
</style>