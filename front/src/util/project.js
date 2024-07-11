import request from './request';

export default {
    getTable, queryDetail
}

export function getTable() {
    return request({
        method :'GET',
        url:'/index/table',
    })
}

export function queryDetail(params) {
    return request({
        method : 'GET',
        url:'/problem/query-detail',
        params: params
    })
}