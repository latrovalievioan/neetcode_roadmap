// https://leetcode.com/problems/container-with-most-water/submissions/
package main

import "fmt"

func main() {
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(input))
}

func maxArea(height []int) int {
	leftP := 0
	rightP := len(height) - 1
	maxArea := 0

	for leftP < rightP {
		bottleNeck := min(height[leftP], height[rightP])
		width := rightP - leftP
		area := bottleNeck * width

		if area > maxArea {
			maxArea = area
		}

		if bottleNeck == height[leftP] {
			leftP++
		} else {
			rightP--
		}
	}

	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
