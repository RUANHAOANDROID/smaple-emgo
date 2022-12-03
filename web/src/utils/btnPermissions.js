import Vue from 'vue'
/**权限指令**/
const has = Vue.directive('has', {
    inserted: function (el, binding, vnode) {
        let path = vnode.context.$route.path;

        if (!Vue.prototype.$_has(binding.value, path)) {
            el.parentNode.removeChild(el);
        }
    }
});

//权限检查方法
Vue.prototype.$_has = function (value, path) {

    let isExist = false;
    let btnPermissionsStr = localStorage.getItem("btnPermissions");
    if (btnPermissionsStr == undefined || btnPermissionsStr == null) {
        return false;
    }

    let btnPermissions = JSON.parse(btnPermissionsStr);

    //let btnPermission = btnPermissions[path.substring(1, path.length) + "/" + value];
    let hasBtns = btnPermissions.filter((item) => {
        return item.resourceCode == path.substring(1, path.length) + "/" + value;
    });

    if (hasBtns && hasBtns.length > 0) {
        isExist = true;
    }

    return isExist;
};

export { has }