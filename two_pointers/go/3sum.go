// https://leetcode.com/problems/3sum/
package main

import (
	"fmt"
	"sort"
)

func main() {
	input := []int{-1, -1, 0, 1}

	fmt.Println(threeSum(input))
}

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	result := [][]int{}

	lookup := make(map[int]bool)

	for i, v := range nums {
		ok := lookup[v]
		if ok {
			continue
		}
		lookup[v] = true

		twoSums := twoSumWithExtra(nums[i+1:], v)

		for _, v1 := range twoSums {
			result = append(result, v1)
		}
	}

	return result
}

func twoSumWithExtra(sortedNums []int, extra int) [][]int {
	leftP := 0
	rightP := len(sortedNums) - 1
	result := [][]int{}
	lookup := make(map[int]bool)

	for leftP < rightP {
		okL := lookup[sortedNums[leftP]]
		if okL {
			leftP++
			continue
		}

		okR := lookup[sortedNums[rightP]]
		if okR {
			rightP--
			continue
		}

		if extra+sortedNums[leftP]+sortedNums[rightP] == 0 {
			result = append(result, []int{extra, sortedNums[leftP], sortedNums[rightP]})
			lookup[sortedNums[leftP]] = true
			lookup[sortedNums[rightP]] = true
			leftP++
			rightP--
			continue
		}

		if extra+sortedNums[leftP]+sortedNums[rightP] < 0 {
			lookup[sortedNums[leftP]] = true
			leftP++
			continue
		}

		if extra+sortedNums[leftP]+sortedNums[rightP] > 0 {
			lookup[sortedNums[rightP]] = true
			rightP--
			continue
		}
	}
	return result
}
