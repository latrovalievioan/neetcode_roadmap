//https://leetcode.com/problems/evaluate-reverse-polish-notation/
//
// TODO: BEtter solution
const evalRPN = (tokens: string[]): number => {
  const operators = new Set(["+", "-", "*", "/"])

  const stack: string[] = []

  tokens.forEach(t => {
    if (operators.has(t)) {
      const n2 = stack.pop()
      const n1 = stack.pop()

      stack.push(eval(`Math.trunc(${n1} ${t} ${n2})`))

    } else {
      stack.push(t)
    }
  })

  return Number(stack[0])
};

console.log(evalRPN(["4", "-2", "/", "2", "-3", "-", "-"]))
