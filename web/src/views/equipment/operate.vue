<template>
  <el-dialog
    class="el-dialog-detail"
    :visible.sync="show"
    @close="cancel"
    :before-close="cancel"
    title="操作&提示"
  >
    <!-- :width="dialogWidth" -->
    <div class="row-line"></div>
    <!-- <div
      style="
        width: 90%;
        display: flex;
        justify-content: flex-start;
        padding: 0 0 24px 0;
      "
    > -->
    <el-tabs v-model="activeName" @tab-click="handleClick">
      <el-tab-pane label="开闸" name="first">
        <div style="display: flex">
          <el-form ref="form1" :model="form" label-width="120px" class="form1">
            <el-form-item label="中继设备号" prop="id">
              <el-input
                style="width: 208px"
                v-model="form1.id"
                disabled
              ></el-input>
            </el-form-item>
          </el-form>
          <el-button
            style="height: 32px; margin-left: 24px"
            type="primary"
            size="small"
            @click="open"
            >开闸</el-button
          >
        </div>
      </el-tab-pane>
      <!-- <el-tab-pane label="关闸" name="second">
        <div style="display: flex">
          <el-form ref="form1" :model="form" label-width="120px" class="form1">
            <el-form-item label="中继设备号" prop="id">
              <el-input
                style="width: 208px"
                v-model="form1.id"
                disabled
              ></el-input>
            </el-form-item>
          </el-form>
          <el-button
            style="height: 32px; margin-left: 24px"
            type="primary"
            size="small"
            @click="close"
            >关闸</el-button
          >
        </div>
      </el-tab-pane> -->
      <el-tab-pane label="检票测试" name="third">
        <el-form ref="form" :model="form" label-width="120px" class="form1">
          <el-form-item label="中继设备号" prop="id">
            <el-input
              style="width: 208px"
              v-model="form.id"
              disabled
            ></el-input>
          </el-form-item>
          <el-form-item label="验票类型" prop="protocolNo">
            <el-select v-model="form.protocolNo" placeholder="请选择">
              <el-option
                v-for="item in Option"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              >
              </el-option>
            </el-select>
          </el-form-item>

          <el-form-item label="票标识" prop="cardNo">
            <el-input
              style="width: 208px"
              v-model="form.cardNo"
              clearable
              placeholder=""
            ></el-input>
          </el-form-item>
        </el-form>
        <div style="width: 100%; text-align: center">
          <el-button type="primary" size="small" @click="test"
            >检票测试</el-button
          >
        </div>
      </el-tab-pane>
      <el-tab-pane label="人数上报" name="fourth">
        <div style="display: flex">
          <el-form ref="form1" :model="form" label-width="120px" class="form1">
            <el-form-item label="中继设备号" prop="id">
              <el-input
                style="width: 208px"
                v-model="form1.id"
                disabled
              ></el-input>
            </el-form-item>
          </el-form>
          <el-button
            type="primary"
            size="small"
            @click="report"
            style="height: 32px; margin-left: 24px"
            >人数上报</el-button
          >
        </div>
      </el-tab-pane>
    </el-tabs>
    <!-- </div> -->
    <div>
      <div class="row-line"></div>
      <div class="right">
        <div class="title">提示信息</div>
        <div class="info-content">
          <div class="flex-between" v-for="(i, index) in infoList" :key="index">
            <div class="item-date">{{ i.addTime }}</div>
            <div class="item-remark">{{ i.content }}</div>
          </div>
        </div>
      </div>
    </div>
  </el-dialog>
</template>
 
