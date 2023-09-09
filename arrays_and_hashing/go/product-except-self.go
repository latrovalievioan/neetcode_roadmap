// https://leetcode.com/problems/product-of-array-except-self/
// TODO: solve again
package main

import (
	"fmt"
)

func main() {
	test := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(test))
}

func productExceptSelf(nums []int) []int {
	prefix := 1
	prefixSlice := make([]int, len(nums))
	postfix := 1
	postfixSlice := make([]int, len(nums))
	result := make([]int, len(nums))

	for i, v := range nums {
		prefixSlice[i] = prefix
		prefix *= v
	}

	for i := len(nums) - 1; i >= 0; i-- {
		postfixSlice[i] = postfix
		postfix *= nums[i]
	}

	for i, v := range prefixSlice {
		result[i] = v * postfixSlice[i]
	}

	return result
}
