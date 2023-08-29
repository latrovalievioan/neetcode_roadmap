// https://leetcode.com/problems/valid-anagram/

const isAnagram = (s: string, t: string): boolean => {
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

console.log(isAnagram('anagram', 'nagaram'))
