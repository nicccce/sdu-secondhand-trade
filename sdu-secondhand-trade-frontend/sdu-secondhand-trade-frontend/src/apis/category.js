import httpInstance from "@/utils/http"

export const getSubCategoryAPI = (data) => {
  return httpInstance({
    url: `/good/temporary/${data.sort_field}`,
    method: 'POST',
    data
  })
}