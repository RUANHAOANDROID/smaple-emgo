<template>
  <div :class="[collapse ? 'collapse-side' : 'sidebar']">
    <el-menu
      class="sidebar-el-menu"
      :default-active="onRoutes"
      :collapse="collapse"
      background-color="#001529"
      text-color="#fff"
      unique-opened
      router
    >
      <template v-for="item in menuList">
        <template v-if="item.subs">
          <el-submenu :index="item.index" :key="item.index">
            <template slot="title">
              <i :class="item.icon"></i
              ><span slot="title">{{ item.title }}</span>
            </template>
            <template v-for="subItem in item.subs">
              <el-submenu
                v-if="subItem.subs"
                :index="subItem.index"
                :key="subItem.index"
              >
                <template slot="title">{{ subItem.title }}</template>
                <el-menu-item
                  v-for="(threeItem, i) in subItem.subs"
                  :key="i"
                  :index="threeItem.index"
                >
                  {{ threeItem.title }}
                </el-menu-item>
              </el-submenu>
              <el-menu-item v-else :index="subItem.index" :key="subItem.index">
                {{ subItem.title }}
              </el-menu-item>
            </template>
          </el-submenu>
        </template>
        <template v-else>
          <el-menu-item :index="item.index" :key="item.index">
            <i :class="item.icon"></i><span slot="title">{{ item.title }}</span>
          </el-menu-item>
        </template>
      </template>
    </el-menu>
  </div>
</template>

<script>
import bus from "../common/bus";
import { getMchId } from "@/utils/auth";
import { urls, method, fetchData } from "../../api/index";
export default {
  data() {
    return {
      collapse: false,
      items: [],
      logo_img: "",
      menuList: [
        {
          icon: "el-icon-lx-home",
          id: "100000",
          index: "dashboard",
          pId: "-1",
          title: "运行概要",
          uri: "dashboard",
        },
        {
          icon: "el-icon-s-management",
          id: "100001",
          index: "equipment",
          pId: "-1",
          // subs: [
          //   {
          //     index: "sys_user",
          //     pId: "100001",
          //     id: "1308728023159144955",
          //     title: "平台用户",
          //     uri: "sys_user",
          //   },
          // ],
          title: "设备调试",
          uri: "100001",
        },
        {
          icon: "el-icon-s-management",
          id: "200000",
          index: "setting",
          pId: "-1",
          title: "参数设置",
          uri: "200000",
        },
        {
          icon: "el-icon-setting",
          id: "300000",
          index: "equipment_bind",
          pId: "-1",
          title: "设备绑定",
          uri: "300000",
        },
      ],
    };
  },
  computed: {
    onRoutes() {
      return this.$route.path.replace("/", "");
    },
  },
  created() {
    // 通过 Event Bus 进行组件间通信，来折叠侧边栏
    bus.$on("collapse", (msg) => {
      this.collapse = msg;
    });
    this.getLogoImg();
  },
  mounted() {
    //获取用户菜单
    // fetchData("/api/menu/getUserMenuList", {}, method.GET).then((res) => {
    //   this.items = res;
    //   localStorage.setItem("menu", JSON.stringify(res));
    // });
    //获取按钮权限
    // fetchData("/api/resource/btnPermissions", {}, method.GET).then((res) => {
    //   localStorage.setItem("btnPermissions", JSON.stringify(res));
    // });
  },
  methods: {
    getLogoImg() {
      let mchId = getMchId();
      if (mchId == "null") {
        mchId = -1;
      }
      let optionCode = "Logo_Path";
      this.logo_img = require("../../assets/img/SCM.png");
    },
  },
};
</script>

<style scoped lang="less">
.sidebar {
  width: 208px;
  height: 100vh;
  /* overflow-y: scroll; */
  background-color: #001529;
  z-index: 9999;
}
.collapse-side {
  width: 64px;
  height: 100vh;
  background-color: #001529;
  z-index: 9999;
}
.sidebar::-webkit-scrollbar {
  /* width: 0; */
}
.sidebar-el-menu:not(.el-menu--collapse) {
  /* width: 180px;
        overflow: auto; */
}
.sidebar > ul {
  height: 100%;
  /* width: 207.5px; */
}
.sidebar /deep/ .el-menu {
  border-right: none;
  width: 208px;
  color: #fff;
}
.collapse-side /deep/ .el-menu {
  border-right: none;
}
.sidebar /deep/ .el-menu-item.is-active {
  background: #1890ff !important;
  color: #fff;
}
//设置选中el-menu-item时的样式
.el-menu-item:hover {
  color: #fff !important;
  background: #001529 !important;
}
</style>
