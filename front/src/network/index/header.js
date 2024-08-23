import request from '../request';

export default {
    insertDifficulty, insertProblem, 
    deleteProblem,
    updataProblem,
    queryAllType, queryAllDiffical, queryAllProblem
}


// 添加困难度
export function insertDifficulty(data) {
    return request({
        method : 'POST',
        url:'/index/insertDifficulty',
        data: data
    })
}
// 添加问题
export function insertProblem(data) {
    return request({
        method : 'POST',
        url:'/index/insertProblem',
        data: data
    })
}

// 删除问题
export function deleteProblem(data) {
    return request({
        method : 'POST',
        url:'/index/deleteProblem',
        data: data
    })
}

// 修改问题
export function updataProblem(data) {
    return request({
        method : 'POST',
        url:'/index/updataProblem',
        data: data
    })
}

// 查所有的题目来源
export function queryAllType () {
    return request({
        method : 'GET',
        url:'/index/queryAllType',
    })
}
// 查所有的题目困难度
export function queryAllDiffical() {
    return request({
        method : 'GET',
        url:'/index/queryAllDiffical',
    })
}
// 查询所有的题目
export function queryAllProblem() {
    return request({
        method :'GET',
        url:'/index/table',
    })
}