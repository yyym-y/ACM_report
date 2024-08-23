import request from '../request';

export default {
    changeSolution, changeDescription,
    queryMdDetail,
    problemCrawler
}

// 修改题解
export function changeSolution(data) {
    return request({
        method : 'POST',
        url:'/problem/changeSolution',
        data: data
    })
}
// 修改题目描述
export function changeDescription(data) {
    return request({
        method : 'POST',
        url:'/problem/changeDescription',
        data: data
    })
}

// 查询问题的题目描述和题解
export function queryMdDetail(params) {
    return request({
        method : 'GET',
        url:'/problem/queryDetail',
        params: params
    })
}

// 题目爬虫
export function problemCrawler(params) {
    return request({
        method : 'GET',
        url:'/problem/problemCrawler',
        params: params
    })
}