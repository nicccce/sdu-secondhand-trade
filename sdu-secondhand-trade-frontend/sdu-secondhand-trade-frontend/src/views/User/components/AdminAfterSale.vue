<script setup>
import { deleteGoodAPI, getGoodDetailAPI, getMyGoodAPI } from '@/apis/good';
import { useRouter } from 'vue-router';
import { ref, onMounted } from 'vue';
import { ElMessage, ElDialog } from 'element-plus';
import { getAfterSaleAPI, getAllAfterSaleAPI, getOrderAPI, updateAfterSaleAPI } from '@/apis/order';

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

// 当前选择的售后问题
const currentProblem = ref(null);

// 弹窗显示状态
const dialogVisible = ref(false);

// 获取售后列表
const getProblemList = async () => {
    const res = await getAllAfterSaleAPI(reqData.value);
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

const updateForm = ref(null);
const updateRules = {
    status: [
        { required: true, message: '请选择售后状态', trigger: 'change' }
    ],
    reply: [
        { required: true, message: '请输入回复内容', trigger: 'blur' }
    ]
};
// 修改售后的状态和回复
const updateAfterSale = async () => {
    updateForm.value.validate(async (valid) => {
        if (!valid) return;
        const res = await updateAfterSaleAPI(currentProblem.data);
        if (res.code === 0) {
            ElMessage.success('售后信息已更新');
            dialogVisible.value = false; // 关闭弹窗
            getProblemList(); // 刷新列表
        } else {
            ElMessage.error('更新失败');
        }
    })
};

// 格式化状态
const formateStatus = (status) => {
    const statusMap = {
        0: "待解决",
        1: "已受理",
        2: "已解决"
    }
    return statusMap[status]
}

// 打开修改弹窗
const openEditDialog = (problem) => {
    currentProblem.value = { ...problem }; // 复制当前问题信息

    dialogVisible.value = true; // 显示弹窗
};

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
                                <div class="column action">
                                    <el-button @click="openEditDialog(problem)" type="primary">编辑</el-button>
                                </div>
                            </div>
                        </div>
                    </div>

                    <!-- 分页 -->
                    <div class="pagination-container">
                        <el-pagination :total="total" :page-size="reqData.page_size" @current-change="pageChange"
                            background layout="prev, pager, next" />
                    </div>
                </div>
            </div>
        </el-tabs>

        <!-- 修改售后的弹窗 -->
        <el-dialog title="编辑售后" v-model="dialogVisible" width="50%" center>
            <el-form ref="updateForm" :rules="updateRules" :model="currentProblem" label-width="100px">
                <el-form-item label="订单编号">
                    <span>{{ currentProblem.order_id }}</span>
                </el-form-item>

                <el-form-item label="问题描述">
                    <span>{{ currentProblem.problem }}</span>
                </el-form-item>

                <el-form-item label="处理状态" prop="status">
                    <el-select v-model="currentProblem.status" placeholder="请选择处理状态">
                        <el-option label="待解决" :value=0 />
                        <el-option label="已受理" :value=1 />
                        <el-option label="已解决" :value=2 />
                    </el-select>
                </el-form-item>

                <el-form-item label="回复内容" prop="reply">
                    <el-input type="textarea" v-model="currentProblem.reply" placeholder="请输入回复内容"
                        :autosize="{ minRows: 4, maxRows: 10 }" />
                </el-form-item>
            </el-form>
            <template #footer>
                <span slot="footer" class="dialog-footer">
                    <el-button @click="dialogVisible = false">取消</el-button>
                    <el-button type="primary" @click="updateAfterSale">保存</el-button>
                </span>
            </template>
        </el-dialog>
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
