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
                    <RouterLink to="/">{{ item.name }}</RouterLink>
                </li>
            </ul>
            <div class="right">
                <li>
                    <RouterLink to="/">品牌</RouterLink>
                </li>
                <li>
                    <RouterLink to="/">专题</RouterLink>
                </li>
            </div>
        </div>
    </div>
</template>
<style scoped lang='scss'>
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

    .app-header-nav {
        width: auto;
        height: 80px;
        display: flex;
        padding-left: 90px;
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
                font-size: 16px;
                line-height: 80px;
                height: 80px;
                display: inline-block;
                color: white;
            }

            .active {
                color: $mainColor;
                border-bottom: 1px solid $mainColor;
            }
        }
    }

    .right {
        width: 220px;
        display: flex;
        text-align: center;
        padding-left: 10px;
        border-left: 10px solid $darkColor;
        text-align: center;

        li {
            width: 100px;
            height: 100%;
            text-align: center;

            &:hover {
                background-color: $darkColor;
            }


            a {
                font-size: 16px;
                line-height: 80px;
                height: 80px;
                display: inline-block;
                color: white;
            }

            .active {
                color: $mainColor;
                border-bottom: 1px solid $mainColor;
            }
        }
    }
}
</style>