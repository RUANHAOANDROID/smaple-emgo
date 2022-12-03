<template>
  <div class="eqiBind_container">
    <div class="table_body">
      <div style="margin-bottom: 12px">
        <!-- <div class="left">      v-has="'add'"-->
        <el-button
          type="primary"
          size="small"
          icon="el-icon-circle-plus-outline"
          @click="handleAdd"
          >新增设备</el-button
        >
      </div>

      <el-table
        :data="tableData"
        :height="tableHeight"
        tooltip-effect="dark"
        :header-cell-style="{ background: 'rgba(0,0,0,0.02)' }"
        :row-style="{ height: '54px' }"
        :header-row-style="{ height: '54px' }"
      >
        <!-- <el-table-column type="selection" width="55"></el-table-column> -->
        <el-table-column prop="id" label="中继设备号" width="160">
        </el-table-column>
        <el-table-column
          prop="deviceNo"
          label="设备编号"
          width="120"
          show-overflow-tooltip
        ></el-table-column>
        <el-table-column prop="deviceIp" label="控制ip"> </el-table-column>
        <el-table-column prop="deviceVersion" label="版本"> </el-table-column>
        <el-table-column prop="status" label="状态">
          <template slot-scope="scope">
            <div v-if="scope.row.status =='01'" style="color: #e02433">
              {{ scope.row.status | isActiveFilter }}
            </div>
            <div v-else style="color: #419eef">
              {{ scope.row.status | isActiveFilter }}
            </div>
          </template></el-table-column
        >
        <el-table-column label="操作">
          <template slot-scope="scope">
            <el-button
              size="mini"
              type="text"
              @click="edit(scope.row)"
              style="color: #419eef"
              >编辑</el-button
            >
            <el-button
              size="mini"
              type="text"
              style="color: red"
              @click="handleDelete(scope.row)"
              >删除</el-button
            >
          </template>
        </el-table-column>
      </el-table>

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

    <!-- 弹出框 -->
    <AddDialog
      ref="addDialog"
      :show="addVisible"
      @closeDialog="closeDialog"
      @getData="getData"
    ></AddDialog>
    <EditDialog
      ref="editDialog"
      :editVisible="editVisible"
      @closeDialog2="closeDialog2"
      @getData="getData"
      :editData="editData"
    ></EditDialog>
  </div>
</template>

<script>
import AddDialog from "./addDialog";
import EditDialog from "./editDialog.vue";
import { urls, method, fetchData, page } from "@/api/index";
import { createSign, deepSortForPostJson } from "@/utils/sign.js";
export default {
  name: "sys_version",
  components: {
    AddDialog,
    EditDialog,
  },
  data() {
    return {
      menu: [],
      currentRoute: "", //当前路由
      inputLen: 2, //搜索框个数
      isOpenSeach: false, //显示隐藏过多的搜索框
      tableHeight: window.innerHeight - 278, //默认table高度
      form: {},
      daterange: [],
      deviceType: "",
      statusOption: [],
      tableData: [],
      multipleSelection: [],

      addVisible: false,
      editVisible: false,
      currentTotal: 0,

      pageNum: 1,
      pageSizes: [10, 20, 50],
      pageSize: 10,
      isStop: false, //停用启用
      editData:{}
    };
  },
  created() {
    this.getData();
  },
    activated() {
    console.log("eqbind激活activated钩子函数");
  },
  filters: {
    isActiveFilter: function (val) {
      if(val == '01'){
        return '在线'
      }else if(val == '02'){
        return '离线'
      }else if(val == '03'){
        return '停用'
      }else if(val == '04'){
        return '报废'
      }else if(val == '05'){
        return '异常'
      }else if(val == '06'){
        return '报修中'
      }else {
        return ''
      }


    },
  },
  watch: {
    $route: {
      handler: function (route) {
        console.log(route);
        this.currentRoute = route;
      },
      immediate: true,
    },
    tableHeight() {},
  },
  computed: {},
  methods: {
    //分页
    handleSizeChange(val) {
      this.pageSize = val;
      this.getData();
    },
    // 分页
    handleCurrentChange(val) {
      this.pageNum = val;
      this.getData();
    },
    // 获取列表数据
    getData() {
      // 待签名的数据
      let testData = {};

      console.log("待排序数据：", JSON.stringify(testData));

      // 1、获取按字典排序后的参数
      let sortedData = deepSortForPostJson(testData);
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
      console.log("最终业务请求参数：", JSON.stringify(bizModel));
      fetchData(urls.deviceList, bizModel, method.POST).then((res) => {
        if (res.code == 1) {
          this.tableData = res.data;
          this.currentTotal = res.data.length;
        }
      });
      this.addVisible = false;
    },

    //新增
    handleAdd() {
      this.addVisible = true;
    },
    //关闭新增
    closeDialog() {
      this.addVisible = false;
    },
    //编辑
    closeDialog2() {
      this.editVisible = false;
    },
    //停用
    edit(row) {

      this.editVisible = true;
      this.editData = row;
    },

    // 确定删除
    handleDelete(row) {
      // 待签名的数据
      console.log("待排序数据：", JSON.stringify({ id: row.id }));

      // 1、获取按字典排序后的参数
      let sortedData = deepSortForPostJson({ id: row.id });
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
      this.$confirm("确认删除此记录吗？")
        .then((_) => {
          fetchData(urls.deleteDevice, bizModel, method.POST).then((res) => {
            this.$message({ message: "删除成功!", type: "success" });

            this.getData();
          });
        })
        .catch((_) => {});
    },
  },
  mounted() {
    window.addEventListener("resize", this.getHeight);
  },
  destroyed() {
    window.removeEventListener("resize", this.getHeight);
  },
};
</script>

<style  lang="less" scoped>
.eqiBind_container {
  width: 100%;
  background: #f8f8f8;
}
.table_body {
  margin: 24px;
  background: #fff;
  padding: 20px;
}
.table_content /deep/ .el-table {
  width: 99.5%;
}
.pagination {
  margin: 24px 24px 0 0;
  background: #fff;
}

// .table_body {
//   margin: 0 16px 16px 16px;
//   height: calc(100vh - 180px);
//   background-color: #fff;
// }
</style>
