/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var threeSum = function(nums) {
  var arrs = [];
  for ( var i = 0; i < nums.length; i++) {
    for (var j = i + 1; j < nums.length; j++) {
      for (var k = j + 1; k < nums.length; k++) {
        var sum  = nums[i] + nums[j] + nums[k];
        if ( sum === 0 ) {
          var arr = [nums[i], nums[j], nums[k]].sort();
          if (alreadyHas(arr, arrs)) continue;
          arrs.push(arr);
        }
      }
    }
  }

  return arrs;
};

var alreadyHas = function (arr, arrs) {
  if (arrs.length === 0) return false;
  for (var i = 0; i < arrs.length; i++) {
    if ( arrs[i][0] == arr[0] && arrs[i][1] == arr[1] && arrs[i][2] == arr[2] ) {
      return true
    }
  }

  return false;
}

console.log(threeSum([-1,0,1,2,-1,-4]));