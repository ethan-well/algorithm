var multiply = function(num1, num2) {
  var num1_arr = num1.split('');
  var num2_arr = num2.split('');
  var sum_arr = [];
  var products_arr = [];
  for (var i = 0; i < num1_arr.length; i++) {
    products_arr.push( get_products_arr( Number(num1_arr[i]), num1_arr.length - 1 - i,  num2 ) );
  }

  for (i = 0; i < products_arr.length; i++) {
    sum_arr = get_sum_arr(sum_arr, products_arr[i]);
  }

  return sum_arr.join('');
};

var get_products_arr = function(num1, exponent, num2) {
  var num2_arr = num2.split('');
  var num2_arr_len = num2_arr.length;
  var to_next = 0;
  for (var i = num2_arr_len - 1; i >= 0; i--) {
    var product = num1 * num2_arr[i];
    var tmp = product % 10 + to_next;
    var to_next = Math.floor( product / 10 );
    num2_arr[i] = tmp;
  }
  if (to_next > 0) num2_arr.unshift(to_next);
  for (var j = 0; j < exponent; j ++) {
    num2_arr.push(0);
  }
  return num2_arr;
}

var get_sum_arr = function(sum_arr, num2_arr) {
  var sum_arr_len =  sum_arr.length;
  var num2_arr_len = num2_arr.length;

  var len = sum_arr_len > num2_arr_len ? sum_arr_len : num2_arr_len;

  for (var i = 0; i < len - sum_arr_len; i++) {
    sum_arr.unshift(0);
  }

  for (var i = 0; i < len - num2_arr_len; i++) {
    num2_arr.unshift(0);
  }

  var to_next = 0;
  for (var i = len - 1; i >= 0; i --) {
    var sum = Number( sum_arr[i] ) + Number( num2_arr[i] ) + to_next;
    to_next = Math.floor( sum/10 );
    sum_arr[i] = sum % 10;
  }

  if (to_next > 0) sum_arr.unshift(to_next);
  return sum_arr;
}

// console.log(multiply('123456789', '987654321'));