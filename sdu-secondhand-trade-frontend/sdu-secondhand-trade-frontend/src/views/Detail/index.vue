<script setup>
import { getGoodDetailAPI } from '@/apis/good'
import { useCategoryStore } from '@/stores/category';
import { computed, onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import GoodItem from '../Category/components/GoodItem.vue';
import ImageView from '@/components/ImageView/index.vue'
import { useStaticStore } from '@/stores/static';
import { initOrderAPI } from '@/apis/order';
const good = ref({})
const route = useRoute()
const getGood = async () => {
  const res = await getGoodDetailAPI(parseInt(route.params.id))
  good.value = res.data
  if (!good.value.is_effective) {
    showErrorMessage()
  }
}
const staticStore = useStaticStore();
const categoryStore = useCategoryStore();

const categoryData = computed(() => {
  // 从 categoryList 中找到符合 route.params.id 的子分类
  let foundCategory = null;

  categoryStore.categoryList.forEach(category => {
    if (category.children && category.children.length > 0) {
      const child = category.children.find(child => child.id == good.value.category);
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



onMounted(() => {
  getGood()
})


const router = useRouter()
// 显示弹窗并返回首页
const showErrorMessage = async () => {
  try {
    await ElMessageBox.confirm(
      '当前商品已被删除或商品已售出，是否返回上一页？',
      '提示',
      {
        confirmButtonText: '返回',
        cancelButtonText: '取消',
        type: 'warning'
      }
    )
    router.go(-1)
  } catch (error) {
  }
}

const buy = async () => {
  try {
    const res = await initOrderAPI(good.value.id)
    if (res.code === 0) {
      router.push(`/order/${res.data}`)
    }
  } catch {
    showErrorMessage()
  }

}

</script>

<template>
  <div class="goods-page">
    <div class="container" v-if="good.is_effective">
      <div class="bread-container" v-if="categoryData.id">
        <el-breadcrumb separator=">">
          <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
          <el-breadcrumb-item :to="{ path: `/category/${categoryData.father.id}` }">{{ categoryData.father.name
            }}</el-breadcrumb-item>
          <el-breadcrumb-item :to="{ path: `/category/sub/${categoryData.id}` }">{{ categoryData.name
            }}</el-breadcrumb-item>
          <el-breadcrumb-item>{{ good.name }}</el-breadcrumb-item>
        </el-breadcrumb>
      </div>
      <!-- 商品信息 -->
      <div class="info-container">
        <div>
          <div class="goods-info">
            <div class="media">
              <!-- 图片预览区 -->
              <ImageView :imageList="good.pictures.slice(0, 5) || []" />

            </div>
            <div class="spec">
              <!-- 商品信息区 -->
              <p class="g-name"> {{ good.name }} </p>
              <p class="g-full-name"> {{ good.full_name }} </p>
              <p class="g-desc">{{ good.description }} </p>
              <p class="g-price">
                <span>{{ good.price }}</span>
              </p>
              <div class="g-service">
                <dl>
                  <dt>商品规格</dt>
                  <dd>{{ good.specification + " " + good.status }}</dd>
                </dl>
                <dl>
                  <dt>服务</dt>
                  <dd>
                    <span>平台保障 </span>
                    <span>平台退货 </span>
                    <span>平台退款 </span>
                    <a href="javascript:;">了解详情 </a>
                  </dd>
                </dl>
              </div>
              <!-- 按钮组件 -->
              <div>
                <el-button size="large" class="btn" @click="buy()">
                  购买
                </el-button>
              </div>

            </div>
          </div>
          <div class="goods-footer">
            <div class="goods-article">
              <!-- 商品详情 -->
              <div class="goods-tabs">
                <nav>
                  <a>商品详情</a>
                </nav>
                <div class="goods-detail">
                  <!-- 属性 -->
                  <ul class="attrs">
                    <li>
                      <span class="dt">商品全称</span>
                      <span class="dd">{{ good.full_name }}</span>
                    </li>
                    <li>
                      <span class="dt">分类</span>
                      <span class="dd">{{ categoryData.name }}</span>
                    </li>
                    <li>
                      <span class="dt">规格</span>
                      <span class="dd">{{ good.specifications }}</span>
                    </li>
                    <li>
                      <span class="dt">交易方式</span>
                      <span class="dd">{{ staticStore.getTransactionMethod(good.transaction_method) }}</span>
                    </li>
                    <li>
                      <span class="dt">品牌</span>
                      <span class="dd">{{ good.brand }}</span>
                    </li>
                    <li>
                      <span class="dt">校区</span>
                      <span class="dd">{{ (staticStore.campusList.find(campus => campus.id === good.campus) ||
                { name: '校外' }).name }}校区</span>
                    </li>
                  </ul>
                  <div class="detail">
                    <h3>交易方式</h3>
                    <div class="transaction-method" v-if="good.transaction_method === 0">
                      <span>其他方式：</span>
                      <span>与卖家商讨</span>
                    </div>
                    <div class="transaction-method" v-if="good.transaction_method === 1">
                      <span>买家自提：</span>
                      <span>买家至卖家处自提（需要与卖家商量好时间）</span>
                    </div>
                    <div class="transaction-method" v-if="good.transaction_method === 1">
                      <span>自提地址：</span>
                      <span>{{ good.seller_address }}</span>
                    </div>
                    <div class="transaction-method" v-if="good.transaction_method === 2">
                      <span>快递邮寄：</span>
                      <span>卖家通过快递、闪送等寄件至以上所填地址</span>
                    </div>
                    <div class="transaction-method" v-if="good.transaction_method === 3">
                      <span>送货上门：</span>
                      <span>卖家亲自将物品送至以上所填地址</span>
                    </div>
                    <div class="transaction-method" v-if="good.transaction_method === 4">
                      <span>当面交易：</span>
                      <span>卖家与买家自行商讨时间地点进行交易</span>
                    </div>
                    <div class="transaction-method" v-if="good.transaction_method === 5">
                      <span>虚拟商品：</span>
                      <span>虚拟商品无需配送</span>
                    </div>
                  </div>
                  <div class="detail">
                    <h3>详细说明</h3>
                    <p>{{ good.detail }}</p>
                  </div>
                  <!-- 图片 -->
                  <div class="goods-images">
                    <h4>商品图片</h4>
                    <div class="image-gallery">
                      <img v-for="img in good.pictures" v-img-lazy="img.url" :key="img.id" alt="商品图片" />
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div class="goods-aside">
              <div class="goods-hot">
                <h3>同类好物</h3>
                <GoodItem v-for="good in (categoryData.goods || []).slice(0, 16)" :good="good" :key="good.id">
                </GoodItem>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div v-else>
      <el-empty description="当前商品已被删除或商品已售出，请返回" />
    </div>
  </div>
</template>


<style scoped lang='scss'>
.goods-page {
  .goods-info {
    min-height: 600px;
    background: #fff;
    display: flex;

    .media {
      width: 680;
      height: 600px;
      padding: 30px 50px;
    }

    .spec {
      flex: 1;
      padding: 60px 40px 60px 30px;
    }
  }

  .goods-footer {
    display: flex;
    margin-top: 40px;

    .goods-article {
      flex: 1;
      margin-right: 25px;
    }

    .goods-aside {
      width: 280px;
      min-height: 1000px;
    }
  }

  .number-box {
    display: flex;
    align-items: center;

    .label {
      width: 60px;
      color: #999;
      padding-left: 10px;
    }
  }

  .g-name {
    font-size: 35px;
    font-weight: bold;
  }

  .g-full-name {
    margin-top: 10px;
    font-size: 25px;
  }

  .g-desc {
    color: #999;
    font-size: 15px;
    margin-top: 15px;
  }

  .g-price {
    margin-top: 10px;

    span {
      &::before {
        content: "¥";
        font-size: 20px;
      }

      &:first-child {
        color: $priceColor;
        margin-right: 10px;
        font-size: 30px;
      }
    }
  }

  .g-service {
    background: #f5f5f5;
    width: 600px;
    padding: 20px 10px 0 10px;
    margin-top: 10px;
    font-size: 15px;

    dl {
      padding-bottom: 30px;
      display: flex;
      align-items: center;

      dt {
        width: 100px;
        color: #999;
      }

      dd {
        color: #666;

        &:last-child {
          span {
            margin-right: 10px;

            &::before {
              content: "•";
              color: $mainColor;
              margin-right: 2px;
            }
          }

          a {
            color: $mainColor;
          }
        }
      }
    }
  }

  .goods-sales {
    display: flex;
    width: 400px;
    align-items: center;
    text-align: center;
    height: 140px;

    li {
      flex: 1;
      position: relative;

      ~li::after {
        position: absolute;
        top: 10px;
        left: 0;
        height: 60px;
        border-left: 1px solid #e4e4e4;
        content: "";
      }

      p {
        &:first-child {
          color: #999;
        }

        &:nth-child(2) {
          color: $priceColor;
          margin-top: 10px;
        }

        &:last-child {
          color: #666;
          margin-top: 10px;

          i {
            color: $mainColor;
            font-size: 14px;
            margin-right: 2px;
          }

          &:hover {
            color: $mainColor;
            cursor: pointer;
          }
        }
      }
    }
  }


  .goods-tabs {
    min-height: 600px;
    background: #fff;
    font-size: 20px;

    nav {
      height: 80px;
      line-height: 80px;
      display: flex;
      border-bottom: 1px solid #f5f5f5;

      a {
        padding: 0 40px;
        font-size: 25px;
        position: relative;

        >span {
          color: $priceColor;
          font-size: 16px;
          margin-left: 10px;
        }
      }
    }
  }


  .goods-detail {
    padding: 60px;
    font-size: 16px;

    .attrs {
      display: flex;
      flex-wrap: wrap;
      margin-bottom: 30px;

      li {
        display: flex;
        margin-bottom: 10px;
        width: 50%;

        .dt {
          width: 100px;
          color: #999;
        }

        .dd {
          flex: 1;
          color: #666;
        }
      }
    }

    .detail {
      margin-bottom: 30px;

      p {
        margin-top: 10px;
        color: #666;
        line-height: 1.8;
        word-wrap: break-word;
        word-break: break-all;
      }

      .transaction-method {
        margin-top: 10px;
        color: #666;
        line-height: 1.8;
        word-wrap: break-word;
        word-break: break-all;
      }
    }

    .goods-images {
      margin-top: 20px;

      h4 {
        font-size: 18px;
        font-weight: bold;
        color: #333;
        margin-bottom: 15px;
      }

      .image-gallery {
        display: flex;
        flex-direction: column;
        /* 竖着排 */
        gap: 20px;
        /* 图片之间的间距 */
        padding: 10px 0;
        /* 控制左右padding */

        img {
          width: 100%;
          /* 图片横向占满 */
          height: auto;
          object-fit: cover;
          /* 保持图片的比例 */
          border-radius: 8px;
        }
      }
    }
  }

  .btn {
    margin-top: 80px;
  }

  .bread-container {
    padding: 25px 0;
  }

  .goods-hot {
    background: white;

    h3 {
      height: 70px;
      background: $helpColor;
      color: #fff;
      font-size: 18px;
      line-height: 70px;
      padding-left: 25px;
      margin-bottom: 10px;
      font-weight: normal;
    }
  }

}
</style>