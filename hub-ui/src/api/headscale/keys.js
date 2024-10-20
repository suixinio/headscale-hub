import request from '@/utils/request'

export function getKeys() {
  return request({
    url: '/api/headscale/pre_auth_key/list',
    method: 'get'
  })
}

export function createKey(data) {
  return request({
    url: '/api/headscale/pre_auth_key/create',
    method: 'post',
    data
  })
}

export function expireKey(data) {
  return request({
    url: '/api/headscale/pre_auth_key/expire',
    method: 'delete',
    data
  })
}

