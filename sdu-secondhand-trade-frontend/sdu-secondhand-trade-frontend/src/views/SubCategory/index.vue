<script setup>
import { getSubCategoryAPI } from '@/apis/category';
import { useCategoryStore } from '@/stores/category';
import { computed, onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import GoodItem from '../Category/components/GoodItem.vue';

const categoryStore = useCategoryStore();
const route = useRoute();

const categoryData = computed(() => {
  // 从 categoryList 中找到符合 route.params.id 的子分类
  let foundCategory = null;

  categoryStore.categoryList.forEach(category => {
    if (category.children && category.children.length > 0) {
      const child = category.children.find(child => child.id == route.params.id);
      if (child) {
        foundCategory = {
          ...child, // 获取子分类的所有数据
          father: { id: category.id, name: category.name } // 附加父分类的信息
        };
      }
    }
  });

  return foundCategory || {}; // 如果没有找到符合条件的分类，返回空对象
});

//获取商品数据
const goodList = ref([])
const reqData = ref({
  category_id: route.params.id,
  page: 1,
  page_size: 20,
  sort_field: 'publish_time'
})
const getGoodList = async () => {
  let res = await getSubCategoryAPI(reqData.value)
  goodList.value = res.data
}

onMounted(() => {
  getGoodList()
})

// tab切换
const tabChange = () => {
  reqData.page = 1
  getGoodList()
}


const load = async () => {
  reqData.value.page++
  const res = await getSubCategoryAPI(reqData.value)
  goodList.value = [...goodList.value, ...res.data]
  // 加载完毕 停止监听
  if (res.data.length === 0) {
    disabled.value = true
  }
}
</script>

<template>
  <div class="container" v-if="categoryData.id">
    <!-- 面包屑 -->
    <div class="bread-container">
      <el-breadcrumb separator=">">
        <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
        <el-breadcrumb-item :to="{ path: `/category/${categoryData.father.id}` }">{{ categoryData.father.name
          }}</el-breadcrumb-item>
        <el-breadcrumb-item>{{ categoryData.name }}</el-breadcrumb-item>
      </el-breadcrumb>
    </div>

    <div class="sub-container">
      <el-tabs v-model="reqData.sort_field" @tab-change="tabChange">
        <el-tab-pane label="最新商品" name="publishTime"></el-tab-pane>
        <el-tab-pane label="我的校区" name="campus"></el-tab-pane>
      </el-tabs>

      <div class="body" v-infinite-scroll="load">
        <!-- 商品列表-->
        <GoodItem v-for="good in goodList" :good="good" :key="good.id"></GoodItem>
      </div>
    </div>
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
