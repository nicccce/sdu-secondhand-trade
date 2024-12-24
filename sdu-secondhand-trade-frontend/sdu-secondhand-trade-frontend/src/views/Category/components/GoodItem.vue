<script setup>
import { useStaticStore } from '@/stores/static';

defineProps({
  good: {
    type: Object,
    default: () => { }
  }
})

const staticStore = useStaticStore()
</script>

<template>
  <RouterLink :to="`/detail/${good.id}`" class="goods-item">
    <div class="image-container">
      <el-tag type="primary" class="campus-tag">
        {{ (staticStore.campusList.find(campus => campus.id === good.campus) || { name: '校外' }).name }}
      </el-tag>
      <img v-img-lazy="good.cover" alt="" />
    </div>
    <p class="name ellipsis">{{ good.name }}</p>
    <p class="desc ellipsis">{{ good.description }}</p>
    <p class="price">&yen;{{ good.price }}</p>
  </RouterLink>
</template>

<style scoped lang="scss">
.goods-item {
  display: block;
  width: 270px;
  height: 350px;
  padding: 25px 49px;
  text-align: center;
  transition: all .5s;

  &:hover {
    transform: translate3d(0, -10px, 0);
    box-shadow: 0 3px 8px rgb(0 0 0 / 20%);
  }

  p {
    padding-top: 10px;
  }

  .name {
    font-size: 16px;
  }

  .desc {
    color: #999;
    height: 29px;
  }

  .price {
    color: $priceColor;
    font-size: 20px;
  }

  .image-container {
    position: relative;
    width: 175px;
    height: 175px;

    img {
      width: 100%;
      height: 100%;
      object-fit: cover;
    }

    .campus-tag {
      position: absolute;
      top: 8px;
      left: 8px;
      z-index: 10;
      font-size: 12px;
      padding: 4px 5px;
      border-radius: 8px;
    }
  }
}
</style>
