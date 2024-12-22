<script setup>
//import LayoutHeaderUl from './LayoutHeaderUl.vue'
import { useScroll } from '@vueuse/core'
import { useCategoryStore } from '@/stores/category';

const { y } = useScroll(window)
const categoryStore = useCategoryStore()
</script>

<template>
  <div class="app-header-sticky" :class="{ show: y > 78 }">
    <div class="container">
      <h1 class="logo">
        <RouterLink to="/">淘山大</RouterLink>
      </h1>
      <ul class="app-header-nav">
        <li class="home" v-for="item in categoryStore.categoryList" :key="item.id">
          <RouterLink active-class="active" :to="`/category/${item.id}`">{{ item.name }}</RouterLink>
        </li>
      </ul>
    </div>
  </div>
</template>

<style scoped lang="scss">
.app-header-sticky {
  width: 100%;
  height: 81px;
  position: fixed;
  left: 0;
  top: 0;
  z-index: 999;
  background: $mainColor;
  border-bottom: 1px solid #e4e4e4;
  transform: translateY(-100%);
  opacity: 0;

  &.show {
    transition: all 0.3s linear;
    transform: none;
    opacity: 1;
  }

  .container {
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  .logo {
    height: 80px;
    width: 160px;
    padding-left: 40px;
    display: flex;
    justify-content: center;

    a {
      display: block;
      height: 80px;
      width: 100%;
      text-indent: -9999px;
      background: url('@/assets/images/logo.png') no-repeat center 11px / contain;
      transform: scale(1.3) translateY(10px);
    }
  }

  .app-header-nav,
  .right ul {
    display: flex;
    margin: 0;
    padding: 0;
    list-style: none; // 去除 li 左侧的小点

    li {
      width: 100px;
      height: 80px;
      text-align: center;

      &:hover {
        background-color: $darkColor;
      }

      a {
        font-size: 16px;
        line-height: 80px;
        display: block;
        color: white;

        &.active {

        background-color: $darkColor;

        }
      }
    }
  }

  .right {
    padding-left: 0; // 移除多余的 padding
  }
}
</style>
