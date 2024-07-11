import axios from 'axios'
const request = axios.create({
    baseURL: "http://localhost:8000",
    withCredentials: true // 表示请求可以携带cookie
})
export default request