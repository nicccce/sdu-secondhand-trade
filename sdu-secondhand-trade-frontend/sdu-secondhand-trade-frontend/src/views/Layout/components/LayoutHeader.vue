<script setup>
import { useCategoryStore } from '@/stores/category';

const categoryStore = useCategoryStore()

</script>

<template>
  <header class='app-header'>
    <div class="container">
      <h1 class="logo">
        <RouterLink to="/">淘山大</RouterLink>
      </h1>
      <ul class="app-header-nav">
        <li class="home" v-for="item in categoryStore.categoryList" :key="item.id"
          :class="{ active: $route.path.includes(`/category/${item.id}`) }">
          <RouterLink active-class="active" :to="`/category/${item.id}`">
            {{ item.name }}
          </RouterLink>
        </li>
      </ul>
      <div class="search">
        <i class="iconfont icon-search"></i>
        <input type="text" placeholder="搜一搜">
      </div>
    </div>
  </header>
</template>

<style scoped lang='scss'>
.app-header {
  background: $mainColor;
  height: 100px;

  .container {
    display: flex;
    align-items: center;
  }

  .logo {
    margin-left: 20px;
    height: 100px;
    width: 200px;
    padding-left: 40px;
    display: flex;
    justify-content: center;

    a {
      display: block;
      height: 132px;
      width: 100%;
      text-indent: -9999px;
      background: url('@/assets/images/logo.png') no-repeat center 18px / contain;
      transform: scale(1.3) translateY(10px);
    }
  }

  .app-header-nav {
    width: auto;
    height: 100px;
    display: flex;
    padding-left: 80px;
    position: relative;
    flex-grow: 1;

    li {
      width: 100px;
      height: 100%;
      text-align: center;
      display: flex;
      align-items: center;
      justify-content: center;
      transition: background-color 0.3s ease;
      /* 平滑的背景颜色过渡 */

      &.active {
        background-color: $darkColor;
        /* 激活状态时的背景颜色 */
      }

      &:hover {
        background-color: $darkColor;
      }

      a {
        font-size: 20px;
        line-height: 100px;
        height: 100px;
        display: inline-block;
        color: white;
      }

      .active {
        background-color: $darkColor;
      }
    }
  }

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

  .search {
    width: 250px;
    height: 40px;
    position: relative;
    background-color: rgba(255, 255, 255, 0.6);
    border-radius: 20px;
    display: flex;
    align-items: center;
    padding-left: 10px;

    .icon-search {
      font-size: 18px;
      margin-right: 10px;
    }

    input {
      width: 100%;
      padding: 0 10px;
      color: #474747;
      background: transparent;
      border: none;
      border-radius: 20px;
      font-size: 20px;

      &::placeholder {
        color: #474747;
      }
    }
  }

  .cart {
    width: 50px;

    .curr {
      height: 32px;
      line-height: 32px;
      text-align: center;
      position: relative;
      display: block;

      .icon-cart {
        font-size: 22px;
      }

      em {
        font-style: normal;
        position: absolute;
        right: 0;
        top: 0;
        padding: 1px 6px;
        line-height: 1;
        background: $helpColor;
        color: #fff;
        font-size: 12px;
        border-radius: 10px;
        font-family: Arial;
      }
    }
  }
}
</style>
