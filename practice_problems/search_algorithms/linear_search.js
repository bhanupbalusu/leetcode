function linearSearch(arr, num) {
  // add whatever parameters you deem necessary - good luck!
  for (i = 0; i < arr.length; i++) {
    if (arr[i] === num) {
      return i
    }
  }
  return -1
}

console.log(linearSearch([10, 15, 20, 25, 30], 15))