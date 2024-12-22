<script setup>
import { getGoodDetailAPI } from '@/apis/good';
import { getOrderAPI } from '@/apis/order';
import { useCountDown } from '@/composables/useCountDown';
import { onMounted, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';

const orderInfo = ref({})
const good = ref({})

const route = useRoute()
const router = useRouter()
const getOderInfo = async () => {
  const res = await getOrderAPI(route.params.id)
  orderInfo.value = res.data
}

const getGood = async () => {
  const res = await getGoodDetailAPI(orderInfo.value.good_id)
  good.value = res.data
}

onMounted(async () => {
  await getOderInfo()
  await getGood()
  startCountDown()
})

//支付宝
const baseURL = 'http://127.0.0.1:5173/'//后端用于支付的服务器
const backURL = `http://127.0.0.1:5173/pay/callback/${route.params.id}`
const redirectUrl = encodeURIComponent(backURL)
const payUrl = `${baseURL}pay/aliPay?order_id=${route.params.id}&redirect=${redirectUrl}`

const countDown = useCountDown()

const startCountDown = () => {
  countDown.start(orderInfo.value.created_at)
}
</script>


<template>
  <div class="pay-page" v-if="good.id">
    <div class="container">
      <!-- 付款信息 -->
      <div class="pay-info">
        <span class="icon iconfont icon-queren2"></span>
        <div class="tip">
          <p>订单提交成功！请尽快完成支付。</p>
          <p v-if="countDown.formatTime.value === '00分00秒' || orderInfo.status === 4">支付还剩 <span
              style="color: red;">00分00秒</span>,订单已自动取消</p>
          <p v-else>支付还剩 <span>{{ countDown.formatTime }}</span>, 超时后将取消订单</p>
        </div>
        <div class="amount">
          <span>应付总额：</span>
          <span>¥{{ parseFloat(good.price).toFixed(2) }}</span>
        </div>
      </div>
      <!-- 付款方式 -->
      <div class="pay-type">
        <p class="head">选择以下支付方式付款</p>
        <div class="item">
          <p>支付平台</p>
          <a class="btn wx" href="javascript:;"></a>
          <a class="btn alipay" :href="backURL"></a>
        </div>
        <div class="item">
          <p>支付方式</p>
          <a class="btn" href="javascript:;">招商银行</a>
          <a class="btn" href="javascript:;">工商银行</a>
          <a class="btn" href="javascript:;">建设银行</a>
          <a class="btn" href="javascript:;">农业银行</a>
          <a class="btn" href="javascript:;">交通银行</a>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss">
.pay-page {
  margin-top: 20px;
}

.pay-info {

  background: #fff;
  display: flex;
  align-items: center;
  height: 240px;
  padding: 50px 131.4px;

  .icon {
    font-size: 100px;
    color: #1dc779;
    width: 100px;
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

.pay-type {
  margin-top: 20px;
  background-color: #fff;
  padding: 50px 80px;
  padding-bottom: 100px;

  p {
    line-height: 70px;
    height: 70px;
    padding-left: 30px;
    font-size: 16px;

    &.head {
      border-bottom: 1px solid #f5f5f5;
    }
  }

  .btn {
    width: 150px;
    height: 50px;
    border: 1px solid #e4e4e4;
    text-align: center;
    line-height: 48px;
    margin-left: 30px;
    color: #666666;
    display: inline-block;
    border-radius: 10px;

    &.active,
    &:hover {
      border-color: $mainColor;
      background: $bgLightColor;
    }

    &.alipay {
      background: url(https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/7b6b02396368c9314528c0bbd85a2e06.png) no-repeat center / contain;
    }

    &.wx {
      background: url(https://cdn.cnbj1.fds.api.mi-img.com/mi-mall/c66f98cff8649bd5ba722c2e8067c6ca.jpg) no-repeat center / contain;
    }
  }
}
</style>