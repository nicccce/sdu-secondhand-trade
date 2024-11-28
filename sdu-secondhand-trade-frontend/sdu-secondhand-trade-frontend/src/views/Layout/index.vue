<script setup>
import LayoutNav from './components/LayoutNav.vue'
import LayoutHeader from './components/LayoutHeader.vue'
import LayoutFooter from './components/LayoutFooter.vue'
</script>

<template>
    <div class="layout-container">
        <LayoutNav />
        <LayoutHeader />
        <RouterView />
        <LayoutFooter />
    </div>
</template>

<style scoped>
.layout-container {
    width: 100vw;
    min-height: 100vh;
  display: flex;
  flex-direction: column;
}
</style>