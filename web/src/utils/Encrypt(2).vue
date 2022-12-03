<template>

  <div>
    <div>
      <button @click="encrypt">测试签名</button>
    </div>
  </div>
</template>

<script>
import md5 from 'js-md5'
export default {
  name: "Encrypt",
  data() {
    return {
        
    };
  },
  methods: {
    encrypt() {
      // 待签名的数据
      let testData = {
        // b1: 'b1',
        // a1: 'a1',
        // c1: [
        //   {
        //     b2: 'b2',
        //     a2: 'a2'
        //   },
        //   {
        //     b3: 'b3',
        //     a3: 'a3'
        //   }
        // ]
      }

      console.log('待排序数据：', JSON.stringify(testData))

      // 1、获取按字典排序后的参数
      let sortedData = this.deepSortForPostJson(testData)
      console.log('已排序数据：', JSON.stringify(sortedData))

      // 2、对已排序的参数进行签名
      let timestamp = parseInt(new Date().getTime() / 1000);
      let sign = this.createSign(sortedData, timestamp)
      console.log('已排序参数签名：', sign)

      // 3、设置最终业务参数
      let bizModel = {
        data: sortedData,
        sign: sign,
        timestamp: timestamp
      }
      console.log('最终业务请求参数：', JSON.stringify(bizModel))

    },

    deepSortForPostJson(data) {
      let result;
      let dataType = this.checkedType(data);
      if (dataType == 'Array') {
        result = []
      } else if (dataType === 'Object') {
        result = {}
      } else {
        return data
      }

      for (let i in data) {
        // 对象、数组里嵌套了对象/数组，递归排序
        if (this.checkedType(data[i]) === 'Object' || this.checkedType(data[i]) === 'Array') { 
            result[i] = this.deepSortForPostJson(data[i]);
        } else { 
          // 转换为字符串
          if (data[i] === null || data[i] === undefined) {
            result[i] = ''
          } else {
            result[i] = String(data[i]);
          }
        }
      }

      return this.sortObj(result);
    },

    checkedType(data) {
      return Object.prototype.toString.call(data).slice(8, -1)
    },

    sortObj(obj) {
      if (this.checkedType(obj) === 'Object') {
        let keys = Object.keys(obj).sort();
        let arr = {};
        for (let i in keys) {
          let a = obj[keys[i]];
          arr[keys[i]] = obj[keys[i]];
        }
        return arr;
      } else {
          return obj;
      }
    },

    createSign(bizModel, timestamp) {
      let md5Key = '#!1EMCXu5eIx190'
      let jsonStr = JSON.stringify(bizModel)
      let md5Str = "data=" + jsonStr + "&key=" + md5Key + "&timestamp=" + timestamp;
      return md5(md5Str).toString();
    },
  }
};
</script>

<!-- 样式，scope 属性表示样式只在当前组件生效 -->
<style lang="css" scoped>
.container {
  padding: 20px;
}
.my-component{
  width: 100%;
  height: 90vh;
  border: 1px solid #000;
}
</style>