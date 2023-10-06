package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	test := []string{"2", "1", "+", "3", "*"}
	fmt.Println(evalRPN(test))
}

func evalRPN(tokens []string) int {
	operators := map[string]struct{}{
		"+": {},
		"-": {},
		"*": {},
		"/": {},
	}
	stackSlice := []int64{}

	for _, v := range tokens {
		_, ok := operators[v]

		if ok {
			stackLength := len(stackSlice)
			a := stackSlice[stackLength-1]
			b := stackSlice[stackLength-2]
			stackSlice = stackSlice[:stackLength-2]

			stackSlice = append(stackSlice, operation(a, b, v))
			continue
		}

		parsed, _ := strconv.ParseInt(v, 10, 64)

		stackSlice = append(stackSlice, parsed)
	}

	return int(stackSlice[0])
}

func operation(a int64, b int64, operator string) int64 {
	switch operator {
	case "+":
		return a + b
	case "*":
		return a * b
	case "-":
		return a - b
	case "/":
		return a / b
	default:
		return int64(math.NaN())
	}
}
