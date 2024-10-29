import axios from "axios";

const request = axios.create({
    baseURL: '/api/',
    timeout: 10000,
    headers: {
        token: localStorage.getItem("token")
    }
});

// 添加请求拦截器
request.interceptors.request.use(
    config => {
      // 获取本地存储中的 token
      const token = localStorage.getItem('token');
      if (token) {
        // 如果有 token，则添加到请求头中
        config.headers['token'] = token;
      }
      return config;
    },
    error => {
      return Promise.reject(error);
    }
  );

export default request