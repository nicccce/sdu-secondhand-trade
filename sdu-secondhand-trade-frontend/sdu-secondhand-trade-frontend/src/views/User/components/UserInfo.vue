<script setup>
import { updatePasswordAPI } from '@/apis/user';
import { addAddressAPI, deleteAddressAPI, updateAddressAPI, updateUserAPI } from '@/apis/user';
import { useStaticStore } from '@/stores/static';
import { useUserStore } from '@/stores/user';
import { ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';

const userStore = useUserStore()
const staticStore = useStaticStore()

const switcher = ref(0)

const route = useRoute()
const router = useRouter()

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
const showAddressDialog = ref(false)
const dialogAddress = ref({
  id: 0,
  receiver: null,
  contact: null,
  address: null,
})
const addressForm = ref(null)

const selectAddress = (address) => {
  dialogAddress.value.address = address.address
  dialogAddress.value.receiver = address.receiver
  dialogAddress.value.contact = address.contact
  dialogAddress.value.id = address.id
  showAddressDialog.value = true
}
const onDeleteAddress = async (id) => {
  const res = await deleteAddressAPI(id)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '地址删除成功！'
    });
    userStore.refreshUserInfo()
  }
}
const submitAddress = async () => {
  addressForm.value.validate(async (valid) => {
    // 表单验证
    if (!valid) return

    if (dialogAddress.value.id == 0) {
      // 调用接口提交新地址
      const res = await addAddressAPI(dialogAddress.value)
      if (res.code === 0) {
        // 成功提示
        ElMessage({
          type: 'success',
          message: '地址添加成功！'
        });
        // 更新地址列表
        userStore.refreshUserInfo()
        // 关闭对话框
        showAddressDialog.value = false;
      } else {
        // 失败提示
        ElMessage({
          type: 'error',
          message: '添加地址失败，请稍后再试！'
        })
      }
    } else {
      // 调用接口提交新地址
      const res = await updateAddressAPI(dialogAddress.value)
      if (res.code === 0) {
        userStore.refreshUserInfo()
        // 成功提示
        ElMessage({
          type: 'success',
          message: '地址修改成功！'
        });
        dialogAddress.value = {
          id: 0,
          receiver: null,
          contact: null,
          address: null,
        }

        // 关闭对话框
        showAddressDialog.value = false
      } else {
        // 失败提示
        ElMessage({
          type: 'error',
          message: '修改地址失败，请稍后再试！'
        })
      }
    }
  })
}

const editingUserInfo = ref({
})
const isEditing = ref(false)
const userForm = ref(null)
const userRules = {
  nickname: [
    { required: true, message: '请输入收货人姓名', trigger: 'blur' }
  ],
}

const onEditInfo = () => {
  editingUserInfo.value = {
    nickname: userStore.userInfo.nickname,
    phone: userStore.userInfo.phone,
    student_id: userStore.userInfo.student_id,
    alipay: userStore.userInfo.alipay,
    gender: userStore.userInfo.gender,
    campus: userStore.userInfo.campus || null,
    introduction: userStore.userInfo.introduction
  }
  isEditing.value = true
}

const saveChanges = () => {
  userForm.value.validate(async (valid) => {
    // 表单验证
    if (!valid) return
    const res = await updateUserAPI(editingUserInfo.value)
    if (res.code === 0) {
      userStore.refreshUserInfo()
      // 成功提示
      ElMessage({
        type: 'success',
        message: '用户信息修改成功！'
      });

      // 关闭对话框
      isEditing.value = false
    } else {
      // 失败提示
      ElMessage({
        type: 'error',
        message: '用户信息修改失败，请稍后再试！'
      })
    }
  })
}

const cancelEdit = () => {
  isEditing.value = false
}


const passwordRules = {
  oldPassword: [
    { required: true, message: '请输入旧密码', trigger: 'blur' },
  ],
  password: [
    { required: true, message: '密码不能为空！', trigger: 'blur' },
    { min: 6, max: 16, message: '密码必须介于6位至16位之间！', trigger: 'blur' }
  ],
  confirmPassword: [
    { required: true, message: '请确认密码', trigger: 'blur' },
    { min: 6, max: 16, message: '密码必须介于6位至16位之间！', trigger: 'blur' },
  ]
};

