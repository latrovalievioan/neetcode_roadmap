//https://leetcode.com/problems/valid-parentheses/
//
//TODO: I don't like line 11
const isValid = (s: string): boolean => {
  const dict = {
    ")": "(",
    "]": "[",
    "}": "{",
  }

  const stack: string[] = [];

  s.split('').forEach(c => {
    if (stack[stack.length - 1] === dict[c] && dict[c] !== undefined) {
      stack.pop()
    } else {
      stack.push(c)
    }
  })

  return !stack.length
};


console.log(isValid("()[]}{"))
