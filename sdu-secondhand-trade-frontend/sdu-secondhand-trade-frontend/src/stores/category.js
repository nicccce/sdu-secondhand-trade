import {
    ref,
    computed
} from 'vue'
import {
    defineStore
} from 'pinia'
import {
    getCategoryAPI,
    getCategoryGoodsAPI
} from '@/apis/layout'

export const useCategoryStore = defineStore('category', () => {
    const categoryList = ref([])

    const getCategory = async () => {
        const res = await getCategoryAPI()
        categoryList.value = res.data

        for (let category of categoryList.value) {
                const goodsRes = await getCategoryGoodsAPI(category.id)
                category.goods = goodsRes.data || []
                for(let child of category.children) {
                    const childGoodsRes = await getCategoryGoodsAPI(child.id)
                    child.goods=childGoodsRes.data
                }
        }
    }

    return {
        categoryList,
        getCategory
    }
})
