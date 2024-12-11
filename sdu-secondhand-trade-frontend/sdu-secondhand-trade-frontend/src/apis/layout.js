import httpInstance from "@/utils/http";

export const getCategoryAPI = () => httpInstance({
    url: '/good/category'
});

export const getCategoryGoodsAPI = (categoryName) => {
    return httpInstance({
        url: `/good/category/${categoryName}`,
        method: 'GET'
    })
}