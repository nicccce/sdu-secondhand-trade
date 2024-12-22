<script setup>
import { getAllUsersAPI, getUserByIdAPI } from '@/apis/user';
import { useStaticStore } from '@/stores/static';
import { ref, onMounted, computed } from 'vue';

// 模拟的静态数据
const staticStore = useStaticStore()

const userList = ref([]);
const searchName = ref('');
const selectedCampus = ref(null);

// 获取用户列表并且加载用户详情
const getUserList = async () => {
  const res = await getAllUsersAPI();
  userList.value = res.data;
  // 获取每个用户的详细信息
  for (let user of userList.value) {
    const userRes = await getUserByIdAPI(user.id);
    user.info = userRes.data;
    // 格式化校区
    user.info.campusName = `${(staticStore.campusList.find(campus => campus.id === user.info.campus) || { name: '校外' }).name}校区`;
  }
};

// 过滤用户列表
const filteredUserList = computed(() => {
  return userList.value.filter(user => {
    return (
      ((user.info.nickname.toLowerCase().includes(searchName.value.toLowerCase())) ||
      (user.info.student_id.toLowerCase().includes(searchName.value.toLowerCase())) ||
      (user.info.phone.includes(searchName.value.toLowerCase()))) &&
      (selectedCampus.value ? user.info.campus === selectedCampus.value : true)
    );
  });
});

onMounted(() => {
  getUserList();
});

// 修改用户的功能方法（暂时空着）
const handleEdit = (userId) => {
  console.log('Edit user with ID:', userId);
};

const clearSearch = () => {
    selectedCampus.value=null
    searchName.value=""
}
</script>

<template>
  <div style="padding: 20px 60px;">
    <!-- 搜索和筛选 -->
    <div class="filter-container">
        <h3>约束条件：</h3>
      <el-input v-model="searchName" placeholder="搜索" style="width: 250px; margin-right: 20px;" />
      <el-select v-model="selectedCampus" placeholder="选择校区" style="width: 250px;">
        <el-option v-for="campus in staticStore.campusList" :key="campus.id" :label="`${campus.name}校区`" :value="campus.id" />
      </el-select>
      <el-button type="primary" @click="clearSearch">清除条件</el-button>
    </div>

    <!-- 表格展示 -->
    <el-table :data="filteredUserList" style="width: 100%" stripe border height="500">
      <el-table-column fixed prop="info.nickname" label="昵称" width="180" align="center"/>
      <el-table-column fixed prop="info.student_id" label="学号" width="180" align="center"/>
      <el-table-column prop="info.gender" label="性别" width="100" align="center"/>
      <el-table-column prop="info.campusName" label="校区" width="220" align="center"/>
      <el-table-column prop="info.phone" label="电话" width="220" align="center"/>
      <el-table-column prop="info.alipay" label="支付宝" width="220" align="center"/>
      <el-table-column prop="info.introduction" label="自我介绍" width="320" align="center"/>
    </el-table>
  </div>
</template>

<style scoped lang="scss">
.filter-container {
  margin-bottom: 30px;
  display: flex;
  align-items: center;
  gap: 20px;
}

.filter-container .el-input,
.filter-container .el-select {
  font-size: 16px;
  padding: 10px 15px;
  height: 50px;
}

.el-table {
  font-size: 18px;
  line-height: 2; 
}

.el-table .cell {
  padding: 20px; /* 增加单元格内边距 */
  text-align: center; /* 让单元格内容居中 */
}

.el-button {
  font-size: 16px; /* 按钮字体 */
  padding: 8px 16px; /* 按钮内边距 */
  text-align: center; /* 确保按钮文本居中 */
}

</style>
