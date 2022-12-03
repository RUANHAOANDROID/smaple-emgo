<template>
  <div class="set-wrapper">
    <div class="set-content">
      <div class="title">
        <div class="rectangle"></div>
        参数设置
      </div>
      <el-form
        ref="form"
        :model="form"
        :rules="rules"
        label-width="140px"
        class="form1"
      >
        <div style="display: flex">
          <el-form-item label="配置ID" prop="id">
            <el-input style="width: 208px" v-model="form.id"></el-input>
          </el-form-item>
          <el-form-item label="参数地址" prop="configUrl">
            <el-input style="width: 556px" v-model="form.configUrl"></el-input>
          </el-form-item>
          <el-button
            type="primary"
            style="height: 32px; margin-left: 24px"
            @click="getInfoData"
            >获取</el-button
          >
        </div>
      </el-form>
    </div>
    <div class="set-content set-content2">
      <el-form
        ref="form2"
        :model="form2"
        :rules="rules2"
        label-width="140px"
        class="form2"
      >
        <div style="display: flex">
          <el-form-item label="厂商标识" prop="manufacturerId1">
            <el-input
              style="width: 208px"
              v-model="form2.manufacturerId1"
            ></el-input>
          </el-form-item>

          <el-form-item label="厂商标识2" prop="manufacturerId2">
            <el-input
              style="width: 208px"
              v-model="form2.manufacturerId2"
              placeholder=""
            ></el-input>
          </el-form-item>
          <el-form-item label="缓冲开闸" style="width: 208px">
            <el-checkbox
              v-model="form2.buffer"
              placeholder=""
              @change="changBuffer"
            ></el-checkbox>
          </el-form-item>
        </div>
        <div style="display: flex">
          <el-form-item label="启动延时（秒）" prop="delayedTime">
            <el-input
              style="width: 208px"
              v-model="form2.delayedTime"
            ></el-input>
          </el-form-item>
          <el-form-item label="验票后未过闸" prop="invalidTime">
            <el-input
              v-model="form2.invalidTime"
              clearable
              style="width: 208px"
            >
              <template slot="append">毫秒后失效</template>
            </el-input>
          </el-form-item>
        </div>
        <div style="display: flex">
          <el-form-item label="人数上传间隔时间" prop="numUpTime">
            <el-input style="width: 208px" v-model="form2.numUpTime"></el-input>
          </el-form-item>

          <el-form-item label="心跳检测地址" prop="heartbeatUrl">
            <el-input
              v-model="form2.heartbeatUrl"
              clearable
              style="width: 208px"
            ></el-input>
          </el-form-item>
          <el-form-item label="心跳检测" prop="heartbeatTime">
            <el-input
              v-model="form2.heartbeatTime"
              clearable
              style="width: 208px"
            >
              <template slot="append">秒</template>
            </el-input>
          </el-form-item>
        </div>
        <div style="display: flex">
          <el-form-item label="门票验票地址" prop="checkUrl">
            <el-input
              style="width: 208px"
              v-model="form2.checkUrl"
              placeholder=""
            ></el-input>
          </el-form-item>
          <el-form-item label="门票核销地址" prop="writeOffUrl">
            <el-input
              style="width: 208px"
              v-model="form2.writeOffUrl"
              placeholder=""
            ></el-input>
          </el-form-item>
          <el-form-item label="人数上报地址" prop="numUpUrl">
            <el-input
              style="width: 208px"
              v-model="form2.numUpUrl"
              placeholder=""
            ></el-input>
          </el-form-item>
        </div>

        <div style="font-weight: bold; margin: 24px 0">验票通过提示</div>
        <div style="display: flex">
          <el-form-item label="普通票" prop="trueVoice4">
            <el-input v-model="form2.trueVoice4" clearable style="width: 384px">
              <template slot="append">门票/套票/合并票</template>
            </el-input>
          </el-form-item>
          <el-form-item label="年卡1" prop="trueVoice1">
            <el-input v-model="form2.trueVoice1" clearable style="width: 384px">
              <template slot="append">普通年卡</template>
            </el-input>
          </el-form-item>
        </div>
        <div style="display: flex">
          <el-form-item label="年卡2" prop="trueVoice2">
            <el-input v-model="form2.trueVoice2" clearable style="width: 384px">
              <template slot="append">员工年卡</template>
            </el-input>
          </el-form-item>
          <el-form-item label="其他" prop="trueVoice3">
            <el-input v-model="form2.trueVoice3" clearable style="width: 384px">
              <template slot="append">其它类型提示</template>
            </el-input>
          </el-form-item>
        </div>
        <div style="font-weight: bold; margin: 24px 0">验票失败提示</div>
        <div style="display: flex">
          <el-form-item label="无效票" prop="falseVoice1">
            <el-input
              v-model="form2.falseVoice1"
              clearable
              style="width: 384px"
            >
              <template slot="append">无效票</template>
            </el-input>
          </el-form-item>
          <el-form-item label="异常" prop="falseVoice2">
            <el-input
              v-model="form2.falseVoice2"
              clearable
              style="width: 384px"
            >
              <template slot="append">其它类型提示</template>
            </el-input>
          </el-form-item>
        </div>
        <div style="display: flex">
          <el-form-item label="已验票" prop="falseVoice3">
            <el-input
              v-model="form2.falseVoice3"
              clearable
              style="width: 384px"
            >
              <template slot="append">已验过的票</template>
            </el-input>
          </el-form-item>
          <el-form-item label="其他" prop="falseVoice4">
            <el-input
              v-model="form2.falseVoice4"
              clearable
              style="width: 384px"
            >
              <template slot="append">其它类型提示</template>
            </el-input>
          </el-form-item>
        </div>
      </el-form>
      <div class="btn-box">
        <el-button type="primary" size="small" @click="submit">保存</el-button>
      </div>
    </div>
  </div>
