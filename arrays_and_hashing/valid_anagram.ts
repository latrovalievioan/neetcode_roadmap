// https://leetcode.com/problems/valid-anagram/
//
// TODO: find a more efficient solution

function isAnagram(s: string, t: string): boolean {
  if (s.length !== t.length) return false

  for (let i = 0; i < s.length; i++) {
    if (s.split(s[i]).length !== t.split(s[i]).length) return false
  }

  return true
};
