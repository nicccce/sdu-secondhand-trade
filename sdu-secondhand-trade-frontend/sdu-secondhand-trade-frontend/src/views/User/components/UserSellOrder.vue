<script setup>
import { getGoodDetailAPI } from '@/apis/good';
import { getSellerOrderListAPI, sendAfterSaleAPI } from '@/apis/order';
import dayjs from 'dayjs'
import { onMounted, ref } from 'vue';
import { useRouter } from 'vue-router';

// tab列表
const tabTypes = [
  { status: -1, label: "全部订单" },
  { status: 0, label: "待提交" },
  { status: 1, label: "待付款" },
  { status: 2, label: "支付成功" },
  { status: 3, label: "订单取消" },
  { status: 4, label: "订单超时" },
]
// 订单列表
const orderList = ref([])
const total = ref(0)
const reqData = ref({
  status: -1,
  page: 1,
  page_size: 2
})
const getSellerOrderList = async () => {
  // 请求订单列表
  const res = await getSellerOrderListAPI(reqData.value);
  orderList.value = res.data.orders
  total.value = res.data.total

  // 遍历订单列表，逐个获取商品详情
  for (let order of orderList.value) {
    const goodRes = await getGoodDetailAPI(order.good_id);  // 获取商品详情
    order.good = goodRes.data;  // 将商品详情填充到订单项
  }
};

onMounted(async () => {
  await getSellerOrderList()
})

//计算截止时间
const computeTime = (targetTime) => {
  // 使用 dayjs 来加上 10 分钟
  const time = dayjs(targetTime).add(10, 'minute')

  // 格式化为 yyyy-mm-dd hh:mm:ss
  const formattedTime = time.format('YYYY-MM-DD HH:mm:ss')

  return formattedTime
}

const tabChange = (status) => {
  reqData.value.status = parseInt(status) - 1
  getSellerOrderList()
}

const pageChange = (page) => {
  reqData.value.page = page
  getSellerOrderList()
}

const formateStatus = (status) => {
  const statusMap = {
    0: "待提交",
    1: "待付款",
    2: "支付成功",
    3: "订单取消",
    4: "订单超时",
  }
  return statusMap[status]
}

const router = useRouter()

//提交售后
// 用于存储弹窗相关的数据
const refundDialogVisible = ref(false);  // 控制弹窗显示
const selectedOrder = ref(null);  // 当前选中的订单
const questionDescription = ref('');  // 存储问题阐述内容

// 弹窗展示的逻辑
const applyRefund = (order) => {
  selectedOrder.value = order;  // 保存当前订单
  refundDialogVisible.value = true;  // 显示弹窗
};

// 关闭弹窗
const closeRefundDialog = () => {
  refundDialogVisible.value = false;
  questionDescription.value = '';  // 清空问题阐述
};

// 处理提交售后请求的逻辑
const submitRefundRequest =async () => {
  if (!questionDescription.value) {
    ElMessage({
      type: 'error',
      message: '问题阐述为必填项！'
    });
    return;
  }
  const res = await sendAfterSaleAPI({
    order_id : selectedOrder.value.id,
    problem: questionDescription.value,
  })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '售后已提交！'
    });
  }
  closeRefundDialog();  // 关闭弹窗
};
</script>

<template>
  <div class="order-container">
    <el-tabs @tab-change="tabChange">
      <!-- tab切换 -->
      <el-tab-pane v-for="item in tabTypes" :key="item.status" :label="item.label" />

      <div class="main-container">
        <div class="holder-container" v-if="orderList.length === 0">
          <el-empty description="暂无订单数据" />
        </div>
        <div v-else>
          <!-- 订单列表 -->
          <div class="order-item" v-for="order in orderList" :key="order.id">
            <div class="head">
              <span>下单时间：{{ order.created_at }}</span>
              <span>订单编号：{{ order.id }}</span>
              <!-- 未付款，倒计时时间还有 -->
              <span class="down-time" v-if="order.status === 1 || order.status === 0">
                <i class="iconfont icon-down-time"></i>
                <b>付款截止: {{ computeTime(order.created_at) }}</b>
              </span>
            </div>
            <div class="body" v-if="order.good">
              <div class="column goods">
                <ul>
                  <li>
                    <a class="image" href="javascript:;">
                      <img :src="order.good.cover" alt="" />
                    </a>
                    <div class="info">
                      <p class="name ellipsis-2">
                        {{ order.good.name }}
                      </p>
                      <p class="attr ellipsis">
                        <span>{{ order.good.description }}</span>
                      </p>
                    </div>
                    <div class="price">¥{{ parseFloat(order.good.price).toFixed(2) }}</div>
                    <div class="specifications">{{ order.good.specifications }}</div>
                  </li>
                </ul>
              </div>
              <div class="column state">
                <p>{{ formateStatus(order.status) }}</p>
              </div>
              <div class="column amount">
                <p class="red">¥{{ parseFloat(order.good.price).toFixed(2) }}</p>
                <p v-if="order.payment === 1">在线支付</p>
                <p v-if="order.payment === 2">货到付款</p>
                <p v-if="order.payment === 3">其他途径</p>
              </div>
              <div class="column action">
                <p><a href="javascript:;" @click="router.push(`/order/${order.id}`)">查看详情</a></p>
                <p v-if="order.status === 2">
                  <a href="javascript:;"  @click="applyRefund(order)">申请售后</a>
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

    </el-tabs>
  </div>

  <!-- 申请售后弹窗 -->
  <el-dialog title="申请售后" v-model="refundDialogVisible" width="600px" @close="closeRefundDialog">
    <p style="font-size: 20;">订单号：{{ selectedOrder?.id }}</p>
    <el-input v-model="questionDescription" type="textarea" placeholder="请输入问题阐述（必填）" :autosize="{ minRows: 4, maxRows: 6 }"
      show-word-limit required  style="padding: 20px 20px;"/>
    <div slot="footer" class="dialog-footer">
      <el-button @click="closeRefundDialog">取消</el-button>
      <el-button type="primary" @click="submitRefundRequest">提交申请</el-button>
    </div>
  </el-dialog>
</template>

<style scoped lang="scss">
.order-container {
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

.order-item {
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

      &.down-time {
        margin-right: 0;
        float: right;

        i {
          vertical-align: middle;
          margin-right: 3px;
        }

        b {
          vertical-align: middle;
          font-weight: normal;
        }
      }
    }

    .del {
      margin-right: 0;
      float: right;
      color: #999;
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

      &.amount {
        width: 200px;

        .red {
          color: $priceColor;
        }
      }

      &.action {
        width: 140px;

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
.dialog-footer {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
  margin-top: 20px;
}
</style>