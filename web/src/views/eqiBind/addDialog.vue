<template>
  <div :class="[collapse ? 'dialog-wrapper-collapse' : 'dialog-wrapper']">
    <el-dialog
      class="el-dialog-detail"
      :width="dialogWidth"
      :visible.sync="show"
      @close="cancel"
      :show-close="false"
      :before-close="cancel"
      title="新增设备"
    >
      <div class="line"></div>
      <el-form
        ref="form"
        :model="form"
        :rules="rules"
        label-width="120px"
        class="form1"
      >
        <div class="title">机器绑定设置</div>
        <el-form-item label="中继设备号" prop="id">
          <el-input
            style="width: 208px"
            v-model="form.id"
            clearable
            placeholder=""
          ></el-input>
        </el-form-item>
        <el-form-item label="设备编号" prop="deviceNo">
          <el-input
            style="width: 208px"
            v-model="form.deviceNo"
            clearable
            placeholder=""
          ></el-input>
        </el-form-item>
        <el-form-item label="控制板IP" prop="deviceIp">
          <el-input
            style="width: 208px"
            v-model="form.deviceIp"
            clearable
            placeholder=""
          ></el-input>
        </el-form-item>

        <el-form-item label="版本" prop="deviceVersion" style="width: 208px">
          <el-input
            style="width: 208px"
            v-model="form.deviceVersion"
            placeholder=""
          ></el-input>
        </el-form-item>
        
        <el-form-item label="设备序列号" prop="serialNumber" style="width: 208px">
          <el-input
            style="width: 208px"
            v-model="form.serialNumber"
            placeholder=""
          ></el-input>
        </el-form-item>
      </el-form>

      <div
        style="
          width: 90%;
          display: flex;
          justify-content: center;
          padding: 24px 0;
        "
      >
        <el-button @click="submit" type="primary">确定</el-button>
        <el-button @click="cancel">取消</el-button>
      </div>
    </el-dialog>
  </div>
</template>
 
<script>
import bus from "@/components/common/bus";
import { urls, method, fetchData, page } from "@/api/index";
import { createSign, deepSortForPostJson } from "@/utils/sign.js";
export default {
  name: "vision-add", //报修详情
  //   components: { equipmentInfo, StepInfo, Handle },
  data() {
    return {
      // visible: this.show,
      modalMask: false,
      dialogWidth: "880",
      rules: {
        id: [{ required: true, message: "请输入", trigger: "blur" }],
        deviceNo: [{ required: true, message: "请输入", trigger: "blur" }],
        deviceIp: [{ required: true, message: "请输入", trigger: "blur" }],
        deviceVersion: [{ required: true, message: "请输入", trigger: "blur" }],
        serialNumber:[{ required: true, message: "请输入", trigger: "blur" }],
      },
      form: {
        id: "",
        deviceNo: "",
        deviceIp: "",
        deviceVersion: "",
        serialNumber:'',
      },
      Option: [],

      collapse: false,
      successImg: [],
      // dialogImageUrl: "",
      dialogVisible: false,
    };
  },
  props: {
    show: {
      type: Boolean,
      default: false,
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
      }
    },
  },
  mounted() {},
  methods: {
    init() {
      this.form.deviceNo = "";
      this.form.id = "";
      this.form.deviceIp = "";
      this.form.deviceVersion = "";
    },
    cancel() {
      this.$emit("closeDialog");
      this.$nextTick(() => {
        this.$refs.form.resetFields();
      });
      // this.visible = false;
    },
    submit() {
      let obj = Object.assign(this.form, {
        deviceStatus: "01"
      });
      // 待签名的数据
      console.log("待排序数据：", JSON.stringify(obj));

      // 1、获取按字典排序后的参数
      let sortedData = deepSortForPostJson(this.form);
      console.log("已排序数据：", JSON.stringify(sortedData));

      // 2、对已排序的参数进行签名
      let timestamp = parseInt(new Date().getTime() / 1000);
      let sign = createSign(sortedData, timestamp);
      console.log("已排序参数签名：", sign);

      // 3、设置最终业务参数
      let bizModel = {
        data: sortedData,
        sign: sign,
        timestamp: timestamp,
      };
      console.log("最终业务请求参数：", JSON.stringify(bizModel));
      this.$refs.form.validate((valid) => {
        if (valid) {
          fetchData(urls.addDevice, bizModel, method.POST).then((res) => {
            this.$message.success("新增成功");
            this.cancel();
            this.$emit("getData");
            this.$nextTick(() => {
              this.$refs.form.resetFields();
            });
          });
        }
      });
    },
  },
};
</script>
<style lang="less" scoped>
.title {
  font-weight: 550;
  padding: 24px 0;
}
.line {
  width: 100%;
  height: 1px;
  background: #ebebeb;
}
.line-box {
  display: flex;
  align-items: center;
  margin-bottom: 24px;
}
.between {
  display: flex;
  justify-content: space-between;
  margin: 12px 0;
}
</style>