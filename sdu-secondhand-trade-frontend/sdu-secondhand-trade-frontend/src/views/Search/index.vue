<script setup>
import { getSubCategoryAPI } from '@/apis/category';
import { useCategoryStore } from '@/stores/category';
import { computed, onMounted, ref } from 'vue';
import { onBeforeRouteUpdate, useRoute } from 'vue-router';
import GoodItem from '../Category/components/GoodItem.vue';
import { searchGoodAPI } from '../../apis/good';

const route = useRoute();
//获取商品数据
const goodList = ref([])
const getGoodList = async () => {
    const res = await searchGoodAPI(route.params.name)
    goodList.value = res.data
}

onBeforeRouteUpdate(async (to) => {
    // 获取当前路由对应的分类数据
    const res = await searchGoodAPI(to.params.name)
    goodList.value = res.data
});

onMounted(() => {
    getGoodList()
})

</script>

<template>
    <div class="container" v-if="goodList.length">
        <!-- 面包屑 -->
        <div class="bread-container">
            <el-breadcrumb separator=">">
                <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
                <el-breadcrumb-item>搜索 {{ route.params.name }}</el-breadcrumb-item>
            </el-breadcrumb>
        </div>
        <div class="sub-container">
            <div class="body">
                <!-- 商品列表-->
                <GoodItem v-for="good in goodList" :good="good" :key="good.id"></GoodItem>
            </div>
        </div>
    </div>
    <div v-else>
        <el-empty description="暂无商品数据" />
    </div>
</template>

<style lang="scss" scoped>
.bread-container {
    padding: 25px 0;
    color: #666;
}

.sub-container {
    padding: 20px 10px;
    background-color: #fff;

    .body {
        display: flex;
        flex-wrap: wrap;
        padding: 0 10px;
        height: 80vh;
        overflow-y: auto;
    }


    .pagination-container {
        margin-top: 20px;
        display: flex;
        justify-content: center;
    }
}
</style>
