import {
    defineStore
} from 'pinia'
import { loginAPI } from '@/apis/user'
import { ref } from 'vue'

export const useUserStore = defineStore('user', () => {
    //用户信息
    const userInfo = ref({})

    //登录
    const getUserInfo = async ({ student_id, password }) => {
        const res = await loginAPI({ student_id, password })
        console.log(res);

        userInfo.value = res.data
    }

    //退出登录
    const clearUserInfo = () => {
        userInfo.value = {}
    }

    return {
        userInfo,
        getUserInfo,
        clearUserInfo
    }
},
    {
        persist: true,
    })