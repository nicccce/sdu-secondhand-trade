import httpInstance from "@/utils/http";

export const getGenderAPI = () => httpInstance({
    url: '/static/gender'
});

export const getCampusAPI = () => httpInstance({
    url: '/static/campus'
});