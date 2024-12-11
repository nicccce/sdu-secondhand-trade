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
        <li class="home" v-for="item in categoryStore.categoryList" :key="item.id">
          <RouterLink to="/">{{ item.name }}</RouterLink>
        </li>
      </ul>
      <div class="search">
        <i class="iconfont icon-search"></i>
        <input type="text" placeholder="搜一搜">
      </div>
      <!-- 头部购物车 -->

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
        color: $mainColor;
        border-bottom: 1px solid $mainColor;
      }
    }
  }

  .search {
    width: 250px;
    height: 40px;
    position: relative;
    background-color: rgba(255, 255, 255, 0.7);
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
        /* 设置 placeholder 颜色为灰色 */
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