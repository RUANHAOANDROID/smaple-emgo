import request from '@/utils/request';
import axios from 'axios';
import { getToken, getMchId } from '@/utils/auth';
import { Loading } from 'element-ui';

export const method = {
    POST: "post",
    GET: "get"
}
export const urls = {
    login: '/admin/login',//登陆URL
    password:'/admin/updatePassword',//修改密码
    deviceList:'/admin/deviceList',//设备列表
    //saveConfig:'/admin/saveConfig', //保存配置
    addDevice:'/admin/addDevice', //新增设备
    updateDevice:'/admin/updateDevice', //编辑设备
    deleteDevice:'/admin/deleteDevice', //删除设备
    getConfig:'/admin/getConfig',//获取配置
    saveConfig:'/admin/saveConfig',//报错配置
    openGateTest :'/admin/openGateTest',//开闸
    checkTicket:'/admin/checkTicketTest',//测试验票
    numUpload:'/admin/numUploadTest',//人数上报
    operateInfo:'/admin/getRecentOperateLog',//设备信息 
    heartbeat:'/admin/heartbeat',//获取状态
    gatSetting:'/admin/getCurrentConfig',//获取参数设置保存后的数据
    numToday:'/admin/getNumToday'//过闸人数
}

export function fetchData(urlStr, param, method) {
    var option = {
        url: urlStr,
        method: method ? method : "post",
        data: param
    }
    if (method == 'get') {
        delete option.data;
        option.params = param; //{abc,123}
    }
    return request(option);
}

/**
 * 公共分页配置
 * @type {{pageSizes: [number,number,number,number]}}
 */
export const page = {
    pageSizes: [10, 20, 50, 100],
    pageSize: 10
}

/**
 * base64串转化为文件
 * @param fileName
 * @param fileBase64Str
 */
export function downloadBlob(fileName, fileBase64Str) {
    var sliceSize = 512;
    var byteCharacters = atob(fileBase64Str);
    var byteArrays = [];

    for (var offset = 0; offset < byteCharacters.length; offset += sliceSize) {
        var slice = byteCharacters.slice(offset, offset + sliceSize);

        var byteNumbers = new Array(slice.length);
        for (var i = 0; i < slice.length; i++) {
            byteNumbers[i] = slice.charCodeAt(i);
        }

        var byteArray = new Uint8Array(byteNumbers);
        byteArrays.push(byteArray);
    }

    var blob = new Blob(byteArrays, { type: "application/vnd.ms-excel" });

    var url = URL.createObjectURL(blob);
    //如果是直接下载,利用a标签来实现下载
    var docEle = document;
    var link = docEle.createElement("a");
    link.innerHTML = fileName;
    link.download = fileName;
    link.href = url;
    docEle.getElementsByTagName("body")[0].appendChild(link);
    link.click();
    //释放内存
    URL.revokeObjectURL(link.href);
}

/**
 * 公共导出方法
 * @param url
 * @param options
 * @returns {Promise}
 */
export function exportExcel(url, options = {}) {
    let loading = Loading.service({ target: '#ark-per-table', text: '正在导出数据,请稍后...', background: 'rgba(0, 0, 0, 0.1)' });

    return new Promise((resolve, reject) => {
        if (getToken()) {
            axios.defaults.headers['authorization'] = getToken()
        }
        if (getMchId()) {
            axios.defaults.headers['mchId'] = getMchId()
        }
        axios.defaults.headers['content-type'] = 'application/json;charset=UTF-8';
        axios({
            method: 'post',
            url: url, // 请求地址
            data: options // 参数
            // responseType: 'blob' // 表明返回服务器返回的数据类型
        }).then(
            response => {
                loading.close();
                resolve(response.data);

                if (response.data.code == 0) {
                    let result = response.data.data;
                    downloadBlob(result.fileName, result.fileBase64Str);
                }
            },
            err => {
                loading.close();
                reject(err);
            }
        )
    })
}

// 防抖
export const debounce = function (func, wait ) {
    let timer = 0;
    
    return function (...args) {
        let self = this;
        if (timer) {
            // 如果在指定时间间隔内又再次触发，则继续延时
            clearTimeout(timer);
        }
        // 重点：使用$nextTick函数
        // self.$nextTick(() => {
            timer = setTimeout(() => {
                func.apply(self, args);
            }, wait);
        // });
    };
};

/**
 * 动态设置table自适应高度
 * @returns {*}
 */export const resize = function (args) {
    return debounce.call(this, function () {
        //分页距离页面窗口底部高度
        let paginationBottom = 10;

        //需要总合计
        if (args) {
            //合计div高度
            let totalHeight = 30;

            this.tableHeight = window.innerHeight - document.getElementsByClassName("table_content")[0].offsetTop - document.getElementsByClassName("module_container")[0].offsetTop - document.getElementsByClassName("content-box")[0].offsetTop - paginationBottom;
            document.getElementsByClassName("el-table__body-wrapper")[0].style.height = this.tableHeight - document.getElementsByClassName("el-table__header-wrapper")[0].offsetHeight - totalHeight + "px";
            // document.getElementsByClassName("el-table__footer-wrapper")[0].style.height = this.tableHeight - document.getElementsByClassName("el-table__header-wrapper")[0].offsetHeight - document.getElementsByClassName("el-table__body-wrapper")[0].offsetHeight + "px";
            document.getElementsByClassName("el-table__footer")[0].style.height = this.tableHeight - document.getElementsByClassName("el-table__header-wrapper")[0].offsetHeight - document.getElementsByClassName("el-table__body-wrapper")[0].offsetHeight + "px";
        }
        //没有合计
        else {
            this.tableHeight = window.innerHeight - document.getElementsByClassName("table_content")[0].offsetTop - document.getElementsByClassName("module_container")[0].offsetTop - document.getElementsByClassName("content-box")[0].offsetTop - paginationBottom;
            document.getElementsByClassName("el-table__body-wrapper")[0].style.height = this.tableHeight - document.getElementsByClassName("el-table__header-wrapper")[0].offsetHeight + "px";
        }
    });
};

/**
 * 总合计
 * @param {*} columns table所有列
 * @param {*} sumColMap 后台返回的对应列的总合计map
 */
export function summaryMethod(columns, sumColMap) {

    const sums = [];

    sums[0] = '合计';

    for (let key in sumColMap) {

        for (let index = 0; index < columns.length; index++) {

            let column = columns[index];

            if (column.property == key) {

                sums[index] = sumColMap[key];
            }

        }
    }

    return sums;
}