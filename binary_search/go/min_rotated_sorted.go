// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/
// THIS WAS FUCKING HARD >>>> I DONT KNOW WHY >>>> SOLVE AGAIN!!!!
package main

import "fmt"

func main() {
	input := []int{4, 5, 6, 7, 0, 1, 2}

	fmt.Println(findMin(input))
}

func findMin(nums []int) int {
	leftP, rightP := 0, len(nums)-1
	min := nums[0]

	for leftP < rightP {
		if nums[rightP] > nums[leftP] {
			min = getMin(nums[leftP], min)
			break
		}

		mid := (leftP + rightP) / 2
		min = getMin(nums[mid], min)

		if nums[mid] >= nums[0] {
			leftP = mid + 1
		} else {
			rightP = mid - 1
		}
	}

	return min
}

func getMin(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
