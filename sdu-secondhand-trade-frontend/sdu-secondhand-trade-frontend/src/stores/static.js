import { ref } from 'vue'
import { defineStore } from 'pinia'
import { getGenderAPI, getCampusAPI } from '@/apis/static'

export const useStaticStore = defineStore('static', () => {
    const genderList = ref([])
    const campusList = ref([])

    const getTransactionMethod = (type) => {
        switch (type) {
            case 1:
                return '买家自提';
            case 2:
                return '快递邮寄';
            case 3:
                return '送货上门';
            case 4:
                return '当面交易';
            case 5:
                return '虚拟商品';
            default:
                return '其他'; // 默认返回其他
        }
    }

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
            for(let campus in campusList.value){
                campus.id = parseInt(campus.id)
            }
        } catch (error) {
            console.error('初始化校区信息失败:', error)
        }
    }

    return {
        genderList,
        campusList,
        getGenderList,
        getCampusList,
        getTransactionMethod,
    }
})
