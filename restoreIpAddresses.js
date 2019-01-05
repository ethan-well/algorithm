/**
 * @param {string} s
 * @return {string[]}
 */
var restoreIpAddresses = function(s) {
  var len = s.length;
  var result =[];
  if ( len <  4 || len > 12 ) return result;

  for ( var i = 1; i < 4; i++) {

    var s1 = s.substr(0, i);
    if ( isInvalidNum(s1) ) continue;

    for (var j = i+1; j < i+4; j++ ) {

      var s2 = s.substring(i, j);
      if ( isInvalidNum(s2) ) continue;

      for (var k = j+1; k < j+4; k++ ) {

        var s3 = s.substring(j, k );
        if ( isInvalidNum(s3) ) continue;
        var s4 = s.substring(k);
        if ( isInvalidNum(s4) ) continue;
        var ip = s1 + '.' + s2 + '.' + s3 + '.' + s4;
        result.push(ip);
      }
    }
  }

  return result;
};


var  isInvalidNum = function(s) {
  if ( ( Number(s) == s ) && 0 < Number(s) && Number(s) <= 255 && String(Number(s)) === s )  return false
  return true
}


console.log(restoreIpAddresses('25525511135'));
// console.log(isInvalidNum('511'));