<script>
import bus from "@/components/common/bus";
import { urls, method, fetchData, page } from "@/api/index";
import { createSign, deepSortForPostJson } from "@/utils/sign.js";
import { debounce } from "../../api/index";
export default {
  name: "vision-add", //报修详情
  //   components: { equipmentInfo, StepInfo, Handle },
  data() {
    return {
      activeName: "first",
      // visible: this.show,
      dialogWidth: "880",
      form1: {
        id: "",
      },
      rules: {
        protocolNo: [
          { required: true, message: "请选择类型", trigger: "change" },
        ],
        cardNo: [{ required: true, message: "请", trigger: "change" }],
      },
      form: {
        protocolNo: "",
        id: "",
        cardNo: "",
      },
      Option: [
        {
          value: "002",
          label: "二维码报文",
        },
        {
          value: "003",
          label: "身份证",
        },
        {
          value: "004",
          label: "IC卡报文（旅游卡）",
        },
        {
          value: "006",
          label: "人脸报文",
        },
      ],
      infoList: [],
      timer: null,
    };
  },
  props: {
    show: {
      type: Boolean,
      default: false,
    },
    deviceId: {
      type: String,
      default: "",
    },
  },
  watch: {
    show(val) {
      // this.visible = this.show;
      if (!val) {
        this.$emit("closeDialog");
        this.init();
      } else {
        this.init();
        this.timer = setInterval(() => {
          this.getInfo();
        }, 5000);
      }
    },
  },
  beforeDestroy() {
    clearInterval(this.timer);
  },
  // activated() {
  //   console.log("dashboard激活activated钩子函数");
  //   this.timer = setInterval(() => {
  //     this.getInfo();
  //   }, 5000);
  // },
  // deactivated() {
  //   console.log("dashboard激活deactivated钩子函数");
  //   clearInterval(this.timer);
  //   if (this.infoList.length > 100) {
  //     this.infoList = [];
  //   }
  // },
  methods: {
    init() {
      this.form1.id = this.deviceId;
      this.form.protocolNo = "";
      this.form.id = this.deviceId;
      this.form.cardNo = "";
    },

    cancel() {
      this.$emit("closeDialog");
      this.$nextTick(() => {
        this.$refs.form.resetFields();
      });
      clearInterval(this.timer);
    },
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
      console.log("最终业务请求参数：", JSON.stringify(bizModel));
      fetchData(urls.operateInfo, bizModel, method.POST).then((res) => {
        //  res.data = [
        //   {
        //     addTime: "101010",
        //     content:
        //       "aaaa按计划发就哈哈警方哈吉哈吉奥干哈给发货啊发就好哈就会发aaaa按计划发就哈哈警方哈吉哈吉奥干哈给发货啊发就好哈就会发aaaa按计划发就哈哈警方哈吉哈吉奥干哈给发货啊发就好哈就会发aaaa按计划发就哈哈警方哈吉哈吉奥干哈给发货啊发就好哈就会发aaaa按计划发就哈哈警方哈吉哈吉奥干哈给发货啊发就好哈就会发aaaa按计划发就哈哈警方哈吉哈吉奥干哈给发货啊发就好哈就会发aaaa按计划发就哈哈警方哈吉哈吉奥干哈给发货啊发就好哈就会发aaaa按计划发就哈哈警方哈吉哈吉奥干哈给发货啊发就好哈就会发aaaa按计划发就哈哈警方哈吉哈吉奥干哈给发货啊发就好哈就会发aaaa按计划发就哈哈警方哈吉哈吉奥干哈给发货啊发就好哈就会发aaaa按计划发就哈哈警方哈吉哈吉奥干哈给发货啊发就好哈就会发aaaa按计划发就哈哈警方哈吉哈吉奥干哈给发货啊发就好哈就会发aaaa按计划发就哈哈警方哈吉哈吉奥干哈给发货啊发就好哈就会发aaaa按计划发就哈哈警方哈吉哈吉奥干哈给发货啊发就好哈就会发aaaa按计划发就哈哈警方哈吉哈吉奥干哈给发货啊发就好哈就会发aaaa按计划发就哈哈警方哈吉哈吉奥干哈给发货啊发就好哈就会发aaaa按计划发就哈哈警方哈吉哈吉奥干哈给发货啊发就好哈就会发aaaa按计划发就哈哈警方哈吉哈吉奥干哈给发货啊发就好哈就会发aaaa按计划发就哈哈警方哈吉哈吉奥干哈给发货啊发就好哈就会发aaaa按计划发就哈哈警方哈吉哈吉奥干哈给发货啊发就好哈就会发aaaa按计划发就哈哈警方哈吉哈吉奥干哈给发货啊发就好哈就会发aaaa按计划发就哈哈警方哈吉哈吉奥干哈给发货啊发就好哈就会发",
        //   },
        // ];
        if (res.code == 1) {
          this.infoList = res.data.concat(this.infoList);
        }
        console.log(this.infoList.length);
      });
    },
    handleClick(tab, event) {
      console.log(tab, event);
    },
    //开闸
    open: debounce(function () {
      // 待签名的数据
      let obj = {
        id: this.deviceId,
        // protocolNo:this.form.protocolNo,
        // cardNo:this.form.cardNo
      };

      // 1、获取按字典排序后的参数
      let sortedData = deepSortForPostJson(obj);

      // 2、对已排序的参数进行签名
      let timestamp = parseInt(new Date().getTime() / 1000);
      let sign = createSign(sortedData, timestamp);

      // 3、设置最终业务参数
      let bizModel = {
        data: sortedData,
        sign: sign,
        timestamp: timestamp,
      };
      // this.$refs.form.validate((valid) => {
      //   if (valid) {
      fetchData(urls.openGateTest, bizModel, method.POST).then((res) => {
        this.$message("操作成功");
        // this.$emit("closeDialog");
        this.getInfo();
        this.$nextTick(() => {
          this.$refs.form.resetFields();
        });
      });
    }, 200),

    //关闸
    close() {
      this.$emit("closeDialog");
      // let obj = {};
      // this.$refs.form.validate((valid) => {
      //   if (valid) {
      //     fetchData(urls., obj, method.POST).then((res) => {
      //       this.$message("操作成功");
      //       this.$emit("closeDialog");
      // this.getInfoList()
      //       this.$nextTick(() => {
      //         this.$refs.form.resetFields();
      //       });
      //     });
      //   }
      // });
    },

    order() {
      this.$emit("closeDialog");
      // let obj = {};
      // this.$refs.form.validate((valid) => {
      //   if (valid) {
      //     fetchData(urls.addNewVision, obj, method.POST).then((res) => {
      //       this.$message("操作成功");
      //       this.$emit("closeDialog");
      //       this.$nextTick(() => {
      //         this.$refs.form.resetFields();
      //       });
      //     });
      //   }
      // });
    },
    //验票测试
    test: debounce(function () {
      // 待签名的数据
      let obj = {
        id: this.deviceId,
        protocolNo: this.form.protocolNo,
        cardNo: this.form.cardNo,
      };
      // 1、获取按字典排序后的参数
      let sortedData = deepSortForPostJson(obj);

      // 2、对已排序的参数进行签名
      let timestamp = parseInt(new Date().getTime() / 1000);
      let sign = createSign(sortedData, timestamp);

      // 3、设置最终业务参数
      let bizModel = {
        data: sortedData,
        sign: sign,
        timestamp: timestamp,
      };
      this.$refs.form.validate((valid) => {
        if (valid) {
          fetchData(urls.checkTicket, bizModel, method.POST).then((res) => {
            this.$message("操作成功");
            // this.$emit("closeDialog");
            this.getInfoList();
            this.$nextTick(() => {
              this.$refs.form.resetFields();
            });
          });
        }
      });
    }, 200),
    //人数上报
    report: debounce(function () {
      this.$emit("closeDialog");
      // 待签名的数据
      let obj = {
        id: this.deviceId,
        // protocolNo:this.form.protocolNo,
        // cardNo:this.form.cardNo
      };

      // 1、获取按字典排序后的参数
      let sortedData = deepSortForPostJson(obj);

      // 2、对已排序的参数进行签名
      let timestamp = parseInt(new Date().getTime() / 1000);
      let sign = createSign(sortedData, timestamp);

      // 3、设置最终业务参数
      let bizModel = {
        data: sortedData,
        sign: sign,
        timestamp: timestamp,
      };

      // this.$refs.form.validate((valid) => {
      //   if (valid) {
      fetchData(urls.numUpload, bizModel, method.POST).then((res) => {
        this.$message("操作成功");
        // this.$emit("closeDialog");
        this.getInfoList();
        this.$nextTick(() => {
          this.$refs.form.resetFields();
        });
      });
      //   }
      // });
    }, 200),
  },
};
</script>
<style lang="less" scoped>
// @import "../../../assets/css/fullDialog.less";
.flex-between {
  display: flex;
  justify-content: space-between;
  // height: 32px;
  align-items: center;
  width: 100%;
  padding: 6px 0;
}
.item-date {
  color: #999;
  font-size: 14px;
  width: 40%;
  text-align: left;
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
}
.item-remark {
  color: #666;
  font-size: 14px;
  width: 60%;
  text-align: left;
  // overflow: hidden;
  // white-space: nowrap;
  // text-overflow: ellipsis;
  // padding-right: 12px;
}
.row-line {
  height: 1px;
  width: 100%;
  background: #ebebeb;
  margin: 10px 0 20px 0;
}
.form1 {
  // width: 320px;
  //  padding-right: 20px;
}

.right {
  // border-left:1px solid #ebebeb ;
  .title {
    font-weight: 600;
    color: #333;
  }
  .info-content {
    height: 240px;
    overflow: hidden;
    overflow-y: auto;
    padding: 20px 0 32px 0;
  }
}
</style>