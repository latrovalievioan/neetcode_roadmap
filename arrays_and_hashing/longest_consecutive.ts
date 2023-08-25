//https://leetcode.com/problems/longest-consecutive-sequence/
// TODO: Ask why is this 0(n) doenst accessing set value add more?
// TODO: Ask why while cycle doesnt add to complexity
// TODO: Solve again

function longestConsecutive(nums: number[]): number {
  if (!nums.length) return 0
  const set = new Set(nums);
  let longest = 0;

  nums.forEach(n => {
    if (set.has(n - 1)) return
    let length = 0

    while (set.has(n + length)) {
      length++
    }

    if (length > longest) longest = length

  })
  return longest
};



const test = [100, 4, 200, 1, 3, 2];
console.log(longestConsecutive(test))

