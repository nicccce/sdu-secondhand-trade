import httpInstance from "@/utils/http";

export const getBannerAPI = () => {
    return httpInstance({
        url: '/good/banner'
    })
}
export const getNewGoodsAPI = () => {
    return httpInstance({
        url: `/good/new`,
        method: 'GET'
    })
}
export const getCampusGoodsAPI = (campus) => {
    return httpInstance({
        url: `/good/campus/${campus}`,
        method: 'GET'
    })
}