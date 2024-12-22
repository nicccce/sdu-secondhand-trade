<script setup>
import { deleteGoodAPI, getGoodDetailAPI, getMyGoodAPI } from '@/apis/good'; // 获取我的商品列表 API
import { useRouter } from 'vue-router';
import { ref, onMounted } from 'vue';
import { ElMessage } from 'element-plus';
import { getAfterSaleAPI, getOrderAPI } from '@/apis/order';

const tabTypes = [
  { status: -1, label: "全部订单" },
  { status: 0, label: "未完成" },
  { status: 1, label: "已完成" },
]

// 售后列表和分页数据
const problemList = ref([]);
const total = ref(0);
const reqData = ref({
  status: -1,
  page: 1,
  page_size: 2
});

// 获取售后列表
const getProblemList = async () => {
  const res = await getAfterSaleAPI(reqData.value);
  problemList.value = res.data.problems;
  total.value = res.data.total;
  for (let problem of problemList.value) {
    const orderRes = await getOrderAPI(problem.order_id);
    problem.order = orderRes.data;
    const goodRes = await getGoodDetailAPI(problem.order.good_id);  // 获取商品详情
    problem.good = goodRes.data;
  }
};

// 组件挂载时获取商品列表
onMounted(() => {
  getProblemList();
});

const tabChange = (status) => {
  reqData.value.status = parseInt(status) - 1
  getProblemList()
}

// 分页改变时重新获取商品列表
const pageChange = (page) => {
  reqData.value.page = page;
  getProblemList();
};

const formateStatus = (status) => {
  const statusMap = {
    0: "待解决",
    1: "已受理",
    2: "已解决"
  }
  return statusMap[status]
}

const router = useRouter();
</script>

<template>
  <div class="problem-container">
    <el-tabs @tab-change="tabChange">
      <el-tab-pane v-for="item in tabTypes" :key="item.status" :label="item.label" />

      <div class="main-container">
        <div class="holder-container" v-if="problemList.length === 0">
          <el-empty description="暂无售后数据" />
        </div>

        <div v-else>
          <div v-for="problem in problemList" :key="problem.id">
            <div class="problem-item" v-if="problem.good">
              <div class="head">
                <span>售后编号：{{ problem.id }}</span>
                <span>订单编号：{{ problem.order.id }}</span>
              </div>
              <div class="body">
                <div class="column goods">
                  <ul>
                    <li>
                      <a class="image" href="javascript:;">
                        <img :src="problem.good.cover" alt="" />
                      </a>
                      <div class="info">
                        <p class="name ellipsis-2">{{ problem.good.name }}</p>
                        <p class="attr ellipsis"><span>{{ problem.good.description }}</span></p>
                      </div>
                      <div class="price">¥{{ parseFloat(problem.good.price).toFixed(2) }}</div>
                    </li>
                  </ul>
                </div>
                <div class="column problem">
                  <p class="problem-text">{{ problem.problem }}</p>
                </div>
                <div class="column reply">
                  <p class="reply-text">{{ problem.reply }}</p>
                </div>
                <div class="column state">
                  <p>{{ formateStatus(problem.status) }}</p>
                </div>
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
    </el-tabs>
  </div>
</template>

<style scoped lang="scss">
.problem-container {
  padding: 10px 20px;

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

.problem-item {
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
        width: 50%;

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

      &.problem,
      &.reply {
        flex: 2;
        padding: 10px;
        width: 40%;
        /* 问题和回复列宽 */
        text-align: left;
        overflow-wrap: break-word;
        word-wrap: break-word;
        word-break: break-word;
        max-width: 500px;
        white-space: normal;
        line-height: 1.6;
      }

      &.state {
        width: 120px;

        .green {
          color: $mainColor;
        }
      }
    }
  }
}
</style>
