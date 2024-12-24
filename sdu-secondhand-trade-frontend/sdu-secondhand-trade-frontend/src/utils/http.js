import router from '@/router'
import { useUserStore } from '@/stores/user'
import axios from 'axios'
import { ElMessage } from 'element-plus'
import 'element-plus/theme-chalk/el-message.css'

export const baseURL = 'http://47.98.214.174:8081'

// 创建axios实例
const httpInstance = axios.create({
  baseURL: baseURL,
  timeout: 5000
})

// axios 请求拦截器
httpInstance.interceptors.request.use(config => {
  // 1. 从pinia获取token数据
  const userStore = useUserStore()
  // 2. 按照后端的要求拼接token数据
  const token = userStore.userInfo.token
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
}, e => Promise.reject(e))

// axios 响应拦截器
httpInstance.interceptors.response.use(res => {
  // 如果code不为0，认为请求失败
  if (res.data.code !== 0) {
    // 显示错误提示
    ElMessage({
      type: 'error',
      message: res.data.msg || '服务器无响应！' // 如果没有msg，则使用默认消息
    })
    return Promise.reject(new Error(res.data.msg || '服务器无响应！'))
  }
  // 如果code为0，表示请求成功，直接返回数据
  return res.data
}, e => {
  const userStore = useUserStore()
  // 处理请求错误
  if (e.response.status === 401) {
    ElMessage({
      type: 'warning',
      message: '请先登录！'
    })
    userStore.clearUserInfo()
    router.push('/login')
  } else if (e.response) {
    ElMessage({
      type: 'warning',
      message: e.response.data.msg || '服务器错误！'
    })
  } else {
    ElMessage({
      type: 'error',
      message: '网络错误或请求超时'
    })
  }
  return Promise.reject(e)
})

export default httpInstance
