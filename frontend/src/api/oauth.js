import request from "@/utils/request";

export function Token(data){
  return request(
    {
      url: '/oauth2/token',
      data,
      method: 'post'

    }
  )
}
