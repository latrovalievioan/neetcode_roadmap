// https://leetcode.com/problems/trapping-rain-water/submissions/
// TODO find a way to simplify the code
package main

import "fmt"

func main() {
	input := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(input))
}

func trap(heights []int) int {
	leftP, rightP, leftMax, rightMax, result := 0, len(heights)-1, 0, 0, 0

	for leftP < rightP {
		if leftMax-heights[leftP] > 0 {
			result += leftMax - heights[leftP]
		}

		if rightMax-heights[rightP] > 0 {
			result += rightMax - heights[rightP]
		}

		if heights[leftP] > leftMax {
			leftMax = heights[leftP]
		}

		if heights[rightP] > rightMax {
			rightMax = heights[rightP]
		}

		if rightMax < leftMax {
			rightP--
		} else {
			leftP++
		}
	}

	return result
}
