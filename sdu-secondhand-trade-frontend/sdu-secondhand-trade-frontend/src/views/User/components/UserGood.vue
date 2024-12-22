<script setup>
import { deleteGoodAPI, getMyGoodAPI } from '@/apis/good'; // 获取我的商品列表 API
import { useRouter } from 'vue-router';
import { ref, onMounted } from 'vue';
import { ElMessage } from 'element-plus';
import { Search } from '@element-plus/icons-vue'

const tabTypes = [
  { status: -1, label: "全部订单" },
  { status: 0, label: "未完成" },
  { status: 1, label: "已完成" },
]

// 商品列表和分页数据
const goodList = ref([]);
const total = ref(0);
const reqData = ref({
  is_effective: -1,
  page: 1,
  page_size: 4,
  search: ''
});

// 获取商品列表
const getMyGoodList = async () => {
  const res = await getMyGoodAPI(reqData.value);
  goodList.value = res.data.goods; // 获取商品列表
  total.value = res.data.total; // 获取商品总数
};

// 组件挂载时获取商品列表
onMounted(() => {
  getMyGoodList();
});

const tabChange = (status) => {
  reqData.value.is_effective = parseInt(status) - 1
  getMyGoodList()
}

// 分页改变时重新获取商品列表
const pageChange = (page) => {
  reqData.value.page = page;
  getMyGoodList();
};

// 删除商品
const deleteGood = async (goodId) => {
  // 假设有一个删除商品的 API
  const res = await deleteGoodAPI(goodId);
  if (res.code === 0) {
    ElMessage.success('商品已删除');
    getMyGoodList(); // 刷新商品列表
  } else {
    ElMessage.error('删除失败');
  }
};

const formateStatus = (status) => {
  const statusMap = {
    false: "已完成",
    true: "未完成",
  }
  return statusMap[status]||"已完成"
}

const onSearch = () => {
  getMyGoodList()
}

const router = useRouter();

const afterSaleMethods = new Map([
    [0, "其他方式"],
    [1, "买家自提"],
    [2, "快递邮寄"],
    [3, "送货上门"],
    [4, "当面交易"],
    [5, "虚拟商品"]
]);

const getAfterSaleMethodDescription = (status) => {
    return afterSaleMethods.get(status) || "未知方式"
};
</script>

<template>
  <div class="good-container">
    <!-- 使用 Flexbox 布局，tabs 和输入框并排显示 -->
    <div class="tabs-and-search">
      <el-tabs @tab-change="tabChange" class="tabs-container" style="flex: 1;">
        <el-tab-pane v-for="item in tabTypes" :key="item.status" :label="item.label" />
      </el-tabs>

      <!-- 右侧的搜索框 -->
      <div class="search-container">
        <el-input v-model="reqData.search" style="width: 240px;" placeholder="请输入" :suffix-icon="Search"
          @change="onSearch" />
      </div>
    </div>

    <div class="main-container">
      <div class="holder-container" v-if="goodList.length === 0">
        <el-empty description="暂无商品数据" />
      </div>

      <div v-else>
        <!-- 商品列表 -->
        <div class="good-item" v-for="good in goodList" :key="good.id">
          <div class="head">
            <span>商品编号：{{ good.id }}</span>
            <span>上架时间：{{ good.created_at }}</span>
          </div>
          <div class="body">
            <div class="column goods">
              <ul>
                <li>
                  <a class="image" href="javascript:;">
                    <img :src="good.cover" alt="" />
                  </a>
                  <div class="info">
                    <p class="name ellipsis-2">{{ good.name }}</p>
                    <p class="attr ellipsis"><span>{{ good.description }}</span></p>
                  </div>
                  <div class="price">¥{{ parseFloat(good.price).toFixed(2) }}</div>
                </li>
              </ul>
            </div>
            <div class="column">
              <p><strong>品牌：</strong>{{ good.brand }}</p>
              <p><strong>规格：</strong>{{ good.specifications }}</p>
              <p><strong>状态：</strong>{{ good.status }}</p>
            </div>
            <div class="column">
              <p><strong>交易方式：</strong>{{ getAfterSaleMethodDescription(good.transaction_method) }}</p>
              <p v-if="good.transaction_method === 1"><strong>自提地址：</strong>{{ good.seller_address }}</p>
            </div>
            <div class="column state">
              <p><strong>商品状态</strong><br/>{{ formateStatus(good.is_effective) }}</p>
            </div>
            <div class="column action">
              <el-button @click="deleteGood(good.id)" type="primary" v-if="good.is_effective">删除商品</el-button>
              <p><a href="javascript:;" v-if="good.is_effective" @click="router.push(`/detail/${good.id}`)">查看详情</a>
              </p>
            </div>
          </div>
        </div>

        <!-- 分页 -->
        <div class="pagination-container">
          <el-pagination :total="total" :page-size="reqData.page_size" @current-change="pageChange" background
            layout="prev, pager, next" />
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss">
.good-container {
  padding: 10px 20px;

  .tabs-and-search {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 20px;
  }

  .tabs-container {
    flex: 1;
  }

  .search-container {
    margin-left: 20px;
  }

  .pagination-container {
    display: flex;
    justify-content: center;
  }

  .main-container {
    min-height: 500px;

    .holder-container {
      min-height: 500px;
      display: flex;
      justify-content: center;
      align-items: center;
    }
  }
}

.good-item {
  margin-bottom: 20px;
  border: 1px solid #f5f5f5;

  .head {
    height: 50px;
    line-height: 50px;
    background: #f5f5f5;
    padding: 0 20px;
    overflow: hidden;

    span {
      margin-right: 20px;
    }
  }

  .body {
    display: flex;
    align-items: stretch;

    .column {
      border-left: 1px solid #f5f5f5;
      text-align: center;
      padding: 20px;

      >p {
        padding-top: 10px;
      }

      &:first-child {
        border-left: none;
      }

      &.goods {
        flex: 1;
        padding: 0;
        align-self: center;
        max-width: 50%;

        ul {
          li {
            border-bottom: 1px solid #f5f5f5;
            padding: 10px;
            display: flex;

            &:last-child {
              border-bottom: none;
            }

            .image {
              width: 70px;
              height: 70px;
              border: 1px solid #f5f5f5;
            }

            .info {
              width: 220px;
              text-align: left;
              padding: 0 10px;

              p {
                margin-bottom: 5px;

                &.name {
                  height: 38px;
                }

                &.attr {
                  color: #999;
                  font-size: 12px;

                  span {
                    margin-right: 5px;
                  }
                }
              }
            }

            .price {
              width: 100px;
            }

            .count {
              width: 80px;
            }
          }
        }
      }

      &.state {
        width: 120px;

        .green {
          color: $mainColor;
        }
      }

      &.action {
        min-width: 20%;

        a {
          display: block;

          &:hover {
            color: $mainColor;
          }
        }
      }
    }
  }
}
</style>