// https://leetcode.com/problems/contains-duplicate/
//
function containsDuplicate(nums: number[]): boolean {
  const set = new Set(nums);

  return nums.length != set.size
};
