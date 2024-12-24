<script setup>
import { getGoodDetailAPI } from '@/apis/good';
import { cancelOrderAPI, getOrderAPI, submitOrderAPI } from '@/apis/order';
import { addAddressAPI } from '@/apis/user';
import { useCountDown } from '@/composables/useCountDown';
import { useUserStore } from '@/stores/user';
import { ElMessage } from 'element-plus';
import { onMounted, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';

const orderInfo = ref({})
const good = ref({})

const userStore = useUserStore()

const route = useRoute()
const router = useRouter()
const getOderInfo = async () => {
  const res = await getOrderAPI(parseInt(route.params.id))
  orderInfo.value = res.data
}

const getGood = async () => {
  const res = await getGoodDetailAPI(orderInfo.value.good_id)
  good.value = res.data
}
onMounted(async () => {
  await getOderInfo()
  await getGood()
  if (orderInfo.value.status === 1) {
    router.replace({ path: `/pay/${orderInfo.value.id}` })
  }
  startCountDown()
})

const showDialog = ref(false)

const activeAddress = ref({
  id: 0,
})
const switchAddress = (item) => {
  activeAddress.value = item
}
const confirmAddress = () => {
  orderInfo.value.buyer_name = activeAddress.value.receiver
  orderInfo.value.buyer_contact = activeAddress.value.contact
  orderInfo.value.buyer_address = activeAddress.value.address
  showDialog.value = false
  activeAddress.value = {
    id: 0,
  }
}
const rejectAddress = () => {
  showDialog.value = false
  activeAddress.value = {
    id: 0,
  }
}

const changePayment = (payment_id) => {
  if (orderInfo.value.payment === payment_id) {
    orderInfo.value.payment = 0
  } else {
    orderInfo.value.payment = payment_id
  }
}

const submitOrder = async () => {
  if (!(orderInfo.value.buyer_name)) {
    ElMessage({
      type: 'warning',
      message: '您需要先添加收货人信息才可提交订单！'
    });
    return
  }
  if (!(orderInfo.value.payment === 1 || orderInfo.value.payment === 2 || orderInfo.value.payment === 3)) {
    ElMessage({
      type: 'warning',
      message: '您您需要选择支付方式才可提交订单！'
    });
    return
  }
  const res = await submitOrderAPI(orderInfo.value)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '订单提交成功！'
    });
    if (orderInfo.value.payment == 1) {
      router.replace({ path: `/pay/${orderInfo.value.id}` })
    } else {
      router.replace({ path: `/pay/callback/${orderInfo.value.id}` })
    }
  }
}

const cancelOrder = async () => {
  const res = await cancelOrderAPI(orderInfo.value.id)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '订单已取消！'
    });
    router.push({ path: `/detail/${good.value.id}` });
  }
}

const countDown = useCountDown()

const startCountDown = () => {
  countDown.start(orderInfo.value.created_at)
}

// 新地址表单的初始值
const newAddress = ref({
  receiver: null,
  contact: null,
  address: null,
});

const addressRules = {
  receiver: [
    { required: true, message: '请输入收货人姓名', trigger: 'blur' }
  ],
  contact: [
    { required: true, message: '请输入联系方式', trigger: 'blur' }
  ],
  address: [
    { required: true, message: '请输入收货地址', trigger: 'blur' }
  ]
};

// 显示添加地址对话框
const showAddAddressDialog = ref(false);

// 提交新地址
const submitNewAddress = async () => {
  addressForm.value.validate(async (valid) => {
    console.log(valid);
    
  // 表单验证
  if (!valid) return

  // 调用接口提交新地址
  const res = await addAddressAPI(newAddress.value)
  if (res.code === 0) {
    // 成功提示
    ElMessage({
      type: 'success',
      message: '地址添加成功！'
    });
    // 更新地址列表
    userStore.refreshUserInfo()
    // 关闭对话框
    showAddAddressDialog.value = false;
  } else {
    // 失败提示
    ElMessage({
      type: 'error',
      message: '添加地址失败，请稍后再试！'
    })
  }
})}

const addressForm = ref(null);
// 校验表单

</script>

