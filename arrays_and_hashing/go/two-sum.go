// https://leetcode.com/problems/two-sum/
package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
	fmt.Println(twoSum2(nums, target))
}

func twoSum(nums []int, target int) []int {

	for i := range nums {
		currentValue := nums[i]
		searchingFor := target - currentValue
		found := includes(nums, searchingFor, i)
		if found != -1 {
			return []int{i, found}
		}
	}

	return []int{}
}

func includes(nums []int, value int, omit int) int {
	for i := range nums {
		if nums[i] == value && i != omit {
			return i
		}
	}

	return -1
}

////////

func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)

	for i := range nums {
		m[nums[i]] = i
	}

	for i := range nums {
		value, ok := m[target-nums[i]]

		if ok && value != i {
			return []int{i, value}
		}
	}

	return []int{}
}
