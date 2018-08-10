arr = [1, 2, 3 , 4, 5, 6, 7, 8];
var target = 5;

function binarySearch(target, arr, left, right) {
  var start = left || 0;
  var end = right || arr.length - 1;
  var mind = parseInt((start + end)/2);
  if(arr[mind] === target) {
    return mind;
  } else if(arr[mind] > target) {
    return binarySearch(target, arr, start, mind - 1);
  } else {
    return binarySearch(target, arr, mind + 1, end);
  }
  return -1;
}

console.log(target == arr[binarySearch(5, arr)]);

function binarySearch(target, arr, start, end) {
  var start = start || 0;
  var end = end || arr.length - 1;

  while(start <= end) {
    let mind = parseInt((start + end)/2);
    if(arr[mind] === target) {
      return mind;
    } else if(target > arr[mind]) {
      start = mind + 1;
    } else {
      end = mind - 1;
    }
  }
  return -1;
}

function binarySearch(arr, target) {
  
}