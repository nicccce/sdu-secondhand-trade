<script setup>
import LayoutNav from './components/LayoutNav.vue'
import LayoutHeader from './components/LayoutHeader.vue'
import LayoutFooter from './components/LayoutFooter.vue'
import LayoutFixed from './components/LayoutFixed.vue';
import { useRouter } from 'vue-router';

const router = useRouter()
</script>

<template>
  <div class="layout-container">
    <LayoutFixed />
    <LayoutNav />
    <LayoutHeader />
    <RouterView />
    <LayoutFooter />
  </div>
  <el-affix position="top" :offset="10">
    <el-button type="primary" class="round-button" @click="router.push({path : '/sell'})"><el-icon size="50px"><Sell /></el-icon></el-button>
  </el-affix>
  <el-backtop :bottom="100">
    <div style="
        height: 100%;
        width: 100%;
        background-color: var(--el-bg-color-overlay);
        box-shadow: var(--el-box-shadow-lighter);
        text-align: center;
        line-height: 40px;
        color: #1989fa;
      ">
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

.round-button {
  position: fixed;
  bottom: 70%;
  right: 20px;
  width: 80px;
  height: 80px;
  border-radius: 50%;
  padding: 0;
  font-size: 24px;
  display: flex;
  justify-content: center;
  align-items: center;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 5);

  .el-icon-sell {
    font-size: 30px;
    color: white;
  }
}
</style>