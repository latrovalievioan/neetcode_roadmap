//https://leetcode.com/problems/top-k-frequent-elements/

function topKFrequent(nums: number[], k: number): number[] {
  const frequencyMap = new Map()

  nums.forEach(n => {
    if (frequencyMap.has(n)) {
      frequencyMap.set(n, frequencyMap.get(n) + 1)
    } else {
      frequencyMap.set(n, 1)
    }
  })

  const sorted = [...frequencyMap.entries()].sort((a, b) => a[1] - b[1]).reverse()

  return sorted.map(tuple => tuple[0]).slice(0, k)
};

console.log(topKFrequent([4, 1, -1, 2, -1, 2, 3], 2))
