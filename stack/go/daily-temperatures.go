// https://leetcode.com/problems/daily-temperatures
// TODO: solve again
package main

import "fmt"

func main() {
	test := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(test))
}

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := []pair{}

	for i, v := range temperatures {
		for len(stack) > 0 && stack[len(stack)-1].v < v {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			res[top.i] = i - top.i
		}

		stack = append(stack, pair{v, i})
	}

	return res
}

type pair struct {
	v int
	i int
}
