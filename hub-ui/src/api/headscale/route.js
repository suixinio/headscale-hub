import request from '@/utils/request'

export function getRouteList() {
  return request({
    url: '/api/headscale/route/list',
    method: 'get'
  })
}

export function enableStatus(data) {
  return request({
    url: '/api/headscale/route/status',
    method: 'post',
    data
  })
}

export function deleteRoute(data) {
  return request({
    url: '/api/headscale/route/delete',
    method: 'delete',
    data
  })
}
