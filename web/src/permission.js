/**
 * 全站权限配置
 *
 */
import router from './router'
import { getToken,removeToken} from '@/utils/auth'
import NProgress from 'nprogress' // progress bar
import 'nprogress/nprogress.css' // progress bar style
NProgress.configure({ showSpinner: false });

router.beforeEach((to, from, next) => {
    //缓冲设置
    NProgress.start()
    //console.log('to.path', to.path+'')
    if(to.path === '/login'){
        removeToken();
        next()
    }else{
        if (getToken()) {
            next()
        }else{
            next({ path: '/login' })
        }
    }
})

router.afterEach(() => {
    NProgress.done();
});