<template>
  <div class="pay-checkout-page" v-if="good.id">
    <div class="container" v-if="orderInfo.status === 0">
      <div class="pay-info">
        <span class="icon iconfont icon-shanchu red" v-if="countDown.formatTime.value === '00分00秒'"></span>
        <span class="icon iconfont icon-queren2 green" v-else></span>
        <div class="tip">
          <p>订单创建成功！请尽快完成提交。</p>
          <p v-if="countDown.formatTime.value === '00分00秒'">支付还剩 <span style="color: red;">00分00秒</span>,订单已自动取消</p>
          <p v-else>支付还剩 <span>{{ countDown.formatTime }}</span>, 超时后将取消订单</p>
        </div>
      </div>
      <div class="wrapper">
        <!-- 收货地址 -->
        <div>
          <h3 class="box-title">收货信息</h3>
          <div class="box-body">
            <div class="address">
              <div class="text">
                <div class="none" v-if="!(orderInfo.buyer_name)">您需要先添加收货人信息才可提交订单。</div>
                <ul v-else>
                  <li><span>收<i />货<i />人：</span>{{ orderInfo.buyer_name }}</li>
                  <li><span>联系方式：</span>{{ orderInfo.buyer_contact }}</li>
                  <li v-if="good.transaction_method === 2 || good.transaction_method === 3"><span>收货地址：</span>{{
                    orderInfo.buyer_address }}</li>
                </ul>
              </div>
              <div class="action">
                <el-button size="large" @click="showDialog = true">切换信息</el-button>
                <el-button size="large" @click="showAddAddressDialog = true">添加信息</el-button>
              </div>
            </div>
          </div>
        </div>
        <!-- 商品信息 -->
        <h3 class="box-title">商品信息</h3>
        <div class="box-body">
          <table class="good">
            <thead>
              <tr>
                <th width="520">商品信息</th>
                <th width="170">规格</th>
                <th width="170">单价</th>
                <th width="170">实付</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td>
                  <a href="javascript:;" class="info">
                    <img :src="good.cover" alt="">
                    <div class="right">
                      <p>{{ good.full_name }}</p>
                      <p>{{ good.description }}</p>
                    </div>
                  </a>
                </td>
                <td>{{ good.specifications }}</td>
                <td>&yen;{{ good.price }}</td>
                <td>&yen;{{ parseFloat(good.price).toFixed(2) }}</td>
              </tr>
            </tbody>
          </table>
        </div>
        <!-- 交易方式 -->
        <h3 class="box-title">交易方式</h3>
        <div class="box-body">
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
        <!-- 支付方式 -->
        <h3 class="box-title">支付方式</h3>
        <span v-if="orderInfo.payment === 2" style="color:#999">卖家与买家现场交易时付款。</span>
        <span v-else-if="orderInfo.payment === 3" style="color:#999">通过其他途径支付平台将无法保障交易的完成。</span>
        <span v-else-if="orderInfo.payment === 1" style="color:#999">买家线上通过平台付款。</span>
        <span v-else style="color:#999">您需要选择支付方式才可提交订单。</span>
        <div class="box-body">
          <a class="my-btn" :class="{ active: orderInfo.payment === 1 }" @click="changePayment(1)"
            href="javascript:;">在线支付</a>
          <a class="my-btn" :class="{ active: orderInfo.payment === 2 }" @click="changePayment(2)"
            href="javascript:;">货到付款</a>
          <a class="my-btn" :class="{ active: orderInfo.payment === 3 }" @click="changePayment(3)"
            href="javascript:;">其他途径</a>
        </div>
        <!-- 金额明细 -->
        <h3 class="box-title">金额明细</h3>
        <div class="box-body">
          <div class="total">
            <dl>
              <dt>商品总价：</dt>
              <dd>¥{{ parseFloat(good.price).toFixed(2) }}</dd>
            </dl>
            <dl>
              <dt>创建时间：</dt>
              <dd>{{ orderInfo.created_at }}</dd>
            </dl>
            <dl>
              <dt>应付总额：</dt>
              <dd class="price">¥{{ parseFloat(good.price).toFixed(2) }}</dd>
            </dl>
          </div>
        </div>
        <!-- 提交订单 -->
        <div class="submit">
          <el-button type="primary" size="large" @click="cancelOrder()">取消订单</el-button>
          <el-button type="primary" size="large" @click="submitOrder()">提交订单</el-button>
        </div>
      </div>
    </div>




    <div class="container" v-else>
      <div class="pay-info">
        <span class="icon iconfont icon-queren2 green" v-if="orderInfo.status === 2"></span>
        <span class="icon iconfont icon-shanchu red" v-else></span>
        <div class="tip">
          <p v-if="orderInfo.status === 2">订单交易成功！</p>
          <p v-if="orderInfo.status === 3">订单已取消！</p>
          <p v-if="orderInfo.status === 4">订单已超时！</p>
        </div>
      </div>
      <div class="wrapper">
        <!-- 收货地址 -->
        <div>
          <h3 class="box-title">收货信息</h3>
          <div class="box-body">
            <div class="address">
              <div class="text">
                <div class="none" v-if="!(orderInfo.buyer_name)">您需要先添加收货人信息才可提交订单。</div>
                <ul v-else>
                  <li><span>收<i />货<i />人：</span>{{ orderInfo.buyer_name }}</li>
                  <li><span>联系方式：</span>{{ orderInfo.buyer_contact }}</li>
                  <li v-if="good.transaction_method === 2 || good.transaction_method === 3"><span>收货地址：</span>{{
                    orderInfo.buyer_address }}</li>
                </ul>
              </div>
            </div>
          </div>
        </div>
        <!-- 商品信息 -->
        <h3 class="box-title">商品信息</h3>
        <div class="box-body">
          <table class="good">
            <thead>
              <tr>
                <th width="520">商品信息</th>
                <th width="170">规格</th>
                <th width="170">单价</th>
                <th width="170">实付</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td>
                  <a href="javascript:;" class="info">
                    <img :src="good.cover" alt="">
                    <div class="right">
                      <p>{{ good.full_name }}</p>
                      <p>{{ good.description }}</p>
                    </div>
                  </a>
                </td>
                <td>{{ good.specifications }}</td>
                <td>&yen;{{ good.price }}</td>
                <td>&yen;{{ parseFloat(good.price).toFixed(2) }}</td>
              </tr>
            </tbody>
          </table>
        </div>
        <!-- 交易方式 -->
        <h3 class="box-title">交易方式</h3>
        <div class="box-body">
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
        <!-- 支付方式 -->
        <h3 class="box-title">支付方式</h3>
        <span v-if="orderInfo.payment === 2" style="color:#999">卖家与买家现场交易时付款。</span>
        <span v-else-if="orderInfo.payment === 3" style="color:#999">通过其他途径支付平台将无法保障交易的完成。</span>
        <span v-else-if="orderInfo.payment === 1" style="color:#999">买家线上通过平台付款。</span>
        <span v-else style="color:#999">未选择支付方式。</span>
        <div class="box-body">
          <a class="my-btn" :class="{ active: orderInfo.payment === 1 }" v-if="orderInfo.payment === 1"
            href="javascript:;">在线支付</a>
          <a class="my-btn" :class="{ active: orderInfo.payment === 2 }" v-if="orderInfo.payment === 2"
            href="javascript:;">货到付款</a>
          <a class="my-btn" :class="{ active: orderInfo.payment === 3 }" v-if="orderInfo.payment === 3"
            href="javascript:;">其他途径</a>
        </div>
        <!-- 金额明细 -->
        <h3 class="box-title">金额明细</h3>
        <div class="box-body">
          <div class="total">
            <dl>
              <dt>商品总价：</dt>
              <dd>¥{{ parseFloat(good.price).toFixed(2) }}</dd>
            </dl>
            <dl>
              <dt>创建时间：</dt>
              <dd>{{ orderInfo.created_at }}</dd>
            </dl>
            <dl>
              <dt>应付总额：</dt>
              <dd class="price">¥{{ parseFloat(good.price).toFixed(2) }}</dd>
            </dl>
          </div>
        </div>
      </div>
    </div>
  </div>
  <div v-else>
    <el-empty description="暂无商品数据" />
  </div>
  <!-- 切换地址 -->
  <el-dialog v-model="showDialog" title="切换收货信息" width="30%" center>
    <div class="addressWrapper">
      <div class="text item" :class="{ active: activeAddress.id === item.id }" @click="switchAddress(item)"
        v-for="item in userStore.userInfo.addresses" :key="item.id">
        <ul>
          <li><span>收<i />货<i />人：</span>{{ item.receiver }} </li>
          <li><span>联系方式：</span>{{ item.contact }}</li>
          <li v-if="good.transaction_method === 2 || good.transaction_method === 3"><span>收货地址：</span>{{ item.address }}
          </li>
        </ul>
      </div>
    </div>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="rejectAddress()">取消</el-button>
        <el-button type="primary" @click="confirmAddress()">确定</el-button>
      </span>
    </template>
  </el-dialog>
  <!-- 添加地址 -->
  <el-dialog v-model="showAddAddressDialog" title="添加收货地址" width="30%" center>
    <el-form :rules="addressRules" :model="newAddress" ref="addressForm" label-width="120px">
      <el-form-item label="收货人" prop="receiver">
        <el-input v-model="newAddress.receiver" placeholder="请输入收货人姓名"></el-input>
      </el-form-item>
      <el-form-item label="联系方式" prop="contact">
        <el-input v-model="newAddress.contact" placeholder="请输入联系方式"></el-input>
      </el-form-item>
      <el-form-item label="收货地址" prop="address">
        <el-input v-model="newAddress.address" placeholder="请输入收货地址"></el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="showAddAddressDialog = false">取消</el-button>
        <el-button type="primary" @click="submitNewAddress">保存</el-button>
      </span>
    </template>
  </el-dialog>


  <!-- 添加地址 -->
