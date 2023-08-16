// https://leetcode.com/problems/product-of-array-except-self/

// TODO solve this again alone 
function productExceptSelf(nums: number[]): number[] {
  const result: number[] = new Array(nums.length);

  for (let i = 0, acc = 1; i < nums.length; i++) {
    result[i] = acc
    acc *= nums[i]
  }

  for (let i = nums.length - 1, acc = 1; i >= 0; i--) {
    result[i] *= acc
    acc *= nums[i]
  }

  return result
};

console.log(productExceptSelf([1, 2, 3, 4]));
