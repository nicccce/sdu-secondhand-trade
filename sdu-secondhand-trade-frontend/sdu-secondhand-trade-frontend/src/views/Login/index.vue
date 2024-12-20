<script setup>
import { ref } from 'vue';
import { ElMessage } from 'element-plus'
import 'element-plus/theme-chalk/el-message.css'
import { useRouter } from 'vue-router';
import { useUserStore } from '@/stores/user';

// 表单对象
const form = ref({
    student_id: '',
    password: '',
    agree: false
});

// 规则对象
const rules = {
    student_id: [
        {
            required: true,
            message: '学号不能为空！',
            trigger: 'blur'
        },
        {
            pattern: /^[0-9]+$/,
            message: '学号必须是数字！',
        }
    ],
    password: [
        {
            required: true,
            message: '密码不能为空！',
            trigger: 'blur'
        },
        {
            min: 6,
            max: 16,
            message: '密码必须介于6位至16位之间！',
            trigger: 'blur'
        }
    ],
    agree: [
        {
            validator: (rule, value, callback) => {
                if (value) {
                    callback();
                } else {
                    callback(new Error("请勾选协议！"));
                }
            }
        }
    ]
};

// 弹框控制
const dialogVisible = ref(false); // 隐私条款弹框
const dialogVisibleService = ref(false); // 服务条款弹框

const showPrivacyPolicy = () => {
    dialogVisible.value = true;
};

const showServiceTerms = () => {
    dialogVisibleService.value = true;
};

//登录
const userStore = useUserStore()
const formRef = ref(null)
const router = useRouter()
const doLogin = () => {
    const { student_id, password } = form.value;
    formRef.value.validate(async (valid) => {
        if (valid) {
            try {
                await userStore.getUserInfo({ student_id, password });  // 获取用户信息
                ElMessage({
                    type: 'success',
                    message: '登录成功'
                });
                router.replace({ path: '/' });  // 跳转到首页
            } catch (error) {}
        }
    });
}

</script>


<template>
    <div>
        <header class="login-header">
            <div class="container m-top-20">
                <h1 class="logo">
                    <RouterLink to="/">淘山大</RouterLink>
                </h1>
                <RouterLink class="entry" to="/">
                    进入网站首页
                    <i class="iconfont icon-angle-right"></i>
                    <i class="iconfont icon-angle-right"></i>
                </RouterLink>
            </div>
        </header>

        <section class="login-section">
            <div class="wrapper">
                <nav>
                    <a href="javascript:;">账户登录</a>
                </nav>
                <div class="account-box">
                    <div class="form">
                        <el-form ref="formRef" :model="form" :rules="rules" label-position="right" label-width="60px"
                            status-icon>
                            <el-form-item prop="student_id" label="学号">
                                <el-input v-model="form.student_id" />
                            </el-form-item>
                            <el-form-item prop="password" label="密码">
                                <el-input v-model="form.password" />
                            </el-form-item>
                            <el-form-item prop="agree" label-width="22px">
                                <el-checkbox size="large" v-model="form.agree">
                                    我已同意
                                    <a href="javascript:void(0);" @click="showPrivacyPolicy">隐私条款</a> 和
                                    <a href="javascript:void(0);" @click="showServiceTerms">服务条款</a>
                                </el-checkbox>
                            </el-form-item>
                            <el-button size="large" class="subBtn" @click="doLogin">点击登录</el-button>
                        </el-form>
                    </div>
                </div>
            </div>
        </section>


        <footer class="login-footer">
            <div class="container">
                <p>
                    <a href="javascript:;">关于我们</a>
                    <a href="javascript:;">帮助中心</a>
                    <a href="javascript:;">售后服务</a>
                    <a href="javascript:;">商务合作</a>
                    <a href="javascript:;">搜索推荐</a>
                    <a href="javascript:;">友情链接</a>
                </p>
                <p>CopyRight &copy; 淘山大</p>
            </div>
        </footer>

        <el-dialog v-model="dialogVisible" width="40%">
            <template #header>
                <div class="el-dialog__header">
                    <span>隐私条款</span>
                </div>
            </template>
            <p class="dialog-content">本隐私条款旨在阐述我们（淘山大 平台运营方）如何收集、使用、存储、保护及分享您的个人信息。我们承诺将严格遵守相关法律法规，确保您的个人信息安全。
                当您注册成为我们平台的用户时，我们需要您提供必要的个人信息，包括但不限于姓名、联系方式（如手机号码、电子邮箱）、学号（用于实名认证）等。
            </p>
            <p class="dialog-content">
                在您使用平台服务过程中，我们可能会收集您的交易记录、浏览记录、设备信息、位置信息等数据，以便为您提供更优质的服务。</p>
            <template #footer class="el-dialog__footer">
                <el-button @click="dialogVisible = false">关闭</el-button>
            </template>
        </el-dialog>

        <el-dialog v-model="dialogVisibleService" width="40%">
            <template #header>
                <div class="el-dialog__header">
                    <span>服务条款</span>
                </div>
            </template>
            <p class="dialog-content">我们（淘山大 平台运营方）为您提供一个二手商品交易平台，您可以在此发布商品信息、浏览商品、进行交易等。</p>
            <p class="dialog-content">
                您应确保提供的商品信息真实、准确、完整，不得发布虚假信息或误导性内容。</p>
            <p class="dialog-content">
                您应遵守相关法律法规和平台规定，不得进行违法、违规或损害他人利益的行为。</p>
            <p class="dialog-content">
                交易双方应自行协商交易价格、支付方式、发货方式等事项。</p>
            <p class="dialog-content">
                交易完成后，双方应及时确认交易结果，以便平台对交易进行记录和统计。</p>
            <template #footer class="el-dialog__footer">
                <el-button @click="dialogVisibleService = false">关闭</el-button>
            </template>
        </el-dialog>
    </div>
