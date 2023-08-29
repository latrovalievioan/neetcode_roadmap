//https://leetcode.com/problems/group-anagrams/
//
// TODO: find a better solution

const _isAnagram = (s: string, t: string): boolean => {
  if (s.length !== t.length) return false;

  const alphabet = new Uint32Array(26);

  const xs = [...s, ...t];

  for (let i = 0, j = xs.length - 1; i < j; i++, j--) {
    const leftCharCodeIndex = xs[i].charCodeAt(0) - 97
    const rightCharCodeIndex = xs[j].charCodeAt(0) - 97

    alphabet[leftCharCodeIndex]++
    alphabet[rightCharCodeIndex]--
  }
  return alphabet.every(v => v === 0)
}

function groupAnagrams(strs: string[]): string[][] {
  const result: string[][] = []
  while (strs.length) {
    const anagramIndexes: number[] = [];
    const str = strs.shift()
    if (str === undefined) break
    strs.forEach((s, i) => {
      if (_isAnagram(s, str)) {
        anagramIndexes.push(i)
      }
    })

    result.push([str, ...anagramIndexes.map(x => strs[x])])
    strs = strs.filter((_, i) => !anagramIndexes.includes(i))
  }

  return result
};

console.log(groupAnagrams([""]))
