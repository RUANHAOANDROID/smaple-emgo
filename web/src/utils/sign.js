import md5 from 'js-md5'

 function deepSortForPostJson(data) {
    let result;
    let dataType = checkedType(data);
    if (dataType == 'Array') {
      result = []
    } else if (dataType === 'Object') {
      result = {}
    } else {
      return data
    }

    for (let i in data) {
      // 对象、数组里嵌套了对象/数组，递归排序
      if (checkedType(data[i]) === 'Object' || checkedType(data[i]) === 'Array') { 
          result[i] = deepSortForPostJson(data[i]);
      } else { 
        // 转换为字符串
        if (data[i] === null || data[i] === undefined) {
          result[i] = ''
        } else {
          result[i] = String(data[i]);
        }
      }
    }

    return sortObj(result);
  }

 function checkedType(data) {
    return Object.prototype.toString.call(data).slice(8, -1)
  }

  function sortObj(obj) {
    if (checkedType(obj) === 'Object') {
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
  }

   function createSign(bizModel, timestamp) {
    let md5Key = '#!1EMCXu5eIx190'
    let jsonStr = JSON.stringify(bizModel)
    let md5Str = "data=" + jsonStr + "&key=" + md5Key + "&timestamp=" + timestamp;
    return md5(md5Str).toString();
  }
  export {
    createSign,
    deepSortForPostJson
  }