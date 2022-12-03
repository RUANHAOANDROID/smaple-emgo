/**
 * 获取当前日期
 * @returns {string}
 */
 export function getCurrentDate() {
    var date = new Date();
    var seperator1 = "-";
    var month = date.getMonth() + 1;
    var strDate = date.getDate();
    if (month >= 1 && month <= 9) {
      month = "0" + month;
    }
    if (strDate >= 0 && strDate <= 9) {
      strDate = "0" + strDate;
    }
    var currentdate = date.getFullYear() + seperator1 + month + seperator1 + strDate;
    return currentdate;
  }

  
/**
* 获取当前时间
* @returns {string}
*/ 
export function getCurrentTime() {
    var date = new Date();
    var seperator1 = "-";
    var seperator2 = ":";
    var month = date.getMonth() + 1;
    var strDate = date.getDate();
    if (month >= 1 && month <= 9) {
      month = "0" + month;
    }
    if (strDate >= 0 && strDate <= 9) {
      strDate = "0" + strDate;
    }
    var currentdate = date.getFullYear() + seperator1 + month + seperator1 + strDate
      + " " + date.getHours() + seperator2 + date.getMinutes()
      + seperator2 + date.getSeconds();
    return currentdate;
  }
  
  //获取本月最后一天日期
  export function getLastDay() {
    let date = new Date();
    let currentMonth = date.getMonth();
    let nextMonth = ++currentMonth;
    let nextMonthFirstDay = new Date(date.getFullYear(), nextMonth, 1);
    let oneDay = 1000 * 60 * 60 * 24;
    let lastTime = new Date(nextMonthFirstDay - oneDay);
    let month = parseInt(lastTime.getMonth() + 1);
    let day = lastTime.getDate();
    if (month < 10) {
      month = '0' + month
    }
    if (day < 10) {
      day = '0' + day
    }
    return date.getFullYear() + '-' + month + '-' + day;
  }
  
  /**
   * 日期加 days 天
   * @param {*} date
   * @param {*} days
   */
  export function addDate(date, days) {
    if (days == undefined || days == '') {
      days = 0;
    }
    var date = new Date(date);
    date.setDate(date.getDate() + days);
    var month = date.getMonth() + 1;
    var day = date.getDate();
    var mm = "'" + month + "'";
    var dd = "'" + day + "'";
  
    //单位数前面加0
    if (mm.length == 3) {
      month = "0" + month;
    }
    if (dd.length == 3) {
      day = "0" + day;
    }
    var time = date.getFullYear() + "-" + month + "-" + day;
    return time;
  }
  
  /**
   * 日期加 n 年
   * @param {*} date
   * @param {*} days
   */
  export function addYear(date, n) {
    if (n == undefined || n == '') {
      n = 0;
    }
    var date = new Date(date);
    date.setFullYear(date.getFullYear() + n);
    date.setDate(date.getDate() - 1);
    var month = date.getMonth() + 1;
    var day = date.getDate();
    var mm = "'" + month + "'";
    var dd = "'" + day + "'";
  
    //单位数前面加0
    if (mm.length == 3) {
      month = "0" + month;
    }
    if (dd.length == 3) {
      day = "0" + day;
    }
    var time = date.getFullYear() + "-" + month + "-" + day;
    return time;
  }