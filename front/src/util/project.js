import request from './request';

export default {
    getTable, queryDetail, changeSolution, changeDescription
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
        url:'/problem/queryDetail',
        params: params
    })
}

export function changeSolution(data) {
    return request({
        method : 'POST',
        url:'/problem/changeSolution',
        data: data
    })
}

export function changeDescription(data) {
    return request({
        method : 'POST',
        url:'/problem/changeDescription',
        data: data
    })
}