</template>
<script>
import { urls, method, fetchData, page } from "@/api/index";
import { createSign, deepSortForPostJson } from "@/utils/sign.js";
export default {
  name: "setting", //参数设置
  data() {
    return {
      rules: {
        id: [{ required: true, message: "请输入配置id", trigger: "blur" }],
        configUrl: [
          { required: true, message: "请输入参数地址", trigger: "blur" },
        ],
      },
      form: {
        id: "",
        configUrl: "",
      },
      form2: {
        manufacturerId1: "",
        manufacturerId2: "",
        buffer: "",
        delayedTime: "",
        invalidTime: "",
        writeOffUrl: "",
        checkUrl: "",
        numUpUrl: "",
        numUpTime: "",
        heartbeatTime: "",
        deTrueVoice: "",
        deFalseVoice: "",
        deFalseText: "",
        deTrueText: "",
        heartbeatUrl: "",
        trueVoice1: "",
        trueVoice2: "",
        trueVoice3: "",
        trueVoice4: "",
        falseVoice1: "",
        falseVoice2: "",
        falseVoice3: "",
        falseVoice4: "",
      },
      rules2: {
        manufacturerId1: [
          { required: true, message: "请输入", trigger: "blur" },
        ],
        manufacturerId2: [
          { required: true, message: "请输入", trigger: "blur" },
        ],
        // buffer: [{ required: true, message: "请输入", trigger: "change" }],
        delayedTime: [{ required: true, message: "请输入", trigger: "blur" }],
        invalidTime: [{ required: true, message: "请输入", trigger: "blur" }],
        writeOffUrl: [{ required: true, message: "请输入", trigger: "blur" }],
        checkUrl: [{ required: true, message: "请输入", trigger: "blur" }],
        numUpUrl: [{ required: true, message: "请输入", trigger: "blur" }],
        numUpTime: [{ required: true, message: "请输入", trigger: "blur" }],
        heartbeatTime: [{ required: true, message: "请输入", trigger: "blur" }],
        deTrueVoice: [{ required: true, message: "请输入", trigger: "blur" }],
        deFalseVoice: [{ required: true, message: "请输入", trigger: "blur" }],
        deFalseText: [{ required: true, message: "请输入", trigger: "blur" }],
        deTrueText: [{ required: true, message: "请输入", trigger: "blur" }],
        heartbeatUrl: [{ required: true, message: "请输入", trigger: "blur" }],
        trueVoice1: [
          { required: true, message: "请输入年卡1通过提示", trigger: "blur" },
        ],
        trueVoice2: [
          { required: true, message: "请输入年卡2通过提示", trigger: "blur" },
        ],
        trueVoice3: [
          { required: true, message: "请输入其他通过提示", trigger: "blur" },
        ],
        trueVoice4: [
          { required: true, message: "请输入普通票通过提示", trigger: "blur" },
        ],
        falseVoice1: [
          { required: true, message: "请输入无效票失败提示", trigger: "blur" },
        ],
        falseVoice2: [
          { required: true, message: "请输入异常失败提示", trigger: "blur" },
        ],
        falseVoice3: [
          { required: true, message: "请输入已验票失败提示", trigger: "blur" },
        ],
        falseVoice4: [
          { required: true, message: "请输入其他失败提示", trigger: "blur" },
        ],
      },
      Option: [
        {
          value: "01",
          label: "身份证",
        },
      ],
    };
  },
  mounted() {
    this.getCurrent();
  },
  activated() {
    console.log("setting激活activated钩子函数");
  },
  methods: {
    init() {
      this.$nextTick(() => {
        this.$refs.form.resetFields();
        this.$refs.form2.resetFields();
      });
    },
    getCurrent() {
      let sortedData = deepSortForPostJson({});
      console.log("已排序数据：", JSON.stringify(sortedData));

      // 2、对已排序的参数进行签名
      let timestamp = parseInt(new Date().getTime() / 1000);
      let sign = createSign(sortedData, timestamp);
      console.log("已排序参数签名：", sign);

      // 3、设置最终业务参数
      let bizModel = {
        data: {},
        sign: sign,
        timestamp: timestamp,
      };

      fetchData(urls.gatSetting, bizModel, method.POST).then((res) => {
        if (res.code == 1) {
          if (res.data) {
            this.form.id = res.data.id;
            this.form.configUrl = res.data.configUrl;
            this.form2.manufacturerId1 = res.data.manufacturerId1;
            this.form2.manufacturerId2 = res.data.manufacturerId2;
            this.form2.buffer = res.data.buffer == 1 ? true : false;
            this.form2.delayedTime = res.data.delayedTime;
            this.form2.invalidTime = res.data.invalidTime;
            this.form2.writeOffUrl = res.data.writeOffUrl;
            this.form2.checkUrl = res.data.checkUrl;
            this.form2.numUpUrl = res.data.numUpUrl;
            this.form2.numUpTime = res.data.numUpTime; //
            this.form2.heartbeatTime = res.data.heartbeatTime; //
            this.form2.heartbeatUrl = res.data.heartbeatUrl;
            this.form2.deTrueVoice = res.data.deTrueVoice; //
            this.form2.deFalseVoice = res.data.deFalseVoice; //
            this.form2.deFalseText = res.data.deFalseText; //
            this.form2.deTrueText = res.data.deTrueText; //
            this.form2.trueVoice1 = res.data.trueVoice1; //年卡语言
            this.form2.trueVoice2 = res.data.trueVoice2;
            this.form2.trueVoice3 = res.data.trueVoice3;
            this.form2.trueVoice4 = res.data.trueVoice4;
            this.form2.falseVoice1 = res.data.falseVoice1;
            this.form2.falseVoice2 = res.data.falseVoice2;
            this.form2.falseVoice3 = res.data.falseVoice3;
            this.form2.falseVoice4 = res.data.falseVoice4;
          } else {
            this.init();
          }
        }
      });
    },
    //获取
    getInfoData() {
      let obj = {
        id: this.form.id,
        configUrl: this.form.configUrl,
      };
      // 1、获取按字典排序后的参数
      let sortedData = deepSortForPostJson(obj);
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
      this.$refs.form.validate((valid) => {
        if (valid) {
          fetchData(urls.getConfig, bizModel, method.POST).then((res) => {
            if (res.code == 1) {
              this.form2.manufacturerId1 = res.data.manufacturerId1;
              this.form2.manufacturerId2 = res.data.manufacturerId2;
              this.form2.buffer = res.data.buffer == 1 ? true : false;
              this.form2.delayedTime = res.data.delayedTime;
              this.form2.invalidTime = res.data.invalidTime;
              this.form2.writeOffUrl = res.data.writeOffUrl;
              this.form2.checkUrl = res.data.checkUrl;
              this.form2.numUpUrl = res.data.numUpUrl;
              this.form2.numUpTime = res.data.numUpTime; //
              this.form2.heartbeatTime = res.data.heartbeatTime; //
              this.form2.heartbeatUrl = res.data.heartbeatUrl;
              this.form2.deTrueVoice = res.data.deTrueVoice; //
              this.form2.deFalseVoice = res.data.deFalseVoice; //
              this.form2.deFalseText = res.data.deFalseText; //
              this.form2.deTrueText = res.data.deTrueText; //
              this.form2.trueVoice1 = res.data.trueVoice1; //年卡语言
              this.form2.trueVoice2 = res.data.trueVoice2;
              this.form2.trueVoice3 = res.data.trueVoice3;
              this.form2.trueVoice4 = res.data.trueVoice4;
              this.form2.falseVoice1 = res.data.falseVoice1;
              this.form2.falseVoice2 = res.data.falseVoice2;
              this.form2.falseVoice3 = res.data.falseVoice3;
              this.form2.falseVoice4 = res.data.falseVoice4;
            }
          });
        }
      });
    },
    //保存配置
    submit() {
      this.form2.buffer = this.form2.buffer ? "1" : "0";
      let obj = Object.assign(this.form2, this.form);
      // 1、获取按字典排序后的参数
      let sortedData = deepSortForPostJson(obj);
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
      this.$refs.form2.validate((valid) => {
        if (valid) {
          fetchData(urls.saveConfig, bizModel, method.POST).then((res) => {
            this.$message("操作成功");
            // this.$nextTick(() => {
            //   this.init();
            // });
          });
        }
      });
    },
    //缓冲开闸
    changBuffer(e) {
      this.form2.buffer = e;
    },
  },
};
</script>
    
<style lang="less" scoped>
.set-wrapper {
  width: 100%;
  background: #f8f8f8;
}
.title {
  padding: 0 0 20px 0;
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
.set-content {
  margin: 24px;
  background: #fff;
  padding: 20px;
  overflow: auto;
}
.set-content2 {
  padding-bottom: 0;
  height: calc(100vh - 316px);
  position: relative;
}
.form2 {
  padding: 20px 0 10px 0;
}
.btn-box {
  width: 100%;
  // background: linear-gradient(120deg, #fff 0%, #c2e9fb 60%, #409eef 95%);
  display: flex;
  justify-content: center;
  padding: 0 0 20px 0;
  // margin: 0 11px;
  //position: absolute;
  bottom: 0;
}
</style>