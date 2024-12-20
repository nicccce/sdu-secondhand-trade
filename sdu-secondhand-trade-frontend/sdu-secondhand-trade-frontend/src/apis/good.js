import httpInstance from "@/utils/http"

export const getGoodDetailAPI = (goodId) => {
    return httpInstance({
        url: `/good/${goodId}`,
        method: 'GET'
    })
}