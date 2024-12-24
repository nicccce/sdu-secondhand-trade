<script setup>
import { useUserStore } from '@/stores/user';
import { computed, onMounted, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { ElForm, ElFormItem, ElInput, ElButton, ElSelect, ElOption, ElRadioGroup, ElRadio } from 'element-plus';
import { useStaticStore } from '@/stores/static';
import { useCategoryStore } from '@/stores/category';
import { addGoodAPI } from '@/apis/good';
import { Plus } from '@element-plus/icons-vue'
import { baseURL } from '@/utils/http';
import axios from 'axios';

// 获取用户信息和路由
const userStore = useUserStore();
const staticStore = useStaticStore();
const categoryStore = useCategoryStore();
const route = useRoute();
const router = useRouter();

// 商品数据
const good = ref({
    name: "", // 商品名称
    description: "", // 商品描述
    price: "", // 商品价格
    campus: userStore.userInfo.campus||null, // 校区信息
    category: null, // 商品分类
    brand: "", // 品牌
    status: "", // 商品状态（如 9 成新）
    full_name: "", // 商品全称
    specifications: "", // 商品规格
    transaction_method: null, // 交易方式
    seller_address: "", // 自提地址
    detail: "", // 商品详细说明
});

// 获取分类选项并填充到categoryOption中
const categoryOption = computed(() => {
    return categoryStore.categoryList.map(category => ({
        label: category.name, // 一级分类的label为name
        value: category.id,   // 一级分类的value为id
        children: category.children ? category.children.map(child => ({
            label: child.name,  // 子分类的label为name
            value: child.id     // 子分类的value为id
        })) : [] // 如果没有子分类，返回空数组
    }));
});

const formRules = {
    name: [
        { required: true, message: '商品名称不能为空', trigger: 'blur' },
        { max: 15, message: '商品名称最多 15 个字符', trigger: 'blur' }
    ],
    full_name: [
        { required: true, message: '商品全称不能为空', trigger: 'blur' },
        { max: 15, message: '商品全称最多 15 个字符', trigger: 'blur' }
    ],
    price: [
        { required: true, message: '价格不能为空', trigger: 'blur' },
    ],
    status: [
        { required: true, message: '状态不能为空', trigger: 'blur' },
        { max: 15, message: '状态最多 15 个字符', trigger: 'blur' }
    ],
    brand: [
        { required: true, message: '品牌不能为空', trigger: 'blur' },
        { max: 15, message: '品牌最多 15 个字符', trigger: 'blur' }
    ],
    category: [
        { required: true, message: '分类不能为空', trigger: 'change' }
    ],
    description: [
        { required: true, message: '商品描述不能为空', trigger: 'blur' },
        { max: 25, message: '商品描述最多 25 个字符', trigger: 'blur' }
    ],
    specifications: [
        { required: true, message: '规格不能为空', trigger: 'blur' },
        { max: 15, message: '规格最多 15 个字符', trigger: 'blur' }
    ],
    detail: [
        { required: true, message: '详细说明不能为空', trigger: 'blur' },
        { max: 50, message: '详细说明最多 50 个字符', trigger: 'blur' }
    ],
    transaction_method: [
        { required: true, message: '交易方式不能为空', trigger: 'change' }
    ],
    seller_address: [
        { required: good.transaction_method === 1, message: '自提地址不能为空', trigger: 'blur' },
        { max: 30, message: '自提地址最多 30 个字符', trigger: 'blur' }
    ],
    campus: [
        { required: true, message: '校区不能为空', trigger: 'change' }
    ],
};

// 格式化价格，添加人民币符号
const formatPrice = (value) => {
    return '￥' + value.replace(/[^\d]/g, ''); // 移除非数字字符
};

// 解析价格，去掉人民币符号
const parsePrice = (value) => {
    if (!value) return '';
    return value.replace(/[^\d]/g, ''); // 移除￥符号
};

const active = ref(0);

const next = () => {
    active.value++;
};

const goodInfoRef = ref(null)
const goodRef = ref(null)
const submitGoodInfo = () => {
    goodInfoRef.value.validate(async (valid) => {
        if (valid) {
            good.value.category = good.value.category[1]
            ElMessage({
                type: 'success',
                message: '商品信息提交成功'
            });
            next();
        }
    });
};
const submitGood = () => {
    goodRef.value.validate(async (valid) => {
        if (valid) {
            try {
                good.value.price = parseInt(good.value.price)
                const res = await addGoodAPI(good.value)
                if (res.code === 0) {
                    good.value.id = res.data.id
                    ElMessage({
                        type: 'success',
                        message: '交易信息提交成功'
                    });
                    next()
                }
            } catch (error) { }
        }
    });

};

const coverUrl = ref("")
const handleCoverSuccess = (response, uploadFile) => {
    // 使用浏览器提供的 URL.createObjectURL 方法创建临时的 URL 来预览图片
    coverUrl.value = URL.createObjectURL(uploadFile.raw);
};

const beforeCoverUpload = (rawFile) => {
    // 校验上传文件的格式，支持 JPEG, PNG 和 GIF 格式
    const allowedTypes = ['image/jpeg', 'image/png', 'image/gif'];
    if (!allowedTypes.includes(rawFile.type)) {
        ElMessage.error('封面格式不合法!');
        return false;  // 阻止文件上传
    } else if (rawFile.size / 1024 / 1024 > 4) {
        ElMessage.error('封面图片大小不得超过 4MB!');
        return false;  // 阻止文件上传
    }
    return true;  // 允许文件上传
};

const pictureList = ref([])
const dialogImageUrl = ref('')
const dialogVisible = ref(false)
const uploadUrl = `${baseURL}/good/picture/${good.value.id}`;

const beforePictureUpload = (rawFile) => {
    // 校验上传文件的格式，支持 JPEG, PNG 和 GIF 格式
    const allowedTypes = ['image/jpeg', 'image/png', 'image/gif'];
    if (!allowedTypes.includes(rawFile.type)) {
        ElMessage.error('图片格式不合法!');
        return false;  // 阻止文件上传
    } else if (rawFile.size / 1024 / 1024 > 5) {
        ElMessage.error('图片大小不得超过 5MB!');
        return false;  // 阻止文件上传
    }
    return true;  // 允许文件上传
};

const handlePictureSuccess = (response, uploadFile) => {
    // 使用浏览器提供的 URL.createObjectURL 方法创建临时的 URL 来预览图片
    pictureList.push(URL.createObjectURL(uploadFile.raw));
};

const handlePictureCardPreview = (uploadFile) => {
    dialogImageUrl.value = uploadFile.url;
    dialogVisible.value = true;
};



const submitPicture = ()=>{
    if(pictureList.value.length===0){
        ElMessage.error('请至少上传一张图片!')
    }else if(!coverUrl.value){
        ElMessage.error('请上传封面!')
    }else{
        ElMessage({
                type: 'success',
                message: '商品创建成功！'
            });
        router.push({path: "/"})
    }
}

onMounted(()=>{
    if(!userStore.userInfo.id){
        router.push({path: "/login"})
    }
})
</script>

<template>
    <div class="sell-page">
        <div class="container">
            <el-steps class="steps" :active="active" finish-status="success">
                <el-step title="上传商品信息" />
                <el-step title="上传交易信息" />
                <el-step title="上传图片" />
            </el-steps>

            <transition name="fade" mode="out-in">
                <div :key="active" class="wrapper">
                    <!-- 商品信息 -->
                    <div v-if="active === 0" class="wrapper">
                        <h3 class="box-title">商品信息</h3>
                        <div class="box-body">
                            <el-form :model="good" :rules="formRules" ref="goodInfoRef" label-width="120px"
                                class="form">
                                <el-form-item label="商品名称" prop="name">
                                    <el-input v-model="good.name" placeholder="请输入商品名称" clearable />
                                </el-form-item>
                                <el-form-item label="商品全称" prop="full_name">
                                    <el-input v-model="good.full_name" placeholder="请输入商品名称" clearable />
                                </el-form-item>
                                <el-form-item label="价格" prop="price">
                                    <el-input v-model="good.price" placeholder="请输入价格" clearable
                                        :formatter="formatPrice" :parser="parsePrice" />
                                </el-form-item>
                                <el-form-item label="状态" prop="brand">
                                    <el-input v-model="good.brand" placeholder="请输入价格（如：9成新）" clearable />
                                </el-form-item>
                                <el-form-item label="品牌" prop="status">
                                    <el-input v-model="good.status" placeholder="请输入价格" clearable />
                                </el-form-item>
                                <el-form-item label="分类" prop="category">
                                    <el-cascader :options="categoryOption" v-model="good.category"
                                        :show-all-levels="false" />
                                </el-form-item>
                                <el-form-item label="商品描述" prop="description">
                                    <el-input v-model="good.description" placeholder="请输入商品描述" clearable type="textarea"
                                        :autosize="{ minRows: 2, maxRows: 5 }" />
                                </el-form-item>
                                <el-form-item label="规格" prop="specifications">
                                    <el-input v-model="good.specifications" placeholder="请输入商品规格" clearable
                                        type="textarea" :autosize="{ minRows: 2, maxRows: 5 }" />
                                </el-form-item>
                                <el-form-item label="详细说明" prop="detail" style="width: 500px;">
                                    <el-input v-model="good.detail" placeholder="请输入商品详细说明" clearable type="textarea"
                                        :autosize="{ minRows: 4, maxRows: 10 }" />
                                </el-form-item>
                            </el-form>
                        </div>
                        <div class="submit">
                            <el-button type="primary" size="large" @click="submitGoodInfo()">提交商品信息</el-button>
                        </div>
                    </div>
                    <!-- 收货信息 -->
                    <div v-if="active === 1" class="wrapper">
                        <h3 class="box-title">交易信息</h3>
                        <div class="box-body">
                            <el-form :model="good" :rules="formRules" ref="goodRef" label-width="120px" class="form">
                                <el-form-item label="交易方式" prop="transaction_method">
                                    <el-select v-model="good.transaction_method" placeholder="请选择交易方式">
                                        <el-option label="其他方式" :value="0" />
                                        <el-option label="买家自提" :value="1" />
                                        <el-option label="快递邮寄" :value="2" />
                                        <el-option label="送货上门" :value="3" />
                                        <el-option label="当面交易" :value="4" />
                                        <el-option label="虚拟商品" :value="5" />
                                    </el-select>
                                </el-form-item>
                                <el-form-item v-if="good.transaction_method === 1" label="自提地址" prop="seller_address">
                                    <el-input v-model="good.seller_address" placeholder="请输入自提地址" />
                                </el-form-item>
                                <el-form-item label="校区" prop="campus">
                                    <el-select v-model="good.campus" placeholder="请选择校区">
                                        <el-option v-for="campus in staticStore.campusList" :key="campus.id"
                                            :label="campus.name" :value="campus.id" />
                                    </el-select>
                                </el-form-item>
                            </el-form>
                        </div>
                        <div class="submit">
                            <el-button type="primary" size="large" @click="submitGood()">提交收货信息</el-button>
                        </div>
                    </div>
                    <!-- 上传图片 -->
                    <div v-if="active === 2" class="wrapper">
                        <h3 class="box-title">商品封面</h3>
                        <el-upload class="cover-uploader" :action="`${baseURL}/good/cover/${good.id}`"
                            :show-file-list="false" :on-success="handleCoverSuccess" :before-upload="beforeCoverUpload"
                            accept="image/*">
                            <img v-if="coverUrl" :src="coverUrl" class="cover" />
                            <el-icon v-else class="cover-uploader-icon">
                                <Plus />
                            </el-icon>
                        </el-upload>
                        <h3 class="box-title">商品图片</h3>
                        <el-upload v-model:file-list="pictureList" :action="`${baseURL}/good/picture/${good.id}`" list-type="picture-card"
                            :on-preview="handlePictureCardPreview" :before-upload="beforePictureUpload"
                            :on-success="handlePictureSuccess" :show-remove="false" :on-remove="!1===1" accept="image/*">
                            <el-icon>
                                <Plus />
                            </el-icon>
                        </el-upload>

                        <!-- 图片预览对话框 -->
                        <el-dialog v-model="dialogVisible">
                            <img class="w-full" :src="dialogImageUrl" alt="Preview Image" />
                        </el-dialog>
                        <div class="submit">
                            <el-button type="primary" size="large" @click="submitPicture()">完成</el-button>
                        </div>
                    </div>
                </div>
            </transition>
        </div>
    </div>
</template>

<style scoped lang="scss">
.sell-page {
    margin-top: 20px;

    .container {
        background: #fff;
        justify-content: center;
        transition: 1s;
    }

    .wrapper {
        width: 100%;
        display: flex;
        flex-wrap: wrap;
        gap: 20px;
        padding: 20px 100px;
    }

    .box-title {
        font-size: 16px;
        font-weight: normal;
        padding-left: 10px;
        line-height: 70px;
        border-bottom: 1px solid #f5f5f5;
        width: 100%;
    }

    .form {
        width: 100%;
        display: flex;
        flex-wrap: wrap;
        gap: 10px;
    }

    .el-form-item {
        width: 450px;
        box-sizing: border-box;
    }

    .submit {
        width: 100%;
        display: flex;
        justify-content: center;
        margin-top: 20px;
    }
}

.steps {
    width: 100%;
    padding-top: 50px;
    padding-left: 30%;
    padding-right: 30%;
}

/* 过渡效果 */
.fade-enter-active,
.fade-leave-active {
    transition: opacity 0.5s ease-in-out, transform 0.5s ease-in-out;
}

.fade-enter,
.fade-leave-to {
    opacity: 0;
    transform: translateX(30px);
}

.cover {
    width: 200px;
    height: 200px;
    object-fit: contain;
}

.cover-uploader .el-upload {
    border: 1px dashed var(--el-border-color);
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
    transition: var(--el-transition-duration-fast);
}

.cover-uploader .el-upload:hover {
    border-color: var(--el-color-primary);
}

.el-icon.cover-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: 200px;
    height: 200px;
    text-align: center;
}

.el-upload-list__item .el-upload-list__item-actions .el-icon-delete {
  display: none ;
}
</style>
