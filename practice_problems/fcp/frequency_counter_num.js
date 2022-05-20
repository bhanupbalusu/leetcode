console.log("Frequency Counter Pattern")

// a) Understand the problem
// - Restate the problem
// -> All values in array one must have corresponding squared values in array two and should
// match the frequency and maintaining order is not necessary
// - Inputs that go into this problem
// -> An array of integers
// - Outputs from the solution
// -> An array of squared integers
// - Is the information enough
// -> yes
// - Labeling the data
// - Create a function that takes two input arrays and returns boolean

// b) Explore concrete example
// - Start with simple examples
// -> arr1=[1, 2, 3, 4] Arr2=[1, 4, 9, 16] o/p=true

// c) Break it down
// - Write steps to solve the problem
// -> create a function that takes two integer arrays as input and returns boolean
// -> check if the length of input arrays are same if yes continue or else return false
// -> loop the first array and check each element in the first array has the corresponding squared element
// -> if yes then remove that element from 1st array and repeat above process until first array is empty
// -> if no then return false
// -> return true after the loop

// d) Slove the problem
function same(arr1, arr2) {
  if (arr1.length != arr2.length) {
    return false
  }
  for (let i = 0; i < arr1.length; i++) {
    let idx = arr2.indexOf(arr1[i] ** 2)
    if (idx == -1) {
      return false
    }
    arr2.splice(idx, 1)
  }
  return true
}

// console.log(same([1, 2, 3, 4], [16, 9, 4, 1]))

// Following frequency counter pattern
function sameFCP(a1, a2) {
  if (a1.length != a2.length) {
    return false
  }

  let fc1 = {}
  let fc2 = {}

  console.log(a1)
  console.log(a2)

  for (let val of a1) {
    fc1[val] = (fc1[val] || 0) + 1
  }
  for (let val of a2) {
    fc2[val] = (fc2[val] || 0) + 1
  }

  console.log(fc1)
  console.log(fc2)

  for (let key in fc1) {
    if (!(key ** 2 in fc2)) {
      return false
    }
    if (fc2[key ** 2] != fc1[key]) {
      return false
    }
  }
  return true
}

console.log(sameFCP([1, 2, 3, 4], [16, 9, 4, 1]))