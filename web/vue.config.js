'use strict'
const path = require('path')

function resolve(dir) {
  return path.join(__dirname, dir)
}

module.exports = {
    publicPath: './',
    assetsDir: 'static',
    productionSourceMap: false,
    devServer: {
        proxy: {
            '/admin':{
                target:'http://127.0.0.1:8888/admin',//http://127.0.0.1:8193https://emcs.webapi.yyxcloud.com
                changeOrigin:true,
                pathRewrite:{
                    '/admin':''
                }
            }
        }
    }, 
    configureWebpack: {
        // provide the app's title in webpack's name field, so that
        // it can be accessed in index.html to inject the correct title.
        resolve: {
          alias: {
            '@': resolve('src')
          }
        }
      }
}