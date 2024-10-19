import request from '@/utils/request'

export function getKeys() {
  return request({
    url: '/api/headscale/pre_auth_key/list',
    method: 'get'
  })
}

