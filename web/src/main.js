import Vue from 'vue'
import App from './App.vue'
import router from './router'
import axios from 'axios';
import ElementUI from 'element-ui';
import VueI18n from 'vue-i18n';
import has from './utils/btnPermissions';
import './permission'; // 权限
import { resize } from './api/index';
import { messages } from './components/common/i18n';

    
import 'element-ui/lib/theme-chalk/index.css'; // 默认主题
// import '../static/css/theme-green/index.css';       // 浅绿色主题
import './assets/css/icon.css';
import './assets/iconfont/iconfont.css';

import './components/common/directives';//iconfont以静态资源方式引入
import "babel-polyfill";



Vue.config.productionTip = false
Vue.use(VueI18n);
Vue.use(ElementUI, {
    size: 'small'
});
Vue.prototype.$axios = axios;
Vue.prototype.$resize = resize;

const i18n = new VueI18n({
    locale: 'zh',
    messages
})

new Vue({
    router,
    i18n,
    render: h => h(App)
}).$mount('#app')