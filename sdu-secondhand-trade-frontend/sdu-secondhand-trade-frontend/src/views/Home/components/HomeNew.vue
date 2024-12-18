<script setup>
import HomePanel from './HomePanel.vue';
import { getNewGoodsAPI } from '@/apis/home';
import { onMounted, ref } from 'vue';
import { useStaticStore } from '@/stores/static';
import { useCategoryStore } from '@/stores/category';

const goodList = ref([]);

const getNewGoods = async () => {
    const res = await getNewGoodsAPI();
    goodList.value = res.data;
};

const staticStore = useStaticStore()
const categoryStore = useCategoryStore()

onMounted(() => getNewGoods());


</script>

<template>
  <HomePanel title="新鲜好物" sub-title="二手宝藏，物超所值">
    <ul class="goods-list">
      <li v-for="item in goodList.slice(0, 10)" :key="item.id">
        <RouterLink :to="`/`">
          <el-tag type="primary" class="campus-tag">
            {{ (staticStore.campusList.find(campus => campus.id === item.campus)||{name : '校外'}).name  }}
          </el-tag>
          <el-tag type="success" class="category-tag">
            {{ (categoryStore.categoryList.find(category => category.id === item.category)||{name : '闲置'}).name }}
          </el-tag>
          <img class="icon" v-img-lazy="item.cover" alt="商品图片" loading="lazy" />
          <p class="name">{{ item.name }}</p>
          <p class="price">&yen;{{ item.price }}</p>
        </RouterLink>
      </li>
    </ul>
  </HomePanel>
</template>

<style scoped lang="scss">
.goods-list {
  display: flex;
  justify-content: space-between;
  flex-wrap: wrap;
  row-gap: 20px;

  li {
    width: 306px;
    height: 406px;
    background: $bgLightColor;
    border-radius: 8px;
    overflow: hidden;
    position: relative;
    transition: transform 0.3s ease, box-shadow 0.3s ease;

    &:hover {
      transform: translateY(-5px);
      box-shadow: 0 6px 12px rgba(0, 0, 0, 0.2);
    }

    .campus-tag {
      position: absolute;
      top: 8px;
      left: 8px;
      z-index: 10;
    }
    .category-tag {
      position: absolute;
      top: 8px;
      left: 70px;
      z-index: 10;
    }

    img {
      width: 100%;
      height: 306px;
      object-fit: cover;
    }

    p {
      font-size: 22px;
      padding-top: 12px;
      text-align: center;
      text-overflow: ellipsis;
      overflow: hidden;
      white-space: nowrap;
    }

    .name {
      color: #333;
      font-weight: bold;
    }

    .price {
      color: $priceColor;
      font-size: 20px;
      font-weight: bold;
    }
  }
}
</style>
