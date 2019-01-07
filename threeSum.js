/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var threeSum = function(nums) {
  var arrs = [];
  var nums = nums.sort(compare);
  for ( var i = 0; i < nums.length; i++) {
    if ( nums[i] > 0 ) break;
    for (var j = i + 1; j < nums.length; j++) {
      if ( nums[i] + nums[j] > 0) break;
      for (var k = j + 1; k < nums.length; k++) {
        var sum  = nums[i] + nums[j] + nums[k];
        if ( sum > 0) break;
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

var compare = function(a, b) {
  return a - b;
}

// arrt = [-13,5,13,12,-2,-11,-1,12,-3,0,-3,-7,-7,-5,-3,-15,-2,14,14,13,6,-11,-11,5,-15,-14,5,-5,-2,0,3,-8,-10,-7,11,-5,-10,-5,-7,-6,2,5,3,2,7,7,3,-10,-2,2,-12,-11,-1,14,10,-9,-15,-8,-7,-9,7,3,-2,5,11,-13,-15,8,-3,-7,-12,7,5,-2,-6,-3,-10,4,2,-5,14,-3,-1,-10,-3,-14,-4,-3,-7,-4,3,8,14,9,-2,10,11,-10,-4,-15,-9,-1,-1,3,4,1,8,1]
// console.log(threeSum(arrt));