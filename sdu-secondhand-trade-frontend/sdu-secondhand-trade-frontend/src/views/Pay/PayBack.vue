<script setup>
import { getGoodDetailAPI } from '@/apis/good';
import { getOrderAPI } from '@/apis/order';
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
})
</script>


<template>
    <div class="pay-page" v-if="good.id">
        <div class="container">
            <!-- 支付结果 -->
            <div class="pay-result">
                <span class="iconfont icon-queren2 green" v-if="orderInfo.status === 2"></span>
                <span class="iconfont icon-shanchu red" v-else></span>
                <p class="tit" v-if="orderInfo.payment === 1">支付{{ orderInfo.status === 2 ? '成功' : '失败' }}</p>
                <p class="tit" v-else>交易{{ orderInfo.status === 2 ? '成功' : '失败' }}</p>
                <p class="tip" v-if="orderInfo.status === 2">卖家将尽快与您联系，请保持手机畅通，如遇到问题，请向平台反馈</p>
                <p class="tip" v-else>如遇到问题，请向平台反馈</p>
                <p>交易方式：
                    <span v-if="orderInfo.payment === 1">支付宝</span>
                    <span v-if="orderInfo.payment === 2">货到付款</span>
                    <span v-else>其他方式</span></p>
                <p>交易金额：<span>¥{{ parseFloat(good.price).toFixed(2) }}</span></p>
                <div class="btn">
                    <el-button type="primary" style="margin-right:20px">查看订单</el-button>
                    <el-button>进入首页</el-button>
                </div>
                <p class="alert">
                    <span class="iconfont icon-tip"></span>
                    温馨提示：淘山大平台不会以订单异常、系统升级为由要求您点击任何网址链接进行退款操作，保护资产、谨慎操作。
                </p>
            </div>
        </div>
    </div>
</template>

<style scoped lang="scss">
.pay-result {
    padding: 100px 0;
    background: #fff;
    text-align: center;
    margin-top: 20px;

    >.iconfont {
        font-size: 100px;
    }

    .green {
        color: #1dc779;
    }

    .red {
        color: $priceColor;
    }

    .tit {
        font-size: 24px;
    }

    .tip {
        color: #999;
    }

    p {
        line-height: 40px;
        font-size: 16px;
    }

    .btn {
        margin-top: 50px;
    }

    .alert {
        font-size: 12px;
        color: #999;
        margin-top: 50px;
    }
}
</style>