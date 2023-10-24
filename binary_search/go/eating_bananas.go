//https://leetcode.com/problems/koko-eating-bananas/
package main

import (
	"fmt"
)

func main() {
	xs := []int{3, 6, 7, 11}
	h := 8
	fmt.Println(minEatingSpeed(xs, h))
}

func minEatingSpeed(piles []int, h int) int {
	maxSpeed := sliceMax(piles)
	minSpeed := 1
	minFinishingSpeed := 0

	for minSpeed <= maxSpeed {
		avgSpeed := (minSpeed + maxSpeed) / 2

		if canEatInTime(piles, avgSpeed, h) {
			minFinishingSpeed = avgSpeed
			maxSpeed = avgSpeed - 1
		} else {
			minSpeed = avgSpeed + 1
		}
	}

	return minFinishingSpeed
}

func canEatInTime(piles []int, speed int, maxTime int) bool {
	timeEating := 0
	for _, v := range piles {
		if timeEating > maxTime {
			return false
		}
		division := v / speed
		timeEating += division
		remainder := v % speed
		if remainder > 0 {
			timeEating++
		}
	}

	return timeEating <= maxTime
}

func sliceMax(slice []int) int {
	max := 0
	for _, v := range slice {
		if v > max {
			max = v
		}
	}

	return max
}
