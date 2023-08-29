// https://leetcode.com/problems/contains-duplicate/description/
package main

import (
	"fmt"
)

func main() {
	test := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(test))
}

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)

	for i := range nums {
		current := nums[i]
		_, ok := m[current]

		if ok {
			return true
		} else {
			m[current] = true
		}
	}

	return false
}
