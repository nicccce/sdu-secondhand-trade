<script setup>
import { useCategoryStore } from '@/stores/category';
import { useStaticStore } from '@/stores/static';
const categoryStore = useCategoryStore()
const staticStore = useStaticStore()
</script>

<template>
  <div class="home-category">
    <ul class="menu">
      <li v-for="item in categoryStore.categoryList" :key="item.id">
        <RouterLink :to="`/category/${item.id}`">{{ item.name }}</RouterLink>
        <RouterLink :to="`/category/${item.id}`">{{ item.introduction }}</RouterLink>
        <div class="layer">
          <h4>分类推荐 <small>根据您的购买或浏览记录推荐</small></h4>
          <ul>
            <li v-for="good in (item.goods || []).slice(0, 16)" :key="good.id">
              <RouterLink :to="`/detail/${good.id}`">
                <img class="icon" v-img-lazy="good.cover" />
                <div class="info">
                  <p class="name ellipsis-2">
                    {{ good.name }}
                  </p>
                  <p class="desc ellipsis">{{ good.description }}</p>
                  <div class="price-tag">
                    <p class="price"><i>¥</i>{{ good.price }}</p>
                    <el-tag type="primary">
                      {{ (staticStore.campusList.find(campus => campus.id === good.campus)||{name : '校外'}).name }}
                    </el-tag>
                  </div>
                </div>
              </RouterLink>
            </li>
          </ul>
        </div>
      </li>
    </ul>
  </div>
</template>



<style scoped lang='scss'>
.home-category {
  width: 300px;
  height: 610px;
  background: rgba(0, 0, 0, 0.8);
  position: relative;
  z-index: 99;

  .menu {
    li {
      padding-left: 40px;
      height: 55px;
      line-height: 55px;

      &:hover {
        background: $lightColor;
      }

      a {
        margin-right: 4px;
        color: #fff;

        &:first-child {
          font-size: 16px;
        }
      }

      .layer {
        width: 1400px;
        height: 610px;
        background: rgba(255, 255, 255, 0.8);
        position: absolute;
        left: 300px;
        top: 0;
        display: none;
        padding: 0 15px;

        h4 {
          font-size: 20px;
          font-weight: normal;
          line-height: 60px;

          small {
            font-size: 16px;
            color: #666;
          }
        }

        ul {
          width: 100%;
          display: flex;
          flex-wrap: wrap;

          li {
            width: 310px;
            height: 120px;
            margin-right: 20px;
            margin-bottom: 15px;
            border: 1px solid #eee;
            border-radius: 4px;
            background: #fff;

            &:nth-child(4n) {
              margin-right: 0;
            }

            a {
              display: flex;
              width: 100%;
              height: 100%;
              align-items: center;
              padding: 10px;

              &:hover {
                background: $bgLightColor;
              }

              img {
                width: 95px;
                height: 95px;
              }

              .info {
                padding-left: 10px;
                line-height: 24px;
                overflow: hidden;

                .name {
                  font-size: 16px;
                  color: #666;
                }

                .desc {
                  color: #999;
                }

                .price-tag {
                  display: flex;
                  align-items: center; // 垂直居中对齐
                  justify-content: space-between;
                  margin-top: 8px;

                  .price {
                    font-size: 22px;
                    color: $priceColor;

                    i {
                      font-size: 16px;
                    }
                  }

                  .el-tag {
                    margin-left: 10px; // 为 tag 和价格之间提供间距
                    font-size: 14px; // 调整 tag 的字体大小
                    line-height: 1.5;
                  }
                }
              }
            }
          }
        }
      }

      &:hover {
        .layer {
          display: block;
        }
      }
    }
  }
}
</style>