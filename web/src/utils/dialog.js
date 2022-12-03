
/**
 * @author 封装 element-ui confirm
 * @param text
 * @param title
 * @param config
 * @returns {Promise}
 */
export function confirm(text, title = '温磬提示', config = {}) {
    return new Promise((resolve, reject) => {
        let confirmButtonLoadingClose = () => {
        }
        let _config = merge({
            showCancelButton: true,
            closeOnClickModal: false,
            center: true
        }, config)
        let afterCloseResolve = () => {
        }
        _config.beforeClose = (action, instance, done) => {
            if (lodash.isFunction(config.beforeClose)) {
                config.beforeClose(action, instance, () => {
                })
            }
            if (lodash.isFunction(config.confirmCallBack)) {
                if (action === 'confirm') {
                    instance.confirmButtonLoading = true
                    confirmButtonLoadingClose = () => {
                        instance.confirmButtonLoading = false
                    }
                    config.confirmCallBack({
                        confirmButtonLoadingClose,
                        close: () => new Promise((resolve, reject) => {
                            done()
                            afterCloseResolve = resolve
                        }),
                        action
                    })
                } else {
                    done()
                }
            }
            if (!config.confirmButtonLoading) {
                done()
            }
        }
        delete _config.confirmButtonLoading
        MessageBox.confirm(text, title, _config).then(_ => {
            afterCloseResolve()
            resolve()
        }).catch(err => {
            afterCloseResolve()
            reject(err)
        })
    })
}