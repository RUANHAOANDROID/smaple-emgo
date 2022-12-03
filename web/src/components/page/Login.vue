<template>
  <div class="login-wraper">
    <div class="login-container">
      <!-- <div class="bgContainer">
        <img src="../../assets/img/bg_banner.jpeg" alt="" />
      </div> -->

      <div class="login-form">
        <p class="welcome">
          <!-- <img class="logo1" src="../../assets/img/logo-1.png" alt="" /> -->
          <span style="color: #333; font-weight: bold">中继服务</span>
        </p>

        <!-- <p class="welcome-title">登录</p> -->
        <div class="tab-title">
          <div class="title-box">
            <div
              @click="tabChange(0)"
              :class="[btnnum == 0 ? 'title-text' : 'opacity']"
            >
              登录
            </div>
            <div :class="[btnnum == 0 ? 'btna' : '']"></div>
          </div>

          　　
          <div class="title-box">
            　
            <div
              @click="tabChange(1)"
              :class="[btnnum == 1 ? 'title-text' : 'opacity']"
            >
              修改密码
            </div>
            <div :class="[btnnum == 1 ? 'btna' : '']"></div>
          </div>
        </div>
        <div v-if="btnnum == 0">
          <el-form
            :model="ruleForm"
            :rules="rules"
            ref="ruleForm"
            size="medium"
          >
            <el-form-item prop="userName">
              <el-input
                v-model="ruleForm.userName"
                placeholder="请输入用户名"
                prefix-icon="el-icon-user"
              >
              </el-input>
            </el-form-item>
            <el-form-item prop="password">
              <el-input
                type="password"
                prefix-icon="el-icon-lock"
                placeholder="请输入密码"
                v-model="ruleForm.password"
                @keyup.enter.native="submitForm('ruleForm')"
              >
              </el-input>
            </el-form-item>
            <!-- :disabled="update" -->
            <div class="login-btn">
              <el-button type="primary" @click="submitForm('ruleForm')"
                >立即登录</el-button
              >
            </div>
          </el-form>
        </div>
        <div v-if="btnnum == 1">
          <el-form
            :model="pswForm"
            status-icon
            :rules="pswrules"
            ref="pswForm"
            size="medium"
          >
            <el-form-item prop="user">
              <el-input
                v-model="pswForm.user"
                placeholder="请输入账号"
              ></el-input>
            </el-form-item>

            <el-form-item prop="oldPsw">
              <el-input
                v-model.number="pswForm.oldPsw"
                type="password"
                placeholder="请输入原密码"
              ></el-input>
            </el-form-item>
            <el-form-item prop="psw">
              <el-input
                type="password"
                v-model="pswForm.psw"
                autocomplete="off"
                placeholder="请输入新密码"
              ></el-input>
            </el-form-item>

            <div class="login-btn2">
              <el-button type="primary" @click="submitForm2('pswForm')"
                >提交</el-button
              >
              <el-button @click="resetForm('pswForm')">重置</el-button>
            </div>
          </el-form>
        </div>
      </div>
    </div>
    <div class="login-footer">
      <!-- © 2019 Copyright 2020 Yuanxin All Rights Reserved -->
      <!-- <div style="color: #fff">{{ updateText }}</div> -->
      <!-- <el-progress
        v-if="update"
        :text-inside="true"
        :stroke-width="10"
        :percentage="percentageLine"
        :color="colorLine"
      ></el-progress> -->
    </div>
  </div>
</template>

<script>
import { urls, method, fetchData, page } from "@/api/index";
import { createSign, deepSortForPostJson } from "@/utils/sign.js";
import { setToken, setMchId } from "@/utils/auth";

