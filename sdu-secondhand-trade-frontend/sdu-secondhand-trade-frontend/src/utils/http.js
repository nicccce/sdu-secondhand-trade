import axios from 'axios'

//创建axios实例
const httpInstance = axios.create({
  baseURL: 'http://127.0.0.1:4523/m1/5399366-0-default',
  timeout: 5000
})

// axios请求拦截器
httpInstance.interceptors.request.use(config => {
  return config
}, e => Promise.reject(e))

// axios响应式拦截器
httpInstance.interceptors.response.use(res => res.data, e => {
  return Promise.reject(e)
})

export default httpInstance
