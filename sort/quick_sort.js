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