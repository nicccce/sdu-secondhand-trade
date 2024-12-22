import {
    defineStore
} from 'pinia'
import { loginAPI, refreshAPI } from '@/apis/user'
import { ref } from 'vue'

export const useUserStore = defineStore('user', () => {
    //用户信息
    const userInfo = ref({})

    //登录
    const getUserInfo = async ({ student_id, password }) => {
        const res = await loginAPI({ student_id, password })

        userInfo.value = res.data
    }

    const refreshUserInfo = async()=>{
        const res = await refreshAPI()
        if(res.code===0){
            userInfo.value = res.data
        }
    }

    //退出登录
    const clearUserInfo = () => {
        userInfo.value = {}
    }

    return {
        userInfo,
        getUserInfo,
        clearUserInfo,
        refreshUserInfo
    }
},
    {
        persist: true,
    })