</template>

<style scoped lang="scss">
.pay-info {

  background: #fff;
  display: flex;
  align-items: center;
  height: 240px;
  padding: 50px 131.4px;

  .icon {
    font-size: 100px;
    width: 100px;
  }

  .green {
    color: #1dc779;
  }

  .red {
    color: $priceColor;
  }

  .tip {
    padding-left: 50px;
    flex: 1;

    p {
      &:nth-child(1) {
        font-size: 20px;
        margin-bottom: 5px;
      }

      &:nth-child(2) {
        color: #999;
        font-size: 16px;
      }

      &:nth-child(3) {
        color: #999;
        font-size: 16px;
      }
    }
  }

  .amount {
    span {
      &:first-child {
        font-size: 16px;
        color: #999;
      }

      &:last-child {
        color: $priceColor;
        font-size: 20px;
      }
    }
  }
}

.pay-checkout-page {
  margin-top: 20px;

  .wrapper {
    background: #fff;
    padding: 0 100px;

    .box-title {
      font-size: 16px;
      font-weight: normal;
      padding-left: 10px;
      line-height: 70px;
      border-bottom: 1px solid #f5f5f5;
    }

    .box-body {
      padding: 20px 0;
    }
  }
}

.address {
  border: 1px solid #f5f5f5;
  display: flex;
  align-items: center;

  .text {
    flex: 1;
    min-height: 90px;
    display: flex;
    align-items: center;

    .none {
      line-height: 90px;
      color: #999;
      text-align: center;
      width: 100%;
    }

    >ul {
      flex: 1;
      padding: 20px;

      li {
        line-height: 30px;

        span {
          color: #999;
          margin-right: 5px;

          >i {
            width: 0.5em;
            display: inline-block;
          }
        }
      }
    }

    >a {
      color: $mainColor;
      width: 160px;
      text-align: center;
      height: 90px;
      line-height: 90px;
      border-right: 1px solid #f5f5f5;
    }
  }

  .action {
    width: 420px;
    text-align: center;

    .btn {
      width: 140px;
      height: 46px;
      line-height: 44px;
      font-size: 14px;

      &:first-child {
        margin-right: 10px;
      }
    }
  }
}

