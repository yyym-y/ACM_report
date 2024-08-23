import request from './request';

export default {
    changeSolution, changeDescription,
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