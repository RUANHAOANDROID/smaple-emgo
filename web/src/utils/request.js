import axios from 'axios';
import { getToken, getMchId } from '@/utils/auth'
import { MessageBox, Message } from 'element-ui'

const service = axios.create({
    timeout: 10000
})

service.interceptors.request.use(config => {
    // if (getToken()) {
    //     config.headers['authorization'] = getToken()
    // }
    // if (getMchId()) {
    //     config.headers['mchId'] = getMchId()
    // }

    return config;
}, error => {
    console.log(error);
    return Promise.reject();
})

service.interceptors.response.use(response => {
    if (response.data.code === 1) {
        return response.data;
    } else {
        const message = response.data.message || '请求失败！';

        MessageBox.alert(message, '提示', {
            confirmButtonText: '确定',
            type: 'error'
        });
        return Promise.reject();
    }
}, error => {

    Message({
        message: error.message,
        type: 'error',
        duration: 5 * 1000
    })
    return Promise.reject(error);
})

export default service;