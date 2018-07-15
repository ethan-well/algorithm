to_be_sorted = [26, 8, 17, 4, 29, 10, 10]

function bubbleSort(arr){
  len = arr.length;

  for(let i = 0; i < len; i++) {
    for(let j = 1; j < len - i; j ++) {
      if(arr[j] < arr[j-1]) {
        tmp = arr[j];
        arr[j] = arr[j-1];
        arr[j-1] = tmp;
      }
    }
  }
  return arr;
}
console.log(bubbleSort(to_be_sorted));

function bubbleSort2(arr){
  pos = arr.length;
  for(let i = 0; i < pos; i++) {
    for(let j = 0; j < pos - i - 1; j++) {
      if(arr[j] > arr[j+1]) {
        pos = j;
        let tmp = arr[j];
        arr[j] = arr[j+1];
        arr[j+1] = tmp;
      }
    }
  }
  return arr;
}
console.log(bubbleSort2(to_be_sorted));