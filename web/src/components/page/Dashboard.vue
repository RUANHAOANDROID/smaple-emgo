<template>
  <div class="wrapper">
    <div class="wrapper-left">
      <div class="title">
        <div class="rectangle"></div>
        信息提示
      </div>
      <div class="left-content">
        <div v-for="(i, index) in infoList" :key="index" class="content-item">
          <div class="item-date">{{ i.addTime }}</div>
          <div class="item-remark">{{ i.content }}</div>
        </div>
      </div>
    </div>
    <div class="wrapper-right">
      <div class="title">
        <div class="rectangle"></div>
        在线设备
      </div>
      <div class="table_content">
        <el-table
          :data="tableData"
          :height="tableHeight"
          tooltip-effect="dark"
          :header-cell-style="{ background: 'rgba(0,0,0,0.02)' }"
          :row-style="{ height: '54px' }"
          :header-row-style="{ height: '54px' }"
        >
          <el-table-column
            prop="id"
            label="中继设备号"
            width="120"
          ></el-table-column>
          <el-table-column label="设备编号" width="120" prop="deviceNo">
          </el-table-column>

          <el-table-column
            prop="deviceIp"
            label="控制板IP"
            show-overflow-tooltip
          ></el-table-column>
          <el-table-column
            prop="deviceVersion"
            label="版本"
            show-overflow-tooltip
          ></el-table-column>
          <el-table-column prop="status" label="状态">
            <template slot-scope="scope">
              <div>
                {{ scope.row.status | isActiveFilter }}
              </div>
            </template></el-table-column
          >
        </el-table>
      </div>
      <div class="pagination">
        <el-pagination
          background
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
          :current-page="pageNum"
          :page-sizes="pageSizes"
          :page-size="pageSize"
          layout="total, sizes, prev, pager, next, jumper"
          :total="currentTotal"
        >
        </el-pagination>
      </div>
    </div>
  </div>
</template>
<script>
import { urls, method, fetchData, page } from "@/api/index";
import { createSign, deepSortForPostJson } from "@/utils/sign.js";
export default {
  name: "dashboard", //运行概要
  data() {
    return {
      infoList: [],
      tableHeight: window.innerHeight - 300, //默认table高度
      tableData: [],
      pageSizes: [10, 20, 50],
      pageSize: 10,
      pageNum: 1,
      currentTotal: 0,
      timer: null,
    };
  },
  created() {
    this.getData();
  },
  beforeDestroy() {
    clearInterval(this.timer);
  },
  beforeRouteEnter(to, form, next) {
    to.meta.keepAlive = true;
    next();
  },
  activated() {
    console.log("dashboard激活activated钩子函数");
    this.getData();
    this.timer = setInterval(() => {
      this.getInfo();
    }, 5000);
  },
  deactivated() {
    console.log("dashboard激活deactivated钩子函数");
    clearInterval(this.timer);
    if (this.infoList.length > 500) {
      this.infoList = [];
    }
  },
  mounted() {
    this.getInfo();

    console.log("mountd");
  },
  filters: {
    isActiveFilter: function (val) {
      if (val == "01") {
        return "在线";
      } else if (val == "02") {
        return "离线";
      } else if (val == "03") {
        return "停用";
      } else if (val == "04") {
        return "报废";
      } else if (val == "05") {
        return "异常";
      } else if (val == "06") {
        return "报修中";
      } else {
        return "";
      }
    },
  },
  methods: {
    //获取信息列表
    getInfo() {
      // 待签名的数据
      let testData = {}; //status:1
      let sortedData = deepSortForPostJson(testData);

      let timestamp = parseInt(new Date().getTime() / 1000);
      let sign = createSign(sortedData, timestamp);

      let bizModel = {
        data: {},
        sign: sign,
        timestamp: timestamp,
      };
      fetchData(urls.operateInfo, bizModel, method.POST).then((res) => {
        // res.data = [
        //   {
        //     addTime: "101010",
        //     content:
        //       "aaaa按计划发就哈哈警方哈吉哈吉奥干哈给发货啊发就好哈就会发",
        //   },
        // ];
        if (res.code == 1) {
          this.infoList = res.data.concat(this.infoList);
        }
        console.log(this.infoList.length);
      });
    },
    handleSizeChange(val) {
      this.pageSize = val;
      this.getData();
    },
    // 分页导航
    handleCurrentChange(val) {
      this.pageNum = val;
      this.getData();
    },
    getData() {
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
      console.log("最终业务请求参数：", JSON.stringify(bizModel));
      fetchData(urls.deviceList, bizModel, method.POST).then((res) => {
        if (res.data) {
          this.tableData = res.data;
          this.currentTotal = parseInt(res.data.length);
        }
      });
    },
  },
};
</script>
<style lang="less" scoped>
.wrapper {
  display: flex;
  width: 100%;
  height: 100%;
  background: #f8f8f8;
  .title {
    font-size: 16px;
    font-weight: 500;
    padding: 20px 0 0 20px;
    display: flex;
    align-items: center;
  }
  .rectangle {
    background: #419eef;
    width: 5px;
    height: 14px;
    margin-right: 8px;
    margin-top: 2px;
    border-radius: 2px;
  }
  .wrapper-left {
    width: 60%;
    background: #fff;
    // box-shadow: 0 0 20px -5px #ebebeb;
    // height: 98vh;
    margin: 24px;
    height: 82%;
    border-radius: 8px;
    height: calc(100vh - 144px);
    // border: 1px solid #ebebeb;
    .left-content {
      padding: 20px;
      height: calc(100vh - 240px);
      overflow-y: auto;
      overflow-x: hidden;
      .content-item {
        // height: 32px;
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding: 6px 0;
      }
      .item-date {
        color: #999;
        font-size: 14px;
        width: 40%;
        text-align: left;
        // overflow: hidden;
        // white-space: nowrap;
        // text-overflow: ellipsis;
      }
      .item-remark {
        color: #666;
        font-size: 14px;
        width: 60%;
        text-align: left;
        // overflow: hidden;
        // white-space: nowrap;
        // text-overflow: ellipsis;
      }
    }
  }
  .wrapper-right {
    flex: 1;
    background: #fff;
    // box-shadow: 0 0 20px -5px #ebebeb;
    height: 82%;
    margin: 24px 24px 24px 0;
    border-radius: 8px;
    // border: 1px solid #ebebeb;
    height: calc(100vh - 144px);
  }
  .table_content {
    padding: 24px;
  }
  //   .table_content /deep/ .el-table {
  //   width: 100%;
  // }
  .pagination {
    margin: 0 24px 0 0;
  }
}
</style>