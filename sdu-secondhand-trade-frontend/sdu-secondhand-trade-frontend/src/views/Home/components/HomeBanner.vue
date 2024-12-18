<script setup>
import { getBannerAPI } from '@/apis/home';
import { onMounted, ref } from 'vue';

const bannerList = ref([])

const getBanner = async () => {
    const res = await getBannerAPI()
    bannerList.value = res.data
}

onMounted(() => getBanner())

</script>



<template>
    <div class="home-banner">
        <el-carousel height="610px">
            <el-carousel-item v-for="item in bannerList" :key="item.id">
                <img v-img-lazy="item.img_url"
                    alt="图片加载失败">
            </el-carousel-item>
        </el-carousel>
    </div>
</template>




<style scoped lang='scss'>
.home-banner {
    margin-left: 0px;
    width: 1650px;
    height: 610px;
    position: absolute;
    left: 0;
    top: 0;
    z-index: 98;

    img {
        width: 100%;
        height: 100%;
        object-fit: fill;
    }
}
</style>