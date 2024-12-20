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