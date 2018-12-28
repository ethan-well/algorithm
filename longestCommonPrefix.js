/**
 * @param {string[]} strs
 * @return {string}
 */
var longestCommonPrefix = function(strs) {
  if (strs.length == 0) { return "" }
  var comm_arr = [];
  var stop = false;
  for (var i = 0; i < strs[0].length; i++) {
      count = 1;
      if (stop) break;

      for (var j = 1; j < strs.length; j++) {
          if (strs[j].length < i+1) {
              stop = true;
              break;
          }
          if (strs[j][i] !== strs[0][i]) {
              stop = true;
              break;
          }
          count += 1;
      }
      if (count === strs.length) comm_arr.push(strs[0][i]);
  }

  return comm_arr.join('');
};