export default {
  data() {
    var checkOld = (rule, value, callback) => {
      if (!value) {
        return callback(new Error("原密码不能为空"));
      }
      //检验原密码是否正确
      let sortedData = deepSortForPostJson({
        userName: "admin",
        password: value,
      });

      let timestamp = parseInt(new Date().getTime() / 1000);
      let sign = createSign(sortedData, timestamp);

      let bizModel = {
        data: sortedData,
        sign: sign,
        timestamp: timestamp,
      };

      fetchData(urls.login, bizModel, method.POST)
        .then((res) => {
          if (res.code == 1) {
            callback();
          }
        })
        .catch((err) => {
          return callback("原密码不正确！");
        });
    };
    return {
      update: false,
      percentageLine: 0,
      updateText: "",
      colorLine: "",
      //   sysTitle: title,
      btnnum: 0, //tab
      ruleForm: {
        userName: "",
        password: "",
      },
      rules: {
        userName: [
          { required: true, message: "请输入用户名", trigger: "blur" },
        ],
        password: [{ required: true, message: "请输入密码", trigger: "blur" }],
      },
      pswForm: {
        user: "",
        psw: null,
        oldPsw: "",
      },
      pswrules: {
        psw: [{ required: true, message: "请输入原密码", trigger: "blur" }],
         user: [{ required: true, message: "请输入用户名", trigger: "blur" }],
        oldPsw: [{ validator: checkOld, trigger: "blur", required: true }],
      },
    };
  },
  created() {},
  components: {},
  mounted() {
    this.getNewVersion();
  },
  beforeDestroy() {
    clearInterval(this.percentageLine);
  },
  watch: {
    updateText(val) {
      if (val == "下载完成！") {
        this.autoUpdate(); //当下载到100%，调用接口自动升级
      }
    },
  },
  methods: {
    //获取最新版本
    getNewVersion() {
      this.update = true;
      //调接口查询如果存在最新版本，则下载最新文件,自动升级成功才能登录
      if (this.update) {
        this.uploadNewVersion();
      }
    },
    //下载最新文件
    uploadNewVersion() {
      //下载文件时调用接口动态更新进度
      // setInterval(() => {
      //   if (this.percentageLine == 100) {
      //     this.update = false;
      //     this.updateText = "下载完成！";
      //     return;
      //   } else {
      //     this.percentageLine += 10; //= Math.round(Math.random() * 100)
      //     this.colorLine = this.getcolorLine();
      //   }
      //   console.log(this.percentageLine);
      // }, 1000);
    },
    //自动升级（下载完成最新文件自动升级）
    autoUpdate() {
      clearInterval(this.percentageLine);
      console.log("自动升级");
      this.update = false;
    },
    //进度条颜色
    getcolorLine() {
      return (
        "#" +
        "0123456789abcdef"
          .split("")
          .map((v, i, a) => {
            return i > 5 ? null : a[Math.floor(Math.random() * 16)];
          })
          .join("")
      );
    },
    tabChange(e) {
      this.btnnum = e;
      if (e == 1) {
        this.$nextTick(() => {
          this.$refs["pswForm"].clearValidate(); // 清除表单验证
        });
      } else {
        this.$nextTick(() => {
          this.$refs["ruleForm"].resetFields(); // 清除表单验证
        });
      }
    },
    submitForm(formName) {
      let sortedData = deepSortForPostJson(this.ruleForm);

      let timestamp = parseInt(new Date().getTime() / 1000);
      let sign = createSign(sortedData, timestamp);

      let bizModel = {
        data: sortedData,
        sign: sign,
        timestamp: timestamp,
      };
      this.$refs[formName].validate((valid) => {
        if (valid) {
          fetchData(urls.login, bizModel, method.POST).then((res) => {
            localStorage.setItem("userInfo", JSON.stringify(res));
            setToken(res.code);
            this.$refs["ruleForm"].resetFields(); // 清除表单验证
            this.$router.push("/");
          });
        }
        
      });
    },
    submitForm2(formName) {
      let obj = {
        password:this.pswForm.psw,
        userName:this.pswForm.user
      }
      let sortedData = deepSortForPostJson(obj);

      let timestamp = parseInt(new Date().getTime() / 1000);
      let sign = createSign(sortedData, timestamp);

      let bizModel = {
        data: sortedData,
        sign: sign,
        timestamp: timestamp,
      };
      this.$refs[formName].validate((valid) => {
        if (valid) {
          fetchData(urls.password, bizModel, method.POST).then((res) => {
            this.$message.success("密码修改成功");
            this.$nextTick(() => {
          this.$refs["pswForm"].resetFields(); // 清除表单验证
        });
          });
        } else {
          console.log("error submit!!");
          return false;
        }
      });
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
  },
};
</script>

<style scoped lang="less">
.login-wraper {
  height: 100%;
  background-size: 100% 100%;
  background-image: url(../../assets/img/bg_banner.jpeg);
  display: flex;
  justify-content: center;
  flex-direction: column;
  align-items: center;
  overflow: hidden;
  position: relative;
}
.login-logo {
  width: 100%;
  padding: 36px 0 0 40px;
  box-sizing: border-box;
  color: #333;
  font-size: 18px;
  display: flex;
  align-items: center;
}
.login-footer {
  width: 650px;
  position: fixed;
  bottom: 30px;
  color: #999;
  text-align: center;
  opacity: 0.6;
}
.login-logo img {
  margin-right: 5px;
  vertical-align: middle;
}
.login-logo span {
  font-size: 18px;
  margin-left: 8px;
}
.login-logo span.line {
  font-size: 14px;
}

.login-container .bgContainer {
  width: 60%;
  height: 100%;
  float: left;
  display: flex;
  justify-items: center;
  align-items: center;
}

.login-container .bgContainer img {
  width: 100%;
}

.login-form {
  width: 650px;
  height: 415px;
  background: rgba(252, 250, 250, 0.5);
  border-radius: 24px;
  padding: 48px 40px;
  box-sizing: border-box;
  // margin-bottom: 40px;
  box-shadow: 0 0 20px -5px #999;
}
.login-form .welcome {
  height: 32px;
  // margin-bottom: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
}
.login-form .welcome .logo1 {
  margin-right: 8px;
}
.login-form .welcome span {
  font-size: 24px;
}
.login-form .welcome-title {
  text-align: center;
  font-size: 14px;
  color: #999;
  margin-bottom: 40px;
}

.login-form .code .el-input {
  width: 240px;
  float: left;
}

.login-btn button {
  width: 100%;
  height: 40px;
  margin-top: 30px;
}
.login-btn2 {
  display: flex;
  button {
    width: 50%;
    height: 40px;
    text-align: center;
    // margin-top: 30px;
  }
}
.tab-title {
  display: flex;
  font-size: 16px;
  font-weight: 500;
  color: #333;
  height: 64px;
  justify-content: center;
  margin-top: 30px;
}
.title-box {
  width: 50%;
  height: 80px;
  display: flex;
  flex-direction: column;
  align-items: center;
  cursor: pointer;
  .title-text {
    opacity: 1;
  }

  .opacity {
    opacity: 0.45;
  }

  .btna {
    width: 100px;
    height: 2px;
    background: #578efe;
    border-radius: 2px;
    margin-top: 10px;
  }
}

.dis {
  display: block;
  flex: 1;
}
.special-item /deep/ .el-form-item__content {
  display: flex;
  justify-content: space-between;
  width: 100%;
}
</style>