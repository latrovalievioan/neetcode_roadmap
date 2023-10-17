// https://leetcode.com/problems/largest-rectangle-in-histogram/
// TODO: solve again
package main

import "fmt"

func main() {
	test := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(largestRectangleArea(test))
}

func largestRectangleArea(heights []int) int {
	maxArea := 0
	stack := []tuple{}

	for i, h := range heights {

		currentLowestIndex := i

		for len(stack) > 0 && stack[len(stack)-1].height > h {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			currentLowestIndex = top.lowestIndex

			currentMaxArea := top.height * (i - top.lowestIndex)

			if currentMaxArea > maxArea {
				maxArea = currentMaxArea
			}
		}

		stack = append(stack, tuple{lowestIndex: currentLowestIndex, height: h})
	}

	for _, t := range stack {
		currentArea := t.height * (len(heights) - t.lowestIndex)
		if currentArea > maxArea {
			maxArea = currentArea
		}
	}

	return maxArea
}

type tuple struct {
	lowestIndex int
	height      int
}
