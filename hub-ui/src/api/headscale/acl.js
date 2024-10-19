import request from '@/utils/request'

export function getACL() {
  return request({
    url: '/api/headscale/policy/get',
    method: 'get'
  })
}

export function setACL(data) {
  return request({
    url: '/api/headscale/policy/set',
    method: 'post',
    data
  })
}
