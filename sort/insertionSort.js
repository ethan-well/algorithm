to_be_sorted = [26, 8, 17, 4, 29, 10, 10]

function insertionSort(arr) {
  for (let i = 1; i < arr.length; i++) {
    let key = arr[i];
    let j = i - 1
    while(j >= 0 && key < arr[j]) {
      arr[j+1] = arr[j];
      j--;
    }
    arr[j+1] = key;
  }
  return arr;
}

console.log(insertionSort(to_be_sorted));
