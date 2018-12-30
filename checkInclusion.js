/**
 * @param {string} s1
 * @param {string} s2
 * @return {boolean}
 */
var checkInclusion = function(s1, s2) {
  if ( s1.length > s2.length ) return false;
  var s1_char_count = {};
  var s2_char_count = {};
  for ( var i = 0; i < s1.length; i++ ) {
      if ( s1_char_count[s1[i]] ) {
          s1_char_count[s1[i]] ++;
      } else {
          s1_char_count[s1[i]] = 1;
      }

      if ( s2_char_count[s2[i]] ) {
          s2_char_count[s2[i]] ++;
      } else {
          s2_char_count[s2[i]] = 1;
      }
  }

  if ( objEqual(s1_char_count, s2_char_count) ) return true;

  var start_index = 0;
  for ( var j = s1.length; j < s2.length; j ++ ) {
      var char_start  = s2[start_index];
      var chart_end = s2[j];
      s2_char_count[char_start] --;
      if ( s2_char_count[chart_end] ) {
          s2_char_count[chart_end] ++;
      } else {
          s2_char_count[chart_end] = 1;
      }

      if ( objEqual(s1_char_count, s2_char_count) ) return true;

      start_index ++;
  }

  return false;
};


function objEqual(obj1, obj2) {
  for ( key in obj1 ) {
      if ( obj1[key] !== obj2[key]) return false;
  }

  return true
}
