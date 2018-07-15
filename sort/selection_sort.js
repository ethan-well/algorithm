to_be_sorted = [26, 8, 17, 4, 29, 10, 10]

function selectionSort(arr) {
  let len = arr.length;
  let minIndex, temp;
  for(let i = 0; i < len - 1; i ++) {
    minIndex = i;
    for(j = i + 1; j < len; j ++ ) {
      if(arr[j] < arr[minIndex]) {
        minIndex = j;
      }
    }
    temp = arr[i];
    arr[i] = arr[minIndex];
    arr[minIndex] = temp;
  }
  return arr;
}

console.log(selectionSort(to_be_sorted));