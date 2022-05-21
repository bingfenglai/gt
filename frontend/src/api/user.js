import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/oauth2/token',
    method: 'post',
    params:data
  })
}

export function getInfo() {
  return request({
    url: '/v1/user/info',
    method: 'get'

  })
}

export function logout() {
  return request({
    url: '/oauth2/logout',
    method: 'post'
  })
}
