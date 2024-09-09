import axios from 'axios'
const request = axios.create({
    baseURL: "http://yyym.nat300.top/",
    withCredentials: true // 表示请求可以携带cookie
})
export default request