const passwordForm = ref(null);
const newPasswordData = ref({
  oldPassword: '',
  password: '',
  confirmPassword: ''
});

// 修改密码逻辑
const submitPasswordChange = async () => {
  passwordForm.value.validate(async (valid) => {
    if (!valid) return;

    // 确保新密码和确认密码一致
    if (newPasswordData.value.password !== newPasswordData.value.confirmPassword) {
      ElMessage({
        type: 'error',
        message: '新密码与确认密码不一致！'
      });
      return;
    }

    // 调用接口提交修改密码请求
    const res = await updatePasswordAPI({
      old_password: newPasswordData.value.oldPassword,
      new_password: newPasswordData.value.password
    });
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '密码修改成功！'
      });
      switcher.value = 0; // 修改密码成功后返回到个人信息页面
    } else {
      ElMessage({
        type: 'error',
        message: '密码修改失败，请稍后再试！'
      });
    }
  });
};


const imageUrl = ref(userStore.userInfo.cover)

const handleAvatarSuccess = (response, uploadFile) => {
  imageUrl.value = URL.createObjectURL(uploadFile.raw)
}

const beforeAvatarUpload = (rawFile) => {
  // 校验上传文件的格式，支持 JPEG, PNG 和 GIF 格式
  const allowedTypes = ['image/jpeg', 'image/png', 'image/gif'];
  if (!allowedTypes.includes(rawFile.type)) {
    ElMessage.error('图片格式不合法!');
    return false;  // 阻止文件上传
  } else if (rawFile.size / 1024 / 1024 > 5) {
    ElMessage.error('图片大小不得超过 5MB!');
    return false;  // 阻止文件上传
  }
  imageUrl.value = URL.createObjectURL(rawFile)
  return true;  // 允许文件上传
};
</script>

