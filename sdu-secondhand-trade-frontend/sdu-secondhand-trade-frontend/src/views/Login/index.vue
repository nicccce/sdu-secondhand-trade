<script setup>
import { computed, ref } from 'vue';
import { ElMessage } from 'element-plus';
import { useRouter } from 'vue-router';
import { useUserStore } from '@/stores/user';
import { useCountDown } from '@/composables/useCountDown';
import { forgetAPI, registerAPI } from '@/apis/user';

// 表单对象
const form = ref({
    student_id: '',
    password: '',
    agree: false,
    confirmPassword: '',
    phone: '',
    code: '',
    nickname: '',
});

// 用于控制显示登录/注册界面
const mod = ref(0); // 0: 登录模式, 1: 注册模式

// 获取当前时间（yyyy-mm-dd hh:mm:ss）
const getCurrentTime = () => {
    const now = new Date();
    const year = now.getFullYear();
    const month = String(now.getMonth() + 1).padStart(2, '0');
    const day = String(now.getDate()).padStart(2, '0');
    const hours = String(now.getHours()).padStart(2, '0');
    const minutes = String(now.getMinutes()).padStart(2, '0');
    const seconds = String(now.getSeconds()).padStart(2, '0');
    return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
};

// 表单规则
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
    confirmPassword: [
        {
            required: true,
            message: '请确认密码！',
            trigger: 'blur'
        },
        {
            validator: (rule, value, callback) => {
                if (value !== form.value.password) {
                    callback(new Error('两次输入的密码不一致！'));
                } else {
                    callback();
                }
            },
            trigger: 'blur'
        }
    ],
    phone: [
        {
            required: true,
            message: '手机号不能为空！',
            trigger: 'blur'
        },
        {
            pattern: /^[0-9]{11}$/,
            message: '请输入有效的手机号！',
            trigger: 'blur'
        }
    ],
    code: [
        {
            required: true,
            message: '验证码不能为空！',
            trigger: 'blur'
        }
    ],
    nickname: [
        {
            required: true,
            message: '用户昵称不能为空！',
            trigger: 'blur'
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

// 登录、注册切换
const showRegister = () => {
    mod.value = 1;
};

const showLogin = () => {
    mod.value = 0;
};

const showForget = () => {
    mod.value = 2;
};

// 登录功能
const userStore = useUserStore();
const formRef = ref(null);
const router = useRouter();
const doLogin = () => {
    const { student_id, password } = form.value;
    formRef.value.validate(async (valid) => {
        if (valid) {
            try {
                await userStore.getUserInfo({ student_id, password });
                ElMessage({
                    type: 'success',
                    message: '登录成功'
                });
                router.replace({ path: '/' });
            } catch (error) {
                ElMessage.error('登录失败');
            }
        }
    });
};

// 注册功能
const doRegister = () => {
    const { student_id, password, phone, code, nickname } = form.value;
    formRef.value.validate(async (valid) => {
        if (valid) {
            try {
                const res = await registerAPI({ student_id, password, phone, code, nickname })
                if (res.code === 0) {
                    ElMessage({
                        type: 'success',
                        message: '注册成功'
                    });
                    console.log(111);
                    showLogin()
                    countVerificationCodeTime.reset()
                }
            } catch (error) {
                ElMessage.error('注册失败');
            }
        }
    });
};

const countVerificationCodeTime = useCountDown()
const countVerificationCode = computed(() => {
    if (countVerificationCodeTime.formatSecond.value == '00') {
        return '发送验证码'
    }
    return `请等待${countVerificationCodeTime.formatSecond.value}s`
})

// 发送验证码（无法实现，只是模拟）
const sendVerificationCode = () => {
    countVerificationCodeTime.start1m()
    ElMessage.success('验证码已发送');
    // 实现倒计时功能
};

const doForget = () => {
    countVerificationCodeTime.reset()
    const { student_id, password, phone, code } = form.value;
    formRef.value.validate(async (valid) => {
        if (valid) {
            try {
                const res = await forgetAPI({ student_id, password, phone, code })
                if (res.code === 0) {
                    ElMessage({
                        type: 'success',
                        message: '重置成功，请返回登录'
                    });
                    console.log(111);
                    showLogin()
                    countVerificationCodeTime.reset()
                }
            } catch (error) {
                ElMessage.error('重置失败');
            }
        }
    });
};
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
                    <a href="javascript:;" v-if="mod === 0">账户登录</a>
                    <a href="javascript:;" v-if="mod === 1">账户注册</a>
                    <a href="javascript:;" v-if="mod === 2">忘记密码</a>
                </nav>
                <div class="account-box" v-if="mod === 0">
                    <div class="form">
                        <el-form ref="formRef" :model="form" :rules="rules" label-position="left" label-width="80px"
                            status-icon>
                            <el-form-item prop="student_id" label="学号">
                                <el-input v-model="form.student_id" />
                            </el-form-item>
                            <el-form-item prop="password" label="密码">
                                <el-input type="password" v-model="form.password" />
                            </el-form-item>
                            <el-form-item prop="agree" label-width="22px">
                                <el-checkbox size="large" v-model="form.agree">
                                    我已同意
                                    <a href="javascript:void(0);" @click="showPrivacyPolicy">隐私条款</a> 和
                                    <a href="javascript:void(0);" @click="showServiceTerms">服务条款</a>
                                </el-checkbox>
                                <a href="javascript:;" @click="showRegister"
                                    style="color: blue; margin-left: 10px;">注册</a>
                                <a href="javascript:;" @click="showForget"
                                    style="color: blue; margin-left: 10px;">忘记密码</a>
                            </el-form-item>
                            <el-button size="large" class="subBtn" @click="doLogin">点击登录</el-button>
                        </el-form>
                    </div>
                </div>

                <div class="register-box" v-if="mod === 1">
                    <a href="javascript:;" @click="showLogin" class="back-to-login"><el-icon>
                            <Back />
                        </el-icon>返回登录</a>
                    <div class="form">
                        <el-form ref="formRef" :model="form" :rules="rules" label-position="right" label-width="80px"
                            status-icon>
                            <el-form-item prop="student_id" label="学号">
                                <el-input v-model="form.student_id" />
                            </el-form-item>
                            <el-form-item prop="nickname" label="昵称">
                                <el-input v-model="form.nickname" />
                            </el-form-item>
                            <el-form-item prop="password" label="密码">
                                <el-input type="password" v-model="form.password" />
                            </el-form-item>
                            <el-form-item prop="confirmPassword" label="确认密码">
                                <el-input type="password" v-model="form.confirmPassword" />
                            </el-form-item>
                            <el-form-item prop="phone" label="手机号">
                                <el-input v-model="form.phone" />
                            </el-form-item>
                            <el-form-item prop="code" label="验证码">
                                <el-input v-model="form.code">
                                    <template #append>
                                        <el-button @click="sendVerificationCode" size="small">{{ countVerificationCode
                                            }}</el-button>
                                    </template>
                                </el-input>
                            </el-form-item>
                            <el-button size="large" class="subBtn" @click="doRegister">注册</el-button>
                        </el-form>
                    </div>
                </div>

                <div class="register-box" v-if="mod === 2">
                    <a href="javascript:;" @click="showLogin" class="back-to-login"><el-icon>
                            <Back />
                        </el-icon>返回登录</a>
                    <div class="form">
                        <el-form ref="formRef" :model="form" :rules="rules" label-position="right" label-width="80px"
                            status-icon>
                            <el-form-item prop="student_id" label="学号">
                                <el-input v-model="form.student_id" />
                            </el-form-item>
                            <el-form-item prop="phone" label="手机号">
                                <el-input v-model="form.phone" />
                            </el-form-item>
                            <el-form-item prop="code" label="验证码">
                                <el-input v-model="form.code">
                                    <template #append>
                                        <el-button @click="sendVerificationCode" size="small">{{ countVerificationCode
                                            }}</el-button>
                                    </template>
                                </el-input>
                            </el-form-item>
                            <el-form-item prop="password" label="新密码">
                                <el-input type="password" v-model="form.password" />
                            </el-form-item>
                            <el-form-item prop="confirmPassword" label="确认密码">
                                <el-input type="password" v-model="form.confirmPassword" />
                            </el-form-item>
                            <el-button size="large" class="subBtn" @click="doForget">重置</el-button>
                        </el-form>
                    </div>
                </div>
            </div>
        </section>
        <footer class="login-footer">
            <div class="container">
                <p>
                    <a href="javascript:;" @click="$router.push('/about')">关于我们</a>
                    <a href="javascript:;">帮助中心</a>
                    <a href="javascript:;">售后服务</a>
                    <a href="javascript:;">商务合作</a>
                    <a href="javascript:;">搜索推荐</a>
                    <a href="https://www.sdu.edu.cn/">友情链接</a>
                </p>
                <p>CopyRight &copy; 淘山大</p>
            </div>
        </footer>
    </div>

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
        transition: all 0.5s ease;

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

.register-box {
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

.back-to-login {
    display: inline-flex;
    align-items: center;
    /* 使图标和文本垂直居中 */
    font-size: 16px;
    /* 设置字体大小 */
    text-decoration: none;
    /* 去除默认的链接下划线 */
    margin-bottom: 20px;
    margin-left: 20px;

    el-icon {
        margin-right: 8px;
        /* 给图标添加右间距 */
        font-size: 18px;
        /* 调整图标大小 */
        color: #409EFF;
        /* 设置图标颜色 */
    }
}
</style>