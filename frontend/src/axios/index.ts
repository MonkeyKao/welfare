import axios from "axios";

const request = axios.create({
    baseURL: '/api/',
    timeout: 10000,
    headers: {
        token: localStorage.getItem("token")
    }
});

export default request