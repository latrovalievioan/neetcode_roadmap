package main

import "fmt"

func main() {
	xs := []int{1, 3, 4, 5, 7, 10, 11}

	fmt.Println(twoSum(xs, 9))
}

func twoSum(numbers []int, target int) []int {
	leftP, rightP := 0, len(numbers)-1
	sum := numbers[leftP] + numbers[rightP]

	for sum != target {
		if sum > target {
			rightP--
		} else {
			leftP++
		}

		sum = numbers[leftP] + numbers[rightP]
	}

	return []int{leftP + 1, rightP + 1}
}
