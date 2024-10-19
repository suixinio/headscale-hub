import request from '@/utils/request'

export function getACL() {
  return request({
    url: '/api/headscale/policy/acl',
    method: 'get'
  })
}

export function setACL(data) {
  return request({
    url: '/api/headscale/policy/acl',
    method: 'post',
    data
  })
}
