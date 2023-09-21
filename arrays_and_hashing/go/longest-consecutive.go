// https://leetcode.com/problems/longest-consecutive-sequence/
// TODO: solve again

package main

import (
	"fmt"
)

func main() {
	test := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(test))
}

func longestConsecutive(nums []int) int {
	m := make(map[int]bool)

	for _, n := range nums {
		m[n] = true
	}

	longest := 0

	for _, n := range nums {
		isBeginning := !m[n-1]

		if !isBeginning {
			continue
		}

		inc := 1
		last := m[n+inc]

		for last {
			inc++
			last = m[n+inc]
		}

		if inc > longest {
			longest = inc
		}

	}

	return longest
}