<template>
  <div class="home-overview">
    <!-- 用户信息 -->
    <div class="user-meta">
      <div class="avatar">
        <el-upload class="avatar-uploader" action="" :show-file-list="false" :on-success="handleAvatarSuccess"
          :before-upload="beforeAvatarUpload">
          <img v-if="imageUrl" :src="imageUrl"  />
          <img v-else src="https://www.bkjx.sdu.edu.cn/dfiles/17838/css/style1/images/jsfw.jpg" />
        </el-upload>
      </div>
      <h4>{{ userStore.userInfo.nickname }}</h4>
    </div>
    <div class="item">
      <a href="javascript:;" @click="switcher = 0">
        <span class="iconfont icon-aq"></span>
        <p @click="switcher = 0">个人信息</p>
      </a>
      <a href="javascript:;">
        <span class="iconfont icon-dw" @click="switcher = 1"></span>
        <p @click="switcher = 1">地址信息</p>
      </a>
      <a href="javascript:;">
        <span class="iconfont icon-aq" @click="switcher = 2"></span>
        <p @click="">修改密码</p>
      </a>
    </div>
  </div>
  <div class="info-container">
    <div class="user-panel" v-if="switcher === 0">
      <div class="header">
        <h4>我的信息</h4>
        <el-button type="primary" style="margin-right: 100px; font-size: 16px;" @click="onEditInfo()" size="large">
          修改信息
        </el-button>
      </div>

      <el-form :rules="userRules" :model="editingUserInfo" ref="userForm" label-width="100px" v-if="isEditing">
        <el-row>
          <el-col :span="12">
            <el-form-item label="昵称" prop="nickname">
              <el-input v-model="editingUserInfo.nickname" placeholder="请输入昵称"></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="电话" prop="phone">
              <el-input v-model="editingUserInfo.phone" placeholder="请输入电话"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row>
          <el-col :span="12">
            <el-form-item label="学号" prop="student_id">
              <el-input v-model="editingUserInfo.student_id" disabled></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="支付宝" prop="alipay">
              <el-input v-model="editingUserInfo.alipay" placeholder="请输入支付宝账号"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row>
          <el-col :span="12">
            <el-form-item label="性别" prop="gender">
              <el-select v-model="editingUserInfo.gender" placeholder="请选择性别">
                <el-option v-for="gender in staticStore.genderList" :key="gender.id" :label="gender.name"
                  :value="gender.name"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="校区" prop="campus">
              <el-select v-model="editingUserInfo.campus" placeholder="请选择校区">
                <el-option v-for="campus in staticStore.campusList" :key="campus.id" :label="campus.name + '校区'"
                  :value="campus.id"></el-option>
              </el-select>
            </el-form-item>
          </el-col>
        </el-row>

        <el-row>
          <el-col :span="12">
            <el-form-item label="自我介绍" prop="introduction">
              <el-input v-model="editingUserInfo.introduction" placeholder="请输入自我介绍" type="textarea"></el-input>
            </el-form-item>
          </el-col>
        </el-row>

        <div class="actions">
          <el-button type="primary" @click="saveChanges">保存修改</el-button>
          <el-button @click="cancelEdit">取消</el-button>
        </div>
      </el-form>

      <div v-else>
        <div class="info-list">
          <div class="left-column">
            <ul>
              <li><span>昵称：</span>{{ userStore.userInfo.nickname }}</li>
              <li><span>学号：</span>{{ userStore.userInfo.student_id }}</li>
              <li><span>性别：</span>{{ userStore.userInfo.gender }}</li>
              <li><span>自我介绍：</span>{{ userStore.userInfo.introduction }}</li>
            </ul>
          </div>
          <div class="right-column">
            <ul>
              <li><span>电话：</span>{{ userStore.userInfo.phone }}</li>
              <li><span>支付宝：</span>{{ userStore.userInfo.alipay }}</li>
              <li><span>校区：</span>{{ (staticStore.campusList.find(campus => campus.id === userStore.userInfo.campus) ||
                { name: '校外' }).name }}校区</li>
            </ul>
          </div>
        </div>
      </div>
    </div>
    <div class="address-panel" v-if="switcher === 1">
      <div class="header">
        <h4 data-v-bcb266e0="">我的地址</h4>
        <el-button type="primary" style="margin-right: 100px;" @click="showAddressDialog = true">添加地址</el-button>
      </div>
      <div class="addressWrapper">
        <div class="text item" @click="selectAddress(item)" v-for="item in userStore.userInfo.addresses" :key="item.id">
          <ul>
            <li><span>收<i />货<i />人：</span>{{ item.receiver }} </li>
            <li><span>联系方式：</span>{{ item.contact }}</li>
            <li><span>收货地址：</span>{{ item.address }} </li>
          </ul>
          <div class="delete-button" @click.stop>
            <el-popconfirm @confirm="onDeleteAddress(item.id)" title="确认删除吗?" confirm-button-text="确认"
              cancel-button-text="取消">
              <template #reference>
                <el-button type="primary">删除地址</el-button>
              </template>
            </el-popconfirm>
          </div>
        </div>
      </div>
    </div>
    <div class="password-panel" v-if="switcher === 2">
      <div class="header">
        <h4>修改密码</h4>
      </div>

      <el-form :model="newPasswordData" ref="passwordForm" :rules="passwordRules" label-width="120px">
        <el-form-item label="旧密码" prop="oldPassword">
          <el-input v-model="newPasswordData.oldPassword" placeholder="请输入旧密码" type="password"></el-input>
        </el-form-item>

        <el-form-item label="新密码" prop="password">
          <el-input v-model="newPasswordData.password" placeholder="请输入新密码" type="password"></el-input>
        </el-form-item>

        <el-form-item label="确认密码" prop="confirmPassword">
          <el-input v-model="newPasswordData.confirmPassword" placeholder="请确认新密码" type="password"></el-input>
        </el-form-item>

        <div class="actions">
          <el-button type="primary" @click="submitPasswordChange">确认修改</el-button>
          <el-button @click="switcher = 0">取消</el-button>
        </div>
      </el-form>
    </div>
  </div>

  <el-dialog v-model="showAddressDialog" title="收货地址" width="30%" center>
    <el-form :rules="addressRules" :model="dialogAddress" ref="addressForm" label-width="120px">
      <el-form-item label="收货人" prop="receiver">
        <el-input v-model="dialogAddress.receiver" placeholder="请输入收货人姓名"></el-input>
      </el-form-item>
      <el-form-item label="联系方式" prop="contact">
        <el-input v-model="dialogAddress.contact" placeholder="请输入联系方式"></el-input>
      </el-form-item>
      <el-form-item label="收货地址" prop="address">
        <el-input v-model="dialogAddress.address" placeholder="请输入收货地址"></el-input>
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="showAddressDialog = false">取消</el-button>
        <el-button type="primary" @click="submitAddress">保存</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<style scoped lang="scss">
