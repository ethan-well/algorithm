/**
 * @param {string} str
 * @returns {string}
 */
var reverseWords = function(str) {
  var str_arr = str.split(' ');
  var arr_filtered = str_arr.filter(function(item) { return item !== ""});
  return arr_filtered.reverse().join(' ');
};

// console.log(reverseWords('test1 test2 test3  test4  test5'));