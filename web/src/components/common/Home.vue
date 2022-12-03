<template>
  <div class="wrapper">
    <v-head></v-head>
    <div class="content-box">
      <v-sidebar></v-sidebar>
      <keep-alive>
        <router-view v-if="$route.meta.keepAlive"></router-view>
      </keep-alive>
      <router-view v-if="!$route.meta.keepAlive"></router-view>
      <!-- <v-tags></v-tags> -->
      <!-- <div class="content" :class="{'content-collapse':collapse}"> -->
      <!-- <div class="content" >
                <transition name="move" mode="out-in">
                    <keep-alive :include="tagsList">
                        <router-view></router-view>
                    </keep-alive>
                </transition>
            </div> -->
      <div class="footer">
        <div class="footer-item">
          <i class="el-icon-help"></i>
          <div style="padding: 0 12px 0 6px">状态</div>
          <div v-if="statusInfo == 0" class="normal"></div>
          <div v-else class="error"></div>
        </div>
        <div class="line"></div>
        <div class="footer-item">
          <i class="el-icon-c-scale-to-original"></i>
          <div style="padding: 0 12px 0 6px">次数</div>
          <span style="color: #419eef">{{ num }}</span>
        </div>
        <div class="line"></div>
        <div class="footer-item">
          <i class="el-icon-basketball" style="color: #999"></i>
          <div style="padding: 0 12px 0 6px">网络</div>
          <div v-if="net == 0" class="normal"></div>
          <div v-else class="error"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import vHead from "./Header.vue";
import vSidebar from "./Sidebar.vue";
import vTags from "./Tags.vue";
import bus from "./bus";
import { urls, method, fetchData, page } from "@/api/index";
import { createSign, deepSortForPostJson } from "@/utils/sign.js";
import axios from "axios";
export default {
  data() {
    return {
      tagsList: [],
      statusInfo: 0,
      num: 0,
      net: 0,
      timer: null,
      // collapse: false
    };
  },
  components: {
    vHead,
    vSidebar,
    vTags,
  },
  created() {
    // bus.$on('collapse', msg => {
    //     this.collapse = msg;
    // })

    // 只有在标签页列表里的页面才使用keep-alive，即关闭标签之后就不保存到内存中了。
    bus.$on("tags", (msg) => {
      let arr = [];
      for (let i = 0, len = msg.length; i < len; i++) {
        msg[i].name && arr.push(msg[i].name);
      }
      this.tagsList = arr;
    });
    this.get_status();
    this.getNum();
  },
  watch: {
    $route(to, from) {
      console.log("路由变化了");
      this.get_status();
      this.getNum();
    },
  },
  mounted() {
    let that = this;
    that.timer = setInterval(() => {
      that.get_status();
      that.getNum();
    }, 10000);
    //注意，初始状态下不会出发，只有当网络状态变化才会出发

    window.addEventListener("offline", function (e) {
      console.log("离线了");
      that.net = 1;
      that.statusInfo = 1;
    });
    window.addEventListener("online", function (e) {
      console.log("在线了");
      that.net = 0;
    });
  },
  beforeDestroy() {
    clearInterval(this.timer);
  },
  methods: {
    get_status() {
      // 待签名的数据
      let testData = {}; //status:1

      // 1、获取按字典排序后的参数
      let sortedData = deepSortForPostJson(testData);

      // 2、对已排序的参数进行签名
      let timestamp = parseInt(new Date().getTime() / 1000);
      let sign = createSign(sortedData, timestamp);

      // 3、设置最终业务参数
      let bizModel = {
        data: {},
        sign: sign,
        timestamp: timestamp,
      };
      let that = this

      fetchData(urls.heartbeat,bizModel,method.POST).then((res)=>{
        if (res.data.code == 0) {
          that.statusInfo = 1;
        } else {
          that.statusInfo = 0;
        }
        if (navigator.onLine) {
          that.net = 0;
        } else {
          that.net = 1;
        }
      }).catch((error)=>{
        that.statusInfo = 1;
        console.log(error);
      })
    },
    getNum() {
      // 待签名的数据
      let testData = {}; //status:1

      // 1、获取按字典排序后的参数
      let sortedData = deepSortForPostJson(testData);

      // 2、对已排序的参数进行签名
      let timestamp = parseInt(new Date().getTime() / 1000);
      let sign = createSign(sortedData, timestamp);

      // 3、设置最终业务参数
      let bizModel = {
        data: {},
        sign: sign,
        timestamp: timestamp,
      };
      fetchData(urls.numToday, bizModel, method.POST).then((res) => {
        if (res.status == 200) {
          this.num = res.data;
        }
      });
    },
  },
};
</script>
<style lang="less" scoped>
.content-box {
  display: flex;
  width: 100%;
  height: 100%;
}
.content {
  flex: 1;
}
.footer {
  position: fixed;
  bottom: 0;
  right: 0;
  width: 100%;
  display: flex;
  justify-content: flex-end;
  align-items: center;
  background: #fff;
  border-top: 1px solid #ebebeb;
  // box-shadow: 0 0 20px -5px #999;
  height: 48px;
  padding: 0 24px;
  font-size: 12px;
  color: #999;

  .footer-item {
    width: 200px;
    justify-content: center;
    display: flex;
    align-items: center;
  }
  .line {
    width: 1px;
    height: 12px;
    background: #ebebeb;
  }
  .normal {
    width: 8px;
    height: 8px;
    background: rgb(81, 248, 184);
    border-radius: 50%;
    margin-left: 12px;
  }
  .error {
    width: 8px;
    height: 8px;
    background: rgb(248, 87, 81);
    border-radius: 50%;
    margin-left: 12px;
  }
}
</style>
