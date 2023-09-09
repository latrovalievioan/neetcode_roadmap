// https://leetcode.com/problems/top-k-frequent-elements/
// Note: this uses a variation of bucket sort as shown in the roadmap video
// TODO: faster solution
package main

import (
	"fmt"
)

func main() {
	nums := []int{1}
	k := 1
	fmt.Println(topKFrequent(nums, k))
}
func topKFrequent(nums []int, k int) []int {
	maxLen := len(nums)
	frequencySlice := make([][]int, maxLen+1)
	frequencyMap := make(map[int]int)
	result := make([]int, k)

	for _, v := range nums {
		_, ok := frequencyMap[v]
		if ok {
			frequencyMap[v]++
		} else {
			frequencyMap[v] = 1
		}
	}

	for k, v := range frequencyMap {
		frequencySlice[v] = append(frequencySlice[v], k)
	}

	for i, position := len(frequencySlice)-1, 0; k > 0; i-- {
		current := frequencySlice[i]
		for _, v := range current {
			result[position] = v
			position++
			k--
		}
	}

	return result
}