.home-overview {
  height: 150px;
  background: url(@/assets/images/center-bg.jpg) no-repeat center / cover;
  display: flex;

  .user-meta {
    flex: 1;
    display: flex;
    align-items: center;

    .avatar {
      width: 85px;
      height: 85px;
      border-radius: 50%;
      overflow: hidden;
      margin-left: 60px;
      background-color: rgba(255, 255, 255, 0);

      img {
        width: 100%;
        height: 100%;
        background-color: rgba(255, 255, 255, 0);
      }
    }

    h4 {
      padding-left: 26px;
      font-size: 18px;
      font-weight: normal;
      color: white;
    }
  }

  .item {
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: space-around;

    &:first-child {
      border-right: 1px solid #f4f4f4;
    }

    a {
      color: white;
      font-size: 16px;
      text-align: center;

      .iconfont {
        font-size: 32px;
      }

      p {
        line-height: 32px;
      }
    }
  }
}

.info-container {
  margin-top: 20px;
  border-radius: 4px;
  background-color: #fff;
}

.user-panel {
  background-color: #fff;
  padding: 0 20px;
  margin-top: 20px;
  height: 500px;

  .header {
    margin-bottom: 30px;
    height: 66px;
    border-bottom: 1px solid #f5f5f5;
    padding: 18px 0;
    display: flex;
    justify-content: space-between;
    align-items: baseline;

    h4 {
      font-size: 22px;
      font-weight: 400;
    }

  }

  .actions {
    margin-left: 40%;
  }

  .el-row {
    margin-bottom: 20px;
  }

  .el-form-item {
    margin-right: 80px;
  }


  .info-list {

    display: flex;
    justify-content: space-between;
    margin-top: 20px;

    .left-column,
    .right-column {
      width: 48%;
    }

    .left-column {
      padding-left: 50px;
    }

    ul {
      list-style-type: none;
      padding: 0;
      font-size: 20px;
    }

    li {
      margin-bottom: 20px;
    }

    span {
      font-weight: bold;
    }
  }
}



.address-panel {
  background-color: #fff;
  padding: 0 20px;
  margin-top: 20px;
  height: 400px;

  .header {
    height: 66px;
    border-bottom: 1px solid #f5f5f5;
    padding: 18px 0;
    display: flex;
    justify-content: space-between;
    align-items: baseline;

    h4 {
      font-size: 22px;
      font-weight: 400;
    }

  }

  .addressWrapper {
    max-height: 500px;
    overflow-y: auto;

    .text {
      flex: 1;
      min-height: 90px;
      display: flex;
      align-items: center;


      &.item {
        border: 1px solid #f5f5f5;
        margin-bottom: 10px;
        cursor: pointer;

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

      .delete-button {
        margin-left: auto;
        margin-right: 100px;
      }
    }

  }
}

.password-panel {
  background-color: #fff;
  padding: 50px 400px;
  margin-top: 20px;

  .header {
    height: 66px;
    border-bottom: 1px solid #f5f5f5;
    padding: 18px 0;
    display: flex;
    justify-content: space-between;
    align-items: baseline;

    h4 {
      font-size: 22px;
      font-weight: 400;
    }
  }

  .actions {
    margin-top: 30px;
    display: flex;
    justify-content: flex-end;

    el-button {
      margin-left: 10px;
    }
  }

  .el-form-item {
    margin-bottom: 20px;
  }
}
</style>
