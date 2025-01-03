import httpInstance from "@/utils/http"

export const getGoodDetailAPI = (goodId) => {
    return httpInstance({
        url: `/good/${goodId}`,
        method: 'GET'
    })
}

export const addGoodAPI = (data) => {
    return httpInstance({
        url: `/good/sell`,
        method: 'POST',
        data
    })
}

export const getMyGoodAPI = (data) => httpInstance({
    url: '/good/my',
    method: 'POST',
    data
})

export const getAllGoodAPI = (data) => httpInstance({
    url: '/good/all',
    method: 'POST',
    data
})

export const deleteGoodAPI = (id) => httpInstance({
    url: `/good/${id}`,
    method: 'DELETE'
})

export const searchGoodAPI = (name) => httpInstance({
    url: `/good/search`,
    method: 'GET',
    params: { name }
})