import request from '@/utils/request'

export function getNodeList() {
  return request({
    url: '/api/headscale/node/list',
    method: 'get'
  })
}

export function registerNode(data) {
  return request({
    url: '/api/headscale/node/register',
    method: 'post',
    data
  })
}

export function deleteNode(data) {
  return request({
    url: '/api/headscale/node/delete',
    method: 'delete',
    data
  })
}
