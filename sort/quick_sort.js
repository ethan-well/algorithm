to_be_sorted = [26, 8, 17, 4, 29, 10, 10]

function sort_and_get_tag(arr, left, right) {
  let tmp = arr[left], i = left, j = right;
  if(i < j) {
    while(i < j && arr[j] >= tmp) {
      j--;
    }
    if(i < j) {
      arr[i] = arr[j];
      i++;
    }
    
    while(i < j && arr[i] <= tmp) {
      i++;
    }
    if(i < j) {
      arr[j] = tmp;
      j--;
    }
  }
  arr[i] = tmp;
  return i;
}

function quickSort(arr, left, right) {
  if (left > right)
    return;
  let i = sort_and_get_tag(arr, left, right);
  quickSort(arr, left, i - 1);
  quickSort(arr, i + 1, right);
}

to_be_sorted = [26, 8, 17, 4, 29, 10, 10]
quickSort(to_be_sorted, 0, to_be_sorted.length -1)
console.log(to_be_sorted);


function quickSort(array, left, right) {
  console.time('1.快速排序耗时');
  if (Object.prototype.toString.call(array).slice(8, -1) === 'Array' && typeof left === 'number' && typeof right === 'number') {
      if (left < right) {
          var x = array[right], i = left - 1, temp;
          for (var j = left; j <= right; j++) {
              if (array[j] <= x) {
                  i++;
                  temp = array[i];
                  array[i] = array[j];
                  array[j] = temp;
              }
          }
          quickSort(array, left, i - 1);
          quickSort(array, i + 1, right);
      }
      console.timeEnd('1.快速排序耗时');
      return array;
  } else {
      return 'array is not an Array or left or right is not a number!';
  }
}