import httpInstance from "@/utils/http";

export const loginAPI = ({ student_id, password }) => {
    return httpInstance({
        url: '/user/login',
        method: 'POST',
        data: {
            student_id,
            password
        },
    })
}

export const refreshAPI = ()=>httpInstance({
    url: '/user/me',
    method: 'GET',
})

export const updateUserAPI = (data)=>httpInstance({
    url: '/user/me',
    method: 'POST',
    data
})

export const addAddressAPI = (data) => httpInstance({
    url: '/user/address',
    method: 'POST',
    data
})

export const deleteAddressAPI = (id) => httpInstance({
    url: `/user/address/${id}`,
    method: 'DELETE'
})

export const updateAddressAPI = (data) => httpInstance({
    url: `/user/address/${data.id}`,
    method: 'POST',
    data
})

export const getUserByIdAPI = (id) => httpInstance({
    url: `/user/${id}`,
    method: 'GET',
})

export const getAllUsersAPI = () =>httpInstance({
    url: `/user`,
    method: 'GET'
})

export const updatePasswordAPI = (data)=>httpInstance({
    url:'/user/password',
    method: 'POST',
    data
})