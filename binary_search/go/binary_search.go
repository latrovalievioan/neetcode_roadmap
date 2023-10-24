// https://leetcode.com/problems/binary-search/
// Question: why don't we check for left and right values the first time?
package main

import "fmt"

func main() {
	xs := []int{-1, 0, 3, 5, 9, 12}
	target := 9

	fmt.Println(search(xs, target))
}

func search(nums []int, target int) int {
	leftP, rightP := 0, len(nums)-1

	for leftP <= rightP {
		mid := (leftP + rightP) / 2
		if nums[mid] < target {
			leftP = mid + 1
		} else if nums[mid] > target {
			rightP = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