.good {
  width: 100%;
  border-collapse: collapse;
  border-spacing: 0;

  .info {
    display: flex;
    text-align: left;

    img {
      width: 70px;
      height: 70px;
      margin-right: 20px;
    }

    .right {
      line-height: 24px;

      p {
        &:last-child {
          color: #999;
        }
      }
    }
  }

  tr {
    th {
      background: #f5f5f5;
      font-weight: normal;
    }

    td,
    th {
      text-align: center;
      padding: 20px;
      border-bottom: 1px solid #f5f5f5;

      &:first-child {
        border-left: 1px solid #f5f5f5;
      }

      &:last-child {
        border-right: 1px solid #f5f5f5;
      }
    }
  }
}

.my-btn {
  width: 228px;
  height: 50px;
  border: 1px solid #e4e4e4;
  text-align: center;
  line-height: 48px;
  margin-right: 25px;
  color: #666666;
  display: inline-block;
  border-radius: 10px;

  &.active,
  &:hover {
    border-color: $mainColor;
    background: $bgLightColor;
  }
}

.total {
  dl {
    display: flex;
    justify-content: flex-end;
    line-height: 50px;

    dt {
      i {
        display: inline-block;
        width: 2em;
      }
    }

    dd {
      width: 240px;
      text-align: right;
      padding-right: 70px;

      &.price {
        font-size: 20px;
        color: $priceColor;
      }
    }
  }
}

.submit {
  text-align: right;
  padding: 60px;
  border-top: 1px solid #f5f5f5;
}

.addressWrapper {
  max-height: 500px;
  overflow-y: auto;
}

.text {
  flex: 1;
  min-height: 90px;
  display: flex;
  align-items: center;

  &.item {
    border: 1px solid #f5f5f5;
    margin-bottom: 10px;
    cursor: pointer;

    &.active,
    &:hover {
      border-color: $mainColor;
      background: $bgLightColor;
    }

    >ul {
      padding: 10px;
      font-size: 14px;
      line-height: 30px;
    }
  }
}

.transaction-method {
  padding: 20px;
  margin-bottom: 15px;
  border: 1px solid #f5f5f5;
  border-radius: 8px;
  background-color: #fafafa;
  font-size: 16px;
  color: #333;
}

.transaction-method p {
  margin: 5px 0;
  line-height: 24px;
}

.transaction-method span:nth-child(1) {
  font-weight: bold;
  color: #333;
}

.transaction-method span:nth-child(2) {
  margin-left: 20px;
  color: #999;
}

.transaction-method span:nth-child(3) {
  font-weight: bold;
  color: #333;
}
</style>