import httpInstance from "@/utils/http";

export const initOrderAPI = (good_id) => httpInstance({
    url: '/order/init',
    method: 'POST',
    data: {good_id}
});

export const getOrderAPI = (order_id) => httpInstance({
    url: `/order/${order_id}`,
    method: 'GET'
});

export const submitOrderAPI = (data) => httpInstance({
    url: `/order/submit`,
    method: 'POST',
    data
})

export const cancelOrderAPI = (order_id) => httpInstance({
    url: `/order/${order_id}`,
    method: 'DELETE'
})

export const getBuyerOrderListAPI = (data) => httpInstance({
    url: '/order/buyer',
    method: 'POST',
    data
})

export const getSellerOrderListAPI = (data) => httpInstance({
    url: '/order/seller',
    method: 'POST',
    data
})

export const sendAfterSaleAPI = (data) => httpInstance({
    url: '/problem/after_sale',
    method: 'POST',
    data
})

export const getAfterSaleAPI = (data) => httpInstance({
    url: '/problem/get',
    method: 'POST',
    data
})

export const getAllAfterSaleAPI = (data) => httpInstance({
    url: '/problem/all',
    method: 'POST',
    data
})

export const updateAfterSaleAPI = (data) => httpInstance({
    url: '/problem/update',
    method: 'POST',
    data
})
