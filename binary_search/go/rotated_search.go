//https://leetcode.com/problems/search-in-rotated-sorted-array/submissions/
// this was terrible needs emergency solving again!
package main

import "fmt"

func main() {
	xs := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0

	fmt.Println(search(xs, target))
}

func search(nums []int, target int) int {
	leftP, rightP := 0, len(nums)-1

	for leftP <= rightP {
		midP := (leftP + rightP) / 2

		if nums[midP] == target {
			return midP
		}

		if nums[leftP] <= nums[midP] {
			// mid is in the left
			if target > nums[midP] || target < nums[leftP] {
				leftP = midP + 1
			} else {
				rightP = midP - 1
			}
		} else {
			// mid is in the right
			if target < nums[midP] || target > nums[len(nums)-1] {
				rightP = midP - 1
			} else {
				leftP = midP + 1
			}
		}
	}

	return -1
}
