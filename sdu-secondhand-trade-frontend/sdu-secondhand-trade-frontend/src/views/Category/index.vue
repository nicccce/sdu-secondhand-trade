<script setup>
import { useCategoryStore } from '@/stores/category';
import { onMounted, ref, watch } from 'vue';
import { onBeforeRouteUpdate, useRoute } from 'vue-router';
import { getBannerAPI } from '@/apis/home';
import GoodItem from './components/GoodItem.vue';

const categoryStore = useCategoryStore()
const route = useRoute()
const categoryData = ref(categoryStore.categoryList.find(category => category.id == route.params.id))

onBeforeRouteUpdate(to => {
    categoryData.value = categoryStore.categoryList.find(category => category.id == to.params.id)
    bannerList.value[bannerList.value.length - 1] = { img_url: categoryData.picture }
})

const bannerList = ref([])

const getBanner = async () => {
    const res = await getBannerAPI()
    bannerList.value = res.data
    bannerList.value.push({ img_url: categoryData.picture })
}

onMounted(() => getBanner())
</script>

<template>
    <div class="top-category">
        <div class="container m-top-20">
            <div class="bread-container">
                <el-breadcrumb separator=">">
                    <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
                    <el-breadcrumb-item>{{ categoryData.name }}</el-breadcrumb-item>
                </el-breadcrumb>
            </div>
            <div class="banner">
                <el-carousel height="610px">
                    <el-carousel-item v-for="item in bannerList" :key="item.id">
                        <img v-img-lazy="item.img_url" alt="图片加载失败">
                    </el-carousel-item>
                </el-carousel>
            </div>
            <div class="sub-list">
                <h3>全部分类</h3>
                <ul>
                    <li v-for="i in categoryData.children" :key="i.id">
                        <RouterLink to="/">
                            <img :src="i.picture" />
                            <p>{{ i.name }}</p>
                        </RouterLink>
                    </li>
                </ul>
            </div>
            <div class="ref-goods" v-for="item in categoryData.children" :key="item.id">
                <div class="head">
                    <h3>- {{ item.name }}-</h3>
                </div>
                <div class="body">
                    <GoodItem v-for="good in (item.goods || []).slice(0, 5)" :good="good" :key="good.id" />
                </div>
            </div>
        </div>
    </div>
</template>


<style scoped lang="scss">
.top-category {
    h3 {
        font-size: 28px;
        color: #666;
        font-weight: normal;
        text-align: center;
        line-height: 100px;
    }

    .sub-list {
        margin-top: 20px;
        background-color: #fff;

        ul {
            display: flex;
            padding: 0 32px;
            flex-wrap: wrap;

            li {
                width: 168px;
                height: 160px;


                a {
                    text-align: center;
                    display: block;
                    font-size: 16px;

                    img {
                        width: 100px;
                        height: 100px;
                    }

                    p {
                        line-height: 40px;
                    }

                    &:hover {
                        color: $mainColor;
                    }
                }
            }
        }
    }

    .ref-goods {
        background-color: #fff;
        margin-top: 20px;
        position: relative;

        .head {
            .xtx-more {
                position: absolute;
                top: 20px;
                right: 20px;
            }

            .tag {
                text-align: center;
                color: #999;
                font-size: 20px;
                position: relative;
                top: -20px;
            }
        }

        .body {
            display: flex;
            justify-content: space-around;
            padding: 0 40px 30px;
        }
    }

    .bread-container {
        padding: 25px 0;
    }
}

.banner {
    margin-left: 0px;
    width: 1650px;
    height: 610px;
    margin: 0 auto;

    img {
        width: 100%;
        height: 100%;
        object-fit: fill;
    }
}
</style>