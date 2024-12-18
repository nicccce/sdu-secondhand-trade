<script setup>
import HomePanel from './HomePanel.vue';
import { getCampusGoodsAPI } from '@/apis/home';
import { onMounted, ref } from 'vue';
import { useCategoryStore } from '@/stores/category';

const goodList = ref([]);

const getCampusGoods = async () => {
    const res = await getCampusGoodsAPI(1);
    goodList.value = res.data;
};
const categoryStore = useCategoryStore()

onMounted(() => getCampusGoods());
</script>

<template>
  <HomePanel title="校区甄选" sub-title="近享美物，触手可及">
    <ul class="goods-list">
      <li v-for="item in goodList.slice(0, 10)" :key="item.id">
        <RouterLink :to="`/`">
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
    position: relative; // 添加相对定位，为 el-tag 提供定位参照
    transition: transform 0.3s ease, box-shadow 0.3s ease;

    &:hover {
      transform: translateY(-5px);
      box-shadow: 0 6px 12px rgba(0, 0, 0, 0.2);
    }

    .category-tag {
      position: absolute;
      top: 8px;
      left: 8px;
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
