<script setup>
import LayoutNav from './components/LayoutNav.vue'
import LayoutHeader from './components/LayoutHeader.vue'
import LayoutFooter from './components/LayoutFooter.vue'
import LayoutFixed from './components/LayoutFixed.vue';

</script>

<template>
    <div class="layout-container">
        <LayoutFixed />
        <LayoutNav />
        <LayoutHeader />
        <RouterView />
        <LayoutFooter />
    </div>
    <el-backtop :bottom="100">
    <div
      style="
        height: 100%;
        width: 100%;
        background-color: var(--el-bg-color-overlay);
        box-shadow: var(--el-box-shadow-lighter);
        text-align: center;
        line-height: 40px;
        color: #1989fa;
      "
    >
    ↑淘
    </div>
  </el-backtop>
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