</template>


<style scoped lang="scss">
.login-header {
    background: #fff;
    border-bottom: 1px solid #e4e4e4;

    .container {
        display: flex;
        align-items: flex-end;
        justify-content: space-between;
    }

    .logo {
        width: 600px;

        a {
            display: block;
            height: 130px;
            width: 100%;
            text-indent: -9999px;
            background: url("@/assets/images/logo11.png") no-repeat center 30px / contain;
            background-size: 360px auto;
        }
    }

    .sub {
        flex: 1;
        font-size: 24px;
        font-weight: normal;
        margin-bottom: 38px;
        margin-left: 20px;
        color: #666;
    }

    .entry {
        width: 330px;
        margin-bottom: 50px;
        font-size: 18px;

        i {
            font-size: 14px;
            color: $mainColor;
            letter-spacing: -5px;
        }
    }
}

.login-section {
    background: url('@/assets/images/login-bg.png') no-repeat center / cover;
    height: 620px;
    position: relative;

    .wrapper {
        border-radius: 20px;
        width: 400px;
        background: #fff;
        position: absolute;
        left: 57%;
        top: 150px;
        transform: translate3d(100px, 0, 0);
        box-shadow: 0 0 20px rgba(0, 0, 0, 0.15);

        nav {
            font-size: 18px;
            height: 60px;
            margin-bottom: 20px;
            border-bottom: 1px solid #f5f5f5;
            display: flex;
            padding: 0 40px;
            text-align: right;
            align-items: center;

            a {
                flex: 1;
                line-height: 1;
                display: inline-block;
                font-size: 20px;
                position: relative;
                text-align: center;
            }
        }
    }
}

.login-footer {
    padding: 30px 0 50px;
    background: #fff;

    p {
        text-align: center;
        color: #999;
        padding-top: 20px;

        a {
            line-height: 1;
            padding: 0 10px;
            color: #999;
            display: inline-block;

            ~a {
                border-left: 1px solid #ccc;
            }
        }
    }
}

.account-box {
    .toggle {
        padding: 15px 40px;
        text-align: right;

        a {
            color: $mainColor;

            i {
                font-size: 14px;
            }
        }
    }

    .form {
        padding: 0 20px 20px 20px;
    }

    .subBtn {
        background: $mainColor;
        width: 100%;
        color: #fff;
    }
}

.dialog-content {
    font-size: 18px;
    /* 设置字体大小 */
    line-height: 2;
    /* 设置行距 */
    text-indent: 2em;
    /* 设置首行缩进 */
    color: #333;
    /* 设置字体颜色 */
    margin-bottom: 15px;
    /* 设置段落间距 */
}

/* 让弹出框居中，宽度更小 */
.el-dialog {
    width: 350px;
    /* 设置弹出框的宽度 */
    margin: auto;
    /* 弹出框居中 */
}

.el-dialog__header {
    text-align: center;
    /* 标题居中 */
    font-weight: bold;
    font-size: 20px;
    /* 设置标题字体大小 */
}

.el-dialog__footer {
    display: flex;
    justify-content: center;
    /* 居中内容 */
    align-items: center;
    /* 垂直居中 */
}

.el-dialog__body {
    padding: 20px;
}
</style>