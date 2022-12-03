<template>
  <div class="eqi-wrapper">
     <div class="table_content">
       <div class="title"><div class="rectangle"></div>设备列表</div>
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
          >
           <!-- <template slot-scope="scope">
              <a
                @click="handleAdd(scope.row.id)"
                style="color: #1890ff; cursor: pointer"
                >{{ scope.row.id }}</a
              >
            </template> -->
            </el-table-column>
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
          <el-table-column
            prop="status"
            label="状态"
            show-overflow-tooltip
          >
          <template slot-scope="scope">
            <div>{{ scope.row.status | isActiveFilter }}
            </div>
            </template>
          </el-table-column>
           <el-table-column label="操作">
          <template slot-scope="scope">
            <el-button
              size="mini"
              type="text"
              @click="handleAdd(scope.row.id)"
              style="color: #419eef"
              >设备调试</el-button
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
      @closeDialog="closeAddDialog"
      :deviceId="deviceId"
    ></AddDialog>
  </div>
</template>
<script>
import AddDialog from "./operate.vue";
import { urls, method, fetchData, page } from "@/api/index";
import { createSign, deepSortForPostJson } from "@/utils/sign.js";

export default {
  name:'eqiupment',//设备调试
  components:{AddDialog},
  data() {
    return{
      tableHeight: window.innerHeight - 278, //默认table高度
      tableData: [
        {
          id:'1',
          ip:'',
          version:'v-1.0',
          status:'01'
        },
        {
          id:'2',
          ip:'',
          version:'v-1.0',
          status:'02'
        },
        {
          id:'3',
          ip:'',
          version:'v-1.0',
          status:'03'
        }
      ],
      pageSizes: [10, 20, 50],
      pageSize: 10,
      pageNum: 1,
      currentTotal:0,
      addVisible:false,
      deviceId:''
    }
  },
  created() {
     this.getData();
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
  mounted(){
      
  },
  methods:{
    //操作弹框
    handleAdd(id) {
      this.addVisible = true;
      this.deviceId = id
    },
    //关闭弹框
    closeAddDialog() {
      this.addVisible = false;
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
    // 获取列表数据
    getData() {
     
     // 待签名的数据
      let testData = {};


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
      fetchData(urls.deviceList, bizModel, method.POST).then((res) => {
        if (res.data ) {
          this.tableData = res.data;
          this.currentTotal = parseInt(res.data.length);
        }
      }).catch(err => {

      })
      this.addVisible = false;
    },
  }
}
</script>
<style lang="less" scoped>
.eqi-wrapper{
  width: 100%;
  background: #f8f8f8;
}
.title{
  padding: 0 0 20px 0;
    display: flex;
    align-items: center;
}
  .rectangle{
    background: #419eef;
    width: 5px;
    height: 14px;
    margin-right: 8px;
    margin-top: 2px;
    border-radius: 2px;
  }
 .table_content {
    margin: 24px;
    background: #fff;
    padding: 20px;
  }
    .table_content /deep/ .el-table {
    width: 99.5%;
  }
  .pagination {
    margin: 24px 24px 0 0 ;
    background: #fff;
  }
</style>