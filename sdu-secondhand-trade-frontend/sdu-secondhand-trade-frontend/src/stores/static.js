import { ref } from 'vue'
import { defineStore } from 'pinia'
import { getGenderAPI, getCampusAPI } from '@/apis/static'

export const useStaticStore = defineStore('static', () => {
    const genderList = ref([])
    const campusList = ref([])

    // 获取性别列表
    const getGenderList = async () => {
        try {
            const res = await getGenderAPI()
            genderList.value = res.data
        } catch (error) {
            console.error('初始化身份信息失败:', error)
        }
    }

    const getCampusList = async () => {
        try {
            const res = await getCampusAPI()
            campusList.value = res.data
        } catch (error) {
            console.error('初始化校区信息失败:', error)
        }
    }

    return {
        genderList,
        campusList,
        getGenderList,
        getCampusList
    }
})
