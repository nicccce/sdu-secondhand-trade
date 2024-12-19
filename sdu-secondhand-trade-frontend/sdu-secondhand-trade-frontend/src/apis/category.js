import httpInstance from "@/utils/http"

export const getSubCategoryAPI = (data) => {
    return httpInstance({
      url:'/good/temporary',
      method:'POST',
      data
    })
  }