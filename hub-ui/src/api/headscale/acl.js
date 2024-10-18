import request from '@/utils/request'

export function getACL() {
  return request({
    url: '/api/headscale/policy/acl',
    method: 'get'
  })
}
