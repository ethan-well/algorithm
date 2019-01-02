var simplifyPath = function(path) {
  var str_arr = path.split('/');
  var arr_filtered = str_arr.filter(filter_blank_char).filter(filter_current_path_char);

  return get_absolute_path(arr_filtered);
};

var filter_blank_char = function(item) {
  return (item !== "");
};

var filter_current_path_char = function(item) {
  return item !== '.';
};

var get_absolute_path = function(path_arr) {
  var path = [];
  for ( var i = 0; i < path_arr.length; i++ ) {
    if (path_arr[i] === '..') {
      if ( path.length !== 0 ) path.pop();
    } else {
      path.push(path_arr[i]);
    }
  }
  return '/' + path.join('/');
};