/**
 * @param {string} num1
 * @param {string} num2
 * @return {string}
 */
var multiply = function(num1, num2) {
  var num1_arr = num1.split();
  var num2_arr = num2.split();
  var sum_arr = [];
  var products_arr = [];
  for (var i = 0; i < num1_arr.length; i++) {
     products_arr.push(get_products_arr(num1_arr[i], num2_arr));
  }

  for (i = 0; i < products_arr.length; i++) {
    sum_arr = get_sum_arr(sum_arr, products_arr[i]);
  }

  return sum_arr;
};

var get_products_arr = function(num1, num2_arr) {
  var sum_arr = [0];
  for (i = 1; i <= num1; i++) {
    sum_arr = get_sum_arr(sum_arr, num2_arr);
  }

  return sum_arr;
}

// 大数相加
var get_sum_arr = function(sum_arr, num2_arr) {
  var len = sum_arr.length > num2_arr.length ? sum_arr.length : num2_arr.length;
  for (var i = 0; i < len - sum_arr.length; i ++) {
    sum_arr.unshift(0);
  }

  for (var i = 0; i < len - num2_arr.length; i ++) {
    sum_arr.unshift(0);
  }

  var to_next = 0;
  for (var i = len - 1; i >= 0; i --) {
    var sum = Number(sum_arr[i]) + Number(num2_arr[i]) + to_next;
    var tmp = sum % 10;
    to_next = sum - tmp;
    sum_arr[i] = tmp;
  }

  if (to_next > 0) sum_arr.unshift(to_next);

  return sum_arr;
}
