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
  var sum_arr = [];
  for (i = 1; i <= num1; i++) {
    sum_arr = get_sum_arr(sum_arr, num2_arr);
  }
}

// 大数相加
var get_sum_arr = function(sum_arr, num2_arr) {

}
