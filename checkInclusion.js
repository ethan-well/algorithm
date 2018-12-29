var checkInclusion = function(s1, s2) {
  if ( s1.length > s2.length ) return false;
  var matched_indexs = {};
  var matched_indexs_arr = []
  for (var i = 0; i < s1.length; i++) {
      var start_index = matched_indexs[s1[i]] ?  matched_indexs[s1[i]][-1] : -1;
      var next_match_index = s2.indexOf(s1[i], start_index + 1);

      if ( next_match_index === -1 ) return false;


      if (matched_indexs[s1[i]]) {
          matched_indexs[s1[i]].push(next_match_index);
      } else {
          matched_indexs[s1[i]] = [next_match_index];
      }
  }

  console.log(matched_indexs);
  for ( char in matched_indexs ) {
      if ( matched_indexs.hasOwnProperty(char) ) {
          matched_indexs_arr = matched_indexs_arr.concat(matched_indexs[char]);
      }
  }

  if ( matched_indexs_arr.length !== s1.length ) return false;
  var matched_indexs_arr_sort = matched_indexs_arr.sort();

  return matched_indexs_arr_sort.every(is_increment_arr)
};

function is_increment_arr(e, index, arr) {
  if( index === 0) return true;
  return ( e - arr[index - 1] === 1 );
}

var s1 = 'adc';
var s2 = 'dcda';
console.log( checkInclusion(s1, s2) );