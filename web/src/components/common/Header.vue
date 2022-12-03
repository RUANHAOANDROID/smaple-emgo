<template>
  <div class="header">
    <div :class="[!collapse ? 'logo-box1' : 'logo-box2']">
      <div class="logo1">
        <img src="../../assets/logo1.png" />
        <div
          v-if="!collapse"
          style="color: #fff; font-size: 14px; padding-left: 10px"
        >
          中继服务
        </div>
      </div>
    </div>
    <!-- 折叠按钮 -->
    <div class="collapse-btn" @click="collapseChage">
      <i class="iconfont icon-daohangshouqi" v-if="!collapse"></i>
      <i class="iconfont icon-daohangzhankai" v-else></i>
    </div>
    <div class="header-right">
      <div style="margin-right: 24px;display:flex;align-items:center;">
        <img class="online-img" :src="img1" style="width
        14px;height:14px;margin-right: 10px" />
        <div class="online-text">{{ FormateTime(nowtime) }}</div>
      </div>

      <div class="logout-box">
        <i class="iconfont icon-tuichu" style="line-height: 3"></i>
        <div class="text" @click="loginout">退出登录</div>
      </div>
    </div>
  </div>
</template>
<script>
import bus from "../common/bus";
import { urls, method, fetchData, page } from "../../api/index";
import { title } from "@/setting";
import ElFormItem from "../../../node_modules/element-ui/packages/form/src/form-item.vue";
export default {
  components: { ElFormItem },
  data() {
    return {
      sysTitle: title,
      collapse: false,
      timer: null,
      nowtime: new Date(),
      img1:require('../../assets/img/b.png')
    };
  },
  computed: {
    username() {
      let user = JSON.parse(localStorage.getItem("userInfo"));
      let username = user.userName;
      return username ? username : this.name;
    },
  },
  mounted() {
    // 显示时间1s1s的显示
    let that = this;
    this.timer = setInterval(function () {
      that.nowtime =  new Date().toLocaleString();
      
    });
     if (document.body.clientWidth < 900) {
      this.collapseChage();
    }
  },
  // 销毁定时器
  beforeDestroy() {
    if (this.timer) {
      clearInterval(this.timer);
    }
  },
  methods: {
    loginout() {
      localStorage.removeItem("ms_username");
      this.$router.push("/login");
    },
    // 侧边栏折叠
    collapseChage() {
      this.collapse;
      this.collapse = !this.collapse;
      bus.$emit("collapse", this.collapse);
    },
    FormateTime() {
      var date = new Date();
      var year = this.formateTime(date.getFullYear());
      var month = this.formateTime(date.getMonth() + 1);
      var day = this.formateTime(date.getDate());
      var hour = this.formateTime(date.getHours());
      var minute = this.formateTime(date.getMinutes());
      var second = this.formateTime(date.getSeconds());
      var weekday = date.getDay();
      var weeks = new Array(
        "星期日",
        "星期一",
        "星期二",
        "星期三",
        "星期四",
        "星期五",
        "星期六"
      );
      var week = weeks[weekday];
      return (
        year +
        "-" +
        month +
        "-" +
        day +
        " " +
        hour +
        ":" +
        minute +
        ":" +
        second
      );
    },
    formateTime(n) {
      if (n < 10) {
        return "0" + n;
      } else {
        return n;
      }
    },
  },

};
</script>
<style scoped lang="less">
.header {
  position: relative;
  //   box-sizing: border-box;
  width: 100vw;
  height: 48px;
  line-height: 48px;
  font-size: 22px;
  //   color: #fff;
  background: #fff;
  display: flex;
  /* float: right; */
  box-shadow: 0 0 20px -5px #cfcfcf;
}
.logo-box1 {
  height: 48px;
  width: 208px;
  background: #001529;
}
.logo-box2 {
  width: 64px;
  background: #001529;
}
.logo1 {
  height: 48px;
  display: flex;
  justify-content: center;
  align-items: center;
}
.collapse-btn {
  float: left;
  padding-left: 12px;
  cursor: pointer;
  line-height: 48px;
  height: 48px;
  color: #333;
  background: #fff;
}

.header-right {
  position: absolute;
  right: 14px;
  height: 48px;
  line-height: 48px;
  display: flex;
  width: 60vw;
  justify-content: flex-end;
  align-items: center;
  font-size: 12px;
}
.header-user-con {
  display: flex;
  height: 20px;
  align-items: center;
}
.btn-fullscreen {
  transform: rotate(45deg);
  margin-right: 5px;
  font-size: 24px;
}
.btn-bell,
.btn-fullscreen {
  position: relative;
  width: 30px;
  height: 30px;
  text-align: center;
  border-radius: 15px;
  cursor: pointer;
  line-height: 2px;
}

.btn-bell-badge {
  position: absolute;
  right: 0;
  top: 0;
  width: 8px;
  height: 8px;
  border-radius: 4px;
  background: #f56c6c;
  color: #fff;
}
.btn-bell .el-icon-bell {
  color: #333;
  font-size: 14px;
}
.user-name {
  margin-left: 10px;
}
.user-avator {
  margin-left: 20px;
}
.user-avator img {
  display: block;
  width: 20px;
  height: 20px;
  border-radius: 50%;
}
.el-dropdown-link {
  //   color: #fff;
  cursor: pointer;
  padding-right: 14px;
}
.el-dropdown-menu__item {
  text-align: center;
}
.logout-box {
  cursor: pointer;
  display: flex;
  height: 48px;
  opacity: 0.85;
  font-size: 12px;
  font-weight: 400;
  text-align: left;
  color: rgba(0, 0, 0, 0.85);
  line-height: 48px;
  .text {
    margin-left: 6px;
  }
}
</style>
