import request from './request';

export default {
    getTable, queryDetail, changeSolution, changeDescription,
    queryAllDiffical, queryAllType, deleteProblem,
    insertProblem, updataProblem, insertDifficulty
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

export function queryAllDiffical() {
    return request({
        method : 'GET',
        url:'/index/queryAllDiffical',
    })
}

export function queryAllType () {
    return request({
        method : 'GET',
        url:'/index/queryAllType',
    })
}

export function deleteProblem(data) {
    return request({
        method : 'POST',
        url:'/index/deleteProblem',
        data: data
    })
}

export function insertProblem(data) {
    return request({
        method : 'POST',
        url:'/index/insertProblem',
        data: data
    })
}

export function updataProblem(data) {
    return request({
        method : 'POST',
        url:'/index/updataProblem',
        data: data
    })
}

export function insertDifficulty(data) {
    return request({
        method : 'POST',
        url:'/index/insertDifficulty',
        data: